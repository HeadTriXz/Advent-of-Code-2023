package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

var directions = map[rune]Point{
	'N': {0, -1},
	'E': {1, 0},
	'S': {0, 1},
	'W': {-1, 0},
}

var pipes = map[rune][]Point{
	'S': {directions['N'], directions['E'], directions['S'], directions['W']},
	'|': {directions['N'], directions['S']},
	'-': {directions['W'], directions['E']},
	'J': {directions['N'], directions['W']},
	'L': {directions['N'], directions['E']},
	'F': {directions['S'], directions['E']},
	'7': {directions['S'], directions['W']},
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getStart(lines []string) *Point {
	for row, line := range lines {
		for col, char := range line {
			if char == 'S' {
				return &Point{col, row}
			}
		}
	}

	return nil
}

func isValid(x int, y int, lines []string) bool {
	return x >= 0 && y >= 0 && y < len(lines) && x < len(lines[y])
}

func getLoopNodes(start Point, lines []string) map[Point]bool {
	visited := make(map[Point]bool)
	visited[start] = true

	queue := []Point{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		allowedDirections := pipes[rune(lines[current.y][current.x])]
		for _, direction := range allowedDirections {
			next := Point{current.x + direction.x, current.y + direction.y}

			if isValid(next.x, next.y, lines) && !visited[next] && lines[next.y][next.x] != '.' {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}

	return visited
}

func part1() int {
	lines := getInputLines()
	start := getStart(lines)
	if start == nil {
		log.Fatal("Could not find start")
		return 0
	}

	nodes := getLoopNodes(*start, lines)
	return (len(nodes) + 1) / 2
}

func part2() int {
	lines := getInputLines()
	start := getStart(lines)
	if start == nil {
		log.Fatal("Could not find start")
		return 0
	}

	nodes := getLoopNodes(*start, lines)
	enclosed := 0

	for row, line := range lines {
		isInside := false
		next := ' '

		for col, char := range line {
			if nodes[Point{col, row}] {
				switch char {
				case '|':
					isInside = !isInside
				case 'L':
					next = '7'
				case 'F':
					next = 'J'
				case '7':
					if next == '7' {
						isInside = !isInside
					}

					next = ' '
				case 'J':
					if next == 'J' {
						isInside = !isInside
					}

					next = ' '
				}
			} else if isInside {
				enclosed++
			}
		}
	}

	return enclosed
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
