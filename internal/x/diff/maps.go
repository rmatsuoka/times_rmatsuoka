package diff

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"

	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xiter"
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
	return orderedSeq2Func(maps.All(m), maps.All(n), cmp.Compare, equal)
}

func orderedSeq2Func[K comparable, V any](m, n iter.Seq2[K, V], cmp func(K, K) int, equal func(V, V) bool) string {
	b := new(strings.Builder)

	mkvs := xiter.Collect2(m)
	nkvs := xiter.Collect2(n)

	slices.SortFunc(mkvs, func(x, y xiter.KV[K, V]) int { return cmp(x.K, y.K) })
	slices.SortFunc(nkvs, func(x, y xiter.KV[K, V]) int { return cmp(x.K, y.K) })

	ni := 0
	for mi := range mkvs {
		for ni < len(nkvs) && cmp(mkvs[mi].K, nkvs[ni].K) > 0 {
			fmt.Fprintf(b, "+ %v: %+v\n", nkvs[ni].K, nkvs[ni].V)
			ni++
		}

		if ni < len(nkvs) && cmp(mkvs[mi].K, nkvs[ni].K) == 0 {
			if equal(mkvs[mi].V, nkvs[ni].V) {
				fmt.Fprintf(b, "  %v: %+v\n", mkvs[mi].K, mkvs[mi].V)
			} else {
				fmt.Fprintf(b, "- %v: %+v\n", mkvs[mi].K, mkvs[mi].V)
				fmt.Fprintf(b, "+ %v: %+v\n", mkvs[mi].K, nkvs[ni].V)
			}
			ni++
		} else {
			fmt.Fprintf(b, "- %v: %+v\n", mkvs[mi].K, mkvs[mi].V)
		}
	}
	for ; ni < len(nkvs); ni++ {
		fmt.Fprintf(b, "+ %v: %+v\n", nkvs[ni].K, nkvs[ni].V)
	}
	return b.String()
}
