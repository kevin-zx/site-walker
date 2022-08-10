package sitewalker

type SiteWalker struct {
	CacheDir string
}

type SiteWalkerOption func(sw *SiteWalker)

// 缓存目录
func CacheDir(dir string) SiteWalkerOption {
	return func(sw *SiteWalker) {
		sw.CacheDir = dir
	}
}

func NewSiteWalker(opts ...SiteWalkerOption) *SiteWalker {
	sw := &SiteWalker{}
	for _, opt := range opts {
		opt(sw)
	}
	return sw
}
