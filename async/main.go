package main

import (
	"log"
	"sort"

	"github.com/stretchr/testify/assert"
)

func testCollect() {
	ch1 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	assertEqual("collect int", collect(ch1), []int{1, 2, 3})

	ch2 := make(chan string, 10)
	ch2 <- "Hello"
	ch2 <- "World"
	close(ch2)
	assertEqual("collect string", collect(ch2), []string{"Hello", "World"})
}

func testInto() {
	ch1 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	res := []int{}
	into(ch1, &res)
	assertEqual("into", res, []int{1, 2, 3})
}

func testMap() {
	a := make(chan int, 10)
	r := mmap(a, func(it int) int {
		return it * 2
	})

	a <- 1
	a <- 2
	a <- 3
	close(a)

	assertEqual("map *2", collect(r), []int{2, 4, 6})
}

func testTake2() {
	b := make(chan int, 5)

	b <- 1
	b <- 2
	b <- 3
	close(b)

	bb := take(b, 2)

	assertEqual("take 2", collect(bb), []int{1, 2})
}

func testDrop2() {
	c := make(chan int, 5)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	cc := drop(c, 2)
	assertEqual("drop 2", collect(cc), []int{3})
}

func testMerge() {
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
	result := sort.IntSlice(collect(ff))
	result.Sort()
	assertEqual("merge", result, []int{1, 2, 3, 4, 5})
}

func testMult() {
	g := make(chan int)
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	m := mult(g)
	tap(m, ch1)
	tap(m, ch2)

	g <- 10

	untap(m, ch1)

	g <- 20
	close(g)
	close(ch1)
	close(ch2)

	assertEqual("mult and tap tap1", collect(ch1), []int{10})
	assertEqual("mult and tap tap2", collect(ch2), []int{10, 20})
}

func testSplit() {
	h := make(chan int)
	c1, c2 := split(h, func(i int) bool {
		return i > 2
	})

	h <- 1
	h <- 2
	h <- 3
	h <- 4
	close(h)

	assertEqual("split p1", collect(c1), []int{3, 4})
	assertEqual("split p2", collect(c2), []int{1, 2})
}

func main() {
	testCollect()
	testInto()
	testMap()
	testTake2()
	testDrop2()
	testMerge()
	testMult()
	testSplit()
}

func assertEqual[T any](name string, value T, expected T) {
	if !assert.ObjectsAreEqualValues(value, expected) {
		log.Printf("Fail: %s, expect %v to be %v.\n", name, value, expected)
	} else {
		log.Printf("Pass: %s.", name)
	}
}
