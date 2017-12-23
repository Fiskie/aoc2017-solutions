package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"fmt"
)

var re = regexp.MustCompile("^(\\w*) (inc|dec) (\\S*) if (\\w*) (\\S*) (\\S*)$")
var registers = map[string]int{}
var highestValue = 0

var comparators = map[string]func(int, int) (bool){
	">":  func(a int, b int) bool { return a > b },
	"<":  func(a int, b int) bool { return a < b },
	">=": func(a int, b int) bool { return a >= b },
	"<=": func(a int, b int) bool { return a <= b },
	"==": func(a int, b int) bool { return a == b },
	"!=": func(a int, b int) bool { return a != b },
}

func eval(str string) {
	matches := re.FindStringSubmatch(str)

	integer1, _ := registers[matches[4]]
	integer2, _ := strconv.Atoi(matches[6])

	if comparators[matches[5]](integer1, integer2) {
		registerArg, _ := strconv.Atoi(matches[3])

		if matches[2] == "dec" {
			registerArg = 0 - registerArg
		}

		registers[matches[1]] += registerArg

		if registers[matches[1]] > highestValue {
			highestValue = registers[matches[1]]
		}
	}
}

func getHighestRegister() string {
	bestReg := ""

	for register, value := range registers {
		if value > registers[bestReg] {
			bestReg = register
		}
	}

	return bestReg
}

func main() {
	dat, _ := os.Open("./day08_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))

	for scanner.Scan() {
		eval(scanner.Text())
	}

	highest := getHighestRegister()

	fmt.Printf("Part 1: Highest register is %s (%d)\n", highest, registers[highest])
	fmt.Printf("Part 2: Highest recorded value was %d\n", highestValue)
}
