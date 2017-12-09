package main

import (
	"io/ioutil"
	"fmt"
)

func process(input []byte) (int, int) {
	depth := 0
	score := 0
	eaten := 0
	isIgnoring := false
	isEating := false

	for _, char := range input {
		if isIgnoring {
			isIgnoring = false
		} else if char == '!' {
			isIgnoring = true
		} else if char == '>' {
			isEating = false
		} else if isEating {
			eaten += 1
		} else if char == '<' {
			isEating = true
		} else if char == '{' {
			depth++
		} else if char == '}' {
			score += depth
			depth--
		}
	}

	return score, eaten
}

func main() {
	dat, _ := ioutil.ReadFile("./day09_input.txt")
	score, eaten := process(dat)
	fmt.Printf("Part 1: score is %d.\n", score)
	fmt.Printf("Part 2: %d characters were eaten.\n", eaten)
}
