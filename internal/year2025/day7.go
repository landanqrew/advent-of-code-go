package year2025

import (
	"fmt"
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
	if n.GetLeft() != nil {
		count += 1 + n.Left.CountChildren()
	}
	if n.GetRight() != nil {
		count += 1 + n.Right.CountChildren()
	}

	//fmt.Printf("node (%d, %d): count: %d\n", n.Coordinate.X, n.Coordinate.Y, count)
	return count
}

func (n *day7Node) CountPermutations(seen map[*day7Node]int) int {
	permutations, ok := seen[n]
	if ok {
		return permutations
	}
	count := 0
	if n.GetLeft() != nil {
		count += n.Left.CountPermutations(seen)
	} else {
		count += 1
	}
	if n.GetRight() != nil {
		count += n.Right.CountPermutations(seen)
	} else {
		count += 1
	}
	seen[n] = count

	// fmt.Printf("node (%d, %d): count: %d\n", n.Coordinate.X, n.Coordinate.Y, count)
	return count
}

func (n *day7Node) CountUniqueChildren(seen map[*day7Node]bool) int {
	if seen[n] {
		return 0
	}
	seen[n] = true
	count := 0
	if n.GetLeft() != nil && !seen[n.GetLeft()] {
		count += 1 + n.GetLeft().CountUniqueChildren(seen)
	}
	if n.GetRight() != nil && !seen[n.GetRight()] {
		count += 1 + n.GetRight().CountUniqueChildren(seen)
	}
	return count
}



func getDay7Part1Example() string {
	return `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`
}


func getSplitters(data string) *day7Node {
	lines := strings.Split(data, "\n")
	slices.Reverse(lines)
	xMap := make([]*day7Node, len(lines[0]))

	for y, line := range lines {
		if line == "" {
			continue
		}
		// countFound := 0
		for x, char := range line {
			charString := string(char)

			if charString == "^" {
				// countFound++
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
			}

			if charString == "S" {
				// countFound++
				// build node
				node := &day7Node{
					Coordinate: Coordinate{
						X: x,
						Y: y,
					},
					Char: charString,
				}

				// check left
				node.AddLeft(xMap[x])
				node.AddRight(nil)

				return node
			}
			
		}
		//fmt.Printf("row (%d): countFound: %d\n", y, countFound)
	}
	return nil
}

func Day7Part1(data string) {
	if data == "" {
		data = getDay7Part1Example()
	}
	head := getSplitters(data)
	fmt.Printf("2025 Day 7 Part 1: %d\n", head.CountUniqueChildren(make(map[*day7Node]bool)))
	fmt.Println(head.CountPermutations(make(map[*day7Node]int)) - 1)
}