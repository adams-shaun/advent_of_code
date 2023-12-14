package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

var (
	inputFlag string
)

func findStart(input []string) (int, int) {
	r, _ := regexp.Compile("S")

	for idx, row := range input {
		if res := r.FindStringIndex(row); res != nil {
			return idx, res[0]
		} else {
			continue
		}
	}
	return 0, 0
}

type Path struct {
	currX int
	currY int
	lastX int
	lastY int
	steps int
}

func (p *Path) intersects(p2 *Path) bool {
	if p.currX == p2.currX && p.currY == p2.currY {
		return true
	}
	return false
}

func (p *Path) goNext(input []string) {
	tmpX, tmpY := p.currX, p.currY
	if input[p.currY][p.currX] == '|' {
		if p.lastY == p.currY-1 {
			p.currY += 1
		} else {
			p.currY -= 1
		}
	} else if input[p.currY][p.currX] == '-' {
		if p.lastX == p.currX-1 {
			p.currX += 1
		} else {
			p.currX -= 1
		}
	} else if input[p.currY][p.currX] == '7' {
		if p.lastY == p.currY {
			p.currY += 1
		} else {
			p.currX -= 1
		}
	} else if input[p.currY][p.currX] == 'J' {
		if p.lastY == p.currY {
			p.currY -= 1
		} else {
			p.currX -= 1
		}
	} else if input[p.currY][p.currX] == 'L' {
		if p.lastY == p.currY {
			p.currY -= 1
		} else {
			p.currX += 1
		}
	} else if input[p.currY][p.currX] == 'F' {
		if p.lastY == p.currY {
			p.currY += 1
		} else {
			p.currX += 1
		}
	}
	p.lastX, p.lastY = tmpX, tmpY

	p.steps += 1
}

func startPaths(input []string, x, y int) []*Path {
	paths := []*Path{}
	maxY := len(input)
	maxX := len(input[0])
	if y < maxY && (input[y+1][x] == '|' || input[y+1][x] == 'J' || input[y+1][x] == 'L') {
		paths = append(paths, &Path{x, y + 1, x, y, 1})
	}
	if y > 0 && (input[y-1][x] == '|' || input[y+1][x] == '7' || input[y+1][x] == 'F') {
		paths = append(paths, &Path{x, y - 1, x, y, 1})
	}
	if x < maxX && (input[y][x+1] == '-' || input[y][x+1] == '7' || input[y][x+1] == 'J') {
		paths = append(paths, &Path{x + 1, y, x, y, 1})
	}
	if x > 0 && (input[y][x-1] == '-' || input[y][x-1] == 'F' || input[y][x-1] == 'L') {
		paths = append(paths, &Path{x - 1, y, x, y, 1})
	}

	return paths
}

func findLoopCnt(input []string) int {
	sY, sX := findStart(input)
	paths := startPaths(input, sX, sY)

	fmt.Printf("P0: %+v, P1: %+v\n", paths[0], paths[1])

	for {
		paths[0].goNext(input)
		paths[1].goNext(input)

		fmt.Printf("P0: %+v, P1: %+v\n", paths[0], paths[1])

		if paths[0].intersects(paths[1]) {
			return paths[0].steps
		}
	}
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)

	count := findLoopCnt(inputs)

	fmt.Printf("Result for infile %s : %d \n", inputFlag, count)
}
