package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag string
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func cardResult(input string) (int, int) {
	split1 := strings.Split(input, ": ")
	split2 := strings.Split(split1[1], "| ")
	winnerStrings := strings.Split(split2[0], " ")
	myNumStrings := strings.Split(split2[1], " ")

	winnerMap := map[int]bool{}
	for _, val := range winnerStrings {
		num, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		winnerMap[num] = true
	}

	matches := 0
	for _, val := range myNumStrings {

		num, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		if _, ok := winnerMap[num]; ok {
			matches += 1
		}
	}

	if matches == 0 {
		return 0, 0
	}
	return 1 << (matches - 1), matches
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)

	totalCards := make([]int, len(inputs))
	for idx, _ := range totalCards {
		totalCards[idx] = 1
	}

	totalWins := int64(0)

	for idx, line := range inputs {
		score, matches := cardResult(line)
		totalWins += int64(score)

		for y := idx + 1; y <= min(idx+matches, len(inputs)-1); y++ {
			totalCards[y] += totalCards[idx]
		}
	}

	totalCardSum := int64(0)
	for _, val := range totalCards {
		totalCardSum += int64(val)
	}
	fmt.Printf("Result for infile %s : %d (score) %d (cards)\n", inputFlag, totalWins, totalCardSum)
}
