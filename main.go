package main

import (
	"aoc2021/solutions"
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	SOLUTION_FNS := map[int][]func([]string) int{
		1: {solutions.Day1Part1, solutions.Day1Part2},
		2: {solutions.Day2Part1, solutions.Day2Part2},
		3: {solutions.Day3Part1, solutions.Day3Part2},
		4: {solutions.Day4Part1, solutions.Day4Part2},
		5: {solutions.Day5Part1, solutions.Day5Part2},
		6: {solutions.Day6Part1, solutions.Day6Part2},
		7: {solutions.Day7Part1, solutions.Day7Part2},
		8: {solutions.Day8Part1, solutions.Day8Part2},
		9: {solutions.Day9Part1, solutions.Day9Part2},
		10: {solutions.Day10Part1, solutions.Day10Part2},
		11: {solutions.Day11Part1, solutions.Day11Part2},
	}

	files, err := ioutil.ReadDir("inputs/")
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		log.Fatal("No inputs found. Please add inputs of your puzzles to the inputs/ directory.")
	}
	for _, fileInfo := range files {
		n, err := strconv.Atoi(fileInfo.Name())
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Open("inputs/" + fileInfo.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		inputs := []string{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			inputs = append(inputs, scanner.Text())
		}
		for _, solutionFn := range SOLUTION_FNS[n] {
			log.Println("Day", n, "answer", solutionFn(inputs))
		}
	}
}
