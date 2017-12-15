package main

import "fmt"

func generator(c chan int, factor int, value int, multiple int) {
	for {
		value = value * factor % 2147483647
		if multiple == 0 || value % multiple == 0 {
			c <- value
		}
	}
}

func main() {
	fmt.Printf("Part 1: %d\n", part1(40000000,0, 0))
	fmt.Printf("Part 2: %d\n", part1(5000000,4, 8))
}

func part1(iterations int, aMult int, bMult int) int {
	a := make(chan int)
	b := make(chan int)
	count := 0

	go generator(a,16807, 873, aMult)
	go generator(b,48271, 583, bMult)

	for i := 0; i < iterations; i++ {
		if <-a &^ 0xFFFF0000 == <-b &^ 0xFFFF0000 {
			count += 1
		}
	}

	return count
}