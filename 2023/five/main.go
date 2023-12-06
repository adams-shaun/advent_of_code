package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag string
)

type CustomRange struct {
	outputMin int
	inputMin  int
	rangeSize int
}

func (c *CustomRange) getOutput(input int) int {
	if input >= c.inputMin && input < c.inputMin+c.rangeSize {
		return c.outputMin + (input - c.inputMin)
	}
	return -1
}

type CustomMap struct {
	desc   string
	ranges []CustomRange
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseInput(input []string) ([]int, []CustomMap) {
	seeds := make([]int, 0)
	maps := make([]CustomMap, 0)

	fullInput := strings.Join(input, "\n")
	splitInput := strings.Split(fullInput, "\n\n")

	seedSplit := strings.Split(splitInput[0], " ")
	for y := 1; y < len(seedSplit); y++ {
		res, _ := strconv.Atoi(seedSplit[y])
		seeds = append(seeds, res)
	}

	for j := 1; j < len(splitInput); j++ {
		new := CustomMap{}
		for idx, val := range strings.Split(splitInput[j], "\n") {
			if idx == 0 {
				new.desc = val
				continue
			}
			lineSplit := strings.Split(val, " ")
			res0, _ := strconv.Atoi(lineSplit[0])
			res1, _ := strconv.Atoi(lineSplit[1])
			res2, _ := strconv.Atoi(lineSplit[2])
			newRange := CustomRange{res0, res1, res2}

			new.ranges = append(new.ranges, newRange)
		}
		maps = append(maps, new)
	}

	return seeds, maps
}

func solveSeedLocation(seed int, maps []CustomMap) int {
	nextInput := seed
	for _, cMap := range maps {
		for _, cRange := range cMap.ranges {
			if out := cRange.getOutput(nextInput); out != -1 {
				nextInput = out
				// found = true
				break
			}
		}
	}
	return nextInput
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	seeds, maps := parseInput(inputs)

	res := make(chan int, 1000)
	var wg sync.WaitGroup

	for idx, s := range seeds {
		if idx%2 == 0 {
			wg.Add(1)
			go func(idx int, s int) {
				defer wg.Done()
				fmt.Printf("starting hunt for seed %d\n", s)
				locationMin := math.MaxInt
				for s2 := s; s2 < s+seeds[idx+1]; s2++ {
					locationMin = min(solveSeedLocation(s2, maps), locationMin)
				}
				res <- locationMin
			}(idx, s)
		}
	}

	wg.Wait()
	close(res)
	locationMin := math.MaxInt
	for r := range res {
		locationMin = min(locationMin, r)
	}

	fmt.Printf("Result for infile %s : %d\n", inputFlag, locationMin)
}
