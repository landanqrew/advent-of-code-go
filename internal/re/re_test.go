package re

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetMatchedGroups(t *testing.T) {
	pattern := `\d+`
	input := `1234567890`
	groups := GetMatchedGroups(pattern, input)
	if len(groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(groups))
	}
}

func TestGetNumbers(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	numbers := GetNumbers(input)
	fmt.Println(numbers)
	if len(numbers) != 10 {
		t.Fatalf("expected 10 numbers, got %d", len(numbers))
	}
	if !slices.Equal(numbers, []string{"68", "30", "48", "5", "60", "55", "1", "99", "14", "82"}) {
		t.Fatalf("expected numbers to be [68, 30, 48, 5, 60, 55, 1, 99, 14, 82], got %v", numbers)
	}
}