package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	ApiKey      string `json:"api_key"`
	Query       string `json:"query"`
	Proxy       string `json:"proxy"`
	Concurrency int    `json:"concurrency"`
	PageSize    int    `json:"page_size"`
}

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		return
	}
	defer file.Close()

	config := Config{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return
	}

	if apiKey == "" {
		apiKey = config.ApiKey
	}
	if query == "" {
		query = config.Query
	}
	if proxy == "" {
		proxy = config.Proxy
	}
	if concurrency == 0 {
		concurrency = config.Concurrency
	}
	if pageSize == 0 {
		pageSize = config.PageSize
	}
}
