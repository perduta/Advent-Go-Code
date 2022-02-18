package solutions

import (
	"log"
	"strconv"
	"strings"
)

type Board struct {
	data [5][5]int
	matches [5][5]bool
	won bool
}

func ParseBoard(input []string) *Board {
	board := Board{}
	i, j := 0, 0
	for _, row := range input {
		for _, number := range strings.Split(row, " ") {
			if number == "" { continue }
			number, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			board.data[i][j] = number
			i++
		}
		j++
		i = 0
	}
	return &board
}

func (b *Board) GetBoardSum() (sum int) {
	for i, row := range b.data {
		for j, number := range row {
			if !b.matches[i][j] {
				sum += number
			}
		}
	}
	return sum
}

func (b *Board) ProcessNumber(n int) int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.data[i][j] == n {
				b.matches[i][j] = true
			}
		}
	}

	for i := 0; i < 5; i++ {
		if b.matches[i][0] && b.matches[i][1] && b.matches[i][2] && b.matches[i][3] && b.matches[i][4] {
			return b.GetBoardSum()
		} else if b.matches[0][i] && b.matches[1][i] && b.matches[2][i] && b.matches[3][i] && b.matches[4][i] {
			return b.GetBoardSum()
		}
	}

	return 0
}

func ParseInput(inputs [] string) (numbers []int, boards []*Board) {
	boardData := make([]string, 0)

	for i, input := range inputs {
		if strings.TrimSpace(input) == "" { continue }
		if i == 0 {
			for _, number := range strings.Split(input, ",") {
				number, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, number)
			}
			continue
		}

		boardData = append(boardData, input)

		if len(boardData) == 5 {
			boards = append(boards, ParseBoard(boardData))
			boardData = nil
		}

	}
	return
}

func Day4Part1(inputs []string) int {
	numbers, boards := ParseInput(inputs)

	for _, number := range numbers {
		for _, board := range boards {
			if score := board.ProcessNumber(number); score != 0 {
				return score * number
			}
		}
	}
	return 0
}

func Day4Part2(inputs []string) int {
	numbers, boards := ParseInput(inputs)
	winOrder := []struct {*Board; int}{}
	for _, number := range numbers {
		for _, board := range boards {
			if score := board.ProcessNumber(number); !board.won && score != 0 {
				board.won = true
				winOrder = append(winOrder, struct{*Board; int}{board, score * number})
			}
		}
	}
	return winOrder[len(winOrder) - 1].int
}