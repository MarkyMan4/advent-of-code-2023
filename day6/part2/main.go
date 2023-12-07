package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]+")

type raceRecord struct {
	timeToRace     int
	recordDistance int
}

func strSliceToIntSlice(vals []string) []int {
	converted := []int{}
	for _, val := range vals {
		intval, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		converted = append(converted, intval)
	}

	return converted
}

func parseFile(filename string) raceRecord {
	file, fileErr := os.Open(filename)
	if fileErr != nil {
		panic(fileErr)
	}
	scanner := bufio.NewScanner(file)
	var err error
	var time int
	var distance int

	scanner.Scan()
	text := scanner.Text()
	time, err = strconv.Atoi(strings.Join(numberRegex.FindAllString(text, -1), ""))
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	text = scanner.Text()
	distance, err = strconv.Atoi(strings.Join(numberRegex.FindAllString(text, -1), ""))
	if err != nil {
		panic(err)
	}

	return raceRecord{time, distance}
}

/*
to calculate how far boat travels based on how long button was held:

d = distance
b = length button held
t = race time

d = (t - b) * b

*/
func main() {
	record := parseFile("input.txt")
	waysToWin := 0
	for i := 1; i < record.timeToRace; i++ {
		distance := (record.timeToRace - i) * i
		if distance > record.recordDistance {
			waysToWin++
		}
	}

	fmt.Println(waysToWin)
}
