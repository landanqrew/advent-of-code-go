package year2025

import (
	"fmt"
	"strings"
)

type validRange struct {
	LowerBound int   `json:"lower_bound"`
	UpperBound int   `json:"upper_bound"`
	ValidIDs   []int `json:"valid_ids"`
}

func getDay5Part1Example() string {
	return `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
}

func getValidRanges(lines []string) []validRange {
	validRanges := make([]validRange, 0)
	for _, line := range lines {
		bounds := strings.Split(line, "-")
		validRange := validRange{
			LowerBound: convertStringToInt(bounds[0]),
			UpperBound: convertStringToInt(bounds[1]),
			ValidIDs:   make([]int, 0),
		}
		validRanges = append(validRanges, validRange)
	}
	return validRanges
}

func getPotentialIDs(lines []string) []int {
	potentialIDs := make([]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		num := convertStringToInt(strings.TrimSpace(line))
		potentialIDs = append(potentialIDs, num)
	}
	return potentialIDs
}

func Day5(data string) {
	if data == "" {
		data = getDay5Part1Example()
	}
	lines := strings.Split(data, "\n")
	breakPoint := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			breakPoint = i
			break
		}
	}
	validRanges := getValidRanges(lines[:breakPoint])
	potentialIDs := getPotentialIDs(lines[breakPoint+1:])
	// fmt.Println(potentialIDs)
	validCount := 0
	for _, potentialID := range potentialIDs {
		for _, validRange := range validRanges {
			if potentialID >= validRange.LowerBound && potentialID <= validRange.UpperBound {
				//fmt.Printf("potentialID: %d is in valid range: %d-%d\n", potentialID, validRange.LowerBound, validRange.UpperBound)
				validCount++
				validRange.ValidIDs = append(validRange.ValidIDs, potentialID)
				break
			}
		}
	}
	fmt.Println(validCount)

	possibleIDs := 0
	for i := 0; i < len(validRanges); i++ {
		vr1 := validRanges[i]
		for j := 0; j < len(validRanges); j++ {
			if i == j {
				continue
			}
			vr2 := validRanges[j]
			if vr2.LowerBound <= vr1.LowerBound && vr2.UpperBound >= vr1.LowerBound {
				// fmt.Printf("vr1: %d-%d, vr2: %d-%d\n", vr1.LowerBound, vr1.UpperBound, vr2.LowerBound, vr2.UpperBound)
				validRanges[i].LowerBound = vr2.UpperBound + 1
				// fmt.Printf("new vr1: %d-%d\n", validRanges[i].LowerBound, validRanges[i].UpperBound)
			}
			if vr2.LowerBound <= vr1.UpperBound && vr2.UpperBound >= vr1.UpperBound {
				// fmt.Printf("vr1: %d-%d, vr2: %d-%d\n", vr1.LowerBound, vr1.UpperBound, vr2.LowerBound, vr2.UpperBound)
				validRanges[i].UpperBound = vr2.LowerBound - 1
				// fmt.Printf("new vr1: %d-%d\n", validRanges[i].LowerBound, validRanges[i].UpperBound)
			}
		}
		addNum := max(validRanges[i].UpperBound - validRanges[i].LowerBound + 1, 0)
		// fmt.Printf("validRanges[%d]: %d-%d -> %d\n", i, validRanges[i].LowerBound, validRanges[i].UpperBound, addNum)
		possibleIDs += addNum
	}
	fmt.Println(possibleIDs)
}
