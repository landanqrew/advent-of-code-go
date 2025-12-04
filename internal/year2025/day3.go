package year2025

import (
	"fmt"
	"strconv"
	"strings"
)

func getDay3Part1Example() string {
	return `987654321111111
811111111111119
234234234234278
818181911112111`
}

type bank struct {
	Digits []int `json:"digits"`
	Length int `json:"length"`
}

func (b *bank) recursiveUpdate(num int, k int) {
	// fmt.Println("BEFORE (num: ", num, "k: ", k, "): ")
	// files.PrintJsonType[bank](*b)
	if k == b.Length - 1 {
		if num > b.Digits[k] {
			b.Digits[k] = num
		}
	} else if k < b.Length {
		if num > b.Digits[k] {
			newNum := b.Digits[k]
			b.Digits[k] = num
			b.recursiveUpdate(newNum, k + 1)
		} else if num == b.Digits[k] && k + 1 < b.Length {
			b.recursiveUpdate(num, k + 1)
		}
	}
	// fmt.Println("AFTER: ")
	// files.PrintJsonType[bank](*b)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// order of the string has to stay in tact
func findHighestNumberInString(s string, digitLength int) int {
	// fmt.Println("findHighestNumberInString (s: ", s, "digitLength: ", digitLength, "): ")
	bank := bank{
		Digits: make([]int, digitLength),
		Length: digitLength,
	}
	reversedString := reverseString(s)
	for i, char := range reversedString {
		num := convertStringToInt(string(char))
		if i < digitLength {
			bank.Digits[digitLength - i - 1] = num
		} else {
			bank.recursiveUpdate(num, 0)
		}
	}
	bankString := ""
	for _, digit := range bank.Digits {
		bankString += strconv.Itoa(digit)
	}
	return convertStringToInt(bankString)
}

func Day3Part1(data string) {
	if data == "" {
		data = getDay3Part1Example()
	}
	lines := strings.Split(data, "\n")
	total := 0
	for _, line := range lines {
		total += findHighestNumberInString(line, 2)
	}

	fmt.Println("total: ", total)
}

func Day3Part2(data string) {
	if data == "" {
		data = getDay3Part1Example()
	}
	lines := strings.Split(data, "\n")
	total := 0
	for _, line := range lines {
		total += findHighestNumberInString(line, 12)
	}

	fmt.Println("total: ", total)
}
