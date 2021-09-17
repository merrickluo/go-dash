package async

import (
	"sync"
)

func Map(ch <-chan int, f func(int) int) <-chan int {
	ret := make(chan int)

	go func() {
		defer close(ret)
		for it := range ch {
			ret <- f(it)
		}
	}()

	return ret
}

func Take(ch <-chan int, count int) <-chan int {
	ret := make(chan int)

	go func() {
		defer close(ret)
		for i := 0; i < count; i++ {
			ret <- <-ch
		}
	}()

	return ret
}

func Drop(ch <-chan int, count int) <-chan int {
	ret := make(chan int)

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

func Merge(chs ...chan int) chan int {
	ret := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(c <-chan int) {
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

func Split(ch <-chan int, pred func(int) bool) (chan int, chan int) {
	c1 := make(chan int, 8)
	c2 := make(chan int, 8)
	go func(c1 chan int, c2 chan int) {
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

func Collect(ch <-chan int) []int {
	ret := make([]int, 0)
	for it := range ch {
		ret = append(ret, it)
	}
	return ret
}

func Into(ch chan int, slice *[]int) {
	for it := range ch {
		*slice = append(*slice, it)
	}
}
