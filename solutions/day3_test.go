package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func getInputs() []string {
	return strings.Split(
`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`, "\n")
}

func TestBin2Dec(t *testing.T) {
	inputs := [][]int{
		{1},
		{1, 1},
		{1, 1, 1},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0},
	}
	expected := []int{1, 3, 7, 22, 42}
	for i, input := range inputs {
		result := Bin2Dec(input)
		if result != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], result)
		}
	}
}

func TestDay3Part1(t *testing.T) {
	inputs := getInputs()
	expected := 198
	fmt.Println(">>> TestDay2Part1")
	result := Day3Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
