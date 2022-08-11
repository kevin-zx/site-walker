package pagetype

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

func (p PageType) CNName() string {
	switch p {
	case Home:
		return "首页"
	case About:
		return "关于我们"
	case Case:
		return "案例"
	case Category:
		return "分类"
	case SubCategory:
		return "子分类"
	case Product:
		return "产品"
	case Post:
		return "文章"
	case FAQ:
		return "常见问题"
	case TAG:
		return "标签"
	case Contact:
		return "联系我们"
	case Join:
		return "加入我们"
	case UnKnown:
		return "404"
	case CustomerService:
		return "客服"
	case Support:
		return "技术支持"
	case Download:
		return "下载"
	case Error:
		return "错误"
	default:
		return ""
	}

}

type PageContentType int

// 是列表页还是详情页
const (
	List   PageContentType = iota + 1 // 列表页
	Detail                            // 详情页
)
