package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

type handStrength int

var (
	inputFlag string

	// type
	fiveok handStrength = 6
	fok    handStrength = 5
	fh     handStrength = 4
	tok    handStrength = 3
	tp     handStrength = 2
	op     handStrength = 1
	hc     handStrength = 0

	cardStr map[rune]int = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
)

type hand struct {
	holding  string
	wager    int
	strength handStrength
	relStr   float64
}

func mustAtoi(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func gradeHand(h *hand) {
	cardCnt := map[string]int{
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"T": 0,
		"J": 0,
		"Q": 0,
		"K": 0,
		"A": 0,
	}

	for k, _ := range cardCnt {
		cardCnt[k] = strings.Count(h.holding, k)
	}

	numPair, numThree, numFour, numFive := 0, 0, 0, 0
	for _, v := range cardCnt {
		if v == 2 {
			numPair += 1
		} else if v == 3 {
			numThree += 1
		} else if v == 4 {
			numFour += 1
		} else if v == 5 {
			numFive += 1
		}
	}

	h.strength = hc
	if numFive > 0 {
		h.strength = fiveok
	} else if numFour > 0 {
		h.strength = fok
	} else if numThree > 0 && numPair > 0 {
		h.strength = fh
	} else if numThree > 0 {
		h.strength = tok
	} else if numPair > 0 {
		h.strength = op
		if numPair == 2 {
			h.strength = tp
		}
	}

	for idx, d := range h.holding {
		h.relStr += float64(6*cardStr[d]) / (math.Pow(float64(100), float64(idx+1)))
	}

}

func parseInput(input []string) []hand {
	hands := make([]hand, 0)

	for _, line := range input {
		split1 := strings.Split(line, " ")
		hand := hand{split1[0], mustAtoi(split1[1]), 0, 0}
		gradeHand(&hand)
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength < hands[j].strength {
			return true
		}
		if hands[i].strength == hands[j].strength {
			return hands[i].relStr < hands[j].relStr
		}
		return false
	})
	return hands
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	hands := parseInput(inputs)
	score := int64(0)

	for idx, h := range hands {
		score += int64((idx + 1) * h.wager)
	}

	fmt.Printf("Result for infile Part2 %s : %d\n", inputFlag, score)
}
