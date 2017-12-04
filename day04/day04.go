package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

type uniqueifier func(string) string

func validator(passphrase string, uniqueifier uniqueifier) (bool) {
	words := strings.Split(passphrase, " ")
	unique := map[string]bool{}

	for _, word := range words {
		word = uniqueifier(word)

		if unique[word] != true {
			unique[word] = true
		}
	}

	return len(words) == len(unique)
}

func wordUnique(word string) (string) {
	return word
}

func anagramUnique(word string) (string) {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func validate(uniqueifier uniqueifier) (int) {
	dat, _ := os.Open("./day04_input.txt")

	counter := 0
	reader := bufio.NewReader(dat)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if validator(scanner.Text(), uniqueifier) {
			counter += 1
		}
	}

	return counter
}

func main() {
	fmt.Printf("Part 1: %d passphrases are valid.\n", validate(wordUnique))
	fmt.Printf("Part 2: %d passphrases are valid.\n", validate(anagramUnique))
}
