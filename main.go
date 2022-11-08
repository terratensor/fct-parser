package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Topic struct {
	Question  Comment
	Secondary Comment
	// LinkedQuestion []Comment
	// Comments       []Comment
}

type Comment struct {
	Username string
	Role     string
	Text     string
	Datetime string
}

func main() {

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	for _, url := range os.Args[1:] {
		doc, err := getTopicBody(url)
		if err != nil {
			log.Fatalf("parse: %v\n", err)
		}
		topic, err := parseTopic(doc)
		if err != nil {
			log.Fatalf("parse: %v\n", err)
		}

		// w := csv.NewWriter(file)
		//
		// for _, comment := range comments {
		// 	if err := w.Write(comment); err != nil {
		// 		log.Fatalln("error writing record to csv:", err)
		// 	}
		// }

		// Write any buffered data to the underlying writer (standard output).
		// w.Flush()
		//
		// if err := w.Error(); err != nil {
		// 	log.Fatal(err)
		// }

		fmt.Printf("Topic: %v", *topic)
	}
}

func getTopicBody(url string) (*html.Node, error) {
	resp, err := http.Get(url)

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

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func parseTopic(doc *html.Node) (*Topic, error) {

	topic := &Topic{}

	comment := Comment{}
	comment.parseQuestionView(doc)

	topic.Question = comment
	topic.Secondary = comment

	return topic, nil
}

func (comment *Comment) parseQuestionView(n *html.Node) {

	if n.Type == html.ElementNode && nodeHasRequiredCssClass("question-view", n) {
		parseComment(n, comment)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		comment.parseQuestionView(c)
	}
}

func parseComment(n *html.Node, comment *Comment) {

	var nAnchor *html.Node
	var bufInnerHtml bytes.Buffer

	w := io.Writer(&bufInnerHtml)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && nodeHasRequiredCssClass("datetime", n) {
			comment.Datetime = getInnerText(n)
		}

		if n.Type == html.ElementNode && nodeHasRequiredCssClass("username", n) {
			comment.Username = getInnerText(n)
		}

		if n.Type == html.ElementNode && nodeHasRequiredCssClass("role", n) {
			comment.Role = getInnerText(n)
		}

		if n.Type == html.ElementNode && nodeHasRequiredCssClass("comment-text", n) {
			nAnchor = n
		}

		if nAnchor != nil {
			if n != nAnchor { // don't write the tag and its attributes
				html.Render(w, n)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

		if n == nAnchor {
			comment.Text = bufInnerHtml.String()

			bufInnerHtml.Reset()
			nAnchor = nil

		}
	}
	f(n)

	// Заканчиваем парсинг комментария и выходим из функции если комментарий заполнен
	if (comment != &Comment{}) {
		return
	}

}

func getInnerText(node *html.Node) string {
	for el := node.FirstChild; el != nil; el = el.NextSibling {
		if el.Type == html.TextNode {
			return el.Data
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
