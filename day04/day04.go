package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

type validator func(string) bool

func part1Validator(passphrase string) (bool) {
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

func part2Validator(passphrase string) (bool) {
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

func validate(validatorFunc validator) (int) {
	dat, _ := os.Open("./day04_input.txt")

	counter := 0
	reader := bufio.NewReader(dat)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if validatorFunc(scanner.Text()) {
			counter += 1
		}
	}

	return counter
}

func main() {
	fmt.Printf("Part 1: %d passphrases are valid.\n", validate(part1Validator))
	fmt.Printf("Part 2: %d passphrases are valid.\n", validate(part2Validator))
}
