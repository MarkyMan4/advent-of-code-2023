package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int

const (
	north = iota
	south
	east
	west
)

const (
	pipeEntry = 'S'
	vertPipe  = '|'
	horizPipe = '-'
	neBend    = 'L'
	nwBend    = 'J'
	swBend    = '7'
	seBend    = 'F'
)

// map to lookup a pipe and the direction we came from
// the value found is the transformation to apply to current x, y
// coords to get our next location and the next direction we came from
var pipeFlow = map[rune]map[direction][]int{
	vertPipe: {
		north: {0, 1, north},
		south: {0, -1, south},
	},
	horizPipe: {
		east: {-1, 0, east},
		west: {1, 0, west},
	},
	neBend: {
		north: {1, 0, west},
		east:  {0, -1, south},
	},
	nwBend: {
		north: {-1, 0, east},
		west:  {0, -1, south},
	},
	swBend: {
		south: {-1, 0, east},
		west:  {0, 1, north},
	},
	seBend: {
		south: {1, 0, west},
		east:  {0, 1, north},
	},
}

// pipes that can be accessed when coming from the south, north, east or west
var fromNorthPipes = []rune{vertPipe, nwBend, neBend}
var fromSouthPipes = []rune{vertPipe, swBend, seBend}
var fromEastPipes = []rune{horizPipe, neBend, seBend}
var fromWestPipes = []rune{horizPipe, nwBend, swBend}

type point struct {
	x int // character position in line
	y int // line number
}

func sliceContains(vals []rune, searchVal rune) bool {
	for _, val := range vals {
		if val == searchVal {
			return true
		}
	}

	return false
}

func findStartingPointIndex(vals []rune) int {
	for i := 0; i < len(vals); i++ {
		if vals[i] == pipeEntry {
			return i
		}
	}

	// -1 if 'S' not found
	return -1
}

func parseFile(filename string) (point, [][]rune) {
	startingPoint := point{}
	pipeMap := [][]rune{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	curLine := 0

	for scanner.Scan() {
		line := []rune(scanner.Text())
		if startIdx := findStartingPointIndex(line); startIdx != -1 {
			startingPoint.x = startIdx
			startingPoint.y = curLine
		}
		pipeMap = append(pipeMap, line)

		curLine++
	}

	return startingPoint, pipeMap
}

// find any pipe connected to the start point as the first move
func findFirstMove(startPoint point, pipeMap [][]rune) (direction, point) {
	var fromDir direction
	var startingPipe point

	if startPoint.y > 0 && sliceContains(fromSouthPipes, pipeMap[startPoint.y-1][startPoint.x]) {
		fromDir = south
		startingPipe = point{startPoint.x, startPoint.y - 1}
	} else if startPoint.y < len(pipeMap)-1 && sliceContains(fromNorthPipes, pipeMap[startPoint.y+1][startPoint.x]) {
		fromDir = north
		startingPipe = point{startPoint.x, startPoint.y + 1}
	} else if startPoint.x > 0 && sliceContains(fromWestPipes, pipeMap[startPoint.y][startPoint.x-1]) {
		fromDir = west
		startingPipe = point{startPoint.x - 1, startPoint.y}
	} else if startPoint.x < len(pipeMap[0])-1 && sliceContains(fromEastPipes, pipeMap[startPoint.y][startPoint.x+1]) {
		fromDir = east
		startingPipe = point{startPoint.x + 1, startPoint.y}
	}

	return fromDir, startingPipe
}

// function that takes the direction we came from, the current point and a pipe part
// and returns the next coordinate
func nextPipe(fromDir direction, curPoint point, pipePart rune) (direction, point) {
	transform := pipeFlow[pipePart][fromDir]
	xTransform := transform[0]
	yTransform := transform[1]
	newFromDir := direction(transform[2])

	return newFromDir, point{curPoint.x + xTransform, curPoint.y + yTransform}
}

func main() {
	startingPoint, pipeMap := parseFile("input.txt")
	fromDir, curPipe := findFirstMove(startingPoint, pipeMap)
	curPipePart := pipeMap[curPipe.y][curPipe.x]
	stepsTaken := 1

	// keep following pipes until we end up back at the start
	for curPipePart != pipeEntry {
		fromDir, curPipe = nextPipe(fromDir, curPipe, curPipePart)
		curPipePart = pipeMap[curPipe.y][curPipe.x]
		stepsTaken++
	}

	maxStepsFromStart := stepsTaken / 2
	fmt.Println(maxStepsFromStart)
}
