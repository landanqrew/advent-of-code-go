package year2025

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/landanqrew/advent-of-code-go/internal/re"
)

type button struct {
	Indices []int `json:"indices"`
}

type diagram struct {
	Goal             string         `json:"goal"`
	State            string         `json:"state"`
	Buttons          []button       `json:"buttons"`
	Jolts            []int          `json:"jolts"`
	MinButtonPresses int            `json:"min_button_presses"`
	Solved           bool           `json:"solved"`
	memo             map[string]int // implemented for performance optimization
}

func (d *diagram) PressButton(inputState string, button button) string {
	runes := []rune(inputState) 
	for _, index := range button.Indices {
		if runes[index] == '.' {
			runes[index] = '#'
		} else {
			runes[index] = '.'
		}
	}
	return string(runes)
}

func (d *diagram) Configure(currentState string, presses int, maxPresses int) {

	if d.Solved && presses >= d.MinButtonPresses {
		return
	}
	if presses > maxPresses {
		return
	}

	if cachedPresses, ok := d.memo[currentState]; ok {
		if presses >= cachedPresses {
			return // Already found an equal or shorter path to this state
		}
	}
	// Update memo with the current (potentially shorter) path to this state
	d.memo[currentState] = presses


	if currentState == d.Goal {
		if presses < d.MinButtonPresses {
			d.MinButtonPresses = presses
			d.Solved = true 
		}
		return
	}

	offIndices := make([]int, 0)
	for i, char := range currentState {
		if char != rune(d.Goal[i]) {
			offIndices = append(offIndices, i)
		}
	}

	for _, button := range d.Buttons {
		if len(intersect(button.Indices, offIndices)) == 0 {
			continue 
		}

		newState := d.PressButton(currentState, button)
		d.Configure(newState, presses+1, maxPresses)
	}
}

func intersect(a []int, b []int) []int {
	intersect := make([]int, 0)
	for _, element := range a {
		if slices.Contains(b, element) {
			intersect = append(intersect, element)
		}
	}
	return intersect
}

func getDay10Part1Example() string {
	return `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`
}

func buildDiagram(line string) *diagram {
	parts := strings.Split(line, " ")
	goal := strings.ReplaceAll(parts[0], "[", "")
	goal = strings.ReplaceAll(goal, "]", "")
	buttons := make([]button, 0)
	for _, buttonString := range parts[1:] {
		if strings.HasPrefix(buttonString, "(") {
			numbers := re.GetNumbers(buttonString)
			button := button{Indices: make([]int, len(numbers))}
			for i, num := range numbers {
				button.Indices[i] = convertStringToInt(num)
			}
			buttons = append(buttons, button)
		} else {
			// handle jolts???
			jolts := re.GetNumbers(buttonString)
			joltsInt := make([]int, len(jolts))
			for i, num := range jolts {
				joltsInt[i] = convertStringToInt(num)
			}
			return &diagram{Goal: goal, State: strings.ReplaceAll(goal, "#", "."), Buttons: buttons, Jolts: joltsInt, MinButtonPresses: 999999999, Solved: false, memo: make(map[string]int)}
		}
	}
	return &diagram{Goal: goal, State: strings.ReplaceAll(goal, "#", "."), Buttons: buttons, Jolts: make([]int, 0), MinButtonPresses: 999999999, Solved: false, memo: make(map[string]int)}
}

func decodeInput(data string) []*diagram {
	diagrams := make([]*diagram, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		diagrams = append(diagrams, buildDiagram(line))
	}
	return diagrams
}

func Day10Part1(data string) {
	if data == "" {
		data = getDay10Part1Example()
	}
	diagrams := decodeInput(data)
	total := 0
	for i, diagram := range diagrams {
		maxPresses := 10
		diagram.Configure(diagram.State, 0, maxPresses)

		if !diagram.Solved {
			// expand max presses if not solved
			maxPresses = 20                      
			diagram.MinButtonPresses = 999999999 
			diagram.Solved = false
			diagram.memo = make(map[string]int) 
			diagram.Configure(diagram.State, 0, maxPresses)
		}

		if !diagram.Solved {
			log.Fatalf("diagram %d not solved in %d tries", i, maxPresses)
		}
		fmt.Printf("diagram %d solved in %d tries\n", i, diagram.MinButtonPresses)
		total += diagram.MinButtonPresses
	}

	fmt.Println(total)

}
