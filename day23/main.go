package main

import (
	_ "embed"
	"log"
	"maps"
	"strings"
)

//go:embed input.txt
var input string

var directions = []Node{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var slopes = map[rune]Node{
	'^': directions[0],
	'>': directions[1],
	'v': directions[2],
	'<': directions[3],
}

type Node struct {
	x int
	y int
}

type PathItem struct {
	distance int
	node     Node
	visited  map[Node]bool
}

type TrailItem struct {
	distance int
	node     Node
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getStartEnd(lines []string) (Node, Node) {
	var start Node
	var end Node

	for i, char := range lines[0] {
		if char == '.' {
			start = Node{i, 0}
		}
	}

	lastIdx := len(lines) - 1
	for i, char := range lines[lastIdx] {
		if char == '.' {
			end = Node{i, lastIdx}
		}
	}

	return start, end
}

func getNeighbours(node Node, lines []string, withSlopes bool) []Node {
	if withSlopes {
		char := lines[node.y][node.x]
		for slope, direction := range slopes {
			if rune(char) == slope {
				return []Node{{
					x: node.x + direction.x,
					y: node.y + direction.y,
				}}
			}
		}
	}

	neighbours := []Node{}
	for _, direction := range directions {
		next := Node{node.x + direction.x, node.y + direction.y}
		if isValid(next, lines) && lines[next.y][next.x] != '#' {
			neighbours = append(neighbours, next)
		}
	}

	return neighbours
}

func isValid(node Node, lines []string) bool {
	return node.x >= 0 && node.y >= 0 && node.y < len(lines) && node.x < len(lines[node.y])
}

func calculatePathLength(edges map[Node][]Node, start Node, end Node) (int, Node) {
	count := 1
	for len(edges[end]) == 2 {
		count++
		var next Node
		for _, node := range edges[end] {
			if node != start {
				next = node
				break
			}
		}

		start = end
		end = next
	}

	return count, end
}

func getTrails(lines []string, withSlopes bool) map[Node][]TrailItem {
	edges := map[Node][]Node{}
	for y, line := range lines {
		for x, char := range line {
			node := Node{x, y}

			if char != '#' {
				edges[node] = append(edges[node], getNeighbours(node, lines, withSlopes)...)
			}
		}
	}

	newEdges := map[Node][]TrailItem{}
	for k, v := range edges {
		if len(edges) == 2 {
			continue
		}

		newEdges[k] = []TrailItem{}
		for _, edge := range v {
			count, end := calculatePathLength(edges, k, edge)
			newEdges[k] = append(newEdges[k], TrailItem{count, end})
		}
	}

	return newEdges
}

func getLongestPath(edges map[Node][]TrailItem, start Node, end Node) int {
	visited := map[Node]bool{}
	visited[start] = true

	queue := []PathItem{{
		distance: 0,
		node:     start,
		visited:  visited,
	}}

	maxDistance := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node == end {
			if current.distance > maxDistance {
				maxDistance = current.distance
			}

			continue
		}

		for _, edge := range edges[current.node] {
			if current.visited[edge.node] {
				continue
			}

			nextVisited := maps.Clone(current.visited)
			nextVisited[edge.node] = true

			queue = append(queue, PathItem{
				distance: current.distance + edge.distance,
				node:     edge.node,
				visited:  nextVisited,
			})
		}
	}

	return maxDistance
}

func part1() int {
	lines := getInputLines()
	trails := getTrails(lines, true)

	start, end := getStartEnd(lines)

	return getLongestPath(trails, start, end)
}

func part2() int {
	lines := getInputLines()
	trails := getTrails(lines, false)

	start, end := getStartEnd(lines)

	return getLongestPath(trails, start, end)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
