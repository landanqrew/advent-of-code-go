package year2025

import (
	"reflect"
	"slices"
	"testing"
)

func TestGetDivisorsForStringLength(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		expected []int
	}{
		{
			name:     "length 1",
			length:   1,
			expected: []int{},
		},
		{
			name:     "length 2",
			length:   2,
			expected: []int{2},
		},
		{
			name:     "length 4",
			length:   4,
			expected: []int{2, 4},
		},
		{
			name:     "length 6",
			length:   6,
			expected: []int{2, 3, 6},
		},
		{
			name:     "length 8",
			length:   8,
			expected: []int{2, 4, 8},
		},
		{
			name:     "length 12",
			length:   12,
			expected: []int{2, 3, 4, 6, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDivisorsForStringLength(tt.length)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("getDivisorsForStringLength(%d) = %v, want %v", tt.length, result, tt.expected)
			}
		})
	}
}

func TestDivideString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		divisor  int
		expected []string
	}{
		{
			name:     "divide 6 chars by 2",
			s:        "123456",
			divisor:  2,
			expected: []string{"123", "456"},
		},
		{
			name:     "divide 6 chars by 3",
			s:        "123456",
			divisor:  3,
			expected: []string{"12", "34", "56"},
		},
		{
			name:     "divide 6 chars by 6",
			s:        "123456",
			divisor:  6,
			expected: []string{"1", "2", "3", "4", "5", "6"},
		},
		{
			name:     "divide 4 chars by 2",
			s:        "1234",
			divisor:  2,
			expected: []string{"12", "34"},
		},
		{
			name:     "divide 4 chars by 4",
			s:        "1234",
			divisor:  4,
			expected: []string{"1", "2", "3", "4"},
		},
		{
			name:     "divide 8 chars by 4",
			s:        "12345678",
			divisor:  4,
			expected: []string{"12", "34", "56", "78"},
		},
		{
			name:     "divide 2 chars by 2",
			s:        "12",
			divisor:  2,
			expected: []string{"1", "2"},
		},
		{
			name:     "divisor larger than length",
			s:        "12",
			divisor:  5,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := divideString(tt.s, tt.divisor)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("divideString(%q, %d) = %v, want %v", tt.s, tt.divisor, result, tt.expected)
			}
		})
	}
}

func TestConvertStringToInt(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "single digit",
			s:        "5",
			expected: 5,
		},
		{
			name:     "two digits",
			s:        "42",
			expected: 42,
		},
		{
			name:     "large number",
			s:        "123456789",
			expected: 123456789,
		},
		{
			name:     "zero",
			s:        "0",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertStringToInt(tt.s)
			if result != tt.expected {
				t.Errorf("convertStringToInt(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}

func TestGetLowestNumberForStringLength(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		expected int
	}{
		{
			name:     "length 1",
			length:   1,
			expected: 1,
		},
		{
			name:     "length 2",
			length:   2,
			expected: 10,
		},
		{
			name:     "length 3",
			length:   3,
			expected: 100,
		},
		{
			name:     "length 4",
			length:   4,
			expected: 1000,
		},
		{
			name:     "length 5",
			length:   5,
			expected: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getLowestNumberForStringLength(tt.length)
			if result != tt.expected {
				t.Errorf("getLowestNumberForStringLength(%d) = %d, want %d", tt.length, result, tt.expected)
			}
		})
	}
}

func TestGetHighestNumberForStringLength(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		expected int
	}{
		{
			name:     "length 1",
			length:   1,
			expected: 9,
		},
		{
			name:     "length 2",
			length:   2,
			expected: 99,
		},
		{
			name:     "length 3",
			length:   3,
			expected: 999,
		},
		{
			name:     "length 4",
			length:   4,
			expected: 9999,
		},
		{
			name:     "length 5",
			length:   5,
			expected: 99999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getHighestNumberForStringLength(tt.length)
			if result != tt.expected {
				t.Errorf("getHighestNumberForStringLength(%d) = %d, want %d", tt.length, result, tt.expected)
			}
		})
	}
}

