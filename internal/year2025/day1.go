package year2025

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/landanqrew/advent-of-code-go/internal/files"
	"github.com/landanqrew/advent-of-code-go/internal/re"
)

func getDay1Part1Example() string {
	return `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
}

type click struct {
	Direction      string `json:"direction"`
	OriginalNumber int64  `json:"original_number"`
	Number         int64  `json:"number"`
	FinalPosition  int64  `json:"final_position"`
}

func Day1(data string) {
	dirMap := map[string]int{
		"L": -1,
		"R": 1,
	}
	if data == "" {
		data = getDay1Part1Example()
	}
	lines := strings.Split(data, "\n")
	numbers := re.GetNumbers(data)
	clicks := make([]click, len(lines))
	curPos := 50
	// counted only when final position is 0
	zerosLanded := 0
	// counted every time the click passes 0
	zerosPassed := 0
	for i, line := range lines {
		// just accounts for very last line of file
		if line == "" {
			continue
		}
		d := dirMap[string(line[0])]
		originalNum, err := strconv.ParseInt(numbers[i], 10, 64)
		num := originalNum
		if err != nil {
			log.Fatalf("number %s cannot be parsed\n", string(numbers[i]))
		}
		if num >= 100 {
			parts := num / 100
			num = (num % 100)
			zerosPassed += int(parts)
		}
		finalNum := d * int(num)
		curPos += finalNum
		if curPos < 0 {
			curPos = 100 + curPos
			zerosPassed++
			if i > 0 && clicks[i-1].FinalPosition == 0 {
				zerosPassed--
			}
		}
		if curPos >= 100 {
			curPos = curPos % 100
			if curPos != 0 {
				zerosPassed++
				if i > 0 && clicks[i-1].FinalPosition == 0 {
					zerosPassed--
				}
			}
		}
		if curPos == 0 {
			zerosLanded++
		}
		clicks[i] = click{
			Direction:      string(line[0]) + " | " + strconv.Itoa(zerosLanded) + " | " + strconv.Itoa(zerosPassed),
			OriginalNumber: originalNum,
			Number:         num,
			FinalPosition:  int64(curPos),
		}
	}
	bytes, err := files.EncodeJsonTypeToBytes(clicks)
	if err != nil {
		log.Fatalf("failed to encode clicks: %v", err)
	}
	err = files.WriteFile("output/2025/day_01_p1.json", bytes)
	if err != nil {
		log.Fatalf("failed to write clicks: %v", err)
	}
	fmt.Printf("2025 Day 1 Part 1: %d\n", zerosLanded)
	fmt.Printf("2025 Day 1 Part 2: %d\n", zerosLanded+zerosPassed)
}
