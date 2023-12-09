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

const (
	start = "AAA"
	goal  = "ZZZ"
)

var alphaRegex = regexp.MustCompile("[A-Z]+")

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
		parts := alphaRegex.FindAllString(text, -1)
		network[parts[0]] = []string{parts[1], parts[2]}
	}

	return instructions, network
}

func main() {
	instructions, network := parseFile("input.txt")
	stepCount := 0
	curInstruction := 0
	curNode := start
	for curNode != "ZZZ" {
		stepCount++
		inst := instructions[curInstruction]
		curNode = network[curNode][inst]
		curInstruction++

		// if reached the last instruction, start back at beginning
		if curInstruction >= len(instructions) {
			curInstruction = 0
		}
	}

	fmt.Println(stepCount)
}
