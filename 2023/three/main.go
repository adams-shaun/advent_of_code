package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag string
	re_number string = "(\\d+)"
	re_symbol string = "[^\\s\\d.]"
	re_gear   string = "\\*"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func searchLine(input string, offset int) (bool, []int) {
	r2, _ := regexp.Compile(re_symbol)
	r3, _ := regexp.Compile(re_gear)
	foundSymbol := false
	if r2.MatchString(input) {
		foundSymbol = true
	}

	gearIdx := []int{}
	for _, res := range r3.FindAllStringIndex(input, -1) {
		gearIdx = append(gearIdx, res[0]+offset)
	}
	return foundSymbol, gearIdx
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)

	r, _ := regexp.Compile(re_number)

	total := int64(0)
	gearTotal := int64(0)

	// Keep track of gears found, add numbers associated with them
	// In this case the key = line#*(lineLength-1) + gear_x_idx
	gears := map[int][]int{}

	processResult := func(symPresent bool, gearIdx []int, idx int, line string, indices []int) {
		res := 0
		if symPresent {
			res, _ = strconv.Atoi(line[indices[0]:indices[1]])
			total += int64(res)
		}
		for _, g := range gearIdx {
			gidx := idx*len(line) + g
			if _, ok := gears[gidx]; !ok {
				gears[gidx] = make([]int, 0)
			}
			gears[gidx] = append(gears[gidx], res)
		}
	}

	for idx, line := range inputs {
		res := r.FindAllStringIndex(line, -1)

		for _, indices := range res {
			// Look at surrounding cells
			// above first.
			lower_bound := max(0, indices[0]-1)
			upper_bound := min(len(line)-1, indices[1])
			if idx >= 1 {
				symPresent, gearIdx := searchLine(inputs[idx-1][lower_bound:upper_bound+1], lower_bound)
				processResult(symPresent, gearIdx, idx-1, line, indices)
			}
			symPresent, gearIdx := searchLine(line[lower_bound:upper_bound+1], lower_bound)
			processResult(symPresent, gearIdx, idx, line, indices)

			if idx < len(inputs)-1 {
				symPresent, gearIdx := searchLine(inputs[idx+1][lower_bound:upper_bound+1], lower_bound)
				processResult(symPresent, gearIdx, idx+1, line, indices)
			}
		}
	}

	for _, gIdx := range gears {
		if len(gIdx) == 2 {
			gearTotal += int64(gIdx[0] * gIdx[1])
		}
	}
	fmt.Printf("Result for infile %s : %d (total), %d (geartotal) \n", inputFlag, total, gearTotal)
}
