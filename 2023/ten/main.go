package main

import (
	"errors"
	"flag"
	"fmt"
	"regexp"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

type Direction int

var (
	inputFlag string
)

const (
	up Direction = iota
	left
	down
	right
)

func findStart(input []string) Point {
	r, _ := regexp.Compile("S")

	for idx, row := range input {
		if res := r.FindStringIndex(row); res != nil {
			return Point{idx, res[0]}
		} else {
			continue
		}
	}
	return Point{0, 0}
}

type Point struct {
	y int
	x int
}

type Transition struct {
	lastDirection Direction
	nextRune      rune
}

func nextDirection(t Transition) Direction {
	transMap := map[Transition]Direction{
		{up, '|'}:    up,
		{up, 'F'}:    right,
		{up, '7'}:    left,
		{down, '|'}:  down,
		{down, 'L'}:  right,
		{down, 'J'}:  left,
		{left, '-'}:  left,
		{left, 'L'}:  up,
		{left, 'F'}:  down,
		{right, '-'}: right,
		{right, 'J'}: up,
		{right, '7'}: down,
	}
	return transMap[t]
}

func (p *Point) moveNext(inputs []string, d Direction) (Point, Direction, error) {
	newP := *p
	var newD Direction
	switch d {
	case up:
		newP.y -= 1
		if newP.y < 0 {
			return newP, up, errors.New("oob")
		}
	case down:
		newP.y += 1
		if newP.y > len(inputs)-1 {
			return newP, up, errors.New("oob")
		}
	case left:
		newP.x -= 1
		if newP.x < 0 {
			return newP, up, errors.New("oob")
		}
	case right:
		newP.x += 1
		if newP.x > len(inputs[0])-1 {
			return newP, up, errors.New("oob")
		}
	default:
		return *p, d, errors.New("bad direction")
	}
	t := Transition{d, rune(inputs[newP.y][newP.x])}
	newD = nextDirection(t)
	return newP, newD, nil
}

func findLoopPath(input []string) []Point {
	p0 := findStart(input)

	for _, d := range []Direction{up, down, left, right} {
		path := []Point{p0}
		start, lastP := p0, p0
		lastD := d
		for {
			nextP, nextD, err := lastP.moveNext(input, lastD)
			if err != nil {
				break
			}
			path = append(path, nextP)
			if start == nextP {
				return path
			}
			lastP = nextP
			lastD = nextD
		}
	}
	return []Point{}
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "input", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)

	path := findLoopPath(inputs)
	// markup(inputs, &inputs2)

	fmt.Printf("Result for infile %s : %d \n", inputFlag, len(path)/2)
}
