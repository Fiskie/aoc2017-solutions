package main

import "fmt"

func indexOf(array []int, value int) int {
	for i, other := range array {
		if other == value {
			return i
		}
	}

	return -1
}

func main() {
	input := 304
	pos := 0
	array := []int{0}

	for i := 1; i < 2018; i++ {
		pos = ((pos + input) % len(array)) + 1

		first := make([]int, len(array[:pos]))
		copy(first, array[:pos])

		array = append(append(first, i), array[pos:]...)
	}

	for i, num := range array {
		fmt.Printf("%d: %d\n", i, num)
	}

	fmt.Printf("Part 1: value is %d\n", array[indexOf(array, 2017)+1])

	pos = 0
	desired := 0

	for i := 1; i < 50000001; i++ {
		pos = ((pos + input) % i) + 1

		if pos == 1 {
			desired = i
		}
	}

	fmt.Printf("Part 2: value is %d\n", desired)
}
