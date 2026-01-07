# Simple Go Web Scraper

Simple CLI web scraper written in **Go**. It:

- downloads and stores the page HTML
- extracts links and writes them to a text file
- takes a full-page screenshot using a real browser (Chrome/Chromium/Edge/Brave via `chromedp`)

This project is intended for learning and local testing. Always respect a website's Terms of Service and `robots.txt`.

## Features

- **Full-page screenshot** (`screenshot.png`)
- **Link extraction** (absolute URLs, written to `links.txt`)
- **HTML backup** (`site.html`)
- **Cross-platform browser discovery**
  - tries environment overrides first
  - then looks in `PATH`
  - then checks common install locations per OS

## Requirements

- Go (modules enabled; modern Go versions recommended)
- A Chromium-based browser installed:
  - Google Chrome
  - Chromium
  - Microsoft Edge
  - Brave

If no browser is detected automatically, you can provide an explicit path via `CHROME_PATH` env variable.

## Install

Clone:

```bash
git clone https://github.com/Ayb3rk38/go-web-scraper.git
cd go-web-scraper
```

Dependencies are tracked in `go.mod`/`go.sum`. You can download them explicitly:

```bash
go mod tidy
```

## Run

Run from the repository root:

```bash
go run . <target_url>
```

Example:

```bash
go run . https://google.com
```

## Browser selection (cross-platform)

The screenshot step uses `chromedp`, which launches a real browser. The project tries to find a browser automatically.

If detection fails, set `CHROME_PATH` to the browser executable.

### Windows (PowerShell)

```powershell
$env:CHROME_PATH = "C:\Program Files\Google\Chrome\Application\chrome.exe"
go run . https://google.com
```

### Linux/macOS (bash/zsh)

```bash
export CHROME_PATH="/usr/bin/google-chrome"
go run . https://google.com
```

Notes:

- If Chrome is not installed but Edge is, the tool may use Edge (config allows alternatives).
- The tool runs the browser in headless mode.

## Output

The tool generates the following files in the project directory:

- `screenshot.png` (full page)
- `links.txt` (one URL per line)
- `site.html` (raw HTML of the visited page)

These files are runtime artifacts and are typically ignored by Git.
