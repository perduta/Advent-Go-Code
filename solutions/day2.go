package solutions

import (
	"log"
	"strconv"
	"strings"
)

type MoveCommand struct {
	Direction string
	Distance  int
}

func MoveCommandFromText(text string) MoveCommand {
	split := strings.Split(text, " ")
	if len(split) != 2 {
		log.Fatal("Invalid move command: ", text)
	}

	distance, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}

	return MoveCommand{split[0], distance}
}

func Day2Part1(inputs []string) int {
	horizontalPosition := 0
	depth := 0
	for _, input := range inputs {
		moveCommand := MoveCommandFromText(input)
		if moveCommand.Direction == "forward" {
			horizontalPosition += moveCommand.Distance
		} else if moveCommand.Direction == "up" {
			depth -= moveCommand.Distance
		} else if moveCommand.Direction == "down" {
			depth += moveCommand.Distance
		} else {
			log.Fatal("Invalid direction: ", moveCommand.Direction)
		}
	}
	return horizontalPosition * depth
}

func Day2Part2(inputs []string) int {
	horizontalPosition := 0
	depth := 0
	aim := 0
	for _, input := range inputs {
		moveCommand := MoveCommandFromText(input)
		if moveCommand.Direction == "forward" {
			horizontalPosition += moveCommand.Distance
			depth += aim * moveCommand.Distance
		} else if moveCommand.Direction == "up" {
			aim -= moveCommand.Distance
		} else if moveCommand.Direction == "down" {
			aim += moveCommand.Distance
		} else {
			log.Fatal("Invalid direction: ", moveCommand.Direction)
		}
	}
	return horizontalPosition * depth
}