package year2025

import (
	"fmt"
	"math"
	"strings"
)


func getDay9Part1Example() string {
	return `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
}

func getPermutations(coordinates []Coordinate) [][]Coordinate {
	permutations := make([][]Coordinate, 0)
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			permutations = append(permutations, []Coordinate{coordinates[i], coordinates[j]})
		}
	}
	return permutations
}

func Day9Part1(data string) {
	if data == "" {
		data = getDay9Part1Example()
	}
	lines := strings.Split(data, "\n")
	coordinates := make([]Coordinate, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x := convertStringToInt(parts[0])
		y := convertStringToInt(parts[1])
		c := Coordinate{
			X: x,
			Y: y,
		}
		coordinates = append(coordinates, c)
	}

	permutations := getPermutations(coordinates)
	maxArea := 0
	for _, permutation := range permutations {
		area := int((math.Abs(float64(permutation[0].X - permutation[1].X)) + 1.0) * (math.Abs(float64(permutation[0].Y - permutation[1].Y)) + 1.0))
		if area > maxArea {
			maxArea = area
		}
	}
	fmt.Println(maxArea)
}