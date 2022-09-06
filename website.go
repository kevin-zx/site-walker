package sitewalker

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/kevin-zx/site-walker/util/urltool"
)

// 网站的页面信息
type Page struct {
	// seo text 信息
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	// h1标签的内容
	H1 string `json:"h1"`

	// 页面的原始url
	RawURL string   `json:"raw_url"`
	URL    *url.URL `json:"url"`

	// 页面中的链接
	Links []*Link `json:"links"`
	// 页面中的外部链接
	ExternalLinks []*Link `json:"external_links"`
	// 网站网页数据
	Html []byte `json:"html"`
	deep int    // 爬取深度
}

// 网站的基本信息
type WebSite struct {
	Protocol string  `json:"protocol"`
	Domain   string  `json:"domain"`
	HomePage *Page   `json:"home_page"`
	Pages    []*Page `json:"pages"`
}

type LinkType int

const (
	LinkTypeText LinkType = iota
	LinkTypeImg
)

// 网站的链接信息
type Link struct {
	Href     string   `json:"href"`
	URL      *url.URL `json:"-"`
	Text     string   `json:"text"`
	LinkType LinkType `json:"link_type"`
}

func ParseATag2Link(a *goquery.Selection, pageURL *url.URL) *Link {
	href, ok := a.Attr("href")
	if !ok {
		return nil
	}
	href = urltool.ClearHref(href)
	if href == "" {
		return nil
	}
	if !urltool.IsValidHref(href) {
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
