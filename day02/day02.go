package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func stringsToInts(strings []string) ([]int) {
	var out []int

	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		out = append(out, num)
	}

	return out
}

func main() {
	dat, _ := os.Open("./day02_input.txt")

	checksum := 0
	sumDivisible := 0

	reader := bufio.NewReader(dat)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		numberStrings := strings.Split(scanner.Text(), "	")
		numbers := stringsToInts(numberStrings)
		lowest := 9999999 // fixme :/
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
