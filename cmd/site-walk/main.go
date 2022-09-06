package main

import (
	"encoding/json"
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
	data, err := json.MarshalIndent(webSite, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(data))
}
