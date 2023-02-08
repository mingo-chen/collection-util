package sets

var emptyV = struct{}{}

// Set结构，不可重复
// 底层基于map来实现
type Set[T comparable] struct {
	holder map[T]struct{}
}

// New 创建一个空集合
func New[T comparable]() Set[T] {
	return Set[T]{
		holder: map[T]struct{}{},
	}
}

// Just 通过现有的slice生成集合
func Just[T comparable](ele ...T) Set[T] {
	var set = Set[T]{holder: map[T]struct{}{}}

	for _, e := range ele {
		set.holder[e] = emptyV
	}

	return set
}

// Add 向集合中添加一个元素
func (set Set[T]) Add(v T) Set[T] {
	set.holder[v] = emptyV
	return set
}

// Adds 向集合中批量添加元素
func (set Set[T]) Adds(vs ...T) Set[T] {
	for _, v := range vs {
		set.holder[v] = emptyV
	}
	return set
}

// Remove 向集合中删除一个元素
func (set Set[T]) Remove(v T) Set[T] {
	delete(set.holder, v)
	return set
}

// Removes 向集合中批量删除元素
func (set Set[T]) Removes(vs ...T) Set[T] {
	for _, v := range vs {
		delete(set.holder, v)
	}
	return set
}

// Contains 判断集合中是否包含元素
func (set Set[T]) Contains(v T) bool {
	if set.holder == nil {
		return false
	}

	_, ok := set.holder[v]
	return ok
}

// Len 获取集合长度
func (set Set[T]) Len() int {
	if set.holder == nil {
		return 0
	}

	return len(set.holder)
}

// ToSlice 集合转为Slice
func (set Set[T]) ToSlice() []T {
	if set.holder == nil {
		return []T{}
	}

	var ret []T
	for k := range set.holder {
		ret = append(ret, k)
	}
	return ret
}
