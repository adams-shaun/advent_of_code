package main

import (
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag string
	re_number string = "(\\d+)"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func mustAtoi(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func parseInput(input []string) ([]int, []int) {
	r, _ := regexp.Compile(re_number)
	times := make([]int, 0)
	distances := make([]int, 0)

	for {
		for _, res := range r.FindAllString(input[0], -1) {
			times = append(times, mustAtoi(res))
		}
		for _, res := range r.FindAllString(input[1], -1) {
			distances = append(distances, mustAtoi(res))
		}
		break
	}
	return times, distances
}

func findLimits(time int, distance int) (int, int) {
	// distance traveled :
	// y = m*(t-m)
	// distance = m*(time - m)
	// 0 = m^2 - time*m + distance
	lower := (float64(time) - math.Sqrt(float64(time*time-4*distance))) / 2
	upper := (float64(time) + math.Sqrt(float64(time*time-4*distance))) / 2

	lower_rnd := int(math.Ceil(lower))
	upper_rnd := int(math.Floor(upper))

	// Must actually beat the record
	if lower_rnd*(time-lower_rnd) <= distance {
		lower_rnd += 1
	}

	if upper_rnd*(time-upper_rnd) <= distance {
		upper_rnd -= 1
	}

	return lower_rnd, upper_rnd
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	times, distances := parseInput(inputs)

	res := 1

	for idx, _ := range times {
		lower, upper := findLimits(times[idx], distances[idx])
		raceSolutions := (upper - lower + 1)
		res *= raceSolutions
	}

	fmt.Printf("Result for infile %s : %d\n", inputFlag, res)
}
