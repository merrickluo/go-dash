package dash

func Map[T any, M any](slice []T, f func(T) M) []M {
	ret := make([]M, len(slice), cap(slice))

	for i, it := range slice {
		ret[i] = f(it)
	}
	return ret
}

func FlatMap[T any, M any](slice []T, f func(T) []M) []M {
	ret := make([]M, 0)
	for _, it := range slice {
		for _, mapped := range f(it) {
			ret = append(ret, mapped)
		}
	}
	return ret
}

func Filter[T any](slice []T, f func(T) bool) []T {
	ret := make([]T, 0)
	for _, it := range slice {
		if f(it) {
			ret = append(ret, it)
		}
	}
	return ret
}

func Reduce[T any, A any](slice []T, f func(A, T) A, acc A) A {
	for _, it := range slice {
		acc = f(acc, it)
	}
	return acc
}

func Take[T any](slice []T, count int) []T {
	if count > len(slice) {
		return slice
	}
	ret := make([]T, count)
	for i := 0; i < count; i++ {
		ret[i] = slice[i]
	}

	return ret
}

func Drop[T any](slice []T, count int) []T {
	if count > len(slice) {
		return []T{}
	}

	ret := make([]T, 0)
	for i := count; i < len(slice); i++ {
		ret = append(ret, slice[i])
	}
	return ret
}

func Intersection[T comparable](slices... []T) (result []T) {
	l := uint(len(slices))
	m := map[T]uint{}
	for _, s := range slices {
		for _, k := range s {
			v, found := m[k]
			if !found {
				m[k] = 1
			} else {
				m[k] = v + 1
			}
		}
	}
	for k, v := range m {
		if v == l {
			result = append(result, k)
		}
	}
	return
}

func Uniq[T comparable](slice []T) (result []T) {
	m := map[T]bool{}
	for _, v := range slice {
		if _, found := m[v]; found {
			continue
		} else {
			result = append(result, v)
			m[v] = true
		}
	}
	return
}

func Every[T any](slice []T, pred func(T) bool) bool {
	for _, v := range slice {
		if !pred(v) {
			return false
		}
	}
	return true
}

func Some[T any](slice []T, pred func(T) bool) bool {
	for _, v := range slice {
		if pred(v) {
			return true
		}
	}
	return false
}

func None[T any](slice []T, pred func(T) bool) bool {
	return !Some(slice, pred)
}

func GroupBy[T any, D comparable](slice []T, f func(T) D) map[D][]T {
	m := map[D][]T{}
	for _, v := range slice {
		d := f(v)
		if _, found := m[d]; found {
			m[d] = append(m[d], v)
		} else {
			m[d] = []T{v}
		}
	}
	return m
}
