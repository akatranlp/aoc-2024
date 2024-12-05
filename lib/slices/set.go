package slices

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
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

func (s Set[T]) Has(k T) bool {
	_, ok := s[k]
	return ok
}
