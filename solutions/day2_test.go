package solutions

import (
	"fmt"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	inputs := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	expected := 150
	fmt.Println(">>> TestDay2Part1")
	result := Day2Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay2Part2(t *testing.T) {
	inputs := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	expected := 900
	fmt.Println(">>> TestDay2Part2")
	result := Day2Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}