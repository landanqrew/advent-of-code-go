package year2025

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Day4Node struct {
	Coordinate        Coordinate  `json:"coordinate"`
	Char              string      `json:"char"`
	Connections       []*Day4Node `json:"connections"`
	AdjacentRollCount int         `json:"adjacent_roll_count"`
	Evaluated         bool        `json:"evaluated"`
}

type Day4Map struct {
	Day4Map      map[Coordinate]*Day4Node `json:"day4_map"`
	MaxX         int                      `json:"max_x"`
	MaxY         int                      `json:"max_y"`
	PassFunc     func(n *Day4Node) bool   `json:"pass_func"`
	SuccessCount int                      `json:"success_count"`
	EvalComplete bool                     `json:"eval_complete"`
}

func (d *Day4Map) SetPassFunc(passFunc func(n *Day4Node) bool) {
	d.PassFunc = passFunc
}

func (d *Day4Map) EvaluateMap() {
	d.EvalComplete = true
	for _, node := range d.Day4Map {
		node.EvaluateNode(*d)
		if d.PassFunc(node) {
			// fmt.Printf("Passing node: (%d, %d, %s)\n", node.Coordinate.X, node.Coordinate.Y, node.Char)
			d.SuccessCount++
		}
	}
}

func (d *Day4Map) AddNodeToMap(node *Day4Node) {
	d.Day4Map[node.Coordinate] = node
}

func (n *Day4Node) AddConnection(connection *Day4Node) {
	n.Connections = append(n.Connections, connection)
}

func (n *Day4Node) EvaluateNode(m Day4Map) {
	debug := false
	if n.Evaluated {
		return
	}
	n.Evaluated = true
	for j := max(n.Coordinate.Y-1, 0); j <= min(n.Coordinate.Y+1, m.MaxY); j++ {
		for i := max(n.Coordinate.X-1, 0); i <= min(n.Coordinate.X+1, m.MaxX); i++ {
			if i == n.Coordinate.X && j == n.Coordinate.Y {
				continue
			}
			coordinate := Coordinate{
				X: i,
				Y: j,
			}
			node, ok := m.Day4Map[coordinate]
			if !ok {
				log.Fatalf("node not found at coordinate (%d, %d) when evaluating node (%d, %d)", i, j, n.Coordinate.X, n.Coordinate.Y)
			}
			char := node.Char
			if char == "@" {
				n.AdjacentRollCount++
			}
			n.AddConnection(node)
			if debug && n.Coordinate.X == 1 && n.Coordinate.Y == 9 {
				fmt.Printf("Found connection: (%d, %d, %s)\n", i, j, node.Char)
			}
		}
	}
	if debug {
		fmt.Printf("Evaluated node: (%d, %d, %s)\n", n.Coordinate.X, n.Coordinate.Y, n.Char)
		n.PrintConnections()
	}
}

func (n *Day4Node) PrintConnections() {
	printString := ""
	for _, connection := range n.Connections {
		printString += connection.Char
	}
	fmt.Println(printString)
}

func BuildDay4Map(data string) Day4Map {
	debug := false
	if debug {
		fmt.Println("Building day 4 map...")
	}
	lines := strings.Split(data, "\n")
	day4Map := Day4Map{
		Day4Map:      make(map[Coordinate]*Day4Node),
		MaxX:         0,
		MaxY:         0,
		EvalComplete: false,
	}
	slices.Reverse(lines)
	for i, line := range lines {
		for j, char := range line {
			coordinate := Coordinate{
				X: j,
				Y: i,
			}
			if debug && i == 0 {
				fmt.Printf("Adding node: (%d, %d, %s)\n", j, i, string(char))
			}
			day4Map.Day4Map[coordinate] = &Day4Node{
				Coordinate:        coordinate,
				Char:              string(char),
				Connections:       make([]*Day4Node, 0),
				AdjacentRollCount: 0,
				Evaluated:         false,
			}
			if j > day4Map.MaxX {
				day4Map.MaxX = j
			}
		}
		if i > day4Map.MaxY {
			day4Map.MaxY = i
		}
	}
	return day4Map
}

func getDay4Part1Example() string {
	return `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
}

func Day4Part1(data string) {
	if data == "" {
		data = getDay4Part1Example()
	}
	day4Map := BuildDay4Map(data)
	day4Map.SetPassFunc(func(n *Day4Node) bool {
		return n.AdjacentRollCount < 4 && n.Char == "@"
	})
	day4Map.EvaluateMap()
	fmt.Printf("2025 Day 4 Part 1: %d\n", day4Map.SuccessCount)
}

func Day4Part2(data string) {
	if data == "" {
		data = getDay4Part1Example()
	}
	day4Map := BuildDay4Map(data)
	day4Map.SetPassFunc(func(n *Day4Node) bool {
		result := n.AdjacentRollCount < 4 && n.Char == "@"
		if result {
			// fmt.Printf("Removing node: (%d, %d, %s)\n", n.Coordinate.X, n.Coordinate.Y, n.Char)
			day4Map.EvalComplete = false
			n.Char = "."
			n.Evaluated = false
			n.AdjacentRollCount = 0
			for _, connection := range n.Connections {
				if connection.Evaluated {
					connection.Evaluated = false
					connection.AdjacentRollCount = 0
				}
			}
		}

		return result
	})
	for !day4Map.EvalComplete {
		day4Map.EvaluateMap()
	}
	fmt.Printf("2025 Day 4 Part 2: %d\n", day4Map.SuccessCount)
}
