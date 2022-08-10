package sitewalker

type PageType int

// 普通网页类型
const (
	Home PageType = iota + 1 // 首页
	About
	Case
	Category
	SubCategory
	Product
	Post
	FAQ
	TAG
	Contact
	Join
	CustomerService
	UnKnown
	Support
	Download
	Error
)

func (p PageType) Name(language int) string {
	// switch language {
	// case 1:
	// todo: complete this
	return ""
}

type PageContentType int

// 是列表页还是详情页
const (
	List   PageContentType = iota + 1 // 列表页
	Detail                            // 详情页
)

type Page struct {
	Title           string
	Desc            string
	Keywords        string
	Link            string
	H1              string
	OutPages        []*Page
	deep            int
	PageType        PageType
	PageContentType PageContentType
}

type WebSite struct {
	HomePage *Page
}
