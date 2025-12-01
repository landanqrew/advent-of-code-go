package main

import (
	"log"

	"github.com/landanqrew/advent-of-code-go/internal/aoc"
	"github.com/landanqrew/advent-of-code-go/internal/year2025"
)

func main() {
	aoc.LoadEnv(".env")
	data, err := aoc.GetInput(2025, 1, ".env")
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}
	year2025.Day1(string(data))
}