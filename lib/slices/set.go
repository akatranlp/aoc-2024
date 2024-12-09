package slices

import "iter"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func NewSetFromIter[T comparable](seq iter.Seq[T]) Set[T] {
	return NewSet[T]().SetIter(seq)
}

func NewSetWithValues[T comparable](slice ...T) Set[T] {
	return NewSet[T]().SetValues(slice...)
}

func (s Set[T]) Outer(other Set[T]) Set[T] {
	result := make(Set[T])
	for k, v := range s {
		if _, ok := other[k]; !ok {
			result[k] = v
		}
	}
	for k, v := range other {
		if _, ok := s[k]; !ok {
			result[k] = v
		}
	}
	return result
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := make(Set[T])
	for k, v := range s {
		if _, ok := other[k]; ok {
			result[k] = v
		}
	}
	return result
}

func (s Set[T]) Set(k T) Set[T] {
	s[k] = struct{}{}
	return s
}

func (s Set[T]) SetIter(seq iter.Seq[T]) Set[T] {
	for v := range seq {
		s.Set(v)
	}
	return s
}

func (s Set[T]) SetValues(slice ...T) Set[T] {
	for _, v := range slice {
		s.Set(v)
	}
	return s
}

func (s Set[T]) Has(k T) bool {
	_, ok := s[k]
	return ok
}
