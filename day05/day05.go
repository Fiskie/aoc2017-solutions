package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)

func incrementalJump(value int) (int) {
	return value + 1
}

func threeSoftCapJump(value int) (int) {
	if value >= 3 {
		return value - 1
	}

	return value + 1
}

func getJumpCount(jumps []int, jumpFunc func(int) int) (int) {
	// clone the list because it is modified in-place
	jumps = append([]int(nil), jumps...)

	index := 0
	size := len(jumps)
	counter := 0

	for {
		if index < 0 || index >= size {
			return counter
		}

		jump := jumps[index]
		jumps[index] = jumpFunc(jumps[index])
		index += jump
		counter++
	}
}

func main() {
	dat, _ := os.Open("./day05_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	var jumps []int

	for scanner.Scan() {
		jump, _ := strconv.Atoi(scanner.Text())
		jumps = append(jumps, jump)
	}

	fmt.Printf("Part 1 sanity test: jump count is %d\n", getJumpCount([]int{0, 3, 0, 1, -3}, incrementalJump))
	fmt.Printf("Part 1: jump count is %d\n", getJumpCount(jumps, incrementalJump))
	fmt.Printf("Part 2 sanity test: jump count is %d\n", getJumpCount([]int{0, 3, 0, 1, -3}, threeSoftCapJump))
	fmt.Printf("Part 2: jump count is %d\n", getJumpCount(jumps, threeSoftCapJump))
}
