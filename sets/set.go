package sets

var emptyV = struct{}{}

type Set[T comparable] struct {
	holder map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		holder: map[T]struct{}{},
	}
}

func Just[T comparable](ele ...T) *Set[T] {
	var set = &Set[T]{holder: map[T]struct{}{}}

	for _, e := range ele {
		set.holder[e] = emptyV
	}

	return set
}

func (set *Set[T]) Contains(v T) bool {
	if set == nil {
		return false
	}

	_, ok := set.holder[v]
	return ok
}

func (set *Set[T]) Len() int {
	if set == nil {
		return 0
	}

	return len(set.holder)
}

func (set *Set[T]) ToSlice() []T {
	if set == nil {
		return []T{}
	}

	var ret []T
	for k := range set.holder {
		ret = append(ret, k)
	}
	return ret
}
