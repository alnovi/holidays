package utils

import "strings"

func NormalizeURL(url string) string {
	if url == "" {
		return ""
	}

	if !strings.HasPrefix(url, "http") {
		url = "https://" + strings.Trim(url, "/")
	}

	return url
}
