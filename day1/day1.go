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

	fmt.Printf("Part 1: Output is %d\n", getSum(dat, 1))
	fmt.Printf("Part 2: Output is %d\n", getSum(dat, len(dat) / 2))
}

func getSum(dat []byte, offset int) (int) {
	sum := 0
	size := len(dat)

	for i := 0; i < size; i++ {
		if dat[(i + offset) % size] == dat[i] {
			sum += int(dat[i] - '0')
		}
	}

	return sum
}