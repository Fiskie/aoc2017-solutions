package main

import (
	"os"
	"bufio"
	"fmt"
)

type coordinates struct {
	x int
	y int
}

var grid map[coordinates]byte
var pos coordinates
var vel coordinates
var w int
var h int
var letters []byte

func main() {
	dat, _ := os.Open("./day19_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	grid = map[coordinates]byte{}
	letters = []byte{}

	y := 0

	for scanner.Scan() {
		for x, char := range scanner.Text() {
			grid[coordinates{x, y}] = byte(char)
		}

		y += 1
	}

	h = y - 1
	w = len(grid) / h

	pos = findStartingPosition()
	vel = coordinates{0, 1}
	steps := 0

	for grid[pos] != ' ' {
		walk()
		steps += 1
	}

	fmt.Printf("Steps: %d\n", steps)
	fmt.Printf("Letters: %s\n", string(letters))
}

func walk() {
	if grid[pos] == '+' {
		if vel.y != 0 {
			// Travelling vertically, look horizontally
			vel.y = 0
			if grid[coordinates{pos.x + 1, pos.y}] == ' ' {
				vel.x = -1
			} else {
				vel.x = 1
			}
		} else {
			// Travelling horizontally, look vertically
			vel.x = 0
			if grid[coordinates{pos.x, pos.y + 1}] == ' ' {
				vel.y = -1
			} else {
				vel.y = 1
			}
		}
	} else if grid[pos] != '-' && grid[pos] != '|' {
		letters = append(letters, grid[pos])
	}

	pos.x += vel.x
	pos.y += vel.y
}

func findStartingPosition() coordinates {
	for i := 0; i < w; i++ {
		coords := coordinates{i, 0}

		if grid[coords] == '|' {
			return coords
		}
	}

	return coordinates{0, 0}
}
