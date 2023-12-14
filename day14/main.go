package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func tiltHorizontal(lines []string, north bool) []string {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines)-1; j++ {
			top := lines[j]
			bottom := lines[j+1]

			for k := 0; k < len(top); k++ {
				if north && bottom[k] == 'O' && top[k] == '.' {
					top = top[:k] + "O" + top[k+1:]
					bottom = bottom[:k] + "." + bottom[k+1:]
				}

				if !north && bottom[k] == '.' && top[k] == 'O' {
					top = top[:k] + "." + top[k+1:]
					bottom = bottom[:k] + "O" + bottom[k+1:]
				}
			}

			lines[j] = top
			lines[j+1] = bottom
		}
	}

	return lines
}

func tiltVertical(lines []string, west bool) []string {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			for k := 0; k < len(lines[i])-1; k++ {
				left := lines[i][k]
				right := lines[i][k+1]

				if west && left == '.' && right == 'O' {
					lines[i] = lines[i][:k] + "O" + "." + lines[i][k+2:]
				}

				if !west && left == 'O' && right == '.' {
					lines[i] = lines[i][:k] + "." + "O" + lines[i][k+2:]
				}
			}
		}
	}

	return lines
}

func spinCycle(lines []string) []string {
	tilted := tiltHorizontal(lines, true)
	tilted = tiltVertical(tilted, true)
	tilted = tiltHorizontal(tilted, false)
	tilted = tiltVertical(tilted, false)

	return tilted
}

func calcLoad(lines []string) int {
	sum := 0
	for row, line := range lines {
		for _, char := range line {
			if char == 'O' {
				sum += len(lines) - row
			}
		}
	}

	return sum
}

func part1() int {
	lines := getInputLines()
	tilted := tiltHorizontal(lines, true)

	return calcLoad(tilted)
}

func part2() int {
	lines := getInputLines()
	seenStates := make(map[string]int)

	for i := 0; i < 1000000000; i++ {
		state := strings.Join(lines, "\n")
		if idx, ok := seenStates[state]; ok {
			remainingIterations := (1000000000 - i) % (i - idx)

			for j := 0; j < remainingIterations; j++ {
				lines = spinCycle(lines)
			}

			break
		}

		seenStates[state] = i
		lines = spinCycle(lines)
	}

	return calcLoad(lines)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
