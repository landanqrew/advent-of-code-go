package year2025

import "testing"


func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "empty string",
			s:        "",
			expected: "",
		},
		{
			name:     "single character",
			s:        "a",
			expected: "a",
		},
		{
			name:     "multiple characters",
			s:        "hello",
			expected: "olleh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseString(tt.s)
			if result != tt.expected {
				t.Fatalf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

/*
In 987654321111111, you can make the largest joltage possible, 98, by turning on the first two batteries.
In 811111111111119, you can make the largest joltage possible by turning on the batteries labeled 8 and 9, producing 89 jolts.
In 234234234234278, you can make 78 by turning on the last two batteries (marked 7 and 8).
In 818181911112111, the largest joltage you can produce is 92.	
*/

func TestFindHighestDoubleDigitNumberInString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1",
			s:        "987654321111111",
			expected: 98,
		},
		{
			name:     "example 2",
			s:        "811111111111119",
			expected: 89,
		},
		{
			name:     "example 3",
			s:        "234234234234278",
			expected: 78,
		},
		{
			name:     "example 4",
			s:        "818181911112111",
			expected: 92,
		},
		{
			name:     "example 5",
			s:        "818181913112111",
			expected: 93,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findHighestNumberInString(tt.s, 2)
			if result != tt.expected {
				t.Fatalf("expected '%d', got '%d'", tt.expected, result)
			}
		})
	}
}

/*
In 987654321111111, the largest joltage can be found by turning on everything except some 1s at the end to produce 987654321111.
In the digit sequence 811111111111119, the largest joltage can be found by turning on everything except some 1s, producing 811111111119.
In 234234234234278, the largest joltage can be found by turning on everything except a 2 battery, a 3 battery, and another 2 battery near the start to produce 434234234278.
In 818181911112111, the joltage 888911112111 is produced by turning on everything except some 1s near the front.

*/
func TestFindHighest12DigitNumberInString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1",
			s:        "987654321111111",
			expected: 987654321111,
		},
		{
			name:     "example 2",
			s:        "811111111111119",
			expected: 811111111119,
		},
		{
			name:     "example 3",
			s:        "234234234234278",
			expected: 434234234278,
		},
		{
			name:     "example 4",
			s:        "818181911112111",
			expected: 888911112111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findHighestNumberInString(tt.s, 12)
			if result != tt.expected {
				t.Fatalf("expected '%d', got '%d'", tt.expected, result)
			}
		})
	}
}
		
