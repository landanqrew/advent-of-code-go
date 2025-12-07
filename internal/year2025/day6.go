package year2025

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

type mathProblem struct {
	Numbers []int `json:"numbers"`
	Part2Numbers []int `json:"part2_numbers"`
	Operator string `json:"operator"`
	Result int `json:"result"`
	Part2Result int `json:"part2_result"`
}

func (m *mathProblem) Solve() {
	switch m.Operator {
	case "*":
		m.Result = multiplyNumbers(m.Numbers)
	case "+":
		m.Result = addNumbers(m.Numbers)
	default:
		log.Fatalf("invalid operator: %s", m.Operator)
	}
}

func multiplyNumbers(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result *= nums[i]
	}
	return result
}

func addNumbers(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

func (m *mathProblem) solvePart2() {
	switch m.Operator {
	case "*":
		m.Part2Result = multiplyNumbers(m.Part2Numbers)
	case "+":
		m.Part2Result = addNumbers(m.Part2Numbers)
	default:
		log.Fatalf("invalid operator: %s", m.Operator)
	}
}


func getExampleDay6Example() string {
	return `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
}

func extractMathProblems(input string) []mathProblem {
	lines := strings.Split(input, "\n")
	problems := make([]mathProblem, 0)
	problem := mathProblem{
		Numbers: make([]int, 0),
		Part2Numbers: make([]int, 0),
		Operator: "",
		Result: 0,
		Part2Result: 0,
	}
	rowStrings := make([]string, 0)
	for x := 0; x < len(lines[0]); x++ {
		columnString := ""
		for y := 0; y < len(lines); y++ {
			if lines[y] == "" {
				continue
			}

			char := string(lines[y][x])
			
			if char == "*" || char == "+" {
				problem.Operator = char
			} else {
				if len(rowStrings) <= y {
					rowStrings = append(rowStrings, "")
				}
				rowStrings[y] += char
				columnString += char
			}
		}

		// push column string to problem if not end of block
		if strings.TrimSpace(columnString) != "" {
			problem.Part2Numbers = append(problem.Part2Numbers, convertStringToInt(strings.ReplaceAll(columnString, " ", "")))
		}
		// end of block or finished last column of input
		if strings.TrimSpace(columnString) == "" || x == len(lines[0]) - 1 {
			// update problem
			for y, rowString := range rowStrings {
				
				if strings.TrimSpace(rowString) != "" {
					// fmt.Println("rowString: ", rowString)
				  // fmt.Println("converted: ", convertStringToInt(strings.ReplaceAll(rowString, " ", "")))
					problem.Numbers = append(problem.Numbers, convertStringToInt(strings.ReplaceAll(rowString, " ", "")))
				}
				// reset row string
				rowStrings[y] = ""
			}
			slices.Reverse(problem.Part2Numbers)
			// fmt.Println("problem: ", problem)
			problems = append(problems, problem)
			// reset everything
			problem = mathProblem{
				Numbers: make([]int, 0),
				Part2Numbers: make([]int, 0),
				Operator: "",
				Result: 0,
				Part2Result: 0,
			}
		}
	}
	// fmt.Println("stringGroups: ", stringGroups)
	return problems
}

func solveMathProblems(mathProblems []mathProblem) int {
	total := 0
	for _, mathProblem := range mathProblems {
		mathProblem.Solve()
		total += mathProblem.Result
	}
	return total
}

func solveMathProblemsPart2(mathProblems []mathProblem) int {
	total := 0
	for _, mathProblem := range mathProblems {
		mathProblem.solvePart2()
		total += mathProblem.Part2Result
	}
	return total
}

func Day6Part1(data string) {
	if data == "" {
		data = getExampleDay6Example()
	}
	mathProblems := extractMathProblems(data)
	total := solveMathProblems(mathProblems)
	part2Total := solveMathProblemsPart2(mathProblems)
	fmt.Println("total: ", total)
	fmt.Println("part2 total: ", part2Total) // (9562820909123) too low
}