func TestAreStringPartsInBounds(t *testing.T) {
	tests := []struct {
		name       string
		base       int
		divisor    int
		lowerBound int
		upperBound int
		expected   bool
	}{
		{
			name:       "base 1, divisor 2, creates 11",
			base:       1,
			divisor:    2,
			lowerBound: 10,
			upperBound: 20,
			expected:   true, 
		},
		{
			name:       "base 11, divisor 2, 1111",
			base:       11,
			divisor:    2,
			lowerBound: 1000,
			upperBound: 2000,
			expected:   true, 
		},
		{
			name:       "base 12, divisor 2, creates 1212",
			base:       12,
			divisor:    2,
			lowerBound: 1000,
			upperBound: 2000,
			expected:   true, 
		},
		{
			name:       "base 100, divisor 2, builds 100100",
			base:       100,
			divisor:    2,
			lowerBound: 100000,
			upperBound: 200000,
			expected:   true,
		},
		{
			name:       "base 100, divisor 2 builds 100100, upper bound too low",
			base:       100,
			divisor:    2,
			lowerBound: 1000,
			upperBound: 9999,
			expected:   false, // "100100" > 1000
		},
		{
			name:       "base 5, divisor 3, creates 555",
			base:       5,
			divisor:    3,
			lowerBound: 100,
			upperBound: 999,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := areStringPartsInBounds(tt.base, tt.divisor, tt.lowerBound, tt.upperBound)
			if result != tt.expected {
				t.Errorf("areStringPartsInBounds(%d, %d, %d, %d) = %v, want %v",
					tt.base, tt.divisor, tt.lowerBound, tt.upperBound, result, tt.expected)
			}
		})
	}
}

