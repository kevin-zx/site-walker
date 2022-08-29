package sitewalker

import (
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const fileReg = ".+?(\\.jpg|\\.png|\\.gif|\\.GIF|\\.PNG|\\.JPG|\\.pdf|\\.PDF|\\.doc|\\.DOC|\\.csv|\\.CSV|\\.xls|\\.XLS|\\.xlsx|\\.XLSX|\\.mp40|\\.lfu|\\.DNG|\\.ZIP|\\.zip)(\\W+?\\w|$)"

type SiteWalker struct {
	collector *colly.Collector
	// colly limit
	limitRule *colly.LimitRule
	uaCustom  bool
}

type SiteWalkerOption func(sw *SiteWalker)

func NewSiteWalker(startUrl string, AllowedDomains []string, opts ...SiteWalkerOption) *SiteWalker {
	sw := &SiteWalker{}
	sw.init()
	for _, opt := range opts {
		opt(sw)
	}
	sw.collector.Limit(sw.limitRule)
	return sw
}

// simulate baidu search engine crawler
const (
	defaultPCUserAgent     = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
	defaultMobileUserAgent = "Mozilla/5.0 (Linux;u;Android 4.2.2;zh-cn;) AppleWebKit/534.46 (KHTML,like Gecko) Version/5.1 Mobile Safari/10600.6.3 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
)

func (sw *SiteWalker) init() {
	collector := colly.NewCollector()
	collector.DetectCharset = true
	collector.Async = true
	collector.UserAgent = defaultPCUserAgent

	// default MaxDepth is 1000,
	// incase crawling is too heavy,
	collector.MaxDepth = 1000

	collector.Async = true
	collector.SetRequestTimeout(time.Second * 20)
	collector.SetRequestTimeout(10 * time.Second)
	collector.DisallowedURLFilters = append(collector.DisallowedURLFilters, regexp.MustCompile(fileReg))

	sw.collector = collector
	sw.limitRule.DomainGlob = "*"
	sw.limitRule.Parallelism = 1
	sw.limitRule.RandomDelay = time.Second
	sw.limitRule.Delay = time.Second

}

func (sw *SiteWalker) Walk(url string) (*WebSite, error) {
	webSite := &WebSite{}
	sw.collector.OnHTML("html", func(e *colly.HTMLElement) {
		page := &Page{}

		// TDK
		page.Title = e.ChildText("title")
		page.Desciption = e.ChildText("meta[name=description]")
		page.Keywords = e.ChildText("meta[name=keywords]")

		// cuurentUrl := e.Request.URL.String()
		e.DOM.Find("a[href]").EachWithBreak(func(i int, s *goquery.Selection) bool {
			link := ParseATag2Link(s, e.Request.URL)
			if link == nil {
				return true
			}
			return true
		})

	})

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
		sw.limitRule.Parallelism = n
	}
}

// UserAgent
// this UserAgent will cover DeviceType UserAgent
func UserAgent(ua string) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.collector.UserAgent = ua
		sw.uaCustom = true
	}
}

// device type will be used to decide UserAgent
func WithDeviceType(isMobile bool) SiteWalkerOption {
	return func(sw *SiteWalker) {
		if sw.uaCustom {
			return
		}

		if isMobile {
			sw.collector.UserAgent = defaultMobileUserAgent
		} else {
			sw.collector.UserAgent = defaultPCUserAgent
		}
	}
}

// withDelay
func WithDelay(randomDelay time.Duration, delay time.Duration) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.limitRule.RandomDelay = randomDelay
		sw.limitRule.Delay = delay
	}
}

// with timeout
func WithTimeout(timeout time.Duration) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.collector.SetRequestTimeout(timeout)
	}
}
