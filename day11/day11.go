package main

import (
	"io/ioutil"
	"strings"
	"aoc2017"
	"fmt"
)

type vector struct {
	x int
	y int
}

type distResult struct {
	shortestPath int
	furthestDistance int
}

var distMap = map[string]vector{
	"nw": {-1, 0},
	"n":  {0, -1},
	"ne": {1, -1},
	"se": {1, 0},
	"s":  {0, 1},
	"sw": {-1, 1},
}

func evaluate(directions []string) distResult {
	dist := vector{0, 0}
	furthestDistance := 0

	for _, direction := range directions {
		dist.x += distMap[direction].x
		dist.y += distMap[direction].y
		shortest := aoc2017.Max(aoc2017.Abs(dist.x), aoc2017.Abs(dist.y))

		if shortest > furthestDistance {
			furthestDistance = shortest
		}
	}

	return distResult{
		aoc2017.Max(aoc2017.Abs(dist.x), aoc2017.Abs(dist.y)),
		furthestDistance,
	}
}

func main() {
	dat, _ := ioutil.ReadFile("./day11_input.txt")
	directions := strings.Split(string(dat), ",")

	fmt.Printf("Ex 3, got %d\n", evaluate([]string{"ne", "ne", "ne"}).shortestPath)
	fmt.Printf("Ex 0, got %d\n", evaluate([]string{"ne", "ne", "sw", "sw"}).shortestPath)
	fmt.Printf("Ex 2, got %d\n", evaluate([]string{"ne", "ne", "s", "s"}).shortestPath)
	fmt.Printf("Ex 3, got %d\n", evaluate([]string{"se", "sw", "se", "sw", "sw"}).shortestPath)

	result := evaluate(directions)

	fmt.Printf("Part 1: Shortest path is %d\n", result.shortestPath)
	fmt.Printf("Part 2: Furthest distance is %d\n", result.furthestDistance)
}
