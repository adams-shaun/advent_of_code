package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

type Direction int

var (
	inputFlag string
)

type Point struct {
	row int
	col int
}

// recursive backtracking to get distinct pairs of galaxies
func getPairs(points []Point) [][]Point {
	res := [][]Point{}

	var bt func(start int, path []Point)
	bt = func(start int, path []Point) {
		if len(path) == 2 {
			res = append(res, []Point{path[0], path[1]})
		} else {
			for idx, p := range points[start:] {
				if idx > 0 && p == points[idx-1] {
					continue
				}

				path = append(path, p)
				bt(idx+1, path)
				path = path[:len(path)-1]

			}
		}
	}
	bt(0, []Point{})
	return res
}

func getPoints(input []string) []Point {
	res := []Point{}

	for row := range input {
		for col := range input[0] {
			if input[row][col] == '#' {
				res = append(res, Point{row, col})
			}
		}
	}
	return res
}

func expand(input []string) (emptyRows []int, emptyCols []int) {
	emptyRows = []int{}
	emptyCols = []int{}

	// expand row
	for idx, row := range input {
		if strings.Count(row, "#") == 0 {
			emptyRows = append(emptyRows, idx)
		}
	}

	// expand col
	for col, _ := range input[0] {
		empty := true
		for row, _ := range input {
			if input[row][col] != '.' {
				empty = false
			}
		}
		if empty {
			emptyCols = append(emptyCols, col)
		}
	}
	return
}

func getDist(p0 Point, p1 Point, emptyRows []int, emptyCols []int, factor int) int {
	res := 0

	// find empty rows b/t p0 and p1
	rowCnt := 0
	for _, val := range emptyRows {
		if (val > p0.row && val < p1.row) || (val < p0.row && val > p1.row) {
			rowCnt += 1
		}
	}

	colCnt := 0
	for _, val := range emptyCols {
		if (val > p0.col && val < p1.col) || (val < p0.col && val > p1.col) {
			colCnt += 1
		}
	}

	deltaRow := p0.row - p1.row
	if deltaRow < 0 {
		deltaRow = -deltaRow
	}
	deltaCol := p0.col - p1.col
	if deltaCol < 0 {
		deltaCol = -deltaCol
	}

	res = deltaRow + (rowCnt * (factor - 1)) + deltaCol + (colCnt * (factor - 1))
	return res
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "input", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	emptyRows, emptyCols := expand(inputs)
	points := getPoints(inputs)
	pairs := getPairs(points)

	sum := 0
	sum2 := 0
	for _, pair := range pairs {
		sum += getDist(pair[0], pair[1], emptyRows, emptyCols, 2)
		sum2 += getDist(pair[0], pair[1], emptyRows, emptyCols, 1000000)
	}

	fmt.Printf("Result for infile %s : %d %d\n", inputFlag, sum, sum2)
}
