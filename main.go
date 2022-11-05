package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

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

			readAttrs(tkn, t)

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

func readAttrs(tkn *html.Tokenizer, t html.Token) {
	for _, attr := range t.Attr {

		if attr.Key == "class" {
			// println(attr.Key)
			// println(attr.Val)
			classes := strings.Split(attr.Val, " ")
			for _, class := range classes {
				// fmt.Printf("Class %d is: %s\n", idx, class)
				if class == "username" {
					fmt.Printf("Username: %s\n", t.Data)
					tkn.Next()
				}
			}
		}

	}
}
