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
	reNodes   string = "[A-Z0-9]{3}"
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func parseInput(input []string) (string, []*Node) {
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

	roots := make([]*Node, 0)
	for k, v := range inNodeMap {
		if strings.HasSuffix(k, "A") {
			roots = append(roots, v)
		}
	}
	return input[0], roots
}

func main() {
	// Read in data set
	flag.StringVar(&inputFlag, "input", "input", "input data set")
	flag.Parse()
	inputs := common.ReadInput(inputFlag)
	pattern, roots := parseInput(inputs)

	currNodes := roots
	count := 0
	for {
		found := false
		for idx, v := range pattern {
			count += 1
			allZ := true
			for idx, n := range currNodes {
				newNode := n.moveNext(v)

				if strings.HasSuffix(newNode.nodeId, "Z") {
					fmt.Printf("roots %d hit Z after %d\n", idx, count)
				}
				if allZ && !strings.HasSuffix(newNode.nodeId, "Z") {
					allZ = false
				}
				currNodes[idx] = newNode
			}
			// fmt.Printf("Curr %s %s\n", currNodes[0].nodeId, currNodes[1].nodeId)
			if allZ && idx == len(pattern)-1 {
				found = true
			}
		}
		if found {
			break
		}
	}

	fmt.Printf("Result for infile %s : %d\n", inputFlag, count)
}
