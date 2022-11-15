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
	Id   int    `json:"id"`
	Num  string `json:"num"`
	Date string `json:"date"`
	Url  string `json:"url"`
}

var configUrl = "https://raw.githubusercontent.com/audetv/fct-parser/main/config.json"

func ReadConfig() Config {
	file, err := os.OpenFile("./config.json", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err == nil {
		DownloadConfigFile(file)
	}

	fileBytes, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Cannot read file", err)
	}

	var conf Config
	err = json.Unmarshal(fileBytes, &conf)
	// TODO возвращаем пустой конфиг при ошибке, надо подумать как сделать по-другому и сделать рефакторинг
	if err != nil {
		conf = Config{List: []Item{}}
		// log.Fatal("неправильный формат config файла: ", err)
	}
	return conf
}

func (c *Config) IsValidConfig() {
	if len(c.List) == 0 {
		log.Fatalf("%v", "неправильный формат конфиг файла, для загрузки конфиг файла используйте опцию fct-parcer -u")
	}
}

func (c *Config) CurrentDiscussion() Item {
	return c.List[len(c.List)-1]
}

func (c *Config) PrintCurrentDiscussion() {
	c.IsValidConfig()
	fmt.Printf("%v\n", c.CurrentDiscussion().Url)
}

func (c *Config) PrintList() {
	c.IsValidConfig()
	for _, item := range c.List {
		fmt.Printf("%v\n", item.Url)
	}
}

func DownloadConfigFile(file *os.File) {
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
