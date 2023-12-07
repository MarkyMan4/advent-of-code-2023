package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func parseFile(filename string) []raceRecord {
	records := []raceRecord{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	times := strSliceToIntSlice(numberRegex.FindAllString(text, -1))

	scanner.Scan()
	text = scanner.Text()
	distances := strSliceToIntSlice(numberRegex.FindAllString(text, -1))

	for i := 0; i < len(times); i++ {
		records = append(records, raceRecord{times[i], distances[i]})
	}

	return records
}

/*
to calculate how far boat travels based on how long button was held:

d = distance
b = length button held
t = race time

d = (t - b) * b

*/
func main() {
	records := parseFile("input.txt")

	product := 1
	for _, rec := range records {
		waysToWin := 0
		for i := 1; i < rec.timeToRace; i++ {
			distance := (rec.timeToRace - i) * i
			if distance > rec.recordDistance {
				waysToWin++
			}
		}

		product *= waysToWin
	}

	fmt.Println(product)
}
