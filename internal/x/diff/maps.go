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
	b := new(strings.Builder)
	mkeys := slices.Collect(maps.Keys(m))
	nkeys := slices.Collect(maps.Keys(n))

	seen := make(map[K]bool, max(len(m), len(n)))

	for _, key := range mkeys {
		seen[key] = true
		if _, ok := n[key]; !ok {
			fmt.Fprintf(b, "- %v: %+v\n", key, m[key])
			continue
		}

		if equal(m[key], n[key]) {
			fmt.Fprintf(b, "  %v: %+v\n", key, m[key])
			continue
		}

		fmt.Fprintf(b, "- %v: %+v\n", key, m[key])
		fmt.Fprintf(b, "+ %v: %+v\n", key, n[key])
	}

	for _, key := range nkeys {
		if seen[key] {
			continue
		}
		fmt.Fprintf(b, "+ %v: %v\n", key, n[key])
	}
	return b.String()
}

func OrderedMaps[M ~map[K]V, K cmp.Ordered, V comparable](m, n M) string {
	return OrderedMapsFunc(m, n, func(v, w V) bool { return v == w })
}

func OrderedFuncMaps[M ~map[K]V, K comparable, V comparable](m, n M, cmp func(K, K) int) string {
	return OrderedFuncMapsFunc(m, n, cmp, func(v, w V) bool { return v == w })
}

func OrderedMapsFunc[M ~map[K]V, K cmp.Ordered, V any](m, n M, equal func(V, V) bool) string {
	return OrderedFuncMapsFunc(m, n, cmp.Compare, equal)
}

func OrderedFuncMapsFunc[M ~map[K]V, K comparable, V any](m, n M, cmp func(K, K) int, equal func(V, V) bool) string {
	b := new(strings.Builder)
	mkeys := slices.SortedFunc(maps.Keys(m), cmp)
	nkeys := slices.SortedFunc(maps.Keys(n), cmp)
	ni := 0
	for _, key := range mkeys {
		for ni < len(nkeys) && key != nkeys[ni] {
			fmt.Fprintf(b, "+ %v: %+v\n", nkeys[ni], n[nkeys[ni]])
			ni++
		}

		if ni < len(nkeys) {
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
