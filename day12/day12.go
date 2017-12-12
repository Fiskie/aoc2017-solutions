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
	numbers := aoc2017.StringsToInts(strings.Split(groups[2], ", "))
	pipes[id] = numbers
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

	// Part 2: this could be optimised greatly by
	// skipping group members we encounter, but oh well
	memberLists := map[int]map[int]bool{}
	unique := map[int]bool{}

	for pipeId := range pipes {
		memberLists[pipeId] = getGroupMembers(pipeId)
	}

	for _, memberList := range memberLists {
		highestId := 0

		for memberId := range memberList {
			if memberId > highestId {
				highestId = memberId
			}
		}

		if !unique[highestId] {
			unique[highestId] = true
		}
	}

	fmt.Printf("Part 2: The number of unique groups is %d\n", len(unique))
}
