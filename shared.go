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

func Knot(lengths []byte) []byte {
	sparse := GetRange(256)
	pos := 0
	skipSize := 0

	lengths = append(lengths, []byte{17, 31, 73, 47, 23}...)

	for i := 0; i < 64; i++ {
		sparse, pos, skipSize = Twist(sparse, lengths, pos, skipSize)
	}

	return makeDense(sparse)
}

func Twist(list []int, lengths []byte, pos int, skipSize int) ([]int, int, int) {
	size := len(list)

	for _, length := range lengths {
		list = subReverse(list, pos % size, int(length))
		pos += int(length) + skipSize
		skipSize += 1
	}

	return list, pos, skipSize
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

func GetRange(length int) []int {
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

func xor(chunk []int, size int) byte {
	sum := chunk[0]

	for i := 1; i < size; i++ {
		sum ^= chunk[i]
	}

	return byte(sum)
}