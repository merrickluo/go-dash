package main

func mmap[T any, M any](slice []T, f func(T) M) []M {
	ret := make([]M, len(slice), cap(slice))

	for i, it := range slice {
		ret[i] = f(it)
	}
	return ret
}

func flatMap[T any, M any](slice []T, f func(T) []M) []M {
	ret := make([]M, 0)
	for _, it := range slice {
		for _, mapped := range f(it) {
			ret = append(ret, mapped)
		}
	}
	return ret
}

func filter[T any](slice []T, f func(T) bool) []T {
	ret := make([]T, 0)
	for _, it := range slice {
		if f(it) {
			ret = append(ret, it)
		}
	}
	return ret
}

func reduce[T any, A any](slice []T, f func(A, T) A, acc A) A {
	for _, it := range slice {
		acc = f(acc, it)
	}
	return acc
}

func take[T any](slice []T, count int) []T {
	if count > len(slice) {
		return slice
	}
	ret := make([]T, count)
	for i := 0; i < count; i++ {
		ret[i] = slice[i]
	}

	return ret
}

func drop[T any](slice []T, count int) []T {
	if count > len(slice) {
		return []T{}
	}

	ret := make([]T, 0)
	for i := count; i < len(slice); i++ {
		ret = append(ret, slice[i])
	}
	return ret
}
