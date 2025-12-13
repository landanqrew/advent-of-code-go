package year2025

import "strings"

type button struct {
	Indices []int `json:"indices"`
}

type diagram struct {
	Goal    string   `json:"goal"`
	State   string   `json:"state"`
	Buttons []button `json:"buttons"`
	Jolts   []int    `json:"jolts"`
	MinButtonPresses int `json:"min_button_presses"`
}

func (d *diagram) PressButton(button button) {
	runes := []rune(d.State)
	for _, index := range button.Indices {
		if string(runes[index]) == "." {
			runes[index] = '#'
		} else {
			runes[index] = '.'
		}
	}
	d.State = string(runes)
}

func (d *diagram) Configure(currentState string, presses int) {
	if presses > d.MinButtonPresses {
		return
	}

	if currentState == d.Goal {
		d.MinButtonPresses = presses
		return
	}

	for _, button := range d.Buttons {
		d.State = currentState
		d.PressButton(button)
		d.Configure(d.State, presses + 1)
	}
}




func getDay10Part1Example() string {
	return `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`
}

func buildDiagram(line string) *diagram {
	parts := strings.Split(line, " ")
	goal := parts[0]
	state := parts[1]
	indices := strings.Split(parts[2], ",")
	indicesInt := make([]int, len(indices))
	for i, index := range indices {
		indicesInt[i] = convertStringToInt(index)
	}
	return &diagram{Goal: goal, State: state, Buttons: []button{{Indices: indicesInt}}}
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
}
