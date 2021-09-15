package main

import (
	"log"

	"github.com/stretchr/testify/assert"
)

func main() {
	arr := []int{1, 2, 3}
	assertEqual("map", mmap(arr, func(i int) int { return i * 2 }), []int{2, 4, 6})
	assertEqual("flatMap", flatMap(arr, func(i int) []int { return []int{i, i * 2} }), []int{1, 2, 2, 4, 3, 6})
	assertEqual("filter", filter(arr, func(i int) bool { return i > 2 }), []int{3})
	assertEqual("reduce", reduce(arr, func(acc int, i int) int { return acc + i }, 0), 6)
	assertEqual("take 2", take(arr, 2), []int{1, 2})
	assertEqual("drop 2", drop(arr, 2), []int{3})
}

func assertEqual[T any](name string, value T, expected T) {
	if !assert.ObjectsAreEqualValues(value, expected) {
		log.Printf("Fail: %s, expect %v to be %v.\n", name, value, expected)
	} else {
		log.Printf("Pass: %s.", name)
	}
}
