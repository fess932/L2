package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Недостаточно аргументов")
	}

	if os.Args[1] == "-r" {
		if len(os.Args) < 4 {
			log.Fatal("Недостаточно аргументов")
		}

		if err := gwget(os.Args[2], os.Args[3]); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := download(os.Args[1], os.Args[2]); err != nil {
		log.Fatal("ошибка", err)
	}
}

// gwget - программа для скачивания сайтов
// expample:
// gwget [-r] url dir
// gwget [-r] https://go.dev/ ./develop/dev09/tmp/go

func download(targetUrl string, target string) error {
	parsedTargetUrl, err := url.ParseRequestURI(targetUrl)
	if err != nil {
		return fmt.Errorf("не удалось получить url путь %s: %w", targetUrl, err)
	}

	target, err = filepath.Abs(target)
	if err != nil {
		return fmt.Errorf("не удалось получить абсолютный путь для %s: %w", target, err)
	}

	f, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %w", err)
	}
	defer f.Close()

	log.Println(parsedTargetUrl, target)

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return fmt.Errorf("не удалось создать запрос %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос %w", err)
	}

	io.Copy(f, resp.Body)

	return nil
}

func gwget(targetUrl string, targetDir string) error {
	targetDir, err := filepath.Abs(targetDir)
	if err != nil {
		return fmt.Errorf("не удалось получить абсолютный путь для директории %w", err)
	}

	parsedTargetUrl, err := url.ParseRequestURI(targetUrl)
	if err != nil {
		return fmt.Errorf("не удалось получить url путь %s: %w", targetUrl, err)
	}

	log.Println(parsedTargetUrl, targetDir)

	c := colly.NewCollector()
	c.OnHTML("link[href]", func(e *colly.HTMLElement) {
		next := e.Attr("href")

		if !toVisit(next) {
			return
		}

		e.Request.Visit(next)
	})
	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		next := e.Attr("src")

		if !toVisit(next) {
			return
		}

		e.Request.Visit(next)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		next := e.Attr("href")

		if !toVisit(next) {
			return
		}

		e.Request.Visit(next)
	})
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		next := e.Attr("src")

		if !toVisit(next) {
			return
		}

		e.Request.Visit(next)
	})
	c.OnResponse(func(r *colly.Response) {
		reqUri := r.Request.URL.RequestURI()

		var absFileName string

		ext := filepath.Ext(reqUri)

		if ext == "" {
			dirpath := filepath.Join(targetDir, reqUri)
			log.Println("DIRPATH without ext", dirpath)

			os.MkdirAll(filepath.Join(dirpath, reqUri), os.ModePerm)
			absFileName = filepath.Join(dirpath, "index.html")
		} else {
			log.Println("EXT", ext)

			dirpath := filepath.Join(targetDir, filepath.Dir(reqUri))
			log.Println("DIRPATH with ext", dirpath)

			os.MkdirAll(dirpath, os.ModePerm)

			absFileName = filepath.Join(dirpath, filepath.Base(reqUri))
		}

		log.Println("ABSFILENAME", absFileName)
		log.Println()

		if err := os.WriteFile(absFileName, r.Body, 0644); err != nil {
			log.Fatal(err)
		}
	})

	c.AllowedDomains = append(c.AllowedDomains, parsedTargetUrl.Host)

	c.Async = true
	if err := c.Visit(parsedTargetUrl.String()); err != nil {
		return fmt.Errorf("не удалось посетить сайт %w", err)
	}
	c.Wait()

	return nil
}

func toVisit(next string) bool {
	if next == "" {
		return false
	}

	u, err := url.Parse(next)
	if err != nil {
		log.Println(err)

		return false
	}

	if u.IsAbs() {
		return false
	}

	return true
}
