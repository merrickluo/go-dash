package main

import "fmt"

func main() {
	// map
	fmt.Println("map(i * 2)")
	a := make(chan int, 10)
	r := mmap(a, func(it int) int {
		return it * 2
	})

	a <- 1
	a <- 2
	a <- 3
	close(a)

	for it := range r {
		fmt.Println(it)
	}

	// take
	fmt.Println("take(2)")
	b := make(chan int, 5)

	b <- 1
	b <- 2
	b <- 3
	close(b)

	bb := take(b, 2)

	for it := range bb {
		fmt.Println(it)
	}

	// drop
	fmt.Println("drop(2)")
	c := make(chan int, 5)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	cc := drop(c, 2)
	for it := range cc {
		fmt.Println(it)
	}

	// merge
	fmt.Println("merge(d, e)")
	d := make(chan int, 5)
	e := make(chan int, 5)

	d <- 1
	e <- 2
	d <- 3
	e <- 4
	d <- 5
	close(d)
	close(e)

	ff := merge(d, e)
	for it := range ff {
		fmt.Println(it)
	}

	// mult
	fmt.Println("start mult")
	s := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)

	m := mult(s)
	tap(m, ch1)
	tap(m, ch2)

	s <- 10

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
