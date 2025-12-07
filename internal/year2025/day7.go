package year2025

import (
	"slices"
	"strings"
)


type day7Node struct {
	Coordinate Coordinate `json:"coordinate"`
	Char string `json:"char"`
	Left *day7Node `json:"left"`
	Right *day7Node `json:"right"`
}

func (n *day7Node) AddLeft(node *day7Node) {
	n.Left = node
}

func (n *day7Node) AddRight(node *day7Node) {
	n.Right = node
}

func (n *day7Node) GetLeft() *day7Node {
	return n.Left
}

func (n *day7Node) GetRight() *day7Node {
	return n.Right
}

func (n *day7Node) CountChildren() int {
	count := 0
	if n.Left != nil {
		count += 1 + n.Left.CountChildren()
	}
	if n.Right != nil {
		count += 1 + n.Right.CountChildren()
	}
	return count
}


func getSplitters(data string) *day7Node {
	lines := strings.Split(data, "\n")
	slices.Reverse(lines)
	xMap := make([]*day7Node, len(lines[0]))

	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, char := range line {
			charString := string(char)

			if charString == "^" || charString == "S" {
				// build node
				node := &day7Node{
					Coordinate: Coordinate{
						X: x,
						Y: y,
					},
					Char: charString,
				}

				xMap[x] = node

				// check left
				if x > 0 {
					node.AddLeft(xMap[x-1])
				}
				// check right
				if x + 1 < len(xMap) {
					node.AddRight(xMap[x+1])
				}

				if charString == "S" {
					return node
				}
			}
		}
	}
	return nil
}

func Day7Part1(data string) {
	if data == "" {
		data = getDay7Part1Example()
	}
	splitter := getSplitters(data)
	fmt.Printf("2025 Day 7 Part 1: %d\n", splitter.CountChildren())
}