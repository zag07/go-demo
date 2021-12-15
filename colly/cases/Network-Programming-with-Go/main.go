package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

var (
	destDomain  = "tumregels.github.io"
	destUrl     = "https://tumregels.github.io/Network-Programming-with-Go/"
	mainDir     = "../../storage/Network-Programming-with-Go/"
	assetsDir   = "../../storage/Network-Programming-with-Go/assets/"
	gitbookDir   = "../../storage/Network-Programming-with-Go/gitbook/"
	savedPath   = ""
	isFirstTime = true
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(destDomain),
		colly.MaxDepth(1),
	)

	err := os.MkdirAll(assetsDir, 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %s", err)
	}
	_ = os.MkdirAll(gitbookDir, 0755)

	c.OnResponse(func(r *colly.Response) {
		fullURL := r.Request.URL.String()
		fmt.Println("URL:", fullURL)

		u, err := url.Parse(fullURL)
		if err != nil {
			log.Fatalf("URL parse error: %v", err)
		}

		paths := strings.Split(u.Path, "/")
		if paths[len(paths)-2] == "Network-Programming-with-Go" {
			// e.g. https://tumregels.github.io/Network-Programming-with-Go/

			savedPath = mainDir + "index.html"
		} else if paths[len(paths)-1] == "" {
			// e.g. https://tumregels.github.io/Network-Programming-with-Go/architecture/

			dirPath := mainDir + paths[len(paths)-2]
			os.Mkdir(dirPath, 0755)
			savedPath = dirPath + "/index.html"
		} else {
			// e.g. https://tumregels.github.io/Network-Programming-with-Go/architecture/protocol_layers.html

			savedPath = mainDir + paths[len(paths)-2] + "/" + paths[len(paths)-1]
		}
		f, err := os.Create(savedPath)
		if err != nil {
			log.Fatalf("Error creating file: %s", err)
		}
		io.WriteString(f, string(r.Body))
	})

	c.OnHTML("link[rel='stylesheet']", func(e *colly.HTMLElement) {
		if !isFirstTime {
			return
		}

		fullUrl := e.Request.AbsoluteURL(e.Attr("href"))
		u, err := url.Parse(fullUrl)
		if err != nil {
			log.Fatalf("URL parse error: %v", err)
		}

		paths := strings.Split(u.Path, "/")
		if paths[len(paths)-2] != "gitbook" {
			// e.g. https://tumregels.github.io/Network-Programming-with-Go/gitbook/gitbook-plugin-search/search.css

			dirPath := mainDir + "gitbook/" + paths[len(paths)-2]
			os.Mkdir(dirPath, 0755)
			savedPath = dirPath + "/" + paths[len(paths)-1]
		} else {
			// e.g. https://tumregels.github.io/Network-Programming-with-Go/gitbook/style.css

			savedPath = mainDir + "gitbook/" + paths[len(paths)-1]
		}

		f, err := os.Create(savedPath)
		if err != nil {
			log.Fatalf("Error creating file: %s", err)
		}

		res, _ := http.Get(fullUrl)
		io.Copy(f, res.Body)
	})

	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		fullUrl := e.Request.AbsoluteURL(e.Attr("src"))

		u, err := url.Parse(fullUrl)
		if err != nil {
			log.Fatalf("URL parse error: %v", err)
		}

		paths := strings.Split(u.Path, "/")
		if paths[len(paths)-2] != "gitbook" {
			// e.g. https://tumregels.github.io/Network-Programming-with-Go/gitbook/gitbook-plugin-search/search.js

			dirPath := mainDir + "gitbook/" + paths[len(paths)-2]
			os.Mkdir(dirPath, 0755)
			savedPath = dirPath + "/" + paths[len(paths)-1]
		} else {
			// e.g. https://tumregels.github.io/Network-Programming-with-Go/gitbook/gitbook.js

			savedPath = mainDir + "gitbook/" + paths[len(paths)-1]
		}

		f, err := os.Create(savedPath)
		if err != nil {
			log.Fatalf("Error creating file: %s", err)
		}

		res, _ := http.Get(fullUrl)
		io.Copy(f, res.Body)
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		fullUrl := e.Request.AbsoluteURL(e.Attr("src"))

		u, err := url.Parse(fullUrl)
		if err != nil {
			log.Fatalf("URL parse error: %v", err)
		}

		paths := strings.Split(u.Path, "/")
		if paths[1] != "Network-Programming-with-Go" {
			return
		}
		// e.g. https://tumregels.github.io/Network-Programming-with-Go/assets/iso.gif
		savedPath = assetsDir + paths[len(paths)-1]

		f, err := os.Create(savedPath)
		if err != nil {
			log.Fatalf("Error creating file: %s", err)
		}

		res, _ := http.Get(fullUrl)
		io.Copy(f, res.Body)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// After the first web is visited and the CSS and JS files are saved we make
		// this bool false, so we will not download those files multiple times
		isFirstTime = false
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.Visit(destUrl)
}
