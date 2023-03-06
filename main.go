package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/audetv/fct-parser/config"
	"github.com/gosimple/slug"
	flag "github.com/spf13/pflag"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Topic struct {
	Question        Comment   `json:"question"`
	LinkedQuestions []Comment `json:"linked_question"`
	Comments        []Comment `json:"comments"`
}

const TypeQuestion = "1"
const TypeLinkedQuestion = "2"
const TypeComment = "3"

type Comment struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Text     string `json:"text"`
	Datetime string `json:"datetime"`
	DataID   string `json:"data_id,omitempty"`
	ParentID string `json:"parent_id"`
	Type     string `json:"type"`
	Count    string `json:"count,omitempty"`
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

var fJson = "json"
var fCsv = "csv"
var format = fCsv

var jsonFormat,
	indent,
	showAll,
	list,
	current,
	htmlTags,
	parseFct,
	updateConfig bool

var outputPath string

func main() {

	flag.BoolVarP(&showAll, "all", "a", false, "сохранение всего списка обсуждений событий с начала СВОДД в отдельные файлы")
	flag.BoolVarP(&list, "list", "l", false, "вывод в консоль списка адресов страниц с обсуждениями событий с начала СВОДД")
	flag.BoolVarP(&current, "current", "c", false, "вывод в консоль адреса ссылки текущего активного обсуждения событий с начала СВОДД")
	flag.BoolVarP(&jsonFormat, "json", "j", false, "вывод в формате json (по умолчанию \"csv\")")
	flag.BoolVarP(&indent, "json-indent", "i", false, "форматированный вывод json с отступами и переносами строк")
	flag.BoolVarP(&htmlTags, "html-tags", "h", false, "вывод с сохранение с html тегов")
	flag.BoolVarP(&parseFct, "parse-fct", "p", false, "парсить все вопросы с сайта")
	flag.BoolVarP(&updateConfig, "update", "u", false, "загрузить конфиг файл")
	flag.StringVarP(&outputPath, "output", "o", "./", "путь сохранения файлов")

	flag.Parse()

	if updateConfig {
		file, err := os.Create("./config.json")
		if err != nil {
			log.Fatalf("update config: %v", err)
		}
		config.DownloadConfigFile(file)
		return
	}
	conf := config.ReadConfig()

	if jsonFormat || indent {
		format = fJson
	}

	if list {
		conf.PrintList()
		return
	}

	if current {
		conf.PrintCurrentDiscussion()
		return
	}

	processAllQuestions(conf)
}

func processAllQuestions(conf config.Config) {

	if showAll {
		conf.IsValidConfig()
		for _, item := range conf.List {
			err := processUrl(item)
			if err != nil {
				log.Printf("skipped: %v", err)
				continue
			}
		}
		return
	}

	if parseFct {
		var item config.Item
		for i := 2419; i < 10000; i++ {
			item.Id = i
			item.Url = fmt.Sprintf("%v%v", "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-", i)
			err := processUrl(item)
			if err != nil {
				log.Printf("skipped: %v", err)
				continue
			}
			log.Printf("Done!")
		}
		return
	}

	if len(flag.Args()) < 1 {
		conf.IsValidConfig()
		err := processUrl(conf.CurrentDiscussion())
		if err != nil {
			log.Printf("skipped: %v", err)
			return
		}
	} else {
		for _, uri := range flag.Args() {
			err := processUrl(config.Item{Url: uri})
			if err != nil {
				log.Printf("skipped: %v", err)
				continue
			}
		}
	}
}

func processUrl(item config.Item) error {

	URI, err := url.ParseRequestURI(item.Url)
	if err != nil {
		return fmt.Errorf("parse request uri: %v\n", err)
	}

	var prefix string
	if item.Num != "" {
		prefix += fmt.Sprintf("%v-", item.Num)
	}

	file := fmt.Sprintf("%v/%v%v.%s", outputPath, prefix, slug.Make(URI.Path), format)
	file = filepath.Clean(file)

	doc, err := getTopicBody(item.Url)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}

	log.Printf("parse %v\n", item.Url)

	parentID := parseViewId(item.Url)
	topic := Topic{}
	topic.parseTopic(parentID, doc)

	if format == fJson {
		writeJsonFile(topic, file, indent)
		log.Printf("file %v was successful writing\n", file)
	}
	if format == fCsv {
		writeCSVFile(topic, file)
		log.Printf("file %v was successful writing\n", file)
	}
	return nil
}

