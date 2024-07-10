package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"os"
	"time"
)

//func sanitizeFilename(url string) string {
//	return strings.ReplaceAll(url, "https://", "")
//}

func GetScreenShot(url string, proxy string, index int) (finalUrl string, title string, filePath string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := append(chromedp.DefaultExecAllocatorOptions[:], //默认配置必须有
		chromedp.ProxyServer(proxy),     //代理服务器
		chromedp.WindowSize(1920, 1080), // 截图大小
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var buf []byte
	// 获取截图并获取将跳转网站的url和标题
	if err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.FullScreenshot(&buf, 100),
			chromedp.Title(&title),
			chromedp.Evaluate("window.location.href", &finalUrl),
		},
	); err != nil {
		return "", "", "", err
	}

	// 确保result目录存在
	if _, err := os.Stat("result"); os.IsNotExist(err) {
		os.Mkdir("result", 0755)
	}

	// 保存到本地
	filePath = fmt.Sprintf("result/%d.png", index)
	if err := os.WriteFile(filePath, buf, 0o644); err != nil {
		return "", "", "", err
	}

	return finalUrl, title, filePath, nil
}
