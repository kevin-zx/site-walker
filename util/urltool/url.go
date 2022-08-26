package urltool

import "strings"

// 判断是否是可用的href
// 排除 javascript, mailto, tel 等其他非http协议的href
func IsValidHref(href string) bool {
	if strings.Contains(href, "javascript") {
		return false
	}

	href = strings.ToLower(href)
	href = strings.TrimSpace(href)
	if strings.Contains(href, ":") && strings.HasPrefix(href, "http") {
		return false
	}
	return true
}
