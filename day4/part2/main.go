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

// parses card text and returns card number, winning numbers, my numbers
func parseCard(card string) (int, []int, []int) {
	// remove the card # and split lists on |
	cardParts := strings.Split(card, ":")
	cardNum, err := strconv.Atoi(numberRegex.FindString(cardParts[0]))
	if err != nil {
		panic(err)
	}

	numberLists := strings.Split(cardParts[1], "|")
	winningNums := numberRegex.FindAllString(numberLists[0], -1)
	myNums := numberRegex.FindAllString(numberLists[1], -1)

	return cardNum, strToIntSlice(winningNums), strToIntSlice(myNums)
}

func sumMapVals(vals map[int]int) int {
	total := 0
	for _, v := range vals {
		total += v
	}

	return total
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
	cardCopies := map[int]int{} // map from

	for scanner.Scan() {
		card := scanner.Text()
		cardNum, winningNums, myNums := parseCard(card)

		// if cardNum is not in map, set it to 1 since we have at least 1 copy of each card
		if _, ok := cardCopies[cardNum]; !ok {
			cardCopies[cardNum] = 1
		}

		copies := cardCopies[cardNum]
		winningNumberCount := 0 // number of winning numbers in this card
		for _, num := range winningNums {
			if sliceContainsVal(myNums, num) {
				winningNumberCount++
			}
		}

		// incremement the number of copies for the next <winningNumberCount> cards
		// number of copies to be added is equal to the number of copies of the current card
		for i := cardNum + 1; i <= cardNum+winningNumberCount; i++ {
			// if the card number doesn't exist yet, set it to the number of copies we're adding
			// plus 1 to account for our default 1 card
			if _, ok := cardCopies[i]; !ok {
				cardCopies[i] = copies + 1
			} else {
				cardCopies[i] += copies
			}
		}
	}

	totalCards := sumMapVals(cardCopies)
	fmt.Println(totalCards)
}
