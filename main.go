package main

import (
	"flag"
	"fmt"
	"log"
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

}
