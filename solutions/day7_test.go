package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func day7GetInputs() []string {
	return strings.Split(`16,1,2,0,4,2,7,1,2,14`, "\n")
}

func TestDay7Part1(t *testing.T) {
	inputs := day7GetInputs()
	expected := 37
	fmt.Println(">>> " + t.Name())
	result := Day7Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay7Part2(t *testing.T) {
	inputs := day7GetInputs()
	expected := 168
	fmt.Println(">>> " + t.Name())
	result := Day7Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}