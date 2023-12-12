package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func indexOf(vals []int, val int) int {
	// find first index of val in slice, return -1 if it doesn't exist
	for i, e := range vals {
		if e == val {
			return i
		}
	}

	return -1
}

func remove(vals []int, val int) []int {
	// remove element from a slice
	idx := indexOf(vals, val)
	return append(vals[:idx], vals[idx+1:]...)
}

func findExpansions(galaxyImg [][]rune) ([]int, []int) {
	// find what rows and columns need to be expanded
	rowsToExpand := []int{}
	colsToExpand := []int{}

	// start with everything in colsToExpand, and remove items as I find #
	for i := 0; i < len(galaxyImg[0]); i++ {
		colsToExpand = append(colsToExpand, i)
	}

	for i := 0; i < len(galaxyImg); i++ {
		expandRow := true
		for j := 0; j < len(galaxyImg[i]); j++ {
			if galaxyImg[i][j] == '#' {
				expandRow = false

				// if index hasn't already been removed
				if indexOf(colsToExpand, j) != -1 {
					colsToExpand = remove(colsToExpand, j)
				}
			}
		}

		if expandRow {
			rowsToExpand = append(rowsToExpand, i)
		}
	}

	return rowsToExpand, colsToExpand
}

func parseFile(filename string) [][]rune {
	// read the input file and expand the galaxy
	// any rows or columns that only contain '.' should be twice as big
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	galaxyImg := [][]rune{}

	for scanner.Scan() {
		galaxyImg = append(galaxyImg, []rune(scanner.Text()))
	}

	return galaxyImg
}

func findGalaxyCoords(galaxyImg [][]rune) [][]int {
	// array of [row, column] pairs telling coordinates of galaxies
	coords := [][]int{}
	for i := 0; i < len(galaxyImg); i++ {
		for j := 0; j < len(galaxyImg[i]); j++ {
			if galaxyImg[i][j] == '#' {
				coords = append(coords, []int{i, j})
			}
		}
	}

	return coords
}

func findManhattanDistance(c1 []int, c2 []int) int {
	return int(math.Abs(float64(c1[0]-c2[0])) + math.Abs(float64(c1[1]-c2[1])))
}

func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}

	return x2
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}

	return x2
}

func main() {
	galaxyImg := parseFile("input.txt")
	rowsToExpand, colsToExpand := findExpansions(galaxyImg)
	coords := findGalaxyCoords(galaxyImg)
	distanceSum := 0

	for i, c := range coords {
		// only need to compare with IDs greater than current ID
		for j := i + 1; j < len(coords); j++ {
			distance := findManhattanDistance(c, coords[j])

			startRow := min(c[0], coords[j][0])
			endRow := max(c[0], coords[j][0])
			startCol := min(c[1], coords[j][1])
			endCol := max(c[1], coords[j][1])

			// add 1 to distance for each expanded row and column crossed
			for k := startRow; k < endRow; k++ {
				if indexOf(rowsToExpand, k) != -1 {
					distance++
				}
			}

			for k := startCol; k < endCol; k++ {
				if indexOf(colsToExpand, k) != -1 {
					distance++
				}
			}

			distanceSum += distance
		}
	}

	fmt.Println(distanceSum)
}
