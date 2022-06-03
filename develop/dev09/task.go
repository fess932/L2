package main

import (
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"path"
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

func wget(url string, p string) {
	w, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	r, err := http.DefaultClient.Do(w)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	os.WriteFile(path.Join(p, "index.html"), body, 0644)

	c := colly.NewCollector()

	log.Prinlnt(c)

	//2 вида ссылок
	//абсолютные
	//https://google.com"
	//
	//относительные
	///js/jquery.js
}
