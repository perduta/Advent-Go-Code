package main

import (
	"aoc2021/solutions"
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

var SOLUTIONS = [][]func([]string) int{
	{solutions.Day1Part1, solutions.Day1Part2},
	{solutions.Day2Part1, solutions.Day2Part2},
	{solutions.Day3Part1, solutions.Day3Part2},
	{solutions.Day4Part1, solutions.Day4Part2},
	{solutions.Day5Part1, solutions.Day5Part2},
	{solutions.Day6Part1, solutions.Day6Part2},
	{solutions.Day7Part1, solutions.Day7Part2},
	{solutions.Day8Part1, solutions.Day8Part2},
	{solutions.Day9Part1, solutions.Day9Part2},
	{solutions.Day10Part1, solutions.Day10Part2},
	{solutions.Day11Part1, solutions.Day11Part2},
}

func parseInput(filename string) ([]string, error) {
	file, err := os.Open("inputs/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	inputs := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs, nil
}

func main() {
	totalTime := time.Now()
	for i, fns := range SOLUTIONS {
		day := i + 1
		inputs, err := parseInput(strconv.Itoa(day))
		if err != nil {
			log.Printf("Could not parse inputs for day %v, skipping...", day)
			continue
		}
		for j, fn := range fns {
			now := time.Now()
			result := fn(inputs)
			elapsed := time.Since(now)
			log.Printf("Day %2v, part %v answer: %16v (took %v)", day, j+1, result, elapsed)
		}
	}
	log.Printf("Calculating all the answers took %v", time.Since(totalTime))
}
