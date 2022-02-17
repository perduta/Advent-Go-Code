package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"aoc2021/solutions"
)

func main() {
	SOLUTION_FNS := map[int][]func([]string) int {
		1: {solutions.Day1Part1, solutions.Day1Part2},
		2: {solutions.Day2Part1, solutions.Day2Part2},
		3: {solutions.Day3Part1, solutions.Day3Part2},
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
			log.Fatal(err);
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
