package maps

import "github.com/mingo-chen/collection-util/sets"

func Keys[K comparable, V any](m map[K]V) sets.Set[K] {
	var ret = sets.New[K]()
	for k := range m {
		ret.Add(k)
	}

	return ret
}

func Values[K comparable, V any](m map[K]V) []V {
	var ret []V
	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}
