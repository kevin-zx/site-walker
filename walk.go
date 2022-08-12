package sitewalker

import "github.com/gocolly/colly"

type SiteWalker struct {
	cacheDir string
}

type SiteWalkerOption func(sw *SiteWalker)

// 缓存目录
func CacheDir(dir string) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.cacheDir = dir
	}
}

func NewSiteWalker(opts ...SiteWalkerOption) *SiteWalker {
	sw := &SiteWalker{}
	for _, opt := range opts {
		opt(sw)
	}
	return sw
}

func (sw *SiteWalker) Walk(url string) (*WebSite, error) {
	collector := colly.NewCollector(
		colly.CacheDir(sw.cacheDir),
	)
	collector.DetectCharset = true
	return nil, nil
}
