package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println("array: [1, 2, 3]")

	// map
	fmt.Print("map i * 2: ")
	fmt.Println(mmap(arr, func(i int) int { return i * 2 }))

	// flatMap
	fmt.Print("flatMap [i, i * 2]: ")
	fmt.Println(flatMap(arr, func(i int) []int { return []int{i, i * 2} }))

	// filter
	fmt.Print("filter i > 2: ")
	fmt.Println(filter(arr, func(i int) bool { return i > 2 }))

	// reduce
	fmt.Print("reduce +: ")
	fmt.Println(reduce(arr, func(acc int, i int) int { return acc + i }, 0))

	// take
	fmt.Print("take 2: ")
	fmt.Println(take(arr, 2))

	// drop
	fmt.Print("drop 2: ")
	fmt.Println(drop(arr, 2))
}
