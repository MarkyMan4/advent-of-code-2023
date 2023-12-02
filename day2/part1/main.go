// Determine which games would have been possible if the bag had
// been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type draw map[string]int

func getGameNumber(text string) int {
	r := regexp.MustCompile("[0-9]+")
	num := string(r.Find([]byte(text)))
	result, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func parseGameData(text string) []draw {
	gameDraws := []draw{}

	raw := strings.Split(strings.Split(text, ":")[1], ";")
	for _, d := range raw {
		cubes := strings.Split(d, ",")
		counts := make(draw)
		for _, cubeCount := range cubes {
			split := strings.Split(strings.Trim(cubeCount, " "), " ")
			color := split[1]
			count, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}

			counts[color] = count
		}

		gameDraws = append(gameDraws, counts)
	}

	return gameDraws
}

func isGamePossible(draws []draw) bool {
	isPossible := true
	for _, d := range draws {
		if count, ok := d["red"]; ok && count > maxRed {
			isPossible = false
		}
		if count, ok := d["green"]; ok && count > maxGreen {
			isPossible = false
		}
		if count, ok := d["blue"]; ok && count > maxBlue {
			isPossible = false
		}
	}

	return isPossible
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		gameNo := getGameNumber(text)
		gameData := parseGameData(text)
		if isGamePossible(gameData) {
			sum += gameNo
		}
	}

	fmt.Println(sum)
}
