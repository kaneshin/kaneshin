package main

import (
	"golang.org/x/exp/constraints"
)

func Square[T constraints.Integer | constraints.Float](list []T) func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		for i := range len(list) {
			if !yield(i, list[i]*list[i]) {
				return
			}
		}
	}
}

func Find[T constraints.Integer | constraints.Float](list []T, f func(T) bool) func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		for i := range len(list) {
			if f(list[i]) {
				if !yield(i, list[i]*list[i]) {
					return
				}
			}
		}
	}
}

func main() {
	v := []int{1, 2, 3, 4, 5}
	for i, x := range Square(v) {
		println(i, v[i], x)
	}

	for i, x := range Find(v, func(t int) bool { return t%2 == 0 }) {
		println(i, v[i], x)
	}
}
