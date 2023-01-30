package slices

// Contains 判断src中是否包含target
func Contains[T comparable](src []T, target T) bool {
	for _, e := range src {
		if e == target {
			return true
		}
	}

	return false
}

// Exist 判断src中是否存在target
func Exist[T comparable](src []T, target T) bool {
	return Contains(src, target)
}

// Add 两个集合相加
func Add[T any](src, dst []T) []T {
	return append(src, dst...)
}

// Sub 两个集合相减，O(N+M)复杂度，N为src长度，M为dst长度
func Sub[T comparable](src, dst []T) []T {
	var existSet = map[T]struct{}{} // 存在的集合
	for _, e := range dst {
		existSet[e] = struct{}{}
	}

	var ret []T
	for _, e := range src {
		if _, ok := existSet[e]; !ok {
			ret = append(ret, e)
		}
	}

	return ret
}
