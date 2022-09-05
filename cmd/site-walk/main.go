package main

import (
	"fmt"
	"os"
	"time"

	sitewalker "github.com/kevin-zx/site-walker"
)

func main() {
	sw := sitewalker.NewSiteWalker(sitewalker.WithTimeout(10 * time.Second))

	webSite, err := sw.Walk("https://www.fxt.cn/", []string{
		"fxt.cn",
	})
	if err != nil {
		panic(err)
	}
	articleFile := "article.txt"
	f, err := os.Create(articleFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// fmt.Printf("%+v", webSite)
	for _, page := range webSite.Pages {
		fmt.Println(page.RawURL)
		if page.RawURL == "https://www.fxt.cn/article" {
			f.Write(page.Html)
			f.Sync()
		}
	}

}
