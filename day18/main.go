package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var directions = map[rune]Node{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
	'0': {1, 0},
	'1': {0, 1},
	'2': {-1, 0},
	'3': {0, -1},
}

type Node struct {
	x int
	y int
}

type Trench struct {
	color Node
	node  Node
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func parseColor(color string) Node {
	distance, _ := strconv.ParseInt(color[:5], 16, 64)
	direction := directions[rune(color[5])]

	return Node{
		x: int(distance) * direction.x,
		y: int(distance) * direction.y,
	}
}

func parseInput() []Trench {
	lines := getInputLines()
	instructions := make([]Trench, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		direction := directions[rune(fields[0][0])]
		length, _ := strconv.Atoi(fields[1])

		instructions[i] = Trench{
			color: parseColor(fields[2][2 : len(fields[2])-1]),
			node: Node{
				x: direction.x * length,
				y: direction.y * length,
			},
		}
	}

	return instructions
}

func getCorners(instructions []Trench) ([]Node, int) {
	corners := make([]Node, len(instructions))
	current := Node{0, 0}
	sum := 0

	for i, instruction := range instructions {
		corners[i] = Node{
			x: current.x + instruction.node.x,
			y: current.y + instruction.node.y,
		}

		sum += abs(instruction.node.x + instruction.node.y)
		current = corners[i]
	}

	sum /= 2
	return corners, sum + 1
}

func getArea(corners []Node) int {
	size := len(corners)
	area := 0

	for i := 0; i < size; i++ {
		j := (i + 1) % size
		area += corners[i].x * corners[j].y
		area -= corners[j].x * corners[i].y
	}

	return abs(area) / 2
}

func part1() int {
	instructions := parseInput()
	corners, length := getCorners(instructions)
	area := getArea(corners)

	return area + length
}

func part2() int {
	instructions := parseInput()
	for i := range instructions {
		instructions[i].node = instructions[i].color
	}

	corners, length := getCorners(instructions)
	area := getArea(corners)

	return area + length
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
