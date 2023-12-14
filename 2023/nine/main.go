package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

type handStrength int

var (
	inputFlag string
	re_number string = "-?\\d*\\.{0,1}\\d+"
)

func mustAtoi(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func parseInput(input []string) [][]int {
	res := [][]int{}

	r, _ := regexp.Compile(re_number)

	for _, line := range input {
		newRes := []int{}
		for _, res := range r.FindAllString(line, -1) {
			newRes = append(newRes, mustAtoi(res))
		}
		res = append(res, newRes)
	}
	return res
}

func extrapolate(input []int) (int, int) {
	extr := [][]int{input}
	currIdx := 0
	for {
		newExtr := []int{}
		zeroCnt := 0
		for i := 0; i < len(extr[currIdx])-1; i++ {
			newVal := extr[currIdx][i+1] - extr[currIdx][i]
			newExtr = append(newExtr, newVal)

			if newVal == 0 {
				zeroCnt += 1
			}
		}
		extr = append(extr, newExtr)
		if zeroCnt == len(extr[currIdx])-1 {
			break
		}
		currIdx += 1
	}

	// extrapolate forward
	extr[len(extr)-1] = append(extr[len(extr)-1], 0)
	for j := len(extr) - 2; j >= 0; j-- {
		newVal := extr[j][len(extr[j])-1] + extr[j+1][len(extr[j])-1]
		extr[j] = append(extr[j], newVal)
	}
	forwardRes := extr[0][len(extr[0])-1]

	// extrapolate backwards
	extr[len(extr)-1] = append([]int{0}, extr[len(extr)-1]...)
	for j := len(extr) - 2; j >= 0; j-- {
		newVal := extr[j][0] - extr[j+1][0]
		extr[j] = append([]int{newVal}, extr[j]...)
	}
	backwardRes := extr[0][0]

	return forwardRes, backwardRes
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	sensors := parseInput(inputs)

	countFwd := 0
	countBack := 0

	for _, data := range sensors {
		fwd, back := extrapolate(data)
		countFwd += fwd
		countBack += back
	}

	fmt.Printf("Result for infile %s : %d (fwd) %d (back) \n", inputFlag, countFwd, countBack)
}
