package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Toggl struct {
		Projects map[string]struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"projects"`
	} `json:"toggl"`
}

func LoadConfig() *Config {
	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	decoder.Decode(&config)
	return &config
}
