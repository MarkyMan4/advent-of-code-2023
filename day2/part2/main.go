package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type draw map[string]int

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

		// put defaults as 0 if color wasn't present
		for _, color := range []string{"red", "green", "blue"} {
			if _, ok := counts[color]; !ok {
				counts[color] = 0
			}
		}

		gameDraws = append(gameDraws, counts)
	}

	return gameDraws
}

func calcCubePower(draws []draw) int {
	minCounts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, d := range draws {
		for _, color := range []string{"red", "green", "blue"} {
			if d[color] > minCounts[color] {
				minCounts[color] = d[color]
			}
		}
	}

	return minCounts["red"] * minCounts["green"] * minCounts["blue"]
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
		gameData := parseGameData(text)
		sum += calcCubePower(gameData)
	}

	fmt.Println(sum)
}
