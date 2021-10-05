// Package async provides functional utilities for channels.
package async

import (
	"sync"
	"time"

	"github.com/merrickluo/go-dash/dash"
)

// Map returns a new channel containing values produced by
// applying function f to every value from source channel ch.
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

// Filter returns a new channel containing values which
// pred(value) returns true.
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

// Reduce returns a channel containing a single value,
// accumulated by applying function f to
// the accumulator value and values from ch.
// The first accumulator value is init, the subsequent
// accumulator value is the return value of f.
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

// Take returns a new channel containing the first n values of ch.
func Take[T any](ch <-chan T, n int) <-chan T {
	ret := make(chan T)

	go func() {
		defer close(ret)
		for i := 0; i < n; i++ {
			ret <- <-ch
		}
	}()

	return ret
}

// Drop returns a new channel containing
// all the values from ch, except the first n.
func Drop[T any](ch <-chan T, n int) <-chan T {
	ret := make(chan T)

	go func() {
		defer close(ret)
		for i := 0; i < n; i++ {
			<-ch
		}

		for it := range ch {
			ret <- it
		}
	}()

	return ret
}

// Merge returns a new channel containing all values from chs.
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

// Split returns two channel.
// The first channel containing values that pred(value) == true.
// The second channel containing values that pred(value) == false.
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

// Collect returns a slice containing all the values from channel ch.
// ch must be closed.
func Collect[T any](ch <-chan T) []T {
	ret := make([]T, 0)
	for it := range ch {
		ret = append(ret, it)
	}
	return ret
}

// Into puts all values from channel ch into slice.
// ch must be closed.
func Into[T any](ch chan T, slice *[]T) {
	for it := range ch {
		*slice = append(*slice, it)
	}
}

// SlidingBuffer returns a new channel with size n.
// When full, the oldest value will be dropped.
func SlidingBuffer[T any](ch chan T, n uint) chan T {
	sb := make(chan T, n)
	full := int(n)
	go func() {
		defer close(sb)
		for i := range ch {
			if len(sb) == full {
				select {
				case <-time.After(10 * time.Millisecond): // best effort
				case <-sb:
				}
			}
			sb <- i
		}
	}()
	return sb
}

// DroppingBuffer returns a new channel with size n.
// When full, new values are ignored.
func DroppingBuffer[T any](ch chan T, n uint) chan T {
	db := make(chan T, n)
	full := int(n)
	go func() {
		defer close(db)
		for i := range ch {
			if len(db) != full {
				select {
				case <-time.After(10 * time.Millisecond): // best effort
				case db <- i:
				}
			}
		}
	}()
	return db
}

// Zip returns a new channel containing values that
// combines each value from ch1 and ch2 as Pair.
// Values in the result channel equals the channel has least values.
func Zip[T any, M any](ch1 chan T, ch2 chan M) chan dash.Pair[T, M] {
	ret := make(chan dash.Pair[T, M])

	go func() {
		defer close(ret)
		for v1 := range ch1 {
			v2, ok := <-ch2
			if !ok {
				return
			}
			ret <- dash.NewPair(v1, v2)
		}
	}()

	return ret
}
