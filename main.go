package main

import (
	"fmt"
	"time"

	"github.com/landanqrew/advent-of-code-go/internal/aoc"
	"github.com/landanqrew/advent-of-code-go/internal/files"
	"github.com/landanqrew/advent-of-code-go/internal/year2025"
)

func main() {
	aoc.LoadEnv(".env")
	// aoc.UpdateInputs(".env")
	//aoc.UpdateInputForDay(2025, 5, ".env")
	data := files.GetInputFromFile(2025, 5)
	fmt.Println(data[0])
	start := time.Now()
	year2025.Day5(data)
	end := time.Now()
	fmt.Printf("time taken: %s\n", end.Sub(start).String())
}