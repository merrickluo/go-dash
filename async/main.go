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
	fmt.Println("mult")
	g := make(chan int)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	m := mult(g)
	tap(m, ch1)
	tap(m, ch2)

	g <- 10

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	close(g)
	close(ch1)
	close(ch2)

	// split
	h := make(chan int)
	c1, c2 := split(h, func(i int) bool {
		return i > 2
	})

	h <- 1
	h <- 2
	h <- 3
	h <- 4
	close(h)

	fmt.Println("spliting i > 2")
	fmt.Println("c1")
	for it := range c1 {
		fmt.Println(it)
	}
	fmt.Println("c2")
	for it := range c2 {
		fmt.Println(it)
	}
}
