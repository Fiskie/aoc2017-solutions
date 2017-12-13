package main

// todo: cleanup
// use Chinese Remainer Theorem

import (
	"bufio"
	"os"
	"aoc2017"
	"strings"
	"fmt"
)

var layers map[int]int

func parse(str string) {
	ints := aoc2017.StringsToInts(strings.Split(str, ": "))
	layers[ints[0]] = ints[1]
}

func isCaught(time int, layer int) bool {
	return layer != 0 && time % (layer + layer - 2) == 0
}

func main() {
	dat, _ := os.Open("./day13_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	layers = map[int]int{}

	for scanner.Scan() {
		parse(scanner.Text())
	}

	totalSeverity := 0

	for i, layer := range layers {
		if isCaught(i, layer) {
			totalSeverity += i * layer
		}
	}

	fmt.Printf("Part 1: severity is %d\n", totalSeverity)

	totalDelay := 0
	caught := true

	for caught {
		caught = false
		totalDelay += 1

		for i, layer := range layers {
			if isCaught(totalDelay + i, layer) {
				caught = true
			}
		}
	}

	fmt.Printf("Part 2: fewest picoseconds is %d\n", totalDelay)
}
