package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag     string
	searchStrings map[string]string = map[string]string{
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
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func setVals(first *string, last *string, new string) {
	if *first == "" {
		*first = new
	}
	*last = new
}

func getCalibrationValue(line string) int {
	// Go back to single pass
	// In this case, if it isn't a digit, lets check for word in our lookup table
	// This is a little tricky -- but our words are only 3 4 or 5 characters, so lets try that.
	first := ""
	last := ""
	lastLineIdx := len(line)
	for idx, r := range line {
		if unicode.IsDigit(r) {
			setVals(&first, &last, string(r))
		} else {
			for _, offset := range []int{3, 4, 5} {
				if val, ok := searchStrings[line[idx:min(idx+offset, lastLineIdx)]]; ok {
					setVals(&first, &last, val)
					continue
				}
			}
		}
	}

	res, err := strconv.Atoi(first + last)
	if err != nil {
		log.Fatalf("error converting %s : %v", line, err)
	}
	return res
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "input", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)

	total := int64(0)
	for _, line := range inputs {
		total += int64(getCalibrationValue(line))
	}

	fmt.Printf("Result for infile %s : %d\n", inputFlag, total)
}
