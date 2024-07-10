package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	ApiKey      string `yaml:"api_key"`
	Query       string `yaml:"query"`
	Proxy       string `yaml:"proxy"`
	Concurrency int    `yaml:"concurrency"`
	PageSize    int    `yaml:"page_size"`
}

func LoadConfig() {
	file, err := os.Open("config.yml")
	if err != nil {
		log.Printf("打开配置文件错误: %v\n", err)
		return
	}
	defer file.Close()

	config := Config{}
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Printf("解析配置文件错误: %v\n", err)
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
