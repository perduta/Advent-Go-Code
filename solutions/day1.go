package solutions

import (
	"log"
	"strconv"
)

func Day1Part1(inputs []string) int {
	previousDepth := 0
	totalIncreases := -1
	for _, input := range inputs {
		input, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if input > previousDepth {
			totalIncreases++
		}
		previousDepth = input
	}
	return totalIncreases
}

func sumStrings(inputs []string) int {
	// Returns the sum of all elements in the list
	sum := 0
	for _, input := range inputs {
		input, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		sum += input
	}
	return sum
}

func Day1Part2(inputs []string) int {
	const WINDOW_SIZE = 3
	previousDepth := 0
	totalIncreases := -1
	for i := 0; i <= len(inputs)-WINDOW_SIZE; i++ {
		window := inputs[i : i+WINDOW_SIZE]
		input := sumStrings(window)
		if input > previousDepth {
			totalIncreases++
		}
		previousDepth = input
	}
	return totalIncreases
}
