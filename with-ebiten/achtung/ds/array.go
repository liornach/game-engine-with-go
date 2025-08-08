package ds

import "slices"

type Array[T comparable] struct {
	arr []T
}

func NewArray[T comparable]() Array[T] {
	return Array[T]{
		arr: []T{},
	}
}

func (a *Array[T]) Add(val T) {
	a.arr = append(a.arr, val)
}

func (a Array[T]) Contains(val T) bool {
	return slices.Contains(a.arr, val)
}

func (a Array[T]) ContainsFunc(f func(T) bool) bool {
	return slices.ContainsFunc(a.arr, f)
}

func (a Array[T]) Len() int {
	return len(a.arr)
}

func (a Array[T]) IndexOf(val T) (int, bool) {
	ret := slices.Index(a.arr, val)
	return ret, ret == -1
}

func (a Array[T]) IndexOfFunc(f func(T) bool) (int, bool) {
	ret := slices.IndexFunc(a.arr, f)
	return ret, ret == -1
}

func (a Array[T]) At(i int) T {
	return a.arr[i]
}

func (a Array[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for _, v := range a.arr {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
