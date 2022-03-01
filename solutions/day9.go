package solutions

import (
	"log"
	"sort"
	"strconv"
)

func getField(inputs []string) (field [][]int) {
	for _, input := range inputs {
		row := []int{}
		for _, c := range input {
			c, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, c)
		}
		field = append(field, row)
	}
	return
}

func getNeighbours(field [][]int, p Point) (neighbours []Point) {
	// Get all four adjacent points
	adjacentPoints := []Point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
	for _, p := range adjacentPoints {
		if p.x >= 0 && p.x < len(field[0]) && p.y >= 0 && p.y < len(field) {
			neighbours = append(neighbours, Point{p.x, p.y})
		}
	}
	return
}

func getLowestPoints(field [][]int) (lowestPoints []Point) {
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			neighbours := getNeighbours(field, Point{x, y})
			isLowestPoint := true
			for _, n := range neighbours {
				if field[n.y][n.x] < field[y][x] {
					isLowestPoint = false
					break
				}
			}
			if isLowestPoint {
				lowestPoints = append(lowestPoints, Point{x, y})
			}
		}
	}
	return
}

func Day9Part1(inputs []string) int {
	field := getField(inputs)

	riskLevelSum := 0
	for _, p := range getLowestPoints(field) {
		riskLevelSum += 1 + field[p.y][p.x]
	}
	return riskLevelSum
}
func Day9Part2(inputs []string) int {
	field := getField(inputs)
	basinSizes := []int{}
	for _, p := range getLowestPoints(field) {
		basin := getBasin(field, p)
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Ints(basinSizes)
	ret := 1
	for _, bs := range basinSizes[len(basinSizes)-3:] {
		ret *= bs
	}
	return ret
}

func getBasin(field [][]int, p Point) (basin []Point) {
	visited := make(map[Point]bool)
	queue := []Point{p}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if visited[p] {
			continue
		}
		visited[p] = true
		basin = append(basin, p)

		for _, n := range getNeighbours(field, p) {
			BASIN_BORDER := 9
			if field[n.y][n.x] < BASIN_BORDER {
				queue = append(queue, n)
			}
		}
	}
	return
}
