package urltool

import "strings"

func handleDoubleSlant(url string) string {
	protocol := ""
	if strings.HasPrefix(url, "https://") {
		protocol = "https://"
	}
	if strings.HasPrefix(url, "http://") {
		protocol = "http://"
	}
	url = strings.ReplaceAll(url, protocol, "")
	pathes := strings.Split(url, "/")
	finalUrl := strings.Join(pathes, "/")
	return protocol + finalUrl
	// var sb strings.Builder
	// sb.WriteString(protocol)
	// for _, path := range pathes {
	// 	if path == "" {
	// 		continue
	// 	}
	// 	sb.WriteString("/")
	// 	sb.WriteString(path)
	// }
	// return sb.String()
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
