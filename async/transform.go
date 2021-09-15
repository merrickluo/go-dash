package async

func Map[T any, R any](chs []chan T) []chan R {
	return nil
}

func Take[T any](ch chan T, n int) chan T {
	return nil
}

type MultChan[T any] chan T

func Mult[T any](ch chan T) MultChan[T] {
	return nil
}

func Tap[T any](ch MultChan[T]) chan T {
	return nil
}
