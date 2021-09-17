package dash

// Temporary implementation just for writing tests

func Map(slice []int, f func(int) int) []int {
	ret := make([]int, len(slice), cap(slice))

	for i, it := range slice {
		ret[i] = f(it)
	}
	return ret
}

func FlatMap(slice []int, f func(int) []int) []int {
	ret := make([]int, 0)
	for _, it := range slice {
		for _, mapped := range f(it) {
			ret = append(ret, mapped)
		}
	}
	return ret
}

func Filter(slice []int, f func(int) bool) []int {
	ret := make([]int, 0)
	for _, it := range slice {
		if f(it) {
			ret = append(ret, it)
		}
	}
	return ret
}

func Reduce(slice []int, f func(int, int) int, acc int) int {
	for _, it := range slice {
		acc = f(acc, it)
	}
	return acc
}

func Take(slice []int, count int) []int {
	if count > len(slice) {
		return slice
	}
	ret := make([]int, count)
	for i := 0; i < count; i++ {
		ret[i] = slice[i]
	}

	return ret
}

func Drop(slice []int, count int) []int {
	if count > len(slice) {
		return []int{}
	}

	ret := make([]int, 0)
	for i := count; i < len(slice); i++ {
		ret = append(ret, slice[i])
	}
	return ret
}
