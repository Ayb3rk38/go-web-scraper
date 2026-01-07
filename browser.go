package main

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/chromedp/chromedp"
)

type BrowserConfig struct {
	EnvKeys []string
	AllowAlternatives bool
	Headless bool
}

func defaultBrowserConfig() BrowserConfig {
	return BrowserConfig{
		EnvKeys:           []string{"CHROME_PATH", "CHROME_BIN", "BROWSER_PATH"},
		AllowAlternatives: true,
		Headless:          true,
	}
}

func fileExists(p string) bool {
	if p == "" {
		return false
	}
	st, err := os.Stat(p)
	return err == nil && !st.IsDir()
}

func FindBrowserExec(cfg BrowserConfig) (string, error) {
	for _, k := range cfg.EnvKeys {
		if p := os.Getenv(k); fileExists(p) {
			return p, nil
		}
	}

	names := []string{
		"google-chrome", "google-chrome-stable",
		"chromium", "chromium-browser",
		"chrome", "chrome.exe",
	}
	if cfg.AllowAlternatives {
		names = append(names,
			"msedge", "msedge.exe",
			"microsoft-edge", "microsoft-edge-stable",
			"brave", "brave.exe", "brave-browser",
		)
	}

	for _, name := range names {
		if p, err := exec.LookPath(name); err == nil && fileExists(p) {
			return p, nil
		}
	}

	switch runtime.GOOS {
	case "windows":
		candidates := []string{
			filepath.Join(os.Getenv("ProgramFiles"), "Google", "Chrome", "Application", "chrome.exe"),
			filepath.Join(os.Getenv("ProgramFiles(x86)"), "Google", "Chrome", "Application", "chrome.exe"),
			filepath.Join(os.Getenv("LocalAppData"), "Google", "Chrome", "Application", "chrome.exe"),
		}
		if cfg.AllowAlternatives {
			candidates = append(candidates,
				filepath.Join(os.Getenv("ProgramFiles(x86)"), "Microsoft", "Edge", "Application", "msedge.exe"),
				filepath.Join(os.Getenv("ProgramFiles"), "Microsoft", "Edge", "Application", "msedge.exe"),
			)
		}
		for _, p := range candidates {
			if fileExists(p) {
				return p, nil
			}
		}

	case "darwin":
		candidates := []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
		}
		if cfg.AllowAlternatives {
			candidates = append(candidates,
				"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
				"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
			)
		}
		for _, p := range candidates {
			if fileExists(p) {
				return p, nil
			}
		}

	default:
		candidates := []string{
			"/usr/bin/google-chrome",
			"/usr/bin/google-chrome-stable",
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
		}
		if cfg.AllowAlternatives {
			candidates = append(candidates,
				"/usr/bin/microsoft-edge",
				"/usr/bin/microsoft-edge-stable",
			)
		}
		for _, p := range candidates {
			if fileExists(p) {
				return p, nil
			}
		}
	}

	return "", errors.New("no Chromium-based browser found. Install Chrome/Chromium (or set CHROME_PATH to the browser executable)")
}

func NewChromeDPContext(cfg BrowserConfig) (context.Context, context.CancelFunc, error) {
	execPath, err := FindBrowserExec(cfg)
	if err != nil {
		return nil, nil, err
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(execPath),
		chromedp.Flag("headless", cfg.Headless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	taskCtx, taskCancel := chromedp.NewContext(allocCtx)

	cancel := func() {
		taskCancel()
		allocCancel()
	}
	return taskCtx, cancel, nil
}
