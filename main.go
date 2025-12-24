package main

import (
	"fmt"
	"log"
	"os"
)

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
