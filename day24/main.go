package main

import (
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const epsilon = 1e-14

// Types
type Node struct {
	x float64
	y float64
	z float64
}

type Stone struct {
	position Node
	velocity Node
}

// Methods
func (h *Stone) slope() float64 {
	return h.velocity.y / h.velocity.x
}

func (h *Stone) intercept() float64 {
	return h.position.y - h.slope()*h.position.x
}

func (h *Stone) intersectX(other Stone) float64 {
	m1, m2 := h.slope(), other.slope()
	b1, b2 := h.intercept(), other.intercept()

	m := m1 - m2
	if m == 0 {
		return math.NaN()
	}

	return (b2 - b1) / m
}

func (h *Stone) intersectY(other Stone) float64 {
	m1, m2 := h.slope(), other.slope()
	b1, b2 := h.intercept(), other.intercept()

	m := m1 - m2
	if m == 0 {
		return math.NaN()
	}

	x := (b2 - b1) / m
	return m1*x + b1
}

func (h *Stone) timeToInterceptX(x float64) float64 {
	return (x - h.position.x) / h.velocity.x
}

func (h *Stone) slowDown(x int, y int, z int) {
	h.velocity.x -= float64(x)
	h.velocity.y -= float64(y)
	h.velocity.z -= float64(z)
}

// Input parsing
func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func parseFloat(s string) float64 {
	trimmed := strings.Trim(s, " ")
	num, _ := strconv.Atoi(trimmed)

	return float64(num)
}

func parseInput() []Stone {
	lines := getInputLines()
	results := make([]Stone, len(lines))

	for i, line := range lines {
		chunks := strings.Split(line, "@")
		positions := strings.Split(chunks[0], ",")
		velocities := strings.Split(chunks[1], ",")

		results[i] = Stone{
			position: Node{
				x: parseFloat(positions[0]),
				y: parseFloat(positions[1]),
				z: parseFloat(positions[2]),
			},
			velocity: Node{
				x: parseFloat(velocities[0]),
				y: parseFloat(velocities[1]),
				z: parseFloat(velocities[2]),
			},
		}
	}

	return results
}

// Utility functions
func isWithinEpsilon(a float64, b float64) bool {
	return a*(1-epsilon) <= b && a*(1+epsilon) >= b
}

func findRockPosition(hailstones []Stone, vx int, vy int) (rx int, ry int, rz int, ok bool) {
	h1, h2, h3 := hailstones[0], hailstones[1], hailstones[2]

	h1.slowDown(vx, vy, 0)
	h2.slowDown(vx, vy, 0)
	h3.slowDown(vx, vy, 0)

	x1 := h1.intersectX(h2)
	x2 := h1.intersectX(h3)
	x3 := h2.intersectX(h3)
	if !isWithinEpsilon(x1, x2) || !isWithinEpsilon(x1, x3) {
		return 0, 0, 0, false
	}

	y1 := h1.intersectY(h2)
	y2 := h1.intersectY(h3)
	y3 := h2.intersectY(h3)
	if !isWithinEpsilon(y1, y2) || !isWithinEpsilon(y1, y3) {
		return 0, 0, 0, false
	}

	t1 := h1.timeToInterceptX(x1)
	t2 := h2.timeToInterceptX(x1)
	t3 := h3.timeToInterceptX(x1)
	if t1 < 0 || t2 < 0 || t3 < 0 {
		return 0, 0, 0, false
	}

	vz1 := (h1.position.z - h2.position.z + t1*h1.velocity.z - t2*h2.velocity.z) / (t1 - t2)
	vz2 := (h1.position.z - h3.position.z + t1*h1.velocity.z - t3*h3.velocity.z) / (t1 - t3)
	vz3 := (h3.position.z - h2.position.z + t3*h3.velocity.z - t2*h2.velocity.z) / (t3 - t2)
	if !isWithinEpsilon(vz1, vz2) || !isWithinEpsilon(vz1, vz3) {
		return 0, 0, 0, false
	}

	z1 := h1.position.z + t1*(h1.velocity.z-vz1)
	z2 := h2.position.z + t2*(h2.velocity.z-vz1)
	z3 := h3.position.z + t3*(h3.velocity.z-vz1)
	if !isWithinEpsilon(z1, z2) || !isWithinEpsilon(z2, z3) {
		return 0, 0, 0, false
	}

	return int(x1), int(y1), int(z1), true
}

// Solutions
func part1() int {
	min := 200000000000000.0
	max := 400000000000000.0

	hailstones := parseInput()

	count := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			x := hailstones[i].intersectX(hailstones[j])
			y := hailstones[i].intersectY(hailstones[j])
			if math.IsNaN(x) || math.IsNaN(y) {
				continue
			}

			t1 := hailstones[i].timeToInterceptX(x)
			t2 := hailstones[j].timeToInterceptX(x)
			if t1 < 0 || t2 < 0 {
				continue
			}

			if x >= min && x <= max && y >= min && y <= max {
				count++
			}
		}
	}

	return count
}

func part2() int {
	hailstones := parseInput()
	combinations := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for sum := 0; sum <= 1000000; sum++ {
		for vx := 0; vx <= sum; vx++ {
			vy := sum - vx

			for _, direction := range combinations {
				x, y, z, found := findRockPosition(hailstones, vx*direction[0], vy*direction[1])
				if found {
					return x + y + z
				}
			}
		}
	}

	return 0
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
