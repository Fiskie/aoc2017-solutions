package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"aoc2017"
)

func redistribute(bank []int) {
	redistributableBlocks := 0
	targetIndex := 0

	for i, blocks := range bank {
		if blocks > redistributableBlocks {
			redistributableBlocks = blocks
			targetIndex = i
		}
	}

	bank[targetIndex] = 0
	size := len(bank)

	for redistributableBlocks > 0 {
		targetIndex += 1
		bank[targetIndex % size] += 1
		redistributableBlocks -= 1
	}
}

func loadBank() []int {
	dat, _ := ioutil.ReadFile("./day06_input.txt")
	return aoc2017.StringsToInts(strings.Split(string(dat), "	"))
}

func key(bank []int) string {
	return fmt.Sprint(bank)
}

func main() {
	bank := loadBank()
	banks := map[string]bool{}

	for !banks[key(bank)] {
		banks[key(bank)] = true
		redistribute(bank)
	}

	fmt.Printf("Part 1: %d redistribution cycles were performed\n", len(banks))

	banks = map[string]bool{}

	for !banks[key(bank)] {
		banks[key(bank)] = true
		redistribute(bank)
	}

	fmt.Printf("Part 2: %d redistribution cycles were performed\n", len(banks))
}
