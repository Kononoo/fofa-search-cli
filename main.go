package main

import (
	"flag"
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

}
