package year2025

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/landanqrew/advent-of-code-go/internal/re"
)

type productRange struct {
	LowerBound int   `json:"lower_bound"`
	UpperBound int   `json:"upper_bound"`
	InvalidIDs []int `json:"invalid_ids"`
}

type numericRange struct {
	LowerBound int   `json:"lower_bound"`
	UpperBound int   `json:"upper_bound"`
	Divisors   []int `json:"divisors"`
	InvalidIDs []int `json:"invalid_ids"`
}

func getValidProductRanges(rangeString string) []productRange {
	productRanges := make([]productRange, 0)
	if rangeString == "" {
		rangeString = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	}
	ranges := strings.Split(rangeString, ",")
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		productRange := productRange{
			LowerBound: 0,
			UpperBound: 0,
			InvalidIDs: make([]int, 0),
		}
		for i, bound := range bounds {
			digits := re.GetNumbers(bound)[0]
			if len(digits) == 0 {
				log.Fatalf("no digits found in bound (%s)", bound)
			}
			boundInt, err := strconv.ParseInt(digits, 10, 64)
			if err != nil {
				log.Fatalf("failed to convert bound (%s) to int: %v", bound, err)
			}
			if i == 0 {
				productRange.LowerBound = int(boundInt)
			} else {
				productRange.UpperBound = int(boundInt)
			}
		}
		if productRange.LowerBound > productRange.UpperBound {
			log.Fatalf("lower bound is greater than upper bound: %v", productRange)
		}
		productRanges = append(productRanges, productRange)
	}
	return productRanges
}

func findInvalidProductIDsInRange(idRange productRange) productRange {
	lowerString := strconv.Itoa(idRange.LowerBound)
	upperString := strconv.Itoa(idRange.UpperBound)
	if len(lowerString) == len(upperString) && math.Mod(float64(len(lowerString)), 2) != 0 {
		return idRange
	}
	halfLength := len(lowerString) / 2
	cur := 0
	if halfLength != 0 {
		initialInt64, err := strconv.ParseInt(lowerString[:halfLength], 10, 64)
		if err != nil {
			log.Fatalf("failed to convert initial to int: %v", err)
		}
		cur = int(initialInt64)
	} else {
		cur = 1
	}
	// fmt.Println("initial: ", cur)
	for {
		// iterate so long as combined string is in bounds
		curString := strconv.Itoa(cur)
		combinedString := curString + curString
		if math.Mod(float64(len(combinedString)), 2) == 0 {
			evalInt64, err := strconv.ParseInt(combinedString, 10, 64)
			if err != nil {
				log.Fatalf("failed to convert combined string to int | upper bound: %d | lower bound: %d | combined string: %s | error: %v", idRange.UpperBound, idRange.LowerBound, combinedString, err)
			}
			evalInt := int(evalInt64)
			// fmt.Println("eval: ", evalInt)
			if idRange.LowerBound <= evalInt && evalInt <= idRange.UpperBound {
				// fmt.Println("adding to invalid IDs: ", evalInt)
				idRange.InvalidIDs = append(idRange.InvalidIDs, evalInt)
			}
			// break if evalInt is greater than upper bound
			if evalInt > idRange.UpperBound {
				break
			}
			cur++
		} else {
			// advance to next number with even length
			newCurString := "1"
			for range curString {
				newCurString = newCurString + "0"
			}
			curInt64, err := strconv.ParseInt(newCurString, 10, 64)
			if err != nil {
				log.Fatalf("failed to convert new cur string to int: %v", err)
			}
			cur = int(curInt64)
		}
	}
	return idRange
}

func getValidNumericRangesForProductRange(productRange productRange) []numericRange {
	if len(strconv.Itoa(productRange.LowerBound)) == len(strconv.Itoa(productRange.UpperBound)) {
		return []numericRange{numericRange{
			LowerBound: productRange.LowerBound,
			UpperBound: productRange.UpperBound,
			Divisors:   getDivisorsForStringLength(len(strconv.Itoa(productRange.LowerBound))),
			InvalidIDs: make([]int, 0),
		}}
	} else {
		numericRanges := make([]numericRange, 0)
		lowerStrLength := len(strconv.Itoa(productRange.LowerBound))
		numericRanges = append(numericRanges, numericRange{
			LowerBound: productRange.LowerBound,
			UpperBound: getHighestNumberForStringLength(lowerStrLength),
			Divisors:   getDivisorsForStringLength(lowerStrLength),
			InvalidIDs: make([]int, 0),
		})
		upperStrLength := len(strconv.Itoa(productRange.UpperBound))
		numericRanges = append(numericRanges, numericRange{
			LowerBound: getLowestNumberForStringLength(upperStrLength),
			UpperBound: productRange.UpperBound,
			Divisors:   getDivisorsForStringLength(upperStrLength),
			InvalidIDs: make([]int, 0),
		})
		return numericRanges
	}
}

