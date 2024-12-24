package diff

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"
)

func Maps[M ~map[K]V, K comparable, V comparable](m, n M) string {
	return MapsFunc(m, n, func(v, w V) bool { return v == w })
}

func MapsFunc[M ~map[K]V, K comparable, V any](m, n M, equal func(V, V) bool) string {
	return OrderedMapsFunc(stringedKeyMap(m), stringedKeyMap(n), equal)
}

func OrderedMaps[M ~map[K]V, K cmp.Ordered, V comparable](m, n M) string {
	return OrderedMapsFunc(m, n, func(v, w V) bool { return v == w })
}

func OrderedMapsFunc[M ~map[K]V, K cmp.Ordered, V any](m, n M, equal func(V, V) bool) string {
	return orderedMapsFunc(m, n, cmp.Compare, equal)
}

func orderedMapsFunc[M ~map[K]V, K comparable, V any](m, n M, cmp func(K, K) int, equal func(V, V) bool) string {
	b := new(strings.Builder)
	mkeys := slices.SortedFunc(maps.Keys(m), cmp)
	nkeys := slices.SortedFunc(maps.Keys(n), cmp)
	ni := 0
	for _, key := range mkeys {
		for ni < len(nkeys) && cmp(key, nkeys[ni]) > 0 {
			fmt.Fprintf(b, "+ %v: %+v\n", nkeys[ni], n[nkeys[ni]])
			ni++
		}

		if ni < len(nkeys) && cmp(key, nkeys[ni]) == 0 {
			if equal(m[key], n[key]) {
				fmt.Fprintf(b, "  %v: %+v\n", key, m[key])
			} else {
				fmt.Fprintf(b, "- %v: %+v\n", key, m[key])
				fmt.Fprintf(b, "+ %v: %+v\n", key, n[key])
			}
			ni++
		} else {
			fmt.Fprintf(b, "- %v: %+v\n", key, m[key])
		}
	}
	for ; ni < len(nkeys); ni++ {
		fmt.Fprintf(b, "+ %v: %+v\n", nkeys[ni], n[nkeys[ni]])
	}
	return b.String()
}
