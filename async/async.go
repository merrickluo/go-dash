package main

// mix admix unmix
// mult tap untap
// split
// pipeline
// sliding-buffer
// split

import (
	"sync"
)

func mmap[T any, M any](ch <-chan T, f func(T) M) <-chan M {
	ret := make(chan M)

	go func() {
		defer close(ret)
		for it := range ch {
			ret <- f(it)
		}
	}()

	return ret
}

func take[T any](ch <-chan T, count int) <-chan T {
	ret := make(chan T)

	go func() {
		defer close(ret)
		for i := 0; i < count; i++ {
			ret <- <-ch
		}
	}()

	return ret
}

func drop[T any](ch <-chan T, count int) <-chan T {
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

func merge[T any](chs ...chan T) chan T {
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