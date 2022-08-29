package sitewalker

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/kevin-zx/site-walker/util/urltool"
)

// 网站的页面信息
type Page struct {
	// seo text 信息
	Title      string
	Desciption string
	Keywords   []string
	H1         string // h1标签的内容

	RawURL string
	URL    *url.URL

	Pages         map[string]*Page // 这个页面链接到的内页页面
	ExternalLinks []Link           // 外链
	Html          []byte           // 网站网页数据

	deep int // 爬取深度
	// SameDomainExternalLinks []Href // 相同域名的外链 比如 image.baidu.com 页面中 含有 zhidao.baidu.com
}

// 网站的基本信息
type WebSite struct {
	Protocol string
	Domain   string
	HomePage *Page
	Pages    []*Page
}

type LinkType int

const (
	LinkTypeText LinkType = iota
	LinkTypeImg
)

// 网站的链接信息
type Link struct {
	Href     string
	URL      *url.URL
	Text     string
	LinkType LinkType
}

func ParseATag2Link(a *goquery.Selection, pageURL *url.URL) *Link {
	href, ok := a.Attr("href")
	if !ok {
		return nil
	}
	if !urltool.IsValidHref(href) {
		return nil
	}
	href = urltool.ClearHref(href)
	if href == "" {
		return nil
	}
	txt := GetATagAnchor(a)
	LinkType := LinkTypeText
	if a.Find("img").Size() > 0 {
		LinkType = LinkTypeImg
	}
	url, _ := urltool.ConvertHref2URL(href, pageURL)
	if url == nil {
		return nil
	}

	return &Link{
		Href:     href,
		URL:      url,
		Text:     txt,
		LinkType: LinkType,
	}
}

func GetATagAnchor(a *goquery.Selection) string {
	txt := ""
	txt = a.Text()
	if txt == "" {
		txt, _ = a.Attr("title")
	}
	if txt == "" {
		txt, _ = a.Attr("alt")
	}
	return txt
}
