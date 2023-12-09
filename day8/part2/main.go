package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	left = iota
	right
)

var alphaNumRegex = regexp.MustCompile("[0-9A-Z]+")

func parseFile(filename string) ([]int, map[string][]string) {
	instructions := []int{}
	network := map[string][]string{}

	mapFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(mapFile)

	// parse the instruction line
	scanner.Scan()
	instLine := []rune(scanner.Text())
	for _, inst := range instLine {
		if inst == 'R' {
			instructions = append(instructions, right)
		} else {
			instructions = append(instructions, left)
		}
	}
	scanner.Scan() // scan past blank line

	// parse the network
	for scanner.Scan() {
		text := scanner.Text()
		parts := alphaNumRegex.FindAllString(text, -1)
		network[parts[0]] = []string{parts[1], parts[2]}
	}

	return instructions, network
}

func findStartingNodes(network map[string][]string) []string {
	// find all nodes ending with "A"
	nodes := []string{}
	for node := range network {
		if []rune(node)[2] == 'A' {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

// greatest common divisor algorithm
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// recursively find the least common multiple of two or more numbers
// https://en.wikipedia.org/wiki/Least_common_multiple
func lcm(vals []int) int {
	if len(vals) == 2 {
		a, b := vals[0], vals[1]
		return (a * b) / gcd(a, b)
	}

	return lcm([]int{vals[0], lcm(vals[1:])})
}

func main() {
	instructions, network := parseFile("input.txt")
	curNodes := findStartingNodes(network)
	curInstruction := 0
	stepsPerNode := []int{}

	// find how many steps it takes for each start node to each an end node
	// then find the least common multiple of those values
	for _, node := range curNodes {
		stepCount := 0
		curNode := node
		for []rune(curNode)[2] != 'Z' {
			stepCount++
			inst := instructions[curInstruction]
			curNode = network[curNode][inst]

			curInstruction++

			// if reached the last instruction, start back at beginning
			if curInstruction >= len(instructions) {
				curInstruction = 0
			}
		}

		stepsPerNode = append(stepsPerNode, stepCount)
	}

	fmt.Println(lcm(stepsPerNode))
}
