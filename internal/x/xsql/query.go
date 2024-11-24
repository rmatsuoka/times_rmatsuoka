package xsql

import (
	"regexp"
	"strings"
)

var listQueryRe = regexp.MustCompile(`\{\{.+?\}\}`)

func ListQuery(query string, ns ...int) string {
	i := 0
	return listQueryRe.ReplaceAllStringFunc(query, func(s string) string {
		s = s[2 : len(s)-2]
		n := ns[i]
		if n <= 0 {
			panic("xsql.ListQuery: zero-length list is not allowed")
		}
		i++
		return strings.Repeat(s+",", n-1) + s
	})
}
