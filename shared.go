package aoc2017

import "strconv"

func StringsToInts(strings []string) []int {
	out := make([]int, len(strings))

	for i, str := range strings {
		num, _ := strconv.Atoi(str)
		out[i] = num
	}

	return out
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}

	return input
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}