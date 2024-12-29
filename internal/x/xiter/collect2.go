package xiter

import "iter"

type KV[K, V any] struct {
	K K
	V V
}

func Collect2[K, V any](seq iter.Seq2[K, V]) []KV[K, V] {
	var kv []KV[K, V]
	for k, v := range seq {
		kv = append(kv, KV[K, V]{K: k, V: v})
	}
	return kv
}