func findInvalidProductIDsInNumericRange(numericRange numericRange) []int {
	invalidIDs := make([]int, 0)
	for _, divisor := range numericRange.Divisors {
		parts := divideString(strconv.Itoa(numericRange.LowerBound), divisor)
		partLength := len(parts[0])
		upperBoundString := strconv.Itoa(numericRange.UpperBound)
		upperBoundPartComp := upperBoundString[:partLength]
		upperBoundPartCompInt := convertStringToInt(upperBoundPartComp)
		cur := convertStringToInt(parts[0])
		for cur <= upperBoundPartCompInt {
			fullID := buildFullID(cur, divisor)
			if fullID >= numericRange.LowerBound && fullID <= numericRange.UpperBound {
				if !slices.Contains(invalidIDs, fullID) {
					invalidIDs = append(invalidIDs, fullID)
				}
			}
			cur++
		}
	}
	return invalidIDs
}

func getDivisorsForStringLength(length int) []int {
	divisors := make([]int, 0)
	for i := 2; i <= length; i++ {
		if math.Mod(float64(length), float64(i)) == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func divideString(s string, divisor int) []string {
	parts := make([]string, 0)
	partLength := len(s) / divisor
	if partLength == 0 {
		return parts
	}
	for i := 0; i < len(s); i += partLength {
		end := i + partLength
		if end > len(s) {
			end = len(s)
		}
		parts = append(parts, s[i:end])
	}
	return parts
}

func convertStringToInt(s string) int {
	evalInt64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("failed to convert string (%s) to int: %v", s, err)
	}
	return int(evalInt64)
}

func areStringPartsInBounds(base int, divisor int, lowerBound int, upperBound int) bool {
	evalString := ""
	for i := 0; i < divisor; i++ {
		evalString += strconv.Itoa(base)
	}
	evalInt := convertStringToInt(evalString)
	return evalInt >= lowerBound && evalInt <= upperBound
}

func buildFullID(base int, divisor int) int {
	baseString := strconv.Itoa(base)
	curString := baseString
	for i := 1; i < divisor; i++ {
		curString = curString + baseString
	}
	return convertStringToInt(curString)
}

func getLowestNumberForStringLength(length int) int {
	evalString := ""
	for i := 0; i < length; i++ {
		if i == 0 {
			evalString += "1"
		} else {
			evalString += "0"
		}
	}
	return convertStringToInt(evalString)
}

func getHighestNumberForStringLength(length int) int {
	evalString := ""
	for i := 0; i < length; i++ {
		evalString += "9"
	}
	return convertStringToInt(evalString)
}

func Day2Part1(data string) {
	productRanges := getValidProductRanges(data)
	totalInvalidSum := 0
	/*bytes, err := files.EncodeJsonTypeToBytes(productRanges)
	if err != nil {
		log.Fatalf("failed to encode product ranges: %v", err)
	}
	fmt.Println("product ranges: \n" + string(bytes) + "\n")*/
	for i, productRange := range productRanges {
		productRanges[i] = findInvalidProductIDsInRange(productRange)
		for _, invalidID := range productRanges[i].InvalidIDs {
			totalInvalidSum += invalidID
		}
	}

	/*bytes, err := files.EncodeJsonTypeToBytes(productRanges)
	if err != nil {
		log.Fatalf("failed to encode product ranges: %v", err)
	}
	fmt.Println("product ranges: \n" + string(bytes) + "\n")*/

	fmt.Println("total invalid sum: ", totalInvalidSum)
}

func Day2Part2(data string) {
	productRanges := getValidProductRanges(data)
	totalInvalidSum := 0
	for _, productRange := range productRanges {
		numericRanges := getValidNumericRangesForProductRange(productRange)
		for j, numericRange := range numericRanges {
			numericRanges[j].InvalidIDs = findInvalidProductIDsInNumericRange(numericRange)
			for _, invalidID := range numericRanges[j].InvalidIDs {
				totalInvalidSum += invalidID
			}
		}
	}
	fmt.Println("total invalid sum: ", totalInvalidSum)
}
