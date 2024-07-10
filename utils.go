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

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer(proxy),
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var buf []byte
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

	filePath = fmt.Sprintf("screenshot-%d.png", index)
	if err := os.WriteFile(filePath, buf, 0644); err != nil {
		return "", "", "", err
	}
	return finalUrl, title, filePath, nil
}
