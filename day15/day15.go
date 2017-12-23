package main

import "fmt"

func generator(factor int, value int, multiple int) func() int {
	return func() int {
		for {
			value = value * factor % 2147483647
			if multiple == 0 || value%multiple == 0 {
				return value
			}
		}
	}
}

func main() {
	fmt.Printf("Part 1: %d\n", iterate(40000000, 0, 0))
	fmt.Printf("Part 2: %d\n", iterate(5000000, 4, 8))
}

func iterate(iterations int, aMult int, bMult int) int {
	count := 0

	a := generator(16807, 873, aMult)
	b := generator(48271, 583, bMult)

	for i := 0; i < iterations; i++ {
		if a()&0xFFFF == b()&0xFFFF {
			count += 1
		}
	}

	return count
}