func TestGetValidNumericRangesForProductRange(t *testing.T) {
	tests := []struct {
		name         string
		productRange productRange
		expected     []numericRange
	}{
		{
			name: "same length bounds",
			productRange: productRange{
				LowerBound: 100,
				UpperBound: 999,
			},
			expected: []numericRange{
				{
					LowerBound: 100,
					UpperBound: 999,
					Divisors:   []int{3}, // divisors of 3 (length of "100")
					InvalidIDs: []int{},
				},
			},
		},
		{
			name: "different length bounds",
			productRange: productRange{
				LowerBound: 95,
				UpperBound: 115,
			},
			expected: []numericRange{
				{
					LowerBound: 95,
					UpperBound: 99,       // highest 2-digit number
					Divisors:   []int{2}, // divisors of 2
					InvalidIDs: []int{},
				},
				{
					LowerBound: 100, // lowest 3-digit number
					UpperBound: 115,
					Divisors:   []int{3}, // divisors of 3
					InvalidIDs: []int{},
				},
			},
		},
		{
			name: "single digit range",
			productRange: productRange{
				LowerBound: 1,
				UpperBound: 9,
			},
			expected: []numericRange{
				{
					LowerBound: 1,
					UpperBound: 9,
					Divisors:   []int{}, // divisors of 1
					InvalidIDs: []int{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getValidNumericRangesForProductRange(tt.productRange)
			if len(result) != len(tt.expected) {
				t.Errorf("getValidNumericRangesForProductRange() returned %d ranges, want %d",
					len(result), len(tt.expected))
				return
			}
			for i, r := range result {
				if r.LowerBound != tt.expected[i].LowerBound ||
					r.UpperBound != tt.expected[i].UpperBound {
					t.Errorf("range[%d]: got LowerBound=%d UpperBound=%d, want LowerBound=%d UpperBound=%d",
						i, r.LowerBound, r.UpperBound, tt.expected[i].LowerBound, tt.expected[i].UpperBound)
				}
				if !reflect.DeepEqual(r.Divisors, tt.expected[i].Divisors) {
					t.Errorf("range[%d]: Divisors = %v, want %v", i, r.Divisors, tt.expected[i].Divisors)
				}
			}
		})
	}
}

/*
11-22 still has two invalid IDs, 11 and 22.
95-115 now has two invalid IDs, 99 and 111.
998-1012 now has two invalid IDs, 999 and 1010.
1188511880-1188511890 still has one invalid ID, 1188511885.
222220-222224 still has one invalid ID, 222222.
1698522-1698528 still contains no invalid IDs.
446443-446449 still has one invalid ID, 446446.
38593856-38593862 still has one invalid ID, 38593859.
565653-565659 now has one invalid ID, 565656.
824824821-824824827 now has one invalid ID, 824824824.
2121212118-2121212124 now has one invalid ID, 2121212121.
*/

func TestFindInvalidProductIDsInNumericRange(t *testing.T) {
	tests := []struct {
		name         string
		numericRange numericRange
		expected     []int
	}{
		{
			name: "range 11-22 with divisor 2",
			numericRange: numericRange{
				LowerBound: 11,
				UpperBound: 22,
				Divisors:   []int{2},
				InvalidIDs: []int{},
			},
			expected: []int{11, 22},
		},
		{
			name: "range 95-115 with divisor 2 (part 1)",
			numericRange: numericRange{
				LowerBound: 95,
				UpperBound: 99,
				Divisors:   []int{2},
				InvalidIDs: []int{},
			},
			expected: []int{99},
		},
		{
			name: "range 95-115 with divisor 2 (part 2)",
			numericRange: numericRange{
				LowerBound: 100,
				UpperBound: 115,
				Divisors:   []int{3},
				InvalidIDs: []int{},
			},
			expected: []int{111},
		},
		{
			name: "range 998-1012 with divisor 2 (part 1)",
			numericRange: numericRange{
				LowerBound: 998,
				UpperBound: 999,
				Divisors:   []int{3},
				InvalidIDs: []int{},
			},
			expected: []int{999},
		},
		{
			name: "range 998-1012 with divisor 2 (part 2)",
			numericRange: numericRange{
				LowerBound: 1000,
				UpperBound: 1012,
				Divisors:   []int{2, 4},
				InvalidIDs: []int{},
			},
			expected: []int{1010},
		},
		{
			name: "range 1188511880-1188511890 with divisor 2, 5, 10",
			numericRange: numericRange{
				LowerBound: 1188511880,
				UpperBound: 1188511889,
				Divisors:   []int{2, 5, 10},
				InvalidIDs: []int{},
			},
			expected: []int{1188511885},
		},
		{
			name: "range 222220-222224 with divisor 2, 3, 6",
			numericRange: numericRange{
				LowerBound: 222220,
				UpperBound: 222224,
				Divisors:   []int{2, 3, 6},
				InvalidIDs: []int{},
			},
			expected: []int{222222},
		},
		{
			name: "range 1698522-1698528 with divisor 2, 3, 6",
			numericRange: numericRange{
				LowerBound: 1698522,
				UpperBound: 1698528,
				Divisors:   []int{2, 3, 6},
				InvalidIDs: []int{},
			},
			expected: []int{},
		},
		{
			name: "range 446443-446449 with divisor 2, 3, 6",
			numericRange: numericRange{
				LowerBound: 446443,
				UpperBound: 446449,
				Divisors:   []int{2, 3, 6},
				InvalidIDs: []int{},
			},
			expected: []int{446446},
		},
		{
			name: "range 38593856-38593862 with divisor 2, 4, 8",
			numericRange: numericRange{
				LowerBound: 38593856,
				UpperBound: 38593862,
				Divisors:   []int{2, 4, 8},
				InvalidIDs: []int{},
			},
			expected: []int{38593859},
		},
		{
			name: "range 565653-565659 with divisor 2, 3, 6",
			numericRange: numericRange{
				LowerBound: 565653,
				UpperBound: 565659,
				Divisors:   []int{2, 3, 6},
				InvalidIDs: []int{},
			},
			expected: []int{565656},
		},
		{
			name: "range 824824821-824824827 with divisor 3, 9",
			numericRange: numericRange{
				LowerBound: 824824821,
				UpperBound: 824824827,
				Divisors:   []int{3, 9},
				InvalidIDs: []int{},
			},
			expected: []int{824824824},
		},
		{
			name: "range 2121212118-2121212124 with divisor 2, 5, 10",
			numericRange: numericRange{
				LowerBound: 2121212118,
				UpperBound: 2121212124,
				Divisors:   []int{2, 5, 10},
				InvalidIDs: []int{},
			},
			expected: []int{2121212121},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findInvalidProductIDsInNumericRange(tt.numericRange)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("findInvalidProductIDsInNumericRange() returned %v, want %v", result, tt.expected)
			}
			// Check that all returned IDs are within bounds
			for _, id := range result {
				if id < tt.numericRange.LowerBound || id > tt.numericRange.UpperBound {
					t.Errorf("findInvalidProductIDsInNumericRange() returned ID %d outside bounds [%d, %d]",
						id, tt.numericRange.LowerBound, tt.numericRange.UpperBound)
				}
			}
		})
	}
}
