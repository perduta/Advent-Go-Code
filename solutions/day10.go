package solutions

import (
	"errors"
	"fmt"
	"sort"
)

var LBRACKETS_MAP = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
	'<': '>',
}
var RBRACKETS_MAP = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
	'>': '<',
}

var SYNTAX_ERROR_SCORE_TABLE = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var AUTOCOMPLETE_SCORE_TABLE = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

type IllegalBracketError struct {
	bracket rune
}

func (e IllegalBracketError) Error() string {
	return fmt.Sprintf("Illegal bracket: %c", e.bracket)
}

func parseLine(input string) (stack []rune, err error) {
	for _, char := range input {
		if LBRACKETS_MAP[char] != 0 {
			stack = append(stack, char)
		} else if RBRACKETS_MAP[char] != 0 {
			if stack[len(stack)-1] == RBRACKETS_MAP[char] {
				stack = stack[:len(stack)-1]
			} else {
				return stack, IllegalBracketError{char}
			}
		} else {
			return stack, errors.New("Unexpected character: " + string(char))
		}
	}
	return stack, nil
}

func Day10Part1(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		_, err := parseLine(input)
		if err == nil {
			continue
		}
		if err, ok := err.(IllegalBracketError); ok {
			sum += SYNTAX_ERROR_SCORE_TABLE[err.bracket]
		}
	}
	return sum
}

func Day10Part2(inputs []string) int {
	scores := []int{}
	for _, input := range inputs {
		stack, err := parseLine(input)
		if err != nil {
			continue
		}
		score := 0
		for len(stack) > 0 {
			score *= 5
			score += AUTOCOMPLETE_SCORE_TABLE[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
