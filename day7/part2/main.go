package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardStrength = map[rune]int{
	'J': 0, // joker
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

const (
	fiveOfKind = iota
	fourOfKind
	fullHouse
	threeOfKind
	twoPair
	onePair
	highCard
)

type hand struct {
	handType int
	cards    []rune
	bid      int
}

func NewHand(cards []rune, bid int) hand {
	// determine the hand type
	cardCount := map[rune]int{} // map to keep track of counts of each card
	jokerCount := 0
	for _, card := range cards {
		if card == 'J' {
			jokerCount++
			continue
		}
		if _, ok := cardCount[card]; !ok {
			cardCount[card] = 1
		} else {
			cardCount[card]++
		}
	}

	// if we have jokers, use them to make the best hand
	if jokerCount > 0 {
		// use the joker to increment the count of the highest card that we have the most of
		initialized := false
		var maxCard rune = 'J' // initialize to joker in case we have 5 jokers, it becomes 5 of a kind
		for card := range cardCount {
			// on first iteration, set max card to whatever we see first
			if !initialized {
				maxCard = card
				initialized = true
			} else if (cardStrength[card] > cardStrength[maxCard] && cardCount[card] >= cardCount[maxCard]) ||
				(cardCount[card] > cardCount[maxCard]) {
				// if we find a stronger card with a count greater or equal to previous max
				// use this as new max card
				// this also handles the case where we only have high card, it will find the highest card
				maxCard = card
			}
		}

		cardCount[maxCard] += jokerCount
	}

	handType := 0

	/*
		 determine hand by number of distinct cards
		 need to handle the following cases:
			2 distinct cards - could be 4 of a kind or full house
			3 distinct cards - could be 2 or 3 of a kind
	*/
	switch len(cardCount) {
	case 1:
		handType = fiveOfKind
	case 2:
		for _, count := range cardCount {
			if count == 1 || count == 4 {
				handType = fourOfKind
				break
			} else {
				handType = fullHouse
				break
			}
		}
	case 3:
		for _, count := range cardCount {
			if count == 3 {
				handType = threeOfKind
				break
			} else if count == 2 {
				handType = twoPair
				break
			}
			// count of 1 doesn't tell us definitely what the type is, so keep checking
		}
	case 4:
		handType = onePair
	case 5:
		handType = highCard
	}

	return hand{handType, cards, bid}
}

func parseFile(filename string) []hand {
	hands := []hand{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		cards := []rune(split[0])
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		hands = append(hands, NewHand(cards, bid))
	}

	return hands
}

func main() {
	hands := parseFile("input.txt")

	// sort hands from weakest to strongest
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for k := 0; k < len(hands[i].cards); k++ {
				if cardStrength[hands[i].cards[k]] < cardStrength[hands[j].cards[k]] {
					return true
				} else if cardStrength[hands[i].cards[k]] > cardStrength[hands[j].cards[k]] {
					return false
				}
			}
		}

		return hands[i].handType > hands[j].handType
	})

	totalWinnings := 0
	for i, hand := range hands {
		rank := i + 1
		totalWinnings += (rank * hand.bid)
	}

	fmt.Println(totalWinnings)
}
