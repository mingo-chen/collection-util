package slices

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("空集合", func(t *testing.T) {
		if got := Contains(nil, 3); got != false { // nil集合
			t.Errorf("Contains() = %v, want %v", got, false)
		}

		if got := Contains([]int{}, 2); got != false { // 空集合
			t.Errorf("Contains() = %v, want %v", got, false)
		}
	})

	t.Run("int类型", func(t *testing.T) {
		if got := Contains([]int{1, 3, 5}, 3); got != true {
			t.Errorf("Contains() = %v, want %v", got, true)
		}

		if got := Contains([]int{1, 3, 5}, 2); got != false {
			t.Errorf("Contains() = %v, want %v", got, false)
		}
	})

	t.Run("string类型", func(t *testing.T) {
		if got := Contains([]string{"tx", "ali", "baidu"}, "tx"); got != true {
			t.Errorf("Contains() = %v, want %v", got, true)
		}

		if got := Contains([]string{"tx", "ali", "baidu"}, "toutiao"); got != false {
			t.Errorf("Contains() = %v, want %v", got, false)
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("src为空", func(t *testing.T) {
		if got := Sub(nil, []int{1, 3, 5}); len(got) != 0 {
			t.Errorf("Sub() = %v, want %v", got, []int{})
		}
	})

	t.Run("dst为空", func(t *testing.T) {
		if got := Sub([]int{1, 3, 5}, nil); len(got) != 3 {
			t.Errorf("Sub() = %v, want %v", got, []int{1, 3, 5})
		}
	})

	t.Run("int集合相减", func(t *testing.T) {
		if got := Sub([]int{1, 3, 5}, []int{1, 2, 3}); len(got) != 1 {
			t.Errorf("Sub() = %v, want %v", got, []int{5})
		}
	})

	t.Run("string集合相减", func(t *testing.T) {
		if got := Sub([]string{"tx", "ali", "baidu"}, []string{"tx", "ali", "baidu"}); len(got) != 0 {
			t.Errorf("Sub() = %v, want %v", got, []string{})
		}
	})
}

func TestSort(t *testing.T) {
	type student struct {
		id   int
		name string
	}

	s1 := student{id: 1, name: "zhaoda"}
	s2 := student{id: 2, name: "liuer"}
	s3 := student{id: 3, name: "zhansan"}
	s4 := student{id: 4, name: "lisi"}
	s5 := student{id: 5, name: "wangwu"}

	ss := []student{s3, s2, s4, s1, s5}
	Sort(ss, func(a, b student) bool {
		return strings.Compare(a.name, b.name) < 0
	})

	t.Logf("output: %+v", ss)
}
