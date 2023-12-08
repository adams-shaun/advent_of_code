package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/adams-shaun/advent_of_code/2023/common"
)

type handStrength int

var (
	inputFlag string
	reNodes   string = "[A-Z]{3}"
)

type Node struct {
	nodeId string
	lEdge  *Node
	rEdge  *Node
}

func (n *Node) moveNext(input rune) *Node {
	if input == 'L' {
		return n.lEdge
	}
	return n.rEdge
}

func parseInput(input []string) (string, *Node) {
	inNodeMap := map[string]*Node{}
	for _, line := range input[2:] {
		split := strings.Split(line, " = ")
		inNodeMap[split[0]] = &Node{nodeId: split[0]}
	}
	// First pass puts all nodes in, second one we creat the edges
	r, _ := regexp.Compile(reNodes)
	for _, line := range input[2:] {
		split := strings.Split(line, " = ")
		thisNode := inNodeMap[split[0]]
		for idx, val := range r.FindAllString(split[1], 2) {
			if idx == 0 {
				thisNode.lEdge = inNodeMap[val]
			} else {
				thisNode.rEdge = inNodeMap[val]
			}
		}
	}
	return input[0], inNodeMap["AAA"]
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "testinput", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	pattern, root := parseInput(inputs)

	currNode := root
	count := 0
	for {
		found := false
		for idx, v := range pattern {
			count += 1
			currNode = currNode.moveNext(v)
			if currNode.nodeId == "ZZZ" && idx == len(pattern)-1 {
				found = true
			}
		}
		if found {
			break
		}
	}

	fmt.Printf("Result for infile %s : %d\n", inputFlag, count)
}
