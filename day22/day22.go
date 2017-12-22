package main

import (
	"os"
	"bufio"
	"fmt"
)

// todo: DRY
// pass a routine function to do()

type vector struct {
	x int
	y int
}

func getMap() map[vector]byte {
	dat, _ := os.Open("./day22_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))

	grid := map[vector]byte{}

	y := 0

	for scanner.Scan() {
		for x, state := range scanner.Text() {
			grid[vector{x,y}] = byte(state)
		}

		y += 1
	}

	return grid
}

func main() {
	fmt.Printf("Part 1: %d new infections.\n", do(10000))
	fmt.Printf("Part 2: %d new infections.\n", do2(10000000))
}

func do(turns int) int {
	grid := getMap()

	vel := vector{0,-1}
	pos := vector{12, 12}

	infections := 0

	for i := 0; i < turns; i++ {
		if grid[pos] == '#' {
			vel.x, vel.y = -vel.y, vel.x
		} else {
			vel.x, vel.y = vel.y, -vel.x
		}

		if grid[pos] == '#' {
			grid[pos] = '.'
		} else {
			infections += 1
			grid[pos] = '#'
		}

		pos.x += vel.x
		pos.y += vel.y
	}

	return infections
}

func do2(turns int) int {
	grid := getMap()

	vel := vector{0,-1}
	pos := vector{12, 12}

	infections := 0

	for i := 0; i < turns; i++ {
		if grid[pos] == '#' {
			vel.x, vel.y = -vel.y, vel.x
		} else if grid[pos] == 'W' {
			// Does not turn
		} else if grid[pos] == 'F' {
			vel.x, vel.y = -vel.x, -vel.y
		} else {
			vel.x, vel.y = vel.y, -vel.x
		}

		if grid[pos] == '#' {
			grid[pos] = 'F'
		} else if grid[pos] == 'W' {
			grid[pos] = '#'
			infections += 1
		} else if grid[pos] == 'F' {
			grid[pos] = '.'
		} else {
			grid[pos] = 'W'
		}

		pos.x += vel.x
		pos.y += vel.y
	}

	return infections
}
