package main

import (
	"os"
	"bufio"
	"strings"
	"aoc2017"
	"fmt"
)

type component struct {
	a int
	b int
}

type bridge []component
type componentSet map[component]bool

var components = componentSet{}
var bridges = []bridge{}

func printBridge(b bridge) {
	fmt.Printf("****\n")

	for i, comp := range b {
		fmt.Printf("%d: %d\n", i, comp)
	}
}

func main() {
	dat, _ := os.Open("day24_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))

	for scanner.Scan() {
		args := aoc2017.StringsToInts(strings.Split(scanner.Text(), "/"))

		components[component{a: args[0], b: args[1]}] = true
		components[component{a: args[1], b: args[0]}] = true
	}

	fmt.Printf("Components: %d\n", len(components))

	for comp := range getComponentsWithPort(0, components) {
		createBridges(bridge{comp}, copyOf(components))
	}

	strongest := 0
	longest := bridge{}

	// TODO: currently using a different set, answers should be 2006 / 1994
	// still trying to figure out hashmap bullshit :(

	for _, b := range bridges {
		s := strengthOf(b)

		fmt.Printf("Bridge str %d len %d\n", s, len(b))

		if s > strongest {
			strongest = s
		}

		if len(b) >= len(longest) {
			if s > strengthOf(longest) || len(b) > len(longest) {
				longest = b
			}
		}
	}

	fmt.Printf("Strongest: %d\n", strongest)
	fmt.Printf("Longest: %d\n", strengthOf(longest))
}

func createBridges(b bridge, available componentSet) {
	head := b[len(b)-1]
	delete(available, head)
	delete(available, component{head.b, head.a})

	next := getComponentsWithPort(head.b, available)

	if len(next) == 0 {
		bridges = append(bridges, b)
	}

	for item := range next {
		createBridges(append(b, item), copyOf(available))
	}
}

func getComponentsWithPort(port int, set componentSet) componentSet {
	filtered := componentSet{}

	for comp := range set {
		if comp.a == port {
			filtered[comp] = true
		}
	}

	return filtered
}

func copyOf(set componentSet) componentSet {
	cloned := componentSet{}

	for comp := range set {
		cloned[comp] = true
	}

	return cloned
}

func strengthOf(b bridge) int {
	sum := 0

	for _, comp := range b {
		sum += comp.a + comp.b
	}

	return sum
}