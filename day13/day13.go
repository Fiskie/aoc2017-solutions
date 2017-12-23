package main

import (
	"bufio"
	"os"
	"aoc2017"
	"strings"
	"fmt"
)

func isCaught(time int, layer int) bool {
	return time%(layer+layer-2) == 0
}

func main() {
	dat, _ := os.Open("./day13_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	layers := map[int]int{}

	for scanner.Scan() {
		ints := aoc2017.StringsToInts(strings.Split(scanner.Text(), ": "))
		layers[ints[0]] = ints[1]
	}

	totalSeverity := 0

	for i, layer := range layers {
		if isCaught(i, layer) {
			totalSeverity += i * layer
		}
	}

	fmt.Printf("Part 1: severity is %d\n", totalSeverity)

	offset := 0
	goodRun := false

	for !goodRun {
		goodRun = true
		offset += 1

		for i, layer := range layers {
			if isCaught(offset+i, layer) {
				goodRun = false
				break
			}
		}
	}

	fmt.Printf("Part 2: fewest picoseconds is %d\n", offset)
}
