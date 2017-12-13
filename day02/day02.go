package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"aoc2017"
	"math"
)

func main() {
	dat, _ := os.Open("./day02_input.txt")

	checksum := 0
	sumDivisible := 0
	scanner := bufio.NewScanner(bufio.NewReader(dat))

	for scanner.Scan() {
		numberStrings := strings.Split(scanner.Text(), "	")
		numbers := aoc2017.StringsToInts(numberStrings)
		lowest := math.MaxInt32
		highest := 0

		for _, num := range numbers {
			if num < lowest {
				lowest = num
			}

			if num > highest {
				highest = num
			}

			for _, cmpNum := range numbers {
				if num > cmpNum && num%cmpNum == 0 {
					sumDivisible += num / cmpNum
				}
			}
		}

		diff := highest - lowest
		checksum += diff
	}

	fmt.Printf("Part 1: Checksum is %d\n", checksum)
	fmt.Printf("Part 2: Sum is %d\n", sumDivisible)
}
