package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func day6GetInputs() []string {
	return strings.Split(`3,4,3,1,2`, "\n")
}

func TestDay6Part1(t *testing.T) {
	inputs := day6GetInputs()
	expected := 5934
	fmt.Println(">>> " + t.Name())
	result := Day6Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay6Part2(t *testing.T) {
	inputs := day6GetInputs()
	expected := 26984457539
	fmt.Println(">>> " + t.Name())
	result := Day6Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}