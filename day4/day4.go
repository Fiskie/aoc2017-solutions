package main

// TODO: dry
// golang has first class functions? use them

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isValidPart1(passphrase string) (bool) {
	words := strings.Split(passphrase, " ")
	unique := map[string]bool{}

	for _, word := range words {
		if unique[word] != true {
			unique[word] = true
		}
	}

	return len(words) == len(unique)
}

func sortWord(word string) (string) {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isValidPart2(passphrase string) (bool) {
	words := strings.Split(passphrase, " ")
	unique := map[string]bool{}

	for _, word := range words {
		word = sortWord(word)

		if unique[word] != true {
			unique[word] = true
		}
	}

	return len(words) == len(unique)
}

func part1() {
	dat, err := os.Open("./day4_input.txt")
	check(err)

	counter := 0
	reader := bufio.NewReader(dat)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if isValidPart1(scanner.Text()) {
			counter += 1
		}
	}

	fmt.Printf("Part 1: %d passphrases are valid.\n", counter)
}

func part2() {
	dat, err := os.Open("./day4_input.txt")
	check(err)

	counter := 0
	reader := bufio.NewReader(dat)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if isValidPart2(scanner.Text()) {
			counter += 1
		}
	}

	fmt.Printf("Part 2: %d passphrases are valid.\n", counter)
}

func main() {
	part1()
	part2()
}
