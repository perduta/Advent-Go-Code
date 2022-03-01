package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func day10GetInputs() []string {
	return strings.Split(`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`, "\n")
}

func TestDay10Part1(t *testing.T) {
	inputs := day10GetInputs()
	expected := 26397
	fmt.Println(">>> " + t.Name())
	result := Day10Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay10Part2(t *testing.T) {
	inputs := day10GetInputs()
	expected := 288957
	fmt.Println(">>> " + t.Name())
	result := Day10Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
