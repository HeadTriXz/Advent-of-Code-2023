package main

import (
	"container/heap"
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var directions = []Node{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type Node struct {
	x int
	y int
}

type Item struct {
	distance int
	state    State
}

type State struct {
	consecutive int
	direction   Node
	node        Node
}

type PriorityQueue []*Item

func getInputGraph() [][]int {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	graph := make([][]int, len(lines))

	for i := range graph {
		graph[i] = make([]int, len(lines[i]))
	}

	for row, line := range lines {
		for col, char := range line {
			graph[row][col], _ = strconv.Atoi(string(char))
		}
	}

	return graph
}

func (queue PriorityQueue) Len() int {
	return len(queue)
}

func (queue PriorityQueue) Less(i int, j int) bool {
	return queue[i].distance < queue[j].distance
}

func (queue PriorityQueue) Swap(i int, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}

func (queue *PriorityQueue) Push(x interface{}) {
	state := x.(*Item)
	*queue = append(*queue, state)
}

func (queue *PriorityQueue) Pop() any {
	old := *queue
	n := len(old)

	state := old[n-1]
	*queue = old[0 : n-1]

	return state
}

func (queue *PriorityQueue) IsEmpty() bool {
	return queue.Len() == 0
}

func isValid(node Node, graph [][]int) bool {
	return node.y >= 0 && node.x >= 0 && node.y < len(graph) && node.x < len(graph[node.y])
}

func getShortestPath(
	graph [][]int,
	start Node,
	end Node,
	minConsecutive int,
	maxConsecutive int,
) int {
	dist := make(map[State]int)
	queue := PriorityQueue{
		&Item{
			distance: 0,
			state: State{
				consecutive: 1,
				direction:   Node{1, 0},
				node:        start,
			},
		},
		&Item{
			distance: 0,
			state: State{
				consecutive: 1,
				direction:   Node{0, 1},
				node:        start,
			},
		},
	}

	heap.Init(&queue)

	for !queue.IsEmpty() {
		item := heap.Pop(&queue).(*Item)

		if item.state.node == end && item.state.consecutive >= minConsecutive {
			return item.distance
		}

		if distance, ok := dist[item.state]; ok {
			if item.distance >= distance {
				continue
			}
		}

		dist[item.state] = item.distance

		for _, direction := range directions {
			if item.state.direction.x == -direction.x && item.state.direction.y == -direction.y {
				continue
			}

			if item.state.consecutive < minConsecutive && item.state.direction != direction {
				continue
			}

			next := Node{
				x: item.state.node.x + direction.x,
				y: item.state.node.y + direction.y,
			}

			if isValid(next, graph) {
				nextConsecutive := 1
				if direction == item.state.direction {
					nextConsecutive = item.state.consecutive + 1
				}

				if nextConsecutive <= maxConsecutive {
					nextDist := item.distance + graph[next.y][next.x]
					nextState := State{
						consecutive: nextConsecutive,
						direction:   direction,
						node:        next,
					}

					heap.Push(&queue, &Item{nextDist, nextState})
				}
			}
		}
	}

	return -1
}

func part1() int {
	graph := getInputGraph()

	lastRow := len(graph) - 1
	lastCol := len(graph[lastRow]) - 1

	start := Node{0, 0}
	end := Node{lastCol, lastRow}

	return getShortestPath(graph, start, end, 1, 3)
}

func part2() int {
	graph := getInputGraph()

	lastRow := len(graph) - 1
	lastCol := len(graph[lastRow]) - 1

	start := Node{0, 0}
	end := Node{lastCol, lastRow}

	return getShortestPath(graph, start, end, 4, 10)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
