package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Toggl struct {
		Projects map[string]string `json:"projects"`
	} `json:"toggl"`
}

func LoadConfig() *Config {
	fmt.Println(os.Getwd())
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	file, err := os.Open("../config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	decoder.Decode(&config)
	return &config
}
