package urltool

import "strings"

func handleDoubleSlant(url string) string {
	protocol := ""
	if strings.Contains(url, "://") {
		parts := strings.Split(url, "://")
		if len(parts) > 1 {
			protocol = parts[0] + "://"
			url = parts[1]
		}
	}

	for strings.Contains(url, "//") {
		url = strings.ReplaceAll(url, "//", "/")
	}

	return protocol + url
}

func handleEndURLUtf8EncodeSpace(url string) string {
	for strings.HasSuffix(url, "%20") {
		url = strings.TrimSuffix(url, "%20")
	}
	return url
}

func handleUnicodeEncodeSpace(url string) string {
	url = strings.ReplaceAll(url, "%u0020", "")
	url = strings.ReplaceAll(url, "&#10;", "")
	url = strings.ReplaceAll(url, "&#9;", "")
	return url
}
