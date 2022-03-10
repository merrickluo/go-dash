package dash

import (
	"math/rand"
	"time"
)

// Map returns a new slice containing values produced
// by applying function f to every value in slice.
func Map[T any, M any](slice []T, f func(T) M) []M {
	ret := make([]M, len(slice), cap(slice))

	for i, it := range slice {
		ret[i] = f(it)
	}
	return ret
}

// FlatMap returns a new slice containing values produced
// by applying function f to every value in slice,
// and flattern to a one-dimensional slice.
func FlatMap[T any, M any](slice []T, f func(T) []M) []M {
	ret := make([]M, 0)
	for _, it := range slice {
		for _, mapped := range f(it) {
			ret = append(ret, mapped)
		}
	}
	return ret
}

// ParallelMap works just like Map including the result order,
// but f is applied in parallel with goroutine.
// This version creates len(slice) of goroutines,
// use ParallelMapN if the amount of goroutines created
// at the same time needs to be limited.
func ParallelMap[T any, M any](slice []T, f func(T) M) []M {
	ret := make([]M, len(slice), len(slice))

	ch := make(chan Pair[int, M])
	defer close(ch)

	for i, it := range slice {
		go func(i int, it T) {
			ch <- NewPair(i, f(it))
		}(i, it)
	}

	c := 0
	for c < len(slice) {
		p := <-ch
		ret[p.First] = p.Second
		c += 1
	}

	return ret
}

// ParallelMapN is a version of ParallelMap,
// but has a limit of goroutine created at the same time.
func ParallelMapN[T any, M any](slice []T, f func(T) M, limit int) []M {
	ret := make([]M, len(slice), len(slice))

	ch := make(chan Pair[int, M], limit)
	defer close(ch)

	sem := make(chan struct{}, limit)
	defer close(sem)

	for i, it := range slice {
		sem <- struct{}{}
		go func(i int, it T) {
			ch <- NewPair(i, f(it))
			<-sem
		}(i, it)
	}

	c := 0
	for c < len(slice) {
		p := <-ch
		ret[p.First] = p.Second
		c += 1
	}

	return ret
}

// Filter returns a new slice containing all values of slice
// that the function when function pred returns true.
func Filter[T any](slice []T, pred func(T) bool) []T {
	ret := make([]T, 0)
	for _, it := range slice {
		if pred(it) {
			ret = append(ret, it)
		}
	}
	return ret
}

// Reduce combines all values of slice into a single value by applying
// function f with the accumulator value and every value in slice.
// First iteration of accumulator value is init, the subsequent
// accumulator value is the return value of f.
func Reduce[T any, A any](slice []T, f func(A, T) A, init A) A {
	acc := init
	for _, it := range slice {
		acc = f(acc, it)
	}
	return acc
}

// Take returns a new slice containing the first n values of slice.
func Take[T any](slice []T, n int) []T {
	if n > len(slice) {
		return slice
	}
	ret := make([]T, n)
	for i := 0; i < n; i++ {
		ret[i] = slice[i]
	}

	return ret
}

// Drop returns a new slice containing the remaining values of slice
// except the first n values.
func Drop[T any](slice []T, n int) []T {
	if n > len(slice) {
		return []T{}
	}

	ret := make([]T, 0)
	for i := n; i < len(slice); i++ {
		ret = append(ret, slice[i])
	}
	return ret
}

// Intersection returns a single slice containing uniq values among slices.
func Intersection[T comparable](slices ...[]T) (result []T) {
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

// Uniq returns a new slice containing all uniq values in slice.
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

// Every returns true if pred never returns false,
// returns false otherwise.
func Every[T any](slice []T, pred func(T) bool) bool {
	for _, v := range slice {
		if !pred(v) {
			return false
		}
	}
	return true
}

// Some returns false if pred never returns true,
// returns true otherwise.
func Some[T any](slice []T, pred func(T) bool) bool {
	for _, v := range slice {
		if pred(v) {
			return true
		}
	}
	return false
}

// None returns true if pred never returns true,
// returns false otherwise.
func None[T any](slice []T, pred func(T) bool) bool {
	return !Some(slice, pred)
}

// GroupBy returns a map.
// The keys are the return values of applying f with the values of slice.
// The values are values of slice which evaluates to it's key.
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

// Shuffle returns a new slice which containing all the values
// of slice, but in a random order.
func Shuffle[T any](slice []T) []T {
	l := len(slice)
	result := make([]T, l)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(l)
	for i, j := range perm {
		result[i] = slice[j]
	}
	return result
}

// Include returns true if value e is in slice.
func Include[T comparable](slice []T, e T) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}
	return false
}

// Cycle returns a new slice by repeating slice n times.
func Cycle[T any](slice []T, n uint) []T {
	l := uint(len(slice))
	ll := n * l

	ret := make([]T, ll)
	var i uint = 0
	var j uint = 0
	for {
		if i == ll {
			break
		}

		ret[i] = slice[j]
		i++
		j++

		if j == l {
			j = 0
		}
	}
	return ret
}

// Reverse return a new slice in reversed order.
func Reverse[T any](slice []T) []T {
	size := len(slice)
	ret := make([]T, size)

	for i, it := range slice {
		ret[size-i-1] = it
	}

	return ret
}

// Partition turn a slice into n-sized(n>0) slices
// the last k elements are dropped if k<n
func Partition[T any](slice []T, n int) [][]T {
	u := len(slice) / n
	ret := make([][]T, u)
	for i := 0; i < u; i++ {
		ret[i] = make([]T, n)
		copy(ret[i], slice[i*n:i*n+n])
	}
	return ret
}
