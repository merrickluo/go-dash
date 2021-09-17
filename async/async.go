package async

// mix admix unmix
// pipeline

import (
	"time"
	"sync"
)

func Map[T any, M any](ch <-chan T, f func(T) M) <-chan M {
	ret := make(chan M)

	go func() {
		defer close(ret)
		for it := range ch {
			ret <- f(it)
		}
	}()

	return ret
}

func Filter[T any](ch <-chan T, pred func(T) bool) <-chan T {
	ret := make(chan T)

	go func() {
		defer close(ret)
		for it := range ch {
			if pred(it) {
				ret <- it
			}
		}
	}()

	return ret
}

func Reduce[T any, M any](ch <-chan T, f func(M, T) M, init M) <-chan M {
	ret := make(chan M)

	go func() {
		defer close(ret)
		var acc = init
		for it := range ch {
			acc = f(acc, it)
		}
		ret <- acc
	}()

	return ret
}

func Take[T any](ch <-chan T, count int) <-chan T {
	ret := make(chan T)

	go func() {
		defer close(ret)
		for i := 0; i < count; i++ {
			ret <- <-ch
		}
	}()

	return ret
}

func Drop[T any](ch <-chan T, count int) <-chan T {
	ret := make(chan T)

	go func() {
		defer close(ret)
		for i := 0; i < count; i++ {
			<-ch
		}

		for it := range ch {
			ret <- it
		}
	}()

	return ret
}

func Merge[T any](chs ...chan T) chan T {
	ret := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(c <-chan T) {
			defer wg.Done()
			for it := range c {
				ret <- it
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(ret)
	}()
	return ret
}

func Split[T any](ch <-chan T, pred func(T) bool) (chan T, chan T) {
	c1 := make(chan T, 8)
	c2 := make(chan T, 8)
	go func(c1 chan T, c2 chan T) {
		defer close(c1)
		defer close(c2)

		for it := range ch {
			if pred(it) {
				c1 <- it
			} else {
				c2 <- it
			}
		}
	}(c1, c2)

	return c1, c2
}

func Collect[T any](ch <-chan T) []T {
	ret := make([]T, 0)
	for it := range ch {
		ret = append(ret, it)
	}
	return ret
}

func Into[T any](ch chan T, slice *[]T) {
	for it := range ch {
		*slice = append(*slice, it)
	}
}

func SlidingBuffer[T any](ch chan T, n uint) (chan T) {
	sb := make(chan T, n)
	full := int(n)
	go func(){
		defer close(sb)
		for i := range ch {
			if len(sb) == full {
				select {
				case <-time.After(10*time.Millisecond): // best effort
				case <-sb:
				}
			}
			sb <- i
		}
	}()
	return sb
}

func DroppingBuffer[T any](ch chan T, n uint) (chan T) {
	db := make(chan T, n)
	full := int(n)
	go func(){
		defer close(db)
		for i := range ch {
			if len(db) != full {
				select {
				case <-time.After(10*time.Millisecond): // best effort
				case db <- i:
				}
			}
		}
	}()
	return db
}
