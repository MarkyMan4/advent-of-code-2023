/*
idea
----
prevLine = nil
cursor = (0, 0)
curLine = scanner.Text()

while scanner.Scan() {
	nextLine = scanner.Text()
}


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

func isNumeric(text string) bool {
	if _, err := strconv.Atoi(text); err != nil {
		return false
	}

	return true
}

func charAt(text string, index int) string {
	return string([]rune(text)[index])
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

// anything besides a "." and a digit is a symbol
func isSymbolAtIndex(text string, index int) bool {
	char := charAt(text, index)
	return char != "." && !isNumeric(char)
}

// check text range based on idxs for symbols
// subtracts one from start and adds one to end (if possible) to check diaganolly as well
func rangeHasSymbols(text string, idxs []int) bool {
	iterStart := max(idxs[0]-1, 0)
	iterEnd := min(idxs[1]+1, len(text))

	for i := iterStart; i < iterEnd; i++ {
		if isSymbolAtIndex(text, i) {
			return true
		}
	}

	return false
}

func getNumFromPos(text string, idxs []int) int {
	numText := ""
	for i := idxs[0]; i < idxs[1]; i++ {
		numText += charAt(text, i)
	}

	num, err := strconv.Atoi(numText)
	if err != nil {
		panic(err)
	}

	return num
}

// sum part numbers in curLine
// i.e. any numbers in the text that are adjacent to a symbol
func findPartNumbers(prevLine, curLine, nextLine string) int {
	numberIndices := numberRegex.FindAllStringIndex(curLine, -1)
	partNumSum := 0

	// each idx is a slice of length 2, start and end positions of regex match
	for _, idxs := range numberIndices {
		if (idxs[0] > 0 && isSymbolAtIndex(curLine, idxs[0]-1)) || // check left of number
			(idxs[1] < len(curLine) && isSymbolAtIndex(curLine, idxs[1])) || // check right of number
			(prevLine != "" && rangeHasSymbols(prevLine, idxs)) || // check prev line and diaganols above
			(nextLine != "" && rangeHasSymbols(nextLine, idxs)) { // check next line and diaganols below

			partNumSum += getNumFromPos(curLine, idxs)
		}
	}

	return partNumSum
}

func main() {
	var prevLine string
	var curLine string
	var nextLine string

	partsFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(partsFile)
	scanner.Scan()

	prevLine = ""
	curLine = scanner.Text()
	partNumSum := 0

	for scanner.Scan() {
		nextLine = scanner.Text()
		partNumSum += findPartNumbers(prevLine, curLine, nextLine)
		prevLine = curLine
		curLine = nextLine
	}

	// check the final line since the loop broke before we could check it
	nextLine = ""
	partNumSum += findPartNumbers(prevLine, curLine, nextLine)

	fmt.Println(partNumSum)
}
