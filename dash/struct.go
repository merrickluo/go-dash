package dash

type Pair[T any, M any] struct {
	First  T
	Second M
}

func NewPair[T any, M any](first T, second M) Pair[T, M] {
	return Pair[T, M]{First: first, Second: second}
}
