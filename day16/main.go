package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

var directions = map[rune]Point{
	'N': {0, -1},
	'E': {1, 0},
	'S': {0, 1},
	'W': {-1, 0},
}

type Point struct {
	x int
	y int
}

type VisitedPoint struct {
	point     Point
	direction Point
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func isValidPoint(x int, y int, lines []string) bool {
	return x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines)
}

func countVisited(visited map[VisitedPoint]bool) int {
	points := make(map[Point]bool)

	for point := range visited {
		points[point.point] = true
	}

	return len(points)
}

func explore(position Point, direction Point, lines []string, visited map[VisitedPoint]bool) map[VisitedPoint]bool {
	nextPosition := Point{position.x + direction.x, position.y + direction.y}
	if !isValidPoint(nextPosition.x, nextPosition.y, lines) {
		return visited
	}

	point := VisitedPoint{nextPosition, direction}
	if visited[point] {
		return visited
	}

	visited[point] = true

	switch lines[nextPosition.y][nextPosition.x] {
	case '|':
		if direction.x != 0 {
			visited = explore(nextPosition, directions['N'], lines, visited)
			visited = explore(nextPosition, directions['S'], lines, visited)
			return visited
		}
	case '-':
		if direction.y != 0 {
			visited = explore(nextPosition, directions['E'], lines, visited)
			visited = explore(nextPosition, directions['W'], lines, visited)
			return visited
		}
	case '/':
		direction = Point{-direction.y, -direction.x}
	case '\\':
		direction = Point{direction.y, direction.x}
	}

	return explore(nextPosition, direction, lines, visited)
}

func exploreWithCount(position Point, direction Point, lines []string) int {
	visited := make(map[VisitedPoint]bool)
	visited = explore(position, direction, lines, visited)

	return countVisited(visited)
}

func part1() int {
	lines := getInputLines()

	currentDirection := Point{1, 0}
	currentPosition := Point{-1, 0}

	return exploreWithCount(currentPosition, currentDirection, lines)
}

func part2() int {
	lines := getInputLines()

	max := 0
	for i := 0; i < len(lines[0]); i++ {
		// Go down
		currentDirection := Point{0, 1}
		currentPosition := Point{i, -1}

		count := exploreWithCount(currentPosition, currentDirection, lines)
		if count > max {
			max = count
		}

		// Go up
		currentDirection = Point{0, -1}
		currentPosition = Point{i, len(lines)}

		count = exploreWithCount(currentPosition, currentDirection, lines)
		if count > max {
			max = count
		}
	}

	for i := 0; i < len(lines); i++ {
		// Go right
		currentDirection := Point{1, 0}
		currentPosition := Point{-1, i}

		count := exploreWithCount(currentPosition, currentDirection, lines)
		if count > max {
			max = count
		}

		// Go left
		currentDirection = Point{-1, 0}
		currentPosition = Point{len(lines[0]), i}

		count = exploreWithCount(currentPosition, currentDirection, lines)
		if count > max {
			max = count
		}
	}

	return max
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
