package urltool

import (
	"net/url"
	"strings"
)

// 判断是否是可用的href
// 排除 javascript, mailto, tel 等其他非http协议的href
func IsValidHref(href string) bool {
	if strings.Contains(href, "javascript") {
		return false
	}

	href = strings.ToLower(href)
	href = strings.TrimSpace(href)
	if strings.Contains(href, ":") && !strings.HasPrefix(href, "http") {
		return false
	}
	return true
}

func ClearHref(href string) string {
	href = strings.ToLower(href)
	href = strings.TrimSpace(href)
	href = handleDoubleSlant(href)
	href = handleEndURLUtf8EncodeSpace(href)
	href = handleUnicodeEncodeSpace(href)
	return href
}

func ConvertHref2URL(href string, currURL *url.URL) (*url.URL, error) {
	return currURL.Parse(href)
}
