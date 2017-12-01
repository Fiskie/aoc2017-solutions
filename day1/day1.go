package main

import (
	"io/ioutil"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day1_input.txt")
	check(err)

	part1(dat)
	part2(dat)
}

func part1(dat []byte) {
	size := len(dat)
	sum := 0

	for i := 0; i < size; i++ {
		if dat[(i+1) % size] == dat[i] {
			sum += int(dat[i] - '0')
		}
	}

	fmt.Printf("Part 1: Output is %d\n", sum)
}

func part2(dat []byte) {
	size := len(dat)

	sum := 0

	for i := 0; i < size; i++ {
		if dat[(i + (size / 2)) % size] == dat[i] {
			sum += int(dat[i] - '0')
		}
	}

	fmt.Printf("Part 2: Output is %d\n", sum)
}
