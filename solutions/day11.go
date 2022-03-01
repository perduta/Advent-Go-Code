package solutions

type OctopusFieldSimulator struct {
	field        [][]int
	hasFlashed   map[Point]bool
	flashesTotal int
}

func (o *OctopusFieldSimulator) getAdjacents(p Point) []Point {
	adjacents := []Point{}
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			adjP := Point{
				p.x + dx,
				p.y + dy,
			}
			if adjP.x >= 0 && adjP.x < len(o.field[0]) &&
				adjP.y >= 0 && adjP.y < len(o.field) &&
				o.field[adjP.y][adjP.x] >= 0 {
				adjacents = append(adjacents, adjP)
			}
		}
	}
	return adjacents
}
func (o *OctopusFieldSimulator) parseInputs(inputs []string) {
	for _, input := range inputs {
		row := []int{}
		for _, char := range input {
			row = append(row, int(char-'0'))
		}
		o.field = append(o.field, row)
	}
}

func (o *OctopusFieldSimulator) singleStep() bool {
	o.hasFlashed = map[Point]bool{}
	queue := make([]Point, 0)
	for y, row := range o.field {
		for x := range row {
			o.field[y][x]++
			if o.field[y][x] > 9 {
				queue = append(queue, Point{x, y})
			}
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		o.field[p.y][p.x]++
		if o.hasFlashed[p] {
			continue
		}
		if o.field[p.y][p.x] > 9 {
			o.hasFlashed[p] = true
			o.flashesTotal++
			queue = append(queue, o.getAdjacents(p)...)
			continue
		}
	}

	stepFlashCounter := 0
	for p := range o.hasFlashed {
		o.field[p.y][p.x] = 0
		stepFlashCounter++
	}
	return stepFlashCounter == 100
}

func Day11Part1(inputs []string) int {
	o := OctopusFieldSimulator{}
	o.parseInputs(inputs)

	for step := 0; step < 100; step++ {
		o.singleStep()
	}
	return o.flashesTotal
}

func Day11Part2(inputs []string) int {
	o := OctopusFieldSimulator{}
	o.parseInputs(inputs)

	i := 1
	for !o.singleStep() {
		i++
	}
	return i
}
