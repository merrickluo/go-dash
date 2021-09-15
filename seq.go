package main

func Map[T any, M any](t []T, f func(T) M) []M {
	n := make([]M, len(t), cap(t))

	for i, e := range t {
		n[i] = f(e)
	}
	return n
}

func Filter[T any](t []T, f func(T) bool) []T {
	n := make([]T, 0)
	for _, e := range t {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func Reduce[T any, A any](t []T, f func(A, T) A, acc A) A {
	for _, e := range t {
		acc = f(acc, e)
	}
	return acc
}

func Take[T any](t []T, count int) []T {
	if count > len(t) {
		count = len(t)
	}
	n := make([]T, count)
	for i := 0; i < count; i++ {
		n[i] = t[i]
	}

	return n
}
