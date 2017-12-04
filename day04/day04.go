package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

type uniqueifier func(string) string

func validator(passphrase string, uniqueFunc uniqueifier) (bool) {
	words := strings.Split(passphrase, " ")
	unique := map[string]bool{}

	for _, word := range words {
		if uniqueFunc != nil {
			word = uniqueFunc(word)
		}

		if unique[word] != true {
			unique[word] = true
		}
	}

	return len(words) == len(unique)
}

func anagramUnique(word string) (string) {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func validate(uniqueFunc uniqueifier) (int) {
	dat, _ := os.Open("./day04_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	counter := 0

	for scanner.Scan() {
		if validator(scanner.Text(), uniqueFunc) {
			counter += 1
		}
	}

	return counter
}

func main() {
	fmt.Printf("Part 1: %d passphrases are valid.\n", validate(nil))
	fmt.Printf("Part 2: %d passphrases are valid.\n", validate(anagramUnique))
}
