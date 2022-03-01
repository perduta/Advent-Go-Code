package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func day9GetInputs() []string {
	return strings.Split(`2199943210
3987894921
9856789892
8767896789
9899965678`, "\n")
}

func TestDay9Part1(t *testing.T) {
	inputs := day9GetInputs()
	expected := 15
	fmt.Println(">>> " + t.Name())
	result := Day9Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay9Part2(t *testing.T) {
	inputs := day9GetInputs()
	expected := 1134
	fmt.Println(">>> " + t.Name())
	result := Day9Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}