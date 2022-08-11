package sitewalker

import "net/url"

// 网站的页面信息
type Page struct {
	Title         string
	Desc          string
	Keywords      []string
	RawURL        string
	URL           *url.URL
	H1            string
	Pages         map[string]*Page // 这个页面链接到的内页页面
	ExternalLinks []Href           // 外链
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

// 网站的链接信息
type Href struct {
	Href string
	URL  *url.URL
	Text string
	// Title string
}
