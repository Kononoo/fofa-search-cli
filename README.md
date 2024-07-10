# fofa-search-cli


## 项目简介
FOFA-命令行工具：给定一个搜索语法，从fofa搜索结果，并对结果中的网站调用浏览器截图， 输出网站列表和截图结果列表。
1. 支持命令行参数输入搜索关键词。
2. 获取 FOFA 搜索结果，包括网站 URL、状态码、网站标题。
3. 对网站进行截图并保存截图文件。 
4. 输出网站信息及截图结果。

## 额外功能项
1. 支持配置代理转发
2. 支持配置并发截图数量
3. 支持配置从fofa查询的结果数量
4. 支持从配置文件读取配置

## 使用说明
```
.\fofa-search-cli.ext -k FOFA_API_KEY -q "title=百度"

.\fofa-search-cli.exe -q "domain=baidu.com" 
```
### 参数配置：
- k：FOFA API Key。
- q：搜索关键词。
- proxy：代理地址。
- concurrency：并发截图数量。
- size：每页查询数量。
### 配置文件
可以通过 config.yml 文件配置参数，示例：
```
api_key: "c9b5570171fccebc5c24ae16b32f9bb4"    # FOFA-API
query: "title=example"                         # 可以在这配置查询语句
proxy: "http://127.0.0.1:7890"                 # 配置代理转发
concurrency: 5                                 # 配置并发截图数量
page_size: 100                                 # 配置从FOFA查询结果数量
```

## 项目构建
```
go get
go build
```


## 学习任务
- golang基础学习，如：基础语法，编译，调试等
- 熟悉 Fofa，理解 Fofa 概念和作用，了解如何编程调用fofa进行搜索

## 参考资料
- Fofa 官方文档：可以访问 Fofa 官方网站，查找官方文档和教程，了解 Fofa 的功能和使用方法。
- 规则手册：   https://ones.baimaohui.net:24688/wiki/#/team/2hTgeDe2/space/Gt4DCMwE/page/NQtUkKPu
- 网络教程： https://nosec.org/home/detail/5060.html
https://tour.go-zh.org/welcome/1