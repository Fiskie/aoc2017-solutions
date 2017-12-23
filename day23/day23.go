package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

type instruction struct {
	op  string
	dst string
	src string
}

func valueOf(registers map[string]int, val string) int {
	num, err := strconv.Atoi(val)

	if err != nil {
		return registers[val]
	}

	return num
}

func run(instructions []instruction, registers map[string]int) int {
	mulCalls := 0

	for i := 0; i >= 0 && i < len(instructions); i++ {
		inst := instructions[i]

		switch inst.op {
		case "set":
			registers[inst.dst] = valueOf(registers, inst.src)
		case "sub":
			registers[inst.dst] -= valueOf(registers, inst.src)
		case "mul":
			registers[inst.dst] *= valueOf(registers, inst.src)
			mulCalls += 1
		case "jnz":
			if valueOf(registers, inst.dst) != 0 {
				i += valueOf(registers, inst.src) - 1
			}
		}
	}

	return mulCalls
}

func getInstructions(filename string) []instruction {
	dat, _ := os.Open(filename)
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	var instructions []instruction

	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		inst := instruction{op: args[0], dst: args[1]}

		if len(args) == 3 {
			inst.src = args[2]
		}

		instructions = append(instructions, inst)
	}

	return instructions
}

func part1() {
	insts := getInstructions("./day23_input.txt")
	fmt.Printf("Part 1: %d mul ops invoked\n", run(insts, map[string]int{}))
}

func main() {
	part1()
	part2()
}

func part2() {
	iterations := 0
	notPrimes := 0
	start := (65 * 100) + 100000
	end := start + 17000

	for ; start <= end; start += 17 {
		isPrime := true // 9

		i := 5
		w := 2

		if start % 2 == 0 || start % 3 == 0 {
			isPrime = false
		} else {
			for i * i <= start {
				if start % i == 0 {
					isPrime = false
				}

				i += w
				w = 6 - w
			}
		}

		if !isPrime { // 25
			notPrimes++ // 26
		}

		iterations++

		fmt.Printf("start: %7d, end: %7d, notPrimes: %7d, iterations: %d\n", start, end, notPrimes, iterations)
	}

	fmt.Printf("Part 2 (specific to my input, sorry): h = %d\n", notPrimes)
}
