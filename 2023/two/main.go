package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag  string
	inputRed   int
	inputBlue  int
	inputGreen int
)

type Game struct {
	id    int
	r_max int
	g_max int
	b_max int
}

func parseGame(line string) Game {
	// Little bit ugly... first attempt is using string splitting
	// TODO: maybe implement regex parsing
	res := Game{}

	split1 := strings.Split(line, ": ")
	split2 := strings.Split(split1[0], " ")
	id, err := strconv.Atoi(split2[1])
	if err != nil {
		log.Fatalf("Error parsing id in line %s, %v", line, err)
	}
	res.id = id

	// Now split each of the observations
	split3 := strings.Split(split1[1], "; ")
	for _, spl := range split3 {
		// And split by color
		for _, spl2 := range strings.Split(spl, ", ") {
			if spl2[len(spl2)-1:] == "n" {
				val, _ := strconv.Atoi(spl2[:strings.Index(spl2, " ")])
				if val > res.g_max {
					res.g_max = val
				}
			}
			if spl2[len(spl2)-1:] == "d" {
				val, _ := strconv.Atoi(spl2[:strings.Index(spl2, " ")])
				if val > res.r_max {
					res.r_max = val
				}
			}
			if spl2[len(spl2)-1:] == "e" {
				val, _ := strconv.Atoi(spl2[:strings.Index(spl2, " ")])
				if val > res.b_max {
					res.b_max = val
				}
			}
		}
	}
	return res
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "input", "input data set")
	flag.IntVar(&inputRed, "red", 12, "max reds allowed")
	flag.IntVar(&inputGreen, "green", 13, "max greens allowed")
	flag.IntVar(&inputBlue, "blue", 14, "max blues allowed")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)

	total := int64(0)
	powerTotal := int64(0)
	for _, line := range inputs {
		g := parseGame(line)

		if g.r_max > inputRed || g.b_max > inputBlue || g.g_max > inputGreen {
			fmt.Printf("Invalid game %+v\n", g)
		} else {
			total += int64(g.id)
		}

		powerTotal += int64(g.b_max * g.r_max * g.g_max)
	}

	fmt.Printf("Result for infile %s : %d (ID SUM), %d(POWER)\n", inputFlag, total, powerTotal)
}
