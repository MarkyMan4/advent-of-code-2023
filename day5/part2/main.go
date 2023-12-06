/*
 time to run
 -----------
 with goroutines: 2m 0s
 without goroutings: 4m 23s
*/

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

func findLowest(vals []int) int {
	var lowest int
	for i, val := range vals {
		if i == 0 || val < lowest {
			lowest = val
		}
	}

	return lowest
}

func main() {
	// goal is to find loweset location number that any seed maps to
	seeds, converter := parseMapFile("input.txt")
	resultChannels := []chan int{}

	for i := 0; i < len(seeds); i += 2 {
		channel := make(chan int)
		resultChannels = append(resultChannels, channel)

		go func(rangeStart, rangeEnd int) {
			var lowestLoc int

			for j := rangeStart; j < rangeEnd; j++ {
				loc := converter.getLocationForSeed(j)
				if loc < lowestLoc || j == rangeStart {
					lowestLoc = loc
				}
			}

			channel <- lowestLoc
		}(seeds[i], seeds[i]+seeds[i+1])
	}

	locsFound := []int{}
	for i := 0; i < len(resultChannels); i++ {
		locsFound = append(locsFound, <-resultChannels[i])
	}

	fmt.Println(findLowest(locsFound))

	// solution without goroutines (twice as slow)
	// initialized := false
	// var lowestLoc int
	// for i := 0; i < len(seeds); i += 2 {
	// 	for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
	// 		loc := converter.getLocationForSeed(j)
	// 		if loc < lowestLoc || !initialized {
	// 			lowestLoc = loc
	// 			initialized = true
	// 		}
	// 	}
	// }

	// fmt.Println(lowestLoc)
}
