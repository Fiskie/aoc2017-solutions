package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
	"aoc2017"
	"encoding/hex"
)

func stringsToBytes(strings []string) []byte {
	out := make([]byte, len(strings))

	for i, str := range strings {
		num, _ := strconv.Atoi(str)
		out[i] = byte(num)
	}

	return out
}

func getLengths() []byte {
	dat, _ := ioutil.ReadFile("./day10_input.txt")
	return stringsToBytes(strings.Split(string(dat), ","))
}

func part1() {
	result, _, _ := aoc2017.Twist(aoc2017.GetRange(256), getLengths(), 0, 0)

	fmt.Printf(
		"Part 1: The answer is %d x %d = %d\n",
		result[0],
		result[1],
		int(result[0]) * int(result[1]))
}

func part2() {
	lengths, _ := ioutil.ReadFile("./day10_input.txt")

	hash := aoc2017.Knot(lengths)

	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash)
	fmt.Printf("Part 2: The knot hash is %s\n", dst)
}

func main() {
	part1()
	part2()
}