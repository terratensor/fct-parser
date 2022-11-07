package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Question struct {
	Question string
	Username string
	Role     string
}

func readHtmlFromFile(fileName string) (string, error) {

	bs, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func parse(text string) (data []string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string

	var isLi bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()

			readAttrs(tkn, t, "question-view")

			if tokenHasRequiredCssClass(t, "question-view") {
				qw := readQuestionViewBlock(tkn, t)
				fmt.Printf("tokenHasRequiredCssClass: %s\n", qw)
			}

			isLi = t.Data == "li"

		case tt == html.TextToken:

			t := tkn.Token()

			if isLi {
				vals = append(vals, t.Data)
			}

			isLi = false
		}
	}
}

func main() {

	fileName := "response.html"
	text, err := readHtmlFromFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	data := parse(text)
	fmt.Println(data)
}

func readAttrs(tkn *html.Tokenizer, t html.Token, value string) {
	for _, attr := range t.Attr {

		if attr.Key == "class" {
			// println(attr.Key)
			// println(attr.Val)
			classes := strings.Split(attr.Val, " ")
			for _, class := range classes {
				// fmt.Printf("Class %d is: %s\n", idx, class)
				if class == value {
					fmt.Printf("Username: %s\n", t.Data)
					tkn.Next()
				}
			}
		}

	}
}

func readQuestionViewBlock(tkn *html.Tokenizer, t html.Token) Question {

	q := Question{
		"Блок вопроса присутствует",
		"username",
		"Пользователь",
	}
	return q
}

// Перебирает аттрибуты токена в цикле и возвращает bool
// если в html token найден переданный css class
func tokenHasRequiredCssClass(t html.Token, rcc string) bool {
	for _, attr := range t.Attr {
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
