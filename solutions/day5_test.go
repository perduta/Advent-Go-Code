package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func day5GetInputs() []string {
	return strings.Split(
		`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`, "\n")
}

func TestDay5Part1(t *testing.T) {
	inputs := day5GetInputs()
	expected := 5
	fmt.Println(">>> TestDay5Part1")
	result := Day5Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay5Part2(t *testing.T) {
	inputs := day5GetInputs()
	expected := 12
	fmt.Println(">>> TestDay5Part2")
	result := Day5Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDrawLine(t *testing.T) {
	type Case struct {
		line    Line
		points []Point
	}
	cases := []Case{
		{Line{Point{0, 0}, Point{4, 0}}, []Point{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
			{4, 0},
		}},

		{Line{Point{0, 0}, Point{0, 4}}, []Point{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
			{0, 4},
		}},

		{Line{Point{9, 4}, Point{3, 4}}, []Point{
			{3, 4},
			{4, 4},
			{5, 4},
			{6, 4},
			{7, 4},
			{8, 4},
			{9, 4},
		}},

		{Line{Point{8, 0}, Point{0, 8}}, []Point{
			{0, 8},
			{1, 7},
			{2, 6},
			{3, 5},
			{4, 4},
			{5, 3},
			{6, 2},
			{7, 1},
			{8, 0},
		}},	
		
		{Line{Point{9, 1}, Point{1, 9}}, []Point{
			{1, 9},
			{2, 8},
			{3, 7},
			{4, 6},
			{5, 5},
			{6, 4},
			{7, 3},
			{8, 2},
			{9, 1},
		}},

		{Line{Point{0, 0}, Point{8, 8}}, []Point{
			{0, 0},
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4},
			{5, 5},
			{6, 6},
			{7, 7},
			{8, 8},
		}},

		{Line{Point{1, 1}, Point{9, 9}}, []Point{
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4},
			{5, 5},
			{6, 6},
			{7, 7},
			{8, 8},
			{9, 9},
		}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("TestDrawLine %s", c.line), func(t *testing.T) {
			result := c.line.GetPoints(true)

			if len(result) != len(c.points) {
				t.Errorf("Result len() == %d, while expected: %d\n", len(result), len(c.points))
			}

			for i := 0; i < len(result); i++ {
				if result[i] != c.points[i] {
					t.Errorf("At index %d result is %s, while expected %s\n", i, result, c.points)
				}
			}
		})
	}
}
