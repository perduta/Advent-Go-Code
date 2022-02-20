package solutions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Line struct {
	a, b Point
}

func (l Line) String() string {
	return fmt.Sprintf("Line{%d %d %d %d}", l.a.x, l.a.y, l.b.x, l.b.y)
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("{x%d/y%d}", p.x, p.y)
}

func LineFromString(s string) Line {
	parts := strings.Split(s, " -> ")
	a := strings.Split(parts[0], ",")
	b := strings.Split(parts[1], ",")
	ax, err := strconv.Atoi(a[0])
	if err != nil {
		log.Fatal(err)
	}
	ay, err := strconv.Atoi(a[1])
	if err != nil {
		log.Fatal(err)
	}
	bx, err := strconv.Atoi(b[0])
	if err != nil {
		log.Fatal(err)
	}
	by, err := strconv.Atoi(b[1])
	if err != nil {
		log.Fatal(err)
	}
	return Line{Point{ax, ay}, Point{bx, by}}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (l Line) GetPoints(drawDiagonals bool) []Point {
	points := []Point{}
	if l.a.x == l.b.x { // vertical line
		min := Min(l.a.y, l.b.y)
		max := Max(l.a.y, l.b.y)
		for i := 0; i <= max-min; i++ {
			points = append(points, Point{l.a.x, min + i})
		}
	} else if l.a.y == l.b.y { // horizontal line
		min := Min(l.a.x, l.b.x)
		max := Max(l.a.x, l.b.x)
		for i := 0; i <= max-min; i++ {
			points = append(points, Point{min + i, l.a.y})
		}
	} else {
		if !drawDiagonals {
			// line nor vertical nor horizontal, ignoring...
			return points
		}
		x := l.b.x - l.a.x
		y := l.b.y - l.a.y
		for i := 0; i <= Abs(x); i++ {
			px, py := 0, 0
			if x < 0 {
				py = Max(l.a.y, l.b.y) - i
			} else if x > 0 {
				py = Min(l.a.y, l.b.y) + i
			}
			if y < 0 {
				px = Max(l.a.x, l.b.x) - i
			} else if y > 0 {
				px = Min(l.a.x, l.b.x) + i
			}
			points = append(points, Point{px, py})
		}
	}
	return points

}

func Day5Common(inputs []string, useDiagonals bool) int {
	lines := make([]Line, len(inputs))
	for i, input := range inputs {
		lines[i] = LineFromString(input)
	}

	points := map[Point]int{}
	for _, line := range lines {
		for _, p := range line.GetPoints(useDiagonals) {
			points[p]++
		}
	}

	collisions := 0
	for _, count := range points {
		if count > 1 {
			collisions++
		}
	}
	return collisions
}

func Day5Part1(inputs []string) int {
	return Day5Common(inputs, false)
}

func Day5Part2(inputs []string) int {
	return Day5Common(inputs, true)
}
