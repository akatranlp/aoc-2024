package its

import (
	"iter"
	"maps"
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

func ZipSlices[T, K any](slice1 []T, slice2 []K) iter.Seq2[T, K] {
	length := max(len(slice1), len(slice2))
	return func(yield func(T, K) bool) {
		for i := range length {
			if !yield(slice1[i], slice2[i]) {
				return
			}
		}
	}
}

type iterPuller[T any] struct {
	next  func() (T, bool)
	value T
}

func (i *iterPuller[T]) Next() bool {
	v, ok := i.next()
	i.value = v
	return ok
}

func (i *iterPuller[T]) Value() T {
	return i.value
}

func PullFromIter[T any](next func() (T, bool)) *iterPuller[T] {
	return &iterPuller[T]{next: next}
}

type Combination[T any] struct {
	L, R T
}

func AllCombinations[T any](slice []T, includeSelf bool) iter.Seq[Combination[T]] {
	return func(yield func(Combination[T]) bool) {
		for i := range len(slice) {
			start := i
			if !includeSelf {
				start += 1
			}
			for j := start; j < len(slice); j++ {
				if !yield(Combination[T]{slice[i], slice[j]}) {
					return
				}
			}
		}
	}
}

func Window[T any](seq iter.Seq[T], n int) iter.Seq[[]T] {
	window := make([]T, 0)
	return func(yield func([]T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for range n {
			v, ok := next()
			if !ok {
				yield(window)
				return
			}
			window = append(window, v)
		}
		ip := PullFromIter(next)
		for ip.Next() {
			window = append(window[1:], ip.value)
			if !yield(window[1:len(window):len(window)]) {
				return
			}
		}
	}
}

func Window2[T any](seq iter.Seq[T]) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		var first, second T
		var ok bool
		second, ok = next()
		if !ok {
			return
		}

		ip := PullFromIter(next)
		for ip.Next() {
			first = second
			second = ip.value
			if !yield(first, second) {
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

func Map1To2[K comparable, V, T any](seq iter.Seq[T], f func(T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

func Map2[K, V, T any](seq iter.Seq2[K, V], f func(K, V) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func MapSlice[T, K any](seq []T, f func(T) K) []K {
	return slices.Collect(Map(slices.Values(seq), f))
}

func Map2Slice[K comparable, V, T any](seq map[K]V, f func(K, V) T) []T {
	return slices.Collect(Map2(maps.All(seq), f))
}

type PredicateFunc[T any] func(T) bool
type PredicateFunc2[K, V any] func(k K, v V) bool

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

func FilterSlice[T any](slice []T, predicate PredicateFunc[T]) []T {
	return slices.Collect(Filter(slices.Values(slice), predicate))
}

func Filter2[K, V any](seq iter.Seq2[K, V], predicate PredicateFunc2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !predicate(k, v) {
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

func All2[K, V any](seq iter.Seq2[K, V], predicate PredicateFunc2[K, V]) bool {
	for k, v := range seq {
		if !predicate(k, v) {
			return false
		}
	}
	return true
}

func Any[T any](seq iter.Seq[T], predicate PredicateFunc[T]) bool {
	for v := range seq {
		if predicate(v) {
			return true
		}
	}
	return false
}

func Any2[K, V any](seq iter.Seq2[K, V], predicate PredicateFunc2[K, V]) bool {
	for k, v := range seq {
		if predicate(k, v) {
			return true
		}
	}
	return false
}

func Reduce[T, K any](seq iter.Seq[T], acc K, reduceFunc func(acc K, value T) K) K {
	for v := range seq {
		acc = reduceFunc(acc, v)
	}
	return acc
}

func Reduce2[K, V, T any](seq iter.Seq2[K, V], acc T, reduceFunc func(acc T, key K, value V) T) T {
	for k, v := range seq {
		acc = reduceFunc(acc, k, v)
	}
	return acc
}

func RemoveIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveIndexNew[T any](slice []T, index int) []T {
	newList := append(make([]T, 0, len(slice)-1), slice[:index]...)
	return append(newList, slice[index+1:]...)
}

func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var idx int
		for v := range seq {
			if !yield(idx, v) {
				return
			}
			idx++
		}
	}
}

func ForEach[T any](seq iter.Seq[T], f func(T)) {
	for v := range seq {
		f(v)
	}
}

func Range(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				return
			}
		}
	}
}
