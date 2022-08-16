package sitewalker

import (
	"time"

	"github.com/gocolly/colly"
)

type SiteWalker struct {
	collector *colly.Collector

	// limit
	parallelism int
	randomDelay time.Duration
}

type SiteWalkerOption func(sw *SiteWalker)

func NewSiteWalker(opts ...SiteWalkerOption) *SiteWalker {
	sw := &SiteWalker{}
	sw.init()
	for _, opt := range opts {
		opt(sw)
	}
	return sw
}

// simulate baidu search engine crawler
const (
	defaultPCUserAgent     = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
	defaultMobileUserAgent = "Mozilla/5.0 (Linux;u;Android 4.2.2;zh-cn;) AppleWebKit/534.46 (KHTML,like Gecko) Version/5.1 Mobile Safari/10600.6.3 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
)

func (sw *SiteWalker) init() {
	sw.collector.UserAgent = defaultPCUserAgent
	sw.parallelism = 1
	// default MaxDepth is 1000,
	// incase crawling is too heavy,
	sw.collector.MaxDepth = 1000
}

func (sw *SiteWalker) Walk(url string) (*WebSite, error) {
	collector := colly.NewCollector()
	collector.DetectCharset = true
	collector.Async = true

	return nil, nil
}

// 缓存目录
func CacheDir(dir string) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.collector.CacheDir = dir
	}
}

// 并发数
func Parallelism(n int) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.parallelism = n
	}
}

// UserAgent
// this UserAgent will cover DeviceType UserAgent
func UserAgent(ua string) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.collector.UserAgent = ua
	}
}

// device type will be used to decide UserAgent
func WithDeviceType(isMobile bool) SiteWalkerOption {
	return func(sw *SiteWalker) {
		if isMobile {
			sw.collector.UserAgent = defaultMobileUserAgent
		} else {
			sw.collector.UserAgent = defaultPCUserAgent
		}
	}
}
