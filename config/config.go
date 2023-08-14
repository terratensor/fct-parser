package config

import (
	"encoding/json"
	"fmt"
	"io"
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

// configUrl deprecated url "https://raw.githubusercontent.com/audetv/fct-parser/main/config.json"
var (
	configUrl  = "https://svodd.ru/api/list"
	configFile = "./config.json"
)

func ReadConfig() Config {
	var body []byte
	_, err := os.Open(configFile)

	if err == nil {
		log.Printf("using config file %v", configFile)
		body, err = os.ReadFile(configFile)
		if err != nil {
			log.Fatal("Cannot read file", err)
		}
	} else {
		body = FetchingConfigFile()
	}

	var conf Config
	err = json.Unmarshal(body, &conf)
	conf.IsValidConfig()
	// // TODO возвращаем пустой конфиг при ошибке, надо подумать как сделать по-другому и сделать рефакторинг
	// if err != nil {
	// 	conf = Config{List: []Item{}}
	// 	// log.Fatal("неправильный формат config файла: ", err)
	// }
	return conf
}

func (c *Config) IsValidConfig() {
	if len(c.List) == 0 {
		log.Fatalf("%v", "неправильный формат конфиг файла, исправьте или удалите файл, для загрузки конфиг файла используйте опцию fct-parser -u")
	}
}

func (c *Config) PreviousDiscussion() Item {
	return c.List[len(c.List)-2]
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

func FetchingConfigFile() []byte {
	log.Printf("fetching config file %v", configUrl)

	respBody := getResponseBody(configUrl)
	defer respBody.Close()

	body, err := io.ReadAll(respBody)
	if err != nil {
		log.Fatalf("cannot read response: %v", err)
	}
	return body
}

func DownloadConfigFile(file *os.File) {
	log.Println("downloading config file config.json")

	respBody := getResponseBody(configUrl)
	defer respBody.Close()

	_, err := file.ReadFrom(respBody)
	if err != nil {
		log.Fatalf("cannot write to the file: %v", err)
	}
}

func getResponseBody(url string) io.ReadCloser {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(fmt.Errorf("%v", err))
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Println(fmt.Errorf("getting %s: %s", configUrl, resp.Status))
	}

	return resp.Body
}
