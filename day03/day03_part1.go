package main

import (
	"fmt"
)

// golang why
func abs(input int) (int) {
	if input < 0 {
		return -input
	}

	return input
}

func getOffsetAndDepthOfSpiralValue(input int) (int, int) {
	// The calculated offset from being in a straight line from the centre.
	offset := 0
	// One halfCycle = one half-trip around the spiral. 2 * halfCycles == depth
	halfCycles := 0

	if input > 1 {
		halfCycles = 1
		corner := halfCycles * halfCycles

		for corner < input {
			halfCycles += 2
			corner = halfCycles * halfCycles
		}

		// This one-liner:
		// - Calculates the 'length' of the spiral section this input lies on.
		// - Calculates the input's position before direction is due to change.
		// - Calculates the maximum possible offset for manhattan distance at this depth.
		// - Uses these to calculate the offset.
		// ex. given a length of 4 and offset of 2, length-offset will be between -2 and 2.
		// we then use the absolute value of the number.
		offset = abs(((corner - input) % (halfCycles - 1)) - ((halfCycles - 1) / 2))
	}

	return offset, halfCycles / 2
}

func spiralPosToManhattan(input int) (int) {
	offset, depth := getOffsetAndDepthOfSpiralValue(input)
	return offset + depth
}

func main() {
	fmt.Printf("Test input 1: expected 0, given %d\n", spiralPosToManhattan(1))
	fmt.Printf("Test input 12: expected 3, given %d\n", spiralPosToManhattan(12))
	fmt.Printf("Test input 23: expected 2, given %d\n", spiralPosToManhattan(23))
	fmt.Printf("Test input 1024: expected 31, given %d\n", spiralPosToManhattan(1024))
	fmt.Printf("Part 1: distance is %d\n", spiralPosToManhattan(361527))
}
