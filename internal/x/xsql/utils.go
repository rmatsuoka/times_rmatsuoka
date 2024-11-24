package xsql

import (
	"iter"
	"slices"
)

func AnySeq[T any](seq iter.Seq[T]) iter.Seq[any] {
	return func(yield func(any) bool) {
		for v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

func CollectAny[T any](seq iter.Seq[T]) []any {
	return slices.Collect(AnySeq(seq))
}

func AnySlice[S ~[]T, T any](s S) []any {
	return CollectAny(slices.Values(s))
}
