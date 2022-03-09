// Package dash brings a number of functional utilities to Go.
package dash

// Pair is a pair of two values.
// Access the values with .First and .Second
type Pair[T any, M any] struct {
	First  T
	Second M
}

// NewPair creates a Pair contains values first and second.
func NewPair[T any, M any](first T, second M) Pair[T, M] {
	return Pair[T, M]{First: first, Second: second}
}
