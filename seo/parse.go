package seo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SEOText struct {
	Title       string
	Description string
	Keywords    []string
	H1          string
}

func ExtractSEOTextInfo(html *goquery.Selection) SEOText {
	title := html.Find("title").Text()
	description := html.Find("meta[name=description]").AttrOr("content", "")
	keywordsText := html.Find("meta[name=keywords]").AttrOr("content", "")
	h1 := html.Find("h1").Text()
	keywords := splitKeywordsStrText(keywordsText)

	return SEOText{
		Title:       title,
		Description: description,
		Keywords:    keywords,
		H1:          h1,
	}
}

// 中文互联网的 keywords 是需要一些特殊处理
var keywordsSplitText = []string{",", "，", "、", "_", ";", "；"}

func splitKeywordsStrText(keywords string) []string {
	var result []string
	for _, text := range keywordsSplitText {
		result = append(result, strings.Split(keywords, text)...)
	}
	return result
}
