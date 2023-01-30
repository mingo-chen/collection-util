package maps

func Keys[K comparable, V any](m map[K]V) []K {
	var ret []K
	for k := range m {
		ret = append(ret, k)
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
