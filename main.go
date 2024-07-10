package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

// 请求参数
var (
	apiKey      string
	query       string
	proxy       string
	concurrency int
	pageSize    int
)

// 解析命令行方法
func init() {
	flag.StringVar(&apiKey, "k", "", "FOFA API Key")
	flag.StringVar(&query, "q", "", "搜索关键词")
	flag.StringVar(&proxy, "proxy", "", "代理地址")
	flag.IntVar(&concurrency, "concurrency", 5, "并发截图数量")
	flag.IntVar(&pageSize, "size", 100, "每页查询数量")
}

func main() {
	flag.Parse()

	if query == "" {
		flag.Usage()
		return
	}

	log.Println("命令解析完毕...")
	// 读取配置文件（如果存在），如果没指定则采用config文件中的配置
	LoadConfig()
	log.Println("配置参数：", apiKey, query, proxy, concurrency, pageSize)

	// 进行FOFA搜索
	results, err := searchFofa(apiKey, query)
	if err != nil {
		fmt.Println("FOFA搜索错误:", err)
		return
	}
	if len(results) == 0 {
		fmt.Println("没有找到结果")
		return
	}

	fmt.Println("正在进行截图操作步骤：", len(results))

	// 对网站进行截图并保存
	takeScreenshots(results)
}

// 网站截图保存
func takeScreenshots(results [][]string) {
	var wg sync.WaitGroup                   // WaitGroup 用于等待所有 goroutine 完成
	sem := make(chan struct{}, concurrency) // sem 是一个带有容量的 channel，用于控制同时运行的 goroutine 数量。

	for i, result := range results {
		wg.Add(1)
		sem <- struct{}{}

		go func(url string, index int) {
			defer wg.Done()
			defer func() { <-sem }()

			finalUrl, title, filePath, err := GetScreenShot(url, proxy, index)
			if err != nil {
				log.Printf("截图失败：%v\n", err)
				return
			}

			fmt.Printf("URL：%s, 状态码：%s, 网站标题：%s, 截图文件：%s\n", finalUrl, "200", title, filePath)
		}(result[0], i)
	}

	wg.Wait()
}
