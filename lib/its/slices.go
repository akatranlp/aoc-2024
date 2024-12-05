package its

import (
	"iter"
	"slices"
)

func Zip[T, K any](seq1 iter.Seq[T], seq2 iter.Seq[K]) iter.Seq2[T, K] {
	return func(yield func(T, K) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		next2, stop2 := iter.Pull(seq2)
		defer stop2()
		for {
			v1, ok1 := next1()
			v2, ok2 := next2()

			if !ok1 || !ok2 {
				return
			}

			if !yield(v1, v2) {
				return
			}
		}
	}
}

func Map[T, K any](seq iter.Seq[T], f func(T) K) iter.Seq[K] {
	return func(yield func(K) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func MapSlice[T, K any](seq []T, f func(T) K) []K {
	return slices.Collect(Map(slices.Values(seq), f))
}

type PredicateFunc[T any] func(T) bool

var FilterEmptyLines PredicateFunc[string] = func(row string) bool { return row != "" }

func Filter[T any](seq iter.Seq[T], predicate PredicateFunc[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !predicate(v) {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Filter2[K, V any](seq iter.Seq2[K, V], predicate PredicateFunc[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !predicate(v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func All[T any](seq iter.Seq[T], predicate PredicateFunc[T]) bool {
	for e := range seq {
		if !predicate(e) {
			return false
		}
	}
	return true
}

func Reduce[T, K any](seq iter.Seq[T], acc K, reduceFunc func(acc K, value T) K) K {
	for v := range seq {
		acc = reduceFunc(acc, v)
	}
	return acc
}

func RemoveIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
