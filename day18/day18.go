package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"sync"
)

type instruction struct {
	op  string
	dst string
	src string
}

type program struct {
	id        int
	last      int
	registers map[string]int
	sent      int
}

var instructions []instruction

func valueOf(p *program, val string) int {
	num, err := strconv.Atoi(val)

	if err != nil {
		return p.registers[val]
	}

	return num
}

func run(p *program, in, out chan int, wg *sync.WaitGroup) {
	for i := 0; i >= 0 && i < len(instructions); i++ {
		inst := instructions[i]

		switch inst.op {
		case "set":
			p.registers[inst.dst] = valueOf(p, inst.src)
		case "snd":
			out <- valueOf(p, inst.dst)
			p.last = valueOf(p, inst.dst)
			p.sent += 1

			if p.id == 1 {
				fmt.Printf("Part 2: PID %d send count: %d\n", p.id, p.sent)
			}
		case "add":
			p.registers[inst.dst] += valueOf(p, inst.src)
		case "mul":
			p.registers[inst.dst] *= valueOf(p, inst.src)
		case "mod":
			p.registers[inst.dst] %= valueOf(p, inst.src)
		case "rcv":
			if p.id == -1 {
				fmt.Printf("Part 1: last 'sound' value on first nonzero rcv call is %d\n", p.last)
				i = 9999
			} else {
				p.registers[inst.dst] = <-in
			}
		case "jgz":
			if valueOf(p, inst.dst) > 0 {
				i += valueOf(p, inst.src) - 1
			}
		}
	}
	wg.Done()
}

func part1() {
	p := program{id: -1, registers: map[string]int{}}
	buf := make(chan int, 500)
	var wg sync.WaitGroup
	wg.Add(1)
	run(&p, buf, buf, &wg)
}

func part2() {
	p0 := program{id: 0, registers: map[string]int{"p": 0}}
	p1 := program{id: 1, registers: map[string]int{"p": 1}}

	ch0 := make(chan int, 500)
	ch1 := make(chan int, 500)

	var wg sync.WaitGroup
	wg.Add(2)
	go run(&p0, ch0, ch1, &wg)
	go run(&p1, ch1, ch0, &wg)
	wg.Wait()
}

func main() {
	dat, _ := os.Open("./day18_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	instructions = []instruction{}

	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		inst := instruction{op:  args[0], dst: args[1]}

		if len(args) == 3 {
			inst.src = args[2]
		}

		instructions = append(instructions, inst)
	}

	part1()
	part2()
}
