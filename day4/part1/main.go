package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]+")

func sliceContainsVal(nums []int, searchVal int) bool {
	for _, num := range nums {
		if num == searchVal {
			return true
		}
	}

	return false
}

func strToIntSlice(vals []string) []int {
	result := []int{}
	for _, val := range vals {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}

	return result
}

func parseCard(card string) ([]int, []int) {
	// remove the card # and split lists on |
	numberLists := strings.Split(strings.Split(card, ":")[1], "|")
	winningNums := numberRegex.FindAllString(numberLists[0], -1)
	myNums := numberRegex.FindAllString(numberLists[1], -1)

	return strToIntSlice(winningNums), strToIntSlice(myNums)
}

// I get 1 point for having one winning number, 2 points for having two winning numbers, 4 points for 3, 8 for 4
// and so on
// so count the number of winning numbers I have, take 2^(count - 1)
func main() {
	cardsFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(cardsFile)
	pointTotal := 0

	for scanner.Scan() {
		card := scanner.Text()
		winningNums, myNums := parseCard(card)
		winningNumberCount := 0 // number of winning numbers in this card
		for _, num := range winningNums {
			if sliceContainsVal(myNums, num) {
				winningNumberCount++
			}
		}

		if winningNumberCount > 0 {
			pointTotal += int(math.Pow(2, float64(winningNumberCount)-1))
		}
	}

	fmt.Println(pointTotal)
}
