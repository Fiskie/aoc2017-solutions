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
	parent   *node
	name     string
	weight   int
	children []*node
}

func addNode(str string, nodes map[string]*node) {
	re := regexp.MustCompile("^([a-z]*) \\((\\d*)\\)(?: -> (.*))?$")
	matches := re.FindStringSubmatch(str)
	weight, _ := strconv.Atoi(matches[2])
	name := matches[1]
	children := strings.Split(matches[3], ", ")

	newNode := nodes[name]

	if newNode == nil {
		newNode = &node{}
	}

	newNode.name = name
	newNode.weight = weight

	for _, child := range children {
		childNode := nodes[child]

		if childNode == nil {
			childNode = &node{}
		}

		childNode.name = child
		childNode.parent = newNode
		nodes[child] = childNode
		newNode.children = append(newNode.children, childNode)
	}

	nodes[name] = newNode
}

func findRoot(nodes map[string]*node) *node {
	target := reflect.ValueOf(nodes).MapKeys()[0].String()

	for nodes[target].parent != nil {
		target = nodes[target].parent.name
	}

	node := nodes[target]
	return node
}

func getStackWeight(node *node) int {
	sum := node.weight

	for _, child := range node.children {
		sum += getStackWeight(child)
	}

	return sum
}

func findUnbalancedNode(root *node) *node {
	weights := map[int][]node{}

	for _, child := range root.children {
		weight := getStackWeight(child)
		weights[weight] = append(weights[weight], *child)
	}

	for _, children := range weights {
		if len(children) == 1 {
			return findUnbalancedNode(&children[0])
		}
	}

	return root
}

func printNodeInfo(node *node) {
	for _, child := range node.children {
		fmt.Printf("%s children: %s (weight %d) (stack weight %d)\n", node.name, child.name, child.weight, getStackWeight(child))
	}

	for _, sibling := range node.parent.children {
		fmt.Printf("%s sibling: %s (weight %d) (stack weight %d)\n", node.name, sibling.name, sibling.weight, getStackWeight(sibling))
	}
}

func main() {
	dat, _ := os.Open("./day07_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	nodes := map[string]*node{}

	for scanner.Scan() {
		addNode(scanner.Text(), nodes)
	}

	root := findRoot(nodes)

	fmt.Printf("Part 1: The bottom node is %s\n", root.name)
	fmt.Printf("Part 2: Make note of the unbalanced weight of the sibling.\n")

	node := findUnbalancedNode(root)
	fmt.Printf("The unbalanced node is %s\n", node.name)

	printNodeInfo(node)
}
