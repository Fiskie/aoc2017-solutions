package main

import (
	"os"
	"bufio"
	"regexp"
	"strings"
	"strconv"
	"reflect"
	"fmt"
)

type node struct {
	parent string
	name string
	weight int
	children []string
}

func addNode(str string, nodes map[string]node) {
	re := regexp.MustCompile("^([a-z]*) \\((\\d*)\\)(?: -> (.*))?$")
	matches := re.FindStringSubmatch(str)
	weight, _ := strconv.Atoi(matches[2])
	name := matches[1]
	children := strings.Split(matches[3], ", ")

	node := nodes[name]
	node.name = name
	node.weight = weight
	node.children = children
	nodes[name] = node

	for _, child := range children {
		node := nodes[child]
		node.parent = name
		nodes[child] = node
	}
}

func findRoot(nodes map[string]node) (node) {
	target := reflect.ValueOf(nodes).MapKeys()[0].String()

	for nodes[target].parent != "" {
		target = nodes[target].parent
	}

	return nodes[target]
}

func getStackWeight(node node, nodes map[string]node) (int) {
	sum := node.weight

	for _, child := range node.children {
		sum += getStackWeight(nodes[child], nodes)
	}

	return sum
}

func main() {
	dat, _ := os.Open("./day07_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	nodes := map[string]node{}

	for scanner.Scan() {
		addNode(scanner.Text(), nodes)
	}

	root := findRoot(nodes)
	fmt.Printf("Part 1: The bottom node is %s\n", root.name)

	//for _, child := range root.children {
	//	fmt.Printf("Weight of %s is %d\n", child, getStackWeight(nodes[child], nodes))
	//}
}