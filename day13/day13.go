package main

// todo: cleanup
// reduce the need for so many global states

import (
	"bufio"
	"os"
	"aoc2017"
	"strings"
	"fmt"
)
var layers map[int]int
var myDepth = -1
var time = 0
var totalSeverity = 0
var totalDepth = 0

func parse(str string) {
	ints := aoc2017.StringsToInts(strings.Split(str, ": "))
	layers[ints[0]] = ints[1]
	totalDepth = ints[0]
}

func step() bool {
	caught := false
	myDepth += 1
	layer := layers[myDepth]

	if layer != 0 && time % (layer + layer - 2) == 0 {
		totalSeverity += myDepth * layer
		caught = true
	}

	time += 1
	return caught
}

func main() {
	dat, _ := os.Open("./day13_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	layers = map[int]int{}

	for scanner.Scan() {
		parse(scanner.Text())
	}

	for myDepth <= totalDepth {
		step()
	}

	fmt.Printf("Part 1: severity is %d\n", totalSeverity)

	totalDelay := 0
	caught := true

	for caught {
		totalDelay += 1
		myDepth = -1
		totalSeverity = 0
		time = totalDelay
		caught = false

		for myDepth <= totalDepth {
			if step() {
				caught = true
			}
		}
	}

	fmt.Printf("Part 2: fewest picoseconds is %d\n", totalDelay)
}
