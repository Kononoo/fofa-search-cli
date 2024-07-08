# fofa-search-cli


## 项目简介
FOFA-命令行工具：给定一个搜索语法，从fofa搜索结果，并对结果中的网站调用浏览器截图， 输出网站列表和截图结果列表。

## 额外功能项
1. 支持配置代理转发
2. 支持配置并发截图数量
3. 支持配置从fofa查询的结果数量
4. 支持从配置文件读取配置

## 使用说明
```
./fofa-search-cli -k FOFA_API_KEY -q "title=\"百度\""
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