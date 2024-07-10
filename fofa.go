package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FofaResponse struct {
	Error   bool       `json:"error"`
	Results [][]string `json:"results"`
}

func searchFofa(apiKey, query string) ([][]string, error) {
	encodedQuery := base64.StdEncoding.EncodeToString([]byte(query))
	log.Println("请求的参数：", apiKey, encodedQuery, pageSize)
	url := fmt.Sprintf("https://fofa.info/api/v1/search/all?key=%s&qbase64=%s&size=%d", apiKey, encodedQuery, pageSize)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fofaResponse FofaResponse
	err = json.Unmarshal(body, &fofaResponse)
	if err != nil {
		return nil, err
	}

	if fofaResponse.Error {
		return nil, fmt.Errorf("FOFA API 返回错误")
	}

	return fofaResponse.Results, nil
}
