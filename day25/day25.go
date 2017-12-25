package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"fmt"
)

var beginRe = regexp.MustCompile("Begin in state (\\w).")
var diagnosticRe = regexp.MustCompile("Perform a diagnostic checksum after (\\d+) steps.")
var stateBlockRe = regexp.MustCompile("In state (\\w):")
var conditionalRe = regexp.MustCompile("If the current value is (\\w):")
var writeRe = regexp.MustCompile("Write the value (\\w).")
var moveRe = regexp.MustCompile("Move one slot to the (\\w*).")
var continueRe = regexp.MustCompile("Continue with state (\\w).")

var states = map[string]map[int]*stateActions{}
var machine = machineT{}

type machineT struct {
	state string
	cursor int
	tape map[int]int
	steps int
}

type stateActions struct {
	write int
	move string
	cont string
}

func step() {
	actions := states[machine.state][machine.tape[machine.cursor]]
	machine.tape[machine.cursor] = actions.write

	if actions.move == "right" {
		machine.cursor += 1
	} else {
		machine.cursor -= 1
	}

	machine.state = actions.cont
}

func main() {
	dat, _ := os.Open("day25_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))

	machine.state = "A"
	totalCycles := 0

	currentState := "A"
	currentConditional := 0

	for scanner.Scan() {
		text := scanner.Text()

		begin := beginRe.FindStringSubmatch(text)

		if len(begin) > 1 {
			machine.state = begin[1]
		}

		diagnostic := diagnosticRe.FindStringSubmatch(text)

		if len(diagnostic) > 1 {
			totalCycles, _ = strconv.Atoi(diagnostic[1])
		}

		stateBlock := stateBlockRe.FindStringSubmatch(text)

		if len(stateBlock) > 1 {
			currentState = stateBlock[1]
			states[currentState] = map[int]*stateActions{}
		}

		conditional := conditionalRe.FindStringSubmatch(text)

		if len(conditional) > 1 {
			currentConditional, _ = strconv.Atoi(conditional[1])
		}

		write := writeRe.FindStringSubmatch(text)

		if len(write) > 1 {
			states[currentState][currentConditional] = &stateActions{}
			states[currentState][currentConditional].write, _ = strconv.Atoi(write[1])
		}

		move := moveRe.FindStringSubmatch(text)

		if len(move) > 1 {
			states[currentState][currentConditional].move = move[1]
		}

		cont := continueRe.FindStringSubmatch(text)

		if len(cont) > 1 {
			states[currentState][currentConditional].cont = cont[1]
		}
	}

	fmt.Printf("Start state: %s, Cycles: %d\n", machine.state, totalCycles)
	machine.tape = map[int]int{}

	for i := 0; i < totalCycles; i++ {
		step()
	}

	sum := 0

	for _, val := range machine.tape {
		sum += val
	}

	fmt.Printf("Diagnostic checksum: %d\n", sum)
}