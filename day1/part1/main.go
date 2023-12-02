// 55108

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

func sumSlice(vals []int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}

	return sum
}

// remove any non-numeric characters from the text
func prepText(text string) string {
	r := regexp.MustCompile("[a-zA-Z]")
	result := r.ReplaceAll([]byte(text), []byte{})

	return string(result)
}

func main() {
	calibrationFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer calibrationFile.Close()
	scanner := bufio.NewScanner(calibrationFile)
	numbers := []int{}

	for scanner.Scan() {
		text := scanner.Text()
		text = prepText(text)
		chars := strings.Split(text, "")
		val, _ := strconv.Atoi(chars[0] + chars[len(chars)-1])
		numbers = append(numbers, val)
	}

	fmt.Println(sumSlice(numbers))
}
