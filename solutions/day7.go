package solutions

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func Median(input []int) int {
	sort.Ints(input)
	return input[len(input)/2]

}

func Day7Part1(inputs []string) int {
	crabs := []int{}
	for _, input := range strings.Split(inputs[0], ",") {
		n, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		crabs = append(crabs, n)
	}

	median := Median(crabs)
	totalFuel := 0
	for _, n := range crabs {
		totalFuel += Abs(n - median)
	}
	return totalFuel
}

func Day7Part2FuelCost(crabs []int, pos int) int {
	totalFuel := 0
	for _, n := range crabs {
		sum := 0
		for i := 1; i <= Abs(n-pos); i++ {
			sum += i
		}
		totalFuel += sum
	}
	return totalFuel
}

func Day7Part2(inputs []string) int {
	crabs := []int{}
	for _, input := range strings.Split(inputs[0], ",") {
		n, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		crabs = append(crabs, n)
	}

	median := Median(crabs)
	bestFuelCost := 0
	for i := 1; i <= median + 1; i++ {
		fuelCost := Min(Day7Part2FuelCost(crabs, median - i), Day7Part2FuelCost(crabs, median + i))
		if fuelCost < bestFuelCost || bestFuelCost == 0 {
			bestFuelCost = fuelCost
		}
	}
	return bestFuelCost
}
