package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func redistribute(bank []int) ([]int) {
	// clone the list because we need to record what we've seen
	bank = append([]int(nil), bank...)

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

	return bank
}

func stringsToInts(strings []string) ([]int) {
	var out []int

	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		out = append(out, num)
	}

	return out
}

func equals(a []int, b []int) (bool) {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func bankSeen(test []int, banks [][]int) (bool) {
	for _, bank := range banks {
		if equals(test, bank) {
			return true
		}
	}

	return false
}

func loadBank() ([]int) {
	dat, _ := ioutil.ReadFile("./day06_input.txt")
	return stringsToInts(strings.Split(string(dat), "	"))
}

func main() {
	var banks [][]int
	bank := loadBank()

	for !bankSeen(bank, banks) {
		banks = append(banks, bank)
		bank = redistribute(bank)
	}

	fmt.Printf("Part 1: %d redistribution cycles were performed\n", len(banks))

	banks = [][]int{}

	for !bankSeen(bank, banks) {
		banks = append(banks, bank)
		bank = redistribute(bank)
	}

	fmt.Printf("Part 2: %d redistribution cycles were performed\n", len(banks))
}
