package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(text string) []int {
	vals := strings.Split(text, " ")
	intVals := []int{}
	for _, val := range vals {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		intVals = append(intVals, intVal)
	}

	return intVals
}

func isAllZeros(vals []int) bool {
	if len(vals) == 0 {
		return false
	}
	for _, val := range vals {
		if val != 0 {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	predictionTotal := 0

	for scanner.Scan() {
		readings := parseLine(scanner.Text())
		firstNumsInSeqs := []int{readings[0]}
		prevDiffs := readings

		for {
			diffs := []int{}
			for i := 0; i < len(prevDiffs)-1; i++ {
				diffs = append(diffs, prevDiffs[i+1]-prevDiffs[i])
			}

			if isAllZeros(diffs) {
				// no need to store the slice of all zeros
				break
			}

			prevDiffs = diffs

			// only need to store the last value in each sequence since that's the only one that's
			// needed to make the prediction
			firstNumsInSeqs = append(firstNumsInSeqs, diffs[0])
		}

		prediction := 0
		for i := len(firstNumsInSeqs) - 1; i >= 0; i-- {
			prediction = firstNumsInSeqs[i] - prediction
		}

		predictionTotal += prediction
	}

	fmt.Println(predictionTotal)
}
