package config

import (
	"encoding/json"
	"fmt"
	"log"
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

func ReadConfig() Config {
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
