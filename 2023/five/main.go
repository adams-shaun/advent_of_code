package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
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
	inputMax  int // incl
}

type Span struct {
	min int
	max int // both incl
}

func calcSpans(span Span, mMap CustomMap) []Span {
	res := make([]Span, 0)

	newSpan := span
	for idx, rng := range mMap.ranges {
		if newSpan.min > rng.inputMax {
			if idx == len(mMap.ranges)-1 {
				res = append(res, newSpan)
			}
			continue
		}

		if newSpan.min > rng.inputMax {
			continue
		}
		if newSpan.min < rng.inputMin {
			if newSpan.max < rng.inputMin {
				res = append(res, newSpan)
				break
			}
			res = append(res, Span{newSpan.min, rng.inputMin - 1})
			newSpan.min = rng.inputMin
		}
		if newSpan.min >= rng.inputMin {
			if newSpan.max <= rng.inputMax {
				res = append(res, Span{rng.outputMin + (newSpan.min - rng.inputMin), rng.outputMin + (newSpan.max - rng.inputMin)})
				break
			}
			res = append(res, Span{rng.outputMin + (newSpan.min - rng.inputMin), rng.outputMin + (rng.inputMax - rng.inputMin)})
			newSpan.min = rng.inputMax + 1
		}
	}
	return res
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
			newRange := CustomRange{res0, res1, res1 + res2 - 1}

			new.ranges = append(new.ranges, newRange)
		}
		sort.Slice(new.ranges, func(i int, j int) bool {
			return new.ranges[i].inputMin < new.ranges[j].inputMin
		})
		maps = append(maps, new)
	}
	return seeds, maps
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	seeds, maps := parseInput(inputs)

	res := make(chan int, len(seeds)/2)
	var wg sync.WaitGroup

	for idx, s := range seeds {
		if idx%2 == 0 {
			wg.Add(1)
			go func(idx int, s int) {
				defer wg.Done()
				currSpans := []Span{Span{s, s + seeds[idx+1] - 1}}
				// fmt.Printf("starting hunt for seed range [%+v]\n", sp)

				for _, m := range maps {
					newSpans := make([]Span, 0)
					for _, cp := range currSpans {
						newSpans = append(newSpans, calcSpans(cp, m)...)
					}
					currSpans = newSpans
				}
				locationMin := math.MaxInt
				for _, sp := range currSpans {
					locationMin = min(sp.min, locationMin)
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
