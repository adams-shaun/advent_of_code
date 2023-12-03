package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	inputFlag string
)

func readInput(infile string) []string {
	data, err := os.ReadFile(infile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	return lines
}

func getCalibrationValue(line string) int {
	// Thinking here we make a single pass.
	// Set both front/back tracker vars to first digit we see
	// Otherwise, just update back on any new digit.
	// Iterating from both sides could yield some improvement.
	first := ""
	last := ""

	for _, r := range line {
		if unicode.IsDigit(r) {
			if first == "" {
				first = string(r)
			}
			last = string(r)
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
	inputs := readInput(inputFlag)

	total := int64(0)
	for _, line := range inputs {
		total += int64(getCalibrationValue(line))
	}

	fmt.Printf("Result for infile %s : %d\n", inputFlag, total)
}
