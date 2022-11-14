package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	List List
}

type List []Item

type Item struct {
	Id       int    `json:"id"`
	Num      string `json:"num"`
	Date     string `json:"date"`
	Url      string `json:"url"`
	Comments int    `json:"comments"`
}

var configUrl = "https://raw.githubusercontent.com/audetv/fct-parser/main/config.json"

func ReadConfig() Config {
	file, err := os.OpenFile("./config.json", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err == nil {
		downloadConfigFile(file)
	}

	fileBytes, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Cannot read file", err)
	}

	var conf Config
	err = json.Unmarshal(fileBytes, &conf)
	if err != nil {
		log.Fatal("Cannot unmarshal json data", err)
	}
	return conf
}

func (c *Config) CurrentDiscussion() Item {
	return c.List[len(c.List)-1]
}

func (c *Config) PrintCurrentActiveQuestion() {
	fmt.Printf("%v\n", c.CurrentDiscussion().Url)
}

func (c *Config) PrintList() {
	for _, item := range c.List {
		fmt.Printf("%v\n", item.Url)
	}
}

func downloadConfigFile(file *os.File) {
	log.Println("downloading config file config.json")

	resp, err := http.Get(configUrl)
	if err != nil {
		log.Println(fmt.Errorf("%v", err))
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Println(fmt.Errorf("getting %s: %s", configUrl, resp.Status))
	}

	defer resp.Body.Close()

	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		log.Fatalf("cannot write to the file: %v", err)
	}
}
