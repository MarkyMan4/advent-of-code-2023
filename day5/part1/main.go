package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile("[0-9]+")

type LookupTable struct {
	destRangeStart int
	srcRangeStart  int
	rangeLength    int
}

type Converter [][]LookupTable

func (c Converter) getLocationForSeed(seed int) int {
	result := seed
	for _, convMap := range c {
		for _, table := range convMap {
			if result >= table.srcRangeStart && result < table.srcRangeStart+table.rangeLength {
				diff := result - table.srcRangeStart
				result = table.destRangeStart + diff
				break
			}
		}
	}

	return result
}

func strSliceToIntSlice(vals []string) []int {
	intSlice := []int{}

	for _, val := range vals {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		intSlice = append(intSlice, intVal)
	}

	return intSlice
}

func parseMapFile(filename string) ([]int, Converter) {
	mapFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(mapFile)
	scanner.Scan()
	seedLine := scanner.Text()
	seeds := strSliceToIntSlice(numberRegex.FindAllString(seedLine, -1))
	converter := Converter{}
	conversionMap := []LookupTable{}

	// scan past first blank line and first map title
	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			// blank like means next line is a title, scan past it
			// add conversionMap to conv
			scanner.Scan()
			converter = append(converter, conversionMap)
			conversionMap = []LookupTable{}
		} else {
			values := strSliceToIntSlice(numberRegex.FindAllString(text, -1))
			lookup := LookupTable{values[0], values[1], values[2]}
			conversionMap = append(conversionMap, lookup)
		}
	}

	// append final map
	converter = append(converter, conversionMap)

	return seeds, converter
}

func main() {
	// goal is to find loweset location number that any seed maps to
	seeds, converter := parseMapFile("input.txt")
	var lowestLoc int
	for i, seed := range seeds {
		loc := converter.getLocationForSeed(seed)
		if loc < lowestLoc || i == 0 {
			lowestLoc = loc
		}
	}

	fmt.Println(lowestLoc)
}
