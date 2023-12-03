package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile("[0-9]+")
var astksRegex = regexp.MustCompile(regexp.QuoteMeta("*"))

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

// multiply a slice of integers together
func multSlice(nums []int) int {
	product := 1

	for _, num := range nums {
		product *= num
	}

	return product
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

func findAdjacentParts(text string, numIndices [][]int, astkStart, astkEnd int) []int {
	partsFound := []int{}

	for _, numIdxs := range numIndices {
		numStart := numIdxs[0]
		numEnd := numIdxs[1] - 1 // end index goes one past end of number, so subtract one from it

		if (numStart >= astkStart && numStart <= astkEnd) ||
			(numEnd >= astkStart && numEnd <= astkEnd) {

			partsFound = append(partsFound, getNumFromPos(text, numIdxs))
		}
	}

	return partsFound
}

/*
 find all asteriks in curLine
 iterate over each asteriks in curLine,
 then go through prevLine, curLine and nextLine, making note of any numbers that are
 adjacent to that asteriks
 if the length of part numbers > 1, multiply them together and add them to the sum
 i.e. any numbers in the text that are adjacent to a symbol

 This finds the sum of gear ratios for one line
*/
func findGearRatios(prevLine, curLine, nextLine string) int {
	// numIndices has start and end indices of all numbers in previous, current and next lines
	// the adjacency check can be done the same way regardless of line, so easier just to have one slice
	prevLineNumIndices := numberRegex.FindAllStringIndex(prevLine, -1)
	curLineNumIndices := numberRegex.FindAllStringIndex(curLine, -1)
	nextLineNumIndices := numberRegex.FindAllStringIndex(nextLine, -1)
	asterikIndices := astksRegex.FindAllStringIndex(curLine, -1)
	gearRatioSum := 0

	// can return 0 right away if the current line doesn't have any gears
	if len(asterikIndices) == 0 {
		return gearRatioSum
	}

	for _, idxs := range asterikIndices {
		astkStart := max(idxs[0]-1, 0)
		astkEnd := min(idxs[1], len(curLine)-1)
		partsFound := []int{} // slice of any parts adjacent

		// iterate over prevLine, curLine and next line number indices and keep track of anything adjacent to the asterik
		partsFound = append(partsFound, findAdjacentParts(prevLine, prevLineNumIndices, astkStart, astkEnd)...)
		partsFound = append(partsFound, findAdjacentParts(curLine, curLineNumIndices, astkStart, astkEnd)...)
		partsFound = append(partsFound, findAdjacentParts(nextLine, nextLineNumIndices, astkStart, astkEnd)...)

		if len(partsFound) > 1 {
			gearRatioSum += multSlice(partsFound)
		}
	}

	return gearRatioSum
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
	ratioSum := 0

	for scanner.Scan() {
		nextLine = scanner.Text()
		ratioSum += findGearRatios(prevLine, curLine, nextLine)
		prevLine = curLine
		curLine = nextLine
	}

	// check the final line since the loop broke before we could check it
	nextLine = ""
	ratioSum += findGearRatios(prevLine, curLine, nextLine)

	fmt.Println(ratioSum)
}
