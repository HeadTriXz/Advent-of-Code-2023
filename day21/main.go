package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

var directions = []Node{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type Node struct {
	x int
	y int
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getStart(lines []string) *Node {
	for row, line := range lines {
		for col, char := range line {
			if char == 'S' {
				return &Node{col, row}
			}
		}
	}

	return nil
}

func isValid(x int, y int, lines []string) bool {
	return x >= 0 && y >= 0 && y < len(lines) && x < len(lines[y])
}

func mod(n int, m int) int {
	r := n % m
	if r < 0 {
		return r + m
	}

	return r
}

type Queue []Node

func (q *Queue) push(n Node) {
	*q = append(*q, n)
}

func (q *Queue) pop() Node {
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

func (q *Queue) len() int {
	return len(*q)
}

func (q *Queue) isEmpty() bool {
	return q.len() == 0
}

func (q *Queue) process(lines []string, borderless bool) {
	newQueue := Queue{}
	visited := make(map[Node]bool)

	for !q.isEmpty() {
		current := q.pop()

		for _, direction := range directions {
			next := Node{current.x + direction.x, current.y + direction.y}
			if (!borderless && !isValid(next.x, next.y, lines)) || visited[next] {
				continue
			}

			y := mod(next.y, len(lines))
			x := mod(next.x, len(lines[y]))

			if lines[y][x] == '#' {
				continue
			}

			newQueue.push(next)
			visited[next] = true
		}
	}

	*q = newQueue
}

func extrapolate(arr []int, n int) int {
	b0 := arr[0]
	b1 := arr[1] - arr[0]
	b2 := arr[2] - arr[1]

	return b0 + (b1 * n) + (n*(n-1)/2)*(b2-b1)
}

func part1() int {
	maxSteps := 64

	lines := getInputLines()
	start := getStart(lines)

	queue := Queue{*start}
	for i := 0; i < maxSteps; i++ {
		queue.process(lines, false)
	}

	return queue.len()
}

func part2() int {
	maxSteps := 26501365
	lines := getInputLines()
	start := getStart(lines)

	center := len(lines)/2 - 1
	increments := []int{}

	queue := Queue{*start}
	for i := 0; i < maxSteps; i++ {
		queue.process(lines, true)

		if i%len(lines) == center {
			increments = append(increments, queue.len())
		}

		if len(increments) == 3 {
			return extrapolate(increments, maxSteps/len(lines))
		}
	}

	return -1
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
