package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(">>> Well hello there")
	file, err := os.Open("inputs/1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputs := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, input)
	}

	log.Println("Day1 answer", Day1Part1(inputs))
	log.Println("Day1 answer part 2", Day1Part2(inputs))
}

func Day1Part1(inputs []int) int {
	previousDepth := 0
	totalIncreases := -1
	for _, input := range inputs {
		if input > previousDepth {
			totalIncreases++
		}
		previousDepth = input
	}
	return totalIncreases
}

func sum(inputs []int) int {
	// Returns the sum of all elements in the list
	sum := 0
	for _, input := range inputs {
		sum += input
	}
	return sum
}

func Day1Part2(inputs []int) int {
	const WINDOW_SIZE = 3
	previousDepth := 0
	totalIncreases := -1
	for i := 0; i <= len(inputs)-WINDOW_SIZE; i++ {
		window := inputs[i : i+WINDOW_SIZE]
		input := sum(window)
		if input > previousDepth {
			totalIncreases++
		}
		previousDepth = input
	}
	return totalIncreases
}
