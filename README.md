# Simple Go Web Scraper

A  web scraper written in **Go**. This tool is designed to perform basic site analysis using **Colly** and **Chromedp**.

## Features

* **Full-Page Screenshot:** Captures the entire page as a PNG file.
* **Link Extraction:** Scrapes all URLs from the page, converts them to absolute paths, and saves them to a `.txt` file.
* **HTML Backup:** Downloads and saves the raw source code (`.html`).


## Built With

* [Go](https://go.dev/)
* [Colly](https://github.com/gocolly/colly) (for HTML parsing)
* [Chromedp](https://github.com/chromedp/chromedp) (for browser screenshots)

## How to Run

1.  Clone the repository:
    ```bash
    git clone [https://github.com/Ayb3rk38/go-web-scraper.git](https://github.com/Ayb3rk38/go-web-scraper.git)
    ```

2.  Install dependencies:
    ```bash
    go mod tidy
    ```

3.  Run the scraper:
    ```bash
    go run main.go
    ```

## Output

The tool generates the following files in the project directory:
* `screenshot.png`: The screenshot of the site.
* `links.txt`: A list of all extracted URLs.
* `cekilen.html`: The source code of the target page.

---
*Created for educational purposes.*
