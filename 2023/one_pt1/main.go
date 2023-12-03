package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag string
)

func getCalibrationValue(line string) int {
	// Two pointer solution
	// Assumes input is all UTF-8, avoid conversion to runes
	// Previous 'one pass' solution guaranteed a full pass
	// This approach could save a bit of time.
	first := 0
	last := len(line) - 1

	var firstByte byte
	var lastByte byte

	for first <= last {
		if (firstByte != 0) && (lastByte != 0) {
			break
		}

		// look for first
		if firstByte == 0 {
			if line[first] >= 48 && line[first] <= 57 {
				firstByte = line[first]
			} else {
				first++
			}
		}

		// look for last
		if lastByte == 0 {
			if line[last] >= 48 && line[last] <= 57 {
				lastByte = line[last]
			} else {
				last--
			}
		}
	}

	res, err := strconv.Atoi(string([]byte{firstByte, lastByte}))
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
