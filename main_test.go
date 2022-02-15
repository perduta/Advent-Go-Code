package main

import (
	"fmt"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	// part 1 answer 1462
	inputs := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	fmt.Println(">>> TestDay1Part1")
	result := Day1Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay1Part2(t *testing.T) {
	// part 1 answer 1462
	inputs := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 5
	fmt.Println(">>> TestDay1Part2")
	result := Day1Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