func parseViewId(s string) string {
	defer duration(track("foo"))

	return strings.ReplaceAll(s, "https://фкт-алтай.рф/qa/question/view-", "")

	// nLen := 0
	// for i := 0; i < len(s); i++ {
	// 	if b := s[i]; '0' <= b && b <= '9' {
	// 		nLen++
	// 	}
	// }
	// var n = make([]int, 0, nLen)
	// for i := 0; i < len(s); i++ {
	// 	if b := s[i]; '0' <= b && b <= '9' {
	// 		n = append(n, int(b)-'0')
	// 	}
	// }
	//
	// return strings.Trim(strings.Replace(fmt.Sprint(n), " ", ",", -1), "[]")
}

func getTopicBody(url string) (*html.Node, error) {
	// Для того чтобы не следовать автоматическим перенаправлениям,
	// создадим свой экземпляр http.Client с методом проверки CheckRedirect.
	// Это поможет нам возвращать код состояния и адрес до перенаправления.
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return doc, nil
}

func (topic *Topic) parseTopic(parentID string, doc *html.Node) {
	parseQuestionView(doc, topic, parentID)
	parseCommentList(doc, topic, parentID)
}

func parseQuestionView(n *html.Node, topic *Topic, parentID string) {

	exit := false

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && nodeHasRequiredCssClass("question-view", n) {
			topic.Question = parseComment(n, TypeQuestion, "")
			topic.Question.DataID = parentID
		}
		if n.Type == html.ElementNode && nodeHasRequiredCssClass("linked-questions", n) {
			topic.LinkedQuestions = parseLinkedQuestions(n, parentID)
			exit = true
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if exit == true {
				break
			}
			f(c)
		}
	}
	f(n)
}

