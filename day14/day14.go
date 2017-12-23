package main

import (
	"strconv"
	"aoc2017"
	"fmt"
	"encoding/hex"
)

func hexToBin(hex string) string {
	ui, _ := strconv.ParseUint(hex, 16, 64)
	return fmt.Sprintf("%016b", ui)
}

type vector struct {
	x int
	y int
}

func getAdjacents(grid map[vector]bool, coords vector) map[vector]bool {
	list := map[vector]bool{}

	neighbours := []vector{
		{coords.x, coords.y + 1},
		{coords.x, coords.y - 1},
		{coords.x + 1, coords.y},
		{coords.x - 1, coords.y},
	}

	for _, neighbour := range neighbours {
		if grid[neighbour] {
			list[neighbour] = true
		}
	}

	return list
}

func getGroup(grid map[vector]bool, coords vector, seen map[vector]bool) map[vector]bool {
	seen[coords] = true

	for neighbour, exists := range getAdjacents(grid, coords) {
		if exists && !seen[neighbour] {
			getGroup(grid, neighbour, seen)
		}
	}

	return seen
}

func main() {
	input := "hxtvlmkl"
	grid := make(map[vector]bool, 128*128)

	for i := 0; i < 128; i++ {
		lengths := []byte(input + "-" + strconv.FormatInt(int64(i), 10))
		hash := aoc2017.Knot(lengths)
		dst := make([]byte, hex.EncodedLen(len(hash)))
		hex.Encode(dst, hash)
		hexString := string(dst)

		for word := 0; word < 8; word++ {
			for offset, bit := range hexToBin(hexString[word*4:word*4+4]) {
				num, _ := strconv.Atoi(string(bit))
				grid[vector{i, (word * 16) + offset}] = num == 1
			}
		}
	}

	count := 0

	// Keeping a map of encountered vectors reduces time complexity to n sets
	encountered := map[vector]bool{}
	unique := map[vector]map[vector]bool{}

	for coords, value := range grid {
		if value {
			count += 1

			if !encountered[coords] {
				unique[coords] = getGroup(grid, coords, map[vector]bool{})

				for coord := range unique[coords] {
					encountered[coord] = true
				}
			}
		}
	}

	fmt.Printf("Part 1: %d squares are used\n", count)
	fmt.Printf("Part 2: The number of unique groups is %d\n", len(unique))
}
