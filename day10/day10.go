package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
	"encoding/hex"
)

func stringsToInts(strings []string) []byte {
	var out []byte

	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		out = append(out, byte(num))
	}

	return out
}

func getLengths() []byte {
	dat, _ := ioutil.ReadFile("./day10_input.txt")
	return stringsToInts(strings.Split(string(dat), ","))
}

func getList(length int) []int {
	list := make([]int, length)
	for i := range list {
		list[i] = i
	}
	return list
}

func subReverse(src []int, start int, length int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	size := len(dst)

	for i := 0; i < length; i++ {
		srcIndex := (start + i) % size
		dstIndex := (start + length - 1 - i) % size
		dst[dstIndex] = src[srcIndex]
	}

	return dst
}

func part1() {
	result, _, _ := twist(getList(256), getLengths(), 0, 0)

	fmt.Printf(
		"Part 1: The answer is %d x %d = %d\n",
		result[0],
		result[1],
		int(result[0]) * int(result[1]))
}

func part2() {
	lengths, _ := ioutil.ReadFile("./day10_input.txt")
	sparse := getList(256)
	pos := 0
	skipSize := 0

	lengths = append(lengths, []byte{17, 31, 73, 47, 23}...)

	for i := 0; i < 64; i++ {
		sparse, pos, skipSize = twist(sparse, lengths, pos, skipSize)
	}

	dense := makeDense(sparse)

	dst := make([]byte, hex.EncodedLen(len(dense)))
	hex.Encode(dst, dense)
	fmt.Printf("Part 2: The knot hash is %s\n", dst)
}

func makeDense(list []int) []byte {
	size := 16
	chunks := len(list) / size
	out := make([]byte, chunks)

	for i := 0; i < chunks; i++ {
		out[i] = xor(list[i * size:i * size + size], size)
	}

	return out
}

func xor(chunk []int, size int) byte {
	sum := chunk[0]

	for i := 1; i < size; i++ {
		sum ^= chunk[i]
	}

	return byte(sum)
}

func twist(list []int, lengths []byte, pos int, skipSize int) ([]int, int, int) {
	size := len(list)

	for _, length := range lengths {
		list = subReverse(list, pos % size, int(length))
		pos += int(length) + skipSize
		skipSize += 1
	}

	return list, pos, skipSize
}

func main() {
	part1()
	part2()
}