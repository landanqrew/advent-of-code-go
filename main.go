package main

import (
	"fmt"
	"log"
	"time"

	"github.com/landanqrew/advent-of-code-go/internal/aoc"
	"github.com/landanqrew/advent-of-code-go/internal/year2025"
)

func main() {
	aoc.LoadEnv(".env")
	// aoc.UpdateInputs(".env")
	data, err := aoc.GetInput(2025, 3, ".env")
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}
	start := time.Now()
	year2025.Day3Part1(string(data))
	year2025.Day3Part2(string(data))
	end := time.Now()
	fmt.Printf("time taken: %s\n", end.Sub(start).String())
}