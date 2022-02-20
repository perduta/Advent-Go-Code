package solutions

import (
	"log"
	"strconv"
	"strings"
)

func Day6Common(inputs []string, daysToSimulate int) int {
	const FISH_TIMER_RESET = 6
	const FISH_TIMER_NEWBORN = 8
	fishes := make([]int, FISH_TIMER_NEWBORN + 1)
	for _, input := range strings.Split(inputs[0], ",") {
		n, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fishes[n]++
	}

	for day := 0; day < daysToSimulate; day++ {
		newFishes := make([]int, FISH_TIMER_NEWBORN + 1)
		for i := 0; i < FISH_TIMER_NEWBORN; i++ {
			newFishes[i] = fishes[i + 1]
		}
		newFishes[FISH_TIMER_NEWBORN] = fishes[0]
		newFishes[FISH_TIMER_RESET] += fishes[0]
		fishes = newFishes
	}
	sum := 0
	for _, n := range fishes {
		sum += n
	}
	return sum
}

func Day6Part1(inputs []string) int {
	return Day6Common(inputs, 80)
}

func Day6Part2(inputs []string) int {
	return Day6Common(inputs, 256)
}