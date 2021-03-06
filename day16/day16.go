package main

import (
	"strings"
	"io/ioutil"
	"fmt"
	"strconv"
)

func indexOf(chars []byte, char byte) int {
	for i, other := range chars {
		if other == char {
			return i
		}
	}

	return -1
}

func contains(array []string, str string) bool {
	for _, other := range array {
		if other == str {
			return true
		}
	}

	return false
}

// ignorePartnerMoves is used in part 2 as they are essentially cancelled
// out by the next iteration, and we're looking for a dance iteration with
// an even number
func perform(programs []byte, move string, ignorePartnerMoves bool) []byte {
	if move[0] == 's' {
		a, _ := strconv.Atoi(string(move[1:]))
		split := len(programs) - a
		programs = append(programs[split:], programs[:split]...)
	} else if move[0] == 'x' {
		inputs := strings.Split(move[1:], "/")
		a, _ := strconv.Atoi(inputs[0])
		b, _ := strconv.Atoi(inputs[1])
		programs[a], programs[b] = programs[b], programs[a]
	} else if !ignorePartnerMoves && move[0] == 'p' {
		a := indexOf(programs, move[1])
		b := indexOf(programs, move[3])
		programs[a], programs[b] = programs[b], programs[a]
	}

	return programs
}

func main() {
	dat, _ := ioutil.ReadFile("./day16_input.txt")
	dance := strings.Split(string(dat), ",")
	programs := []byte("abcdefghijklmnop")

	for _, move := range dance {
		programs = perform(programs, move, false)
	}

	fmt.Printf("Part 1: Final order is %s\n", string(programs))

	programs = []byte("abcdefghijklmnop")
	var permutations []string

	for {
		for _, move := range dance {
			programs = perform(programs, move, true)
		}

		if contains(permutations, string(programs)) {
			break
		}

		permutations = append(permutations, string(programs))
	}

	answer := permutations[(1000000000%len(permutations))-1]
	fmt.Printf("Part 2: Final order is %s\n", answer)
}
