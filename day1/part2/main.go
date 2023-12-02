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

var textDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func sumSlice(vals []int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}

	return sum
}

// remove any non-numeric characters from the text
// replace spelled out digits with actual digits
// only the first character of a substring is replaced with the actual digit to
// handle the case where digits overlap, e.g. twone becomes 2w1ne, and the final
// text returned is 21
func prepText(text string) string {
	// build a map where keys are indices of spelled out digits and the value is the spelled out digit
	// i.e. a key in textDigits
	stringsToReplace := map[int]string{}
	for k := range textDigits {
		// find all occurences of a text string
		r := regexp.MustCompile(k)
		locs := r.FindAllStringIndex(text, -1)

		// if substring was found, add all start indices of it to the map
		if locs != nil {
			for _, loc := range locs {
				startIdx := loc[0]
				stringsToReplace[startIdx] = k
			}
		}
	}

	// sort the indices so we can replace them in order (left to right)
	charList := strings.Split(text, "")
	for k, v := range stringsToReplace {
		charList[k] = textDigits[v]
	}

	newText := strings.Join(charList, "")
	r := regexp.MustCompile("[a-zA-Z]")
	result := r.ReplaceAll([]byte(newText), []byte{})

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
