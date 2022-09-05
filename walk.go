package sitewalker

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/kevin-zx/site-walker/seo"
)

const fileReg = ".+?(\\.jpg|\\.png|\\.gif|\\.GIF|\\.PNG|\\.JPG|\\.pdf|\\.PDF|\\.doc|\\.DOC|\\.csv|\\.CSV|\\.xls|\\.XLS|\\.xlsx|\\.XLSX|\\.mp40|\\.lfu|\\.DNG|\\.ZIP|\\.zip)(\\W+?\\w|$)"

type SiteWalker struct {
	collector *colly.Collector
	// colly limit
	limitRule *colly.LimitRule
	uaCustom  bool

	maxPages int

	// loack
	lock sync.Mutex

	linkMap map[string]bool
}

type SiteWalkerOption func(sw *SiteWalker)

func NewSiteWalker(opts ...SiteWalkerOption) *SiteWalker {
	sw := &SiteWalker{}
	sw.init()
	for _, opt := range opts {
		opt(sw)
	}
	sw.collector.Limit(sw.limitRule)
	sw.lock = sync.Mutex{}
	sw.linkMap = make(map[string]bool)
	return sw
}

// simulate baidu search engine crawler
const (
	defaultPCUserAgent     = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
	defaultMobileUserAgent = "Mozilla/5.0 (Linux;u;Android 4.2.2;zh-cn;) AppleWebKit/534.46 (KHTML,like Gecko) Version/5.1 Mobile Safari/10600.6.3 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
)

func (sw *SiteWalker) init() {

	sw.maxPages = 100
	collector := colly.NewCollector()
	collector.DetectCharset = true
	collector.Async = true
	collector.UserAgent = defaultPCUserAgent

	// default MaxDepth is 100,
	// incase crawling is too heavy,
	collector.MaxDepth = 100

	collector.Async = true
	collector.SetRequestTimeout(time.Second * 20)
	collector.SetRequestTimeout(10 * time.Second)
	collector.DisallowedURLFilters = append(collector.DisallowedURLFilters, regexp.MustCompile(fileReg))

	sw.collector = collector

	sw.limitRule = &colly.LimitRule{}

	sw.limitRule.DomainGlob = "*"
	sw.limitRule.Parallelism = 1
	sw.limitRule.RandomDelay = time.Second
	sw.limitRule.Delay = time.Second

	sw.collector.Limit(sw.limitRule)
}

func (sw *SiteWalker) Walk(homeUrl string, allowedDomains []string) (*WebSite, error) {

	homeURL, err := url.Parse(homeUrl)
	if err != nil {
		return nil, err
	}
	sameHostFilter := NewDomainFilter()
	sameHostFilter.Add(homeURL.Host)
	for _, domain := range allowedDomains {
		sameHostFilter.Add(domain)
	}
	webSite := &WebSite{}
	sw.collector.OnHTML("html", func(e *colly.HTMLElement) {
		page := &Page{}
		// TDK
		seoText := seo.ExtractSEOTextInfo(e.DOM)
		page.Title = seoText.Title
		page.Description = seoText.Description
		page.Keywords = seoText.Keywords
		page.H1 = seoText.H1
		page.deep = e.Request.Depth

		// url
		currentUrl := e.Request.URL
		page.URL = currentUrl
		page.RawURL = currentUrl.String()

		// html raw data
		page.Html = e.Response.Body
		page.deep = e.Request.Depth

		e.DOM.Find("a[href]").EachWithBreak(func(i int, s *goquery.Selection) bool {
			link := ParseATag2Link(s, e.Request.URL)
			if link == nil {
				return true
			}
			if sameHostFilter.IsAllowed(link.URL.Host) {
				page.Links = append(page.Links, link)
				e.Request.Visit(link.URL.String())
			} else {
				page.ExternalLinks = append(page.ExternalLinks, link)
			}
			return true
		})
		if e.Request.ID == 1 {
			webSite.HomePage = page
		} else {
			sw.lock.Lock()
			webSite.Pages = append(webSite.Pages, page)
			sw.lock.Unlock()
		}
	})
	sw.collector.OnRequest(func(r *colly.Request) {
		sw.lock.Lock()
		if sw.linkMap[r.URL.String()] {
			log.Printf("skip url: %s", r.URL.String())
			r.Abort()
		} else {
			sw.linkMap[r.URL.String()] = true
		}
		sw.lock.Unlock()
		if r.ID > uint32(sw.maxPages) {
			r.Abort()
		}
	})

	sw.collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("%s: %+v\n", r.Request.URL.String(), err)
	})

	err = sw.collector.Visit(homeUrl)
	sw.collector.Wait()
	return webSite, err
}

// 缓存目录
func WithCacheDir(dir string) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.collector.CacheDir = dir
	}
}

// 并发数
func WithParallelism(n int) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.limitRule.Parallelism = n
	}
}

// WithUserAgent
// this WithUserAgent will cover DeviceType WithUserAgent
func WithUserAgent(ua string) SiteWalkerOption {
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
