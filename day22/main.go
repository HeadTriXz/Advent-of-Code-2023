package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	x int
	y int
	z int
}

type Brick struct {
	start Node
	end   Node
}

func (b *Brick) countSupporting(bricks []Brick) int {
	without := []Brick{}
	for _, brick := range bricks {
		if brick != *b {
			without = append(without, brick)
		}
	}

	_, count := abuseGravity(without)
	return count
}

func (b *Brick) isColliding(other *Brick) bool {
	return b.start.x <= other.end.x && b.end.x >= other.start.x &&
		b.start.y <= other.end.y && b.end.y >= other.start.y &&
		b.start.z <= other.end.z && b.end.z >= other.start.z
}

func (b *Brick) isCollidingWithAny(bricks []Brick) bool {
	for _, brick := range bricks {
		if b.isColliding(&brick) {
			return true
		}
	}

	return false
}

func (b *Brick) isSupporting(other *Brick) bool {
	return b.start.x <= other.end.x && b.end.x >= other.start.x &&
		b.start.y <= other.end.y && b.end.y >= other.start.y &&
		b.end.z+1 == other.start.z
}

func (b *Brick) getSupporting(bricks []Brick) []Brick {
	supporting := []Brick{}

	for _, brick := range bricks {
		if b.isSupporting(&brick) {
			supporting = append(supporting, brick)
		}
	}

	return supporting
}

func (b *Brick) getSupportedBy(bricks []Brick) []Brick {
	supported := []Brick{}

	for _, brick := range bricks {
		if brick.isSupporting(b) {
			supported = append(supported, brick)
		}
	}

	return supported
}

func (b *Brick) moveDown() {
	if b.start.z == 1 {
		return
	}

	b.start.z--
	b.end.z--
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func parseInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseInput() []Brick {
	lines := getInputLines()
	bricks := make([]Brick, len(lines))

	for i, line := range lines {
		splits := strings.Split(line, "~")
		start := strings.Split(splits[0], ",")
		end := strings.Split(splits[1], ",")

		bricks[i] = Brick{
			start: Node{
				x: parseInt(start[0]),
				y: parseInt(start[1]),
				z: parseInt(start[2]),
			},
			end: Node{
				x: parseInt(end[0]),
				y: parseInt(end[1]),
				z: parseInt(end[2]),
			},
		}
	}

	return bricks
}

func sortBricks(bricks []Brick) []Brick {
	for i := 0; i < len(bricks); i++ {
		for j := i + 1; j < len(bricks); j++ {
			if bricks[i].start.z > bricks[j].start.z {
				bricks[i], bricks[j] = bricks[j], bricks[i]
			}
		}
	}

	return bricks
}

func abuseGravity(bricks []Brick) ([]Brick, int) {
	movedMap := map[int]bool{}

	for i, brick := range bricks {
		for brick.start.z > 1 {
			brick.moveDown()
			if brick.isCollidingWithAny(bricks[:i]) {
				break
			}

			bricks[i] = brick
			movedMap[i] = true
		}
	}

	return bricks, len(movedMap)
}

func part1() int {
	bricks := parseInput()
	bricks = sortBricks(bricks)
	bricks, _ = abuseGravity(bricks)

	count := 0
	for _, brick := range bricks {
		supporting := brick.getSupporting(bricks)
		isSafe := true

		for _, b := range supporting {
			supported := b.getSupportedBy(bricks)
			if len(supported) == 1 {
				isSafe = false
				break
			}
		}

		if isSafe {
			count++
		}
	}

	return count
}

func part2() int {
	bricks := parseInput()
	bricks = sortBricks(bricks)
	bricks, _ = abuseGravity(bricks)

	sum := 0
	for _, brick := range bricks {
		sum += brick.countSupporting(bricks)
	}

	return sum
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