func parseLinkedQuestions(n *html.Node, parentID string) []Comment {
	var comments []Comment

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && nodeHasRequiredCssClass("linked-question", n) {
			comments = append(comments, parseComment(n, TypeLinkedQuestion, parentID))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return comments
}

func parseCommentList(n *html.Node, topic *Topic, parentID string) {
	var comments []Comment

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && nodeHasRequiredCssClass("comment-list", n) {
			// проходим по узлу с атрибутом class block comment-item}
			for cl := n.FirstChild; cl != nil; cl = cl.NextSibling {
				if cl.Type == html.ElementNode && nodeHasRequiredCssClass("comment-item", cl) {
					comments = append(comments, parseComment(cl, TypeComment, parentID))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	log.Printf("всего комментариев %v\n", len(comments))
	topic.Comments = comments
}

func parseComment(n *html.Node, t string, parentID string) Comment {

	var nAnchor *html.Node
	var bufInnerHtml bytes.Buffer

	w := io.Writer(&bufInnerHtml)

	comment := Comment{}
	comment.Type = t
	comment.ParentID = parentID

	exit := false

	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.ElementNode && nodeHasRequiredCssClass("username", n) {
			comment.Username = getInnerText(n)
		}

		if n.Type == html.ElementNode && nodeHasRequiredCssClass("role", n) {
			comment.Role = getInnerText(n)
		}

		// Находим ноду с типом ElementNode и атрибутом со значением comment-text
		// Переменной nAnchor присваиваем эту ноду
		if n.Type == html.ElementNode && nodeHasRequiredCssClass("comment-text", n) {
			comment.DataID = getRequiredDataAttr("data-id", n)
			nAnchor = n
		}

		if n.Type == html.ElementNode && nodeHasRequiredCssClass("datetime", n) {
			comment.Datetime = strings.TrimSpace(getInnerText(n))
			exit = true
		}

		// don't write the tag and its attributes
		if nAnchor != nil {
			// Fixed bug. Html render without nested siblings.
			// Не пишем тег и его атрибуты.
			// Но нам нужны все потомки nextSibling, без своих потомков от blockquote и link.
			// проверяем по родительской ноде с атрибутом class comment-text.
			if n != nAnchor && nodeHasRequiredCssClass("comment-text", n.Parent) {
				err := html.Render(w, n)
				if err != nil {
					return
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if exit {
				break
			}
			f(c)
		}

		if n == nAnchor {
			if htmlTags {
				comment.Text = bufInnerHtml.String()
			} else {
				comment.Text = processCommentText(n)
			}

			bufInnerHtml.Reset()
			nAnchor = nil
		}
	}
	f(n)

	return comment
}

func processCommentText(node *html.Node) string {
	var text string
	for el := node.FirstChild; el != nil; el = el.NextSibling {
		if el.Type == html.TextNode {
			text += el.Data
		}
		if el.Type == html.ElementNode && el.Data == "blockquote" {
			text += fmt.Sprintf("%v\n", processBlockquote(el))
		}
		if el.Type == html.ElementNode && nodeHasRequiredCssClass("link", el) {
			text += strings.TrimSpace(getInnerText(el))
		}
	}

	return strings.TrimSpace(text)
}

func processBlockquote(node *html.Node) string {
	var text string
	newline := ""
	for el := node.FirstChild; el != nil; el = el.NextSibling {
		if el.Type == html.TextNode {
			// UnescapeString для Data нужен, чтобы избавляться от &quot; в цитатах
			// для последующего корректного чтения в exel, кстати гугл таблицы корректно обрабатывали эти цитаты и не ломали csv
			text += fmt.Sprintf("%v%v", newline, strings.TrimSpace(html.UnescapeString(el.Data)))
			newline = fmt.Sprintf("\n%v", "")
		}
		if el.Type == html.ElementNode && nodeHasRequiredCssClass("author", el) {
			text += fmt.Sprintf("%v: «", strings.TrimSpace(getInnerText(el)))
		}
		if el.Type == html.ElementNode && nodeHasRequiredCssClass("link", el) {
			text += fmt.Sprintf("\n%v\n", strings.TrimSpace(getInnerText(el)))
		}
	}

	return fmt.Sprintf("%v»", strings.TrimSpace(text))
}

func getInnerText(node *html.Node) string {
	for el := node.FirstChild; el != nil; el = el.NextSibling {
		if el.Type == html.TextNode {
			return el.Data
		}
	}
	return ""
}

func getRequiredDataAttr(rda string, n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == rda {
			return attr.Val
		}
	}
	return ""
}

// Перебирает аттрибуты токена в цикле и возвращает bool
// если в html token найден переданный css class
func nodeHasRequiredCssClass(rcc string, n *html.Node) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			classes := strings.Split(attr.Val, " ")
			for _, class := range classes {
				if class == rcc {
					return true
				}
			}
		}
	}
	return false
}

func writeCSVFile(topic Topic, outputPath string) {
	// Define header row
	headerRow := []string{
		"Username", "Role", "Text", "Datetime", "DataID",
	}

	// Data array to write to CSV
	data := [][]string{
		headerRow,
	}

	data = append(data, []string{
		// Make sure the property order here matches
		// the one from 'headerRow' !!!
		topic.Question.Username,
		topic.Question.Role,
		topic.Question.Text,
		topic.Question.Datetime,
		topic.Question.DataID,
	})

	// Add linked question to output data
	for _, comment := range topic.LinkedQuestions {
		data = addCommentData(data, comment)
	}

	// Add comment list to output data
	for _, comment := range topic.Comments {
		data = addCommentData(data, comment)
	}

	// Create file
	file, err := os.Create(outputPath)
	checkError("Cannot create file", err)
	defer file.Close()

	// Create writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write rows into file
	for _, value := range data {
		err = writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func addCommentData(data [][]string, comment Comment) [][]string {
	return append(data, []string{
		// Make sure the property order here matches
		// the one from 'headerRow' !!!
		comment.Username,
		comment.Role,
		comment.Text,
		comment.Datetime,
		comment.DataID,
	})
}

func writeJsonFile(topic Topic, outputPath string, indent bool) {

	// Create file
	file, err := os.Create(outputPath)
	checkError("Cannot create file", err)
	defer file.Close()

	if indent {
		aJson, err := json.MarshalIndent(topic, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(aJson)
		checkError("Cannot write to the file", err)
	} else {
		aJson, err := json.Marshal(topic)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(aJson)
		checkError("Cannot write to the file", err)
	}
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
