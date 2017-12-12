package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"aoc2017"
	"strings"
	"fmt"
)

var re = regexp.MustCompile("^(\\d*) <-> (.*)$")
var pipes map[int][]int

func parse(text string) {
	groups := re.FindStringSubmatch(text)
	id, _ := strconv.Atoi(groups[1])
	pipes[id] = aoc2017.StringsToInts(strings.Split(groups[2], ", "))
}

func findNeighbours(id int, seen map[int]bool) map[int]bool {
	seen[id] = true

	for _, neighbour := range pipes[id] {
		if !seen[neighbour] {
			findNeighbours(neighbour, seen)
		}
	}

	return seen
}

func getGroupMembers(id int) map[int]bool {
	return findNeighbours(id, map[int]bool{})
}

func main() {
	dat, _ := os.Open("./day12_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	pipes = map[int][]int{}

	for scanner.Scan() {
		parse(scanner.Text())
	}

	fmt.Printf("Part 1: The number of programs connected to PID 0 is %d\n", len(getGroupMembers(0)))

	// Keeping a map of encountered PIDs reduces time complexity to n sets
	encountered := map[int]bool{}
	memberLists := map[int]map[int]bool{}

	for pipeId := range pipes {
		if !encountered[pipeId] {
			memberLists[pipeId] = getGroupMembers(pipeId)

			for members := range memberLists[pipeId] {
				encountered[members] = true
			}
		}
	}

	fmt.Printf("Part 2: The number of unique groups is %d\n", len(memberLists))
}
