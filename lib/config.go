package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Toggl struct {
		Projects map[string]string `json:"projects"`
	} `json:"toggl"`
}

func LoadConfig() *Config {
	// fmt.Println(os.Getwd())
	file, err := os.Open("../../../../config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	decoder.Decode(&config)
	return &config
}
