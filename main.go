package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

func collector(web string) error {
	var htmlContent []byte
	var links string

	file, err := os.Create("links.txt")

	if err != nil {
		return err
	}
	defer file.Close()

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("İstek Gönderiliyor...")

		if r.StatusCode == 200 {
			fmt.Println("Siteye Erişildi!")
		} else {
			fmt.Println("Siteye Erişilemedi,Hata Kodu:", r.StatusCode)
		}
		htmlContent = (r.Body)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		links = e.Attr("href")
		links = e.Request.AbsoluteURL(links)
		_, err := file.WriteString(links + "\n")
		if err != nil {
			return
		}
	})

	c.Visit(web)

	err = os.WriteFile("site.html", []byte(htmlContent), 0644)
	if err != nil {
		return err
	}

	return nil

}

func captureScreenshot(ctx context.Context, targetUrl string) error {
	var screenShotBuffer []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(targetUrl),
		chromedp.Sleep(5*time.Second),
		chromedp.FullScreenshot(&screenShotBuffer, 100),
	)

	if err != nil {
		return err
	}

	err = os.WriteFile("screenshot.png", screenShotBuffer, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanim: go run main.go <hedef_url>")
		os.Exit(1)
	}
	targetUrl := os.Args[1]
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	if err := collector(targetUrl); err != nil {
		log.Fatal("Siteden Bilgi Çekerken Hata Meydana Geldi!", err)
	}
	if err := captureScreenshot(ctx, targetUrl); err != nil {
		log.Fatal("Ekran Görüntüsü Alinirken Hata Meydana Geldi!", err)
	}
	fmt.Println("Scraper Başari ile Çalişti")
}


