package aoc

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/landanqrew/advent-of-code-go/internal/files"
)

var years = []int{2020, 2021, 2022, 2023, 2024, 2025}
var days = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

func UpdateInputs(envPath string) ([]byte, error) {
	for _, year := range years {
		err := os.MkdirAll(fmt.Sprintf("input/%d", year), 0755)
		if err != nil {
			fmt.Printf("failed to create directory for year %d: %v\n", year, err)
			continue
		}
		for _, day := range days {
			dt := time.Date(year, time.Month(day), 1, 0, 0, 0, 0, time.UTC)
			if dt.After(time.Now()) {
				continue
			}
			input, err := GetInput(year, day, envPath)
			if err != nil {
				return nil, err
			}
			padDay := fmt.Sprintf("%02d", day)
			err = files.WriteFile(fmt.Sprintf("input/%d/day_%s.txt", year, padDay), input)
			if err != nil {
				fmt.Printf("failed to write input for year %d, day %d: %v\n", year, day, err)
			}
		}
	}
	fmt.Println("inputs updated")
	return nil, nil
}

func UpdateInputForDay(year int, day int, envPath string)  {
	input, err := GetInput(year, day, envPath)
	if err != nil {
		log.Fatalf("failed to get input for year %d, day %d: %v", year, day, err)
	}
	padDay := fmt.Sprintf("%02d", day)
	err = files.WriteFile(fmt.Sprintf("input/%d/day_%s.txt", year, padDay), input)
	if err != nil {
		log.Fatalf("failed to write input for year %d, day %d: %v", year, day, err)
	}
	fmt.Printf("input for year %d, day %d updated", year, day)
}