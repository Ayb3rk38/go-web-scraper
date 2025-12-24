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
		fmt.Println("Sending request...")

		if r.StatusCode == 200 {
			fmt.Println("Accessed the site!")
		} else {
			fmt.Println("Couldn't access the site.\n Error code:", r.StatusCode)
		}
		htmlContent = r.Body
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

	err = os.WriteFile("site.html", htmlContent, 0644)
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
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <target_url>")
		os.Exit(1)
	}

	targetUrl := os.Args[1]
	cfg := defaultBrowserConfig()
	ctx, cancel, err := NewChromeDPContext(cfg)
	if err != nil {
		log.Fatal("Browser not found: ", err)
	}
	defer cancel()

	if err := collector(targetUrl); err != nil {
		log.Fatal("An error occured while taking content from the site: ", err)
	}
	if err := captureScreenshot(ctx, targetUrl); err != nil {
		log.Fatal("An error occured while taking screenshot from the site: ", err)
	}

	fmt.Println("Scraping completed.")
}
