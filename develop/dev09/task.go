package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"net/url"

	"github.com/gocolly/colly/v2"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	wget("https://golang.org/", "./develop/dev09/tmp/")
}

// flag -r - рекурсивно скачивать все ссылки на сайте с тем же доменом

func wget(turl string, p string) {
	p, err := filepath.Abs(p)
	if err != nil {
		log.Fatal(err)
	}
	// w, err := http.NewRequest("GET", urgol, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// r, err := http.DefaultClient.Do(w)
	// if err != nil {
	// 	panic(err)
	// }

	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// os.WriteFile(path.Join(p, "index.html"), body, 0644)

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

	c.OnResponse(func(r *colly.Response) {
		u := r.Request.URL
		log.Println("Visiting", u, u.Path)

		dirPath := path.Join(p, u.Path)
		os.MkdirAll(dirPath, os.ModePerm)

		// TODO:
		// check stylesheet (end on .css)
		// js (end on .js)
		// check any with .* on end
		// default index.html

		log.Println("FILENAME", filepath.Ext(u.String()))

		if filepath.Ext(u.String()) == "" {

		}

		if err = r.Save(dirPath + "/index.html"); err != nil {
			log.Fatal(err)
		}

		// if err := os.WriteFile(dirPath+"/index.html", r.Body, os.ModePerm); err != nil {
		// log.Println(err)
		// }
	})

	c.Visit(turl)

	//2 вида ссылок
	//абсолютные
	//https://google.com"
	//
	//относительные
	///js/jquery.js
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
