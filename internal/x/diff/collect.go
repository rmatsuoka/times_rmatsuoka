package diff

import "fmt"

func stringedKeyMap[M ~map[K]V, K comparable, V any](m M) map[string]V {
	strm := make(map[string]V, len(m))
	for k, v := range m {
		strm[fmt.Sprint(k)] = v
	}
	return strm
}
