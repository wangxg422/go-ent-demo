package strutil

import "net/url"

// 对字符串进行url编码
func URLEncoding(s string) string {
	return url.QueryEscape(s)
}
