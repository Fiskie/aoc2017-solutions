package main

import (
	"fmt"
)

type vector struct {
	x int
	y int
}

func sumAdjacents(x int, y int, grid map[vector]int) int {
	return grid[vector{x - 1, y - 1}] +
		grid[vector{x - 1, y}] +
		grid[vector{x - 1, y + 1}] +
		grid[vector{x, y - 1}] +
		grid[vector{x, y}] +
		grid[vector{x, y + 1}] +
		grid[vector{x + 1, y - 1}] +
		grid[vector{x + 1, y}] +
		grid[vector{x + 1, y + 1}]
}

func main() {
	// grid will be a one-dimensional map keyed by x,y as a string
	grid := map[vector]int{}
	grid[vector{0, 0}] = 1

	input := 361527
	lastAdjacentSum := 0

	x := 0
	y := 0
	length := 1
	progress := 0
	offset := 1
	isX := true

	for lastAdjacentSum < input {
		if isX {
			x += offset
		} else {
			y += offset
		}

		progress += 1

		if progress == length {
			progress = 0
			isX = !isX

			if isX {
				offset = -offset
				length += 1
			}
		}

		lastAdjacentSum = sumAdjacents(x, y, grid)
		fmt.Printf("%d, %d: %d\n", x, y, lastAdjacentSum)
		grid[vector{x, y}] = lastAdjacentSum
	}

	fmt.Printf("Part 2: First larger value is %d\n", lastAdjacentSum)
}
