package aoc

import (
	"fmt"
	"testing"
)

func TestGetSessionCookie(t *testing.T) {
	cookie, err := GetSessionCookie("../../.env")
	if err != nil {
		t.Fatalf("failed to get session cookie: %v", err)
	}
	fmt.Println(cookie)
}
func TestGetInput(t *testing.T) {
	input, err := GetInput(2025, 1, "../../.env")
	if err != nil {
		t.Fatalf("failed to get input: %v", err)
	}
	fmt.Println(string(input))
}

/*
func TestGetSolution(t *testing.T) {
	input, err := GetInput(2025, 1, "../../.env")
	if err != nil {
		t.Fatalf("failed to get input: %v", err)
	}
	solution, err := GetSolution(2025, 1, 1, string(input))
	if err != nil {
		t.Fatalf("failed to get solution: %v", err)
	}
	fmt.Println(solution)
}
*/