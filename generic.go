package main

import (
	"sort"
)

type Number interface {
	~uint | ~uint32 | ~uint64 | ~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Min 最小数
func Min[T Number](base, comp T) T {
	res := base
	if base > comp {
		res = comp
	}
	return res
}

// Max 最大数
func Max[T Number](base, comp T) T {
	res := base
	if base < comp {
		res = comp
	}
	return res
}

// Cond Ternary Operator 三元运算
func Cond[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// SortSlice 切片排序方法
func SortSlice[T any](slice []T, less func(T, T) bool) {
	sort.Sort(SliceFn[T]{slice, less})
}

// SliceFn implements sort.Interface for a slice of T.
type SliceFn[T any] struct {
	s    []T
	less func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}
func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T]) Less(i, j int) bool {
	return s.less(s.s[i], s.s[j])
}
