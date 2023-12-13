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

func splitByEmptyLine(input []string) (chunks [][]string) {
	chunk := []string{}

	for _, line := range input {
		if line == "" {
			chunks = append(chunks, chunk)
			chunk = []string{}
			continue
		}

		chunk = append(chunk, line)
	}

	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}

	return chunks
}

func findHorizontalReflection(lines []string, maxDiff int) int {
	for i := 0; i < len(lines)-1; i++ {
		top := i
		bottom := i + 1

		diff := 0
		for top >= 0 && bottom < len(lines) {
			for j := 0; j < len(lines[top]); j++ {
				if lines[top][j] != lines[bottom][j] {
					diff++
				}
			}

			top--
			bottom++
		}

		if diff == maxDiff {
			return i
		}
	}

	return -1
}

func findVerticalReflection(lines []string, maxDiff int) int {
	for i := 0; i < len(lines[0])-1; i++ {
		left := i
		right := i + 1

		diff := 0
		for left >= 0 && right < len(lines[0]) {
			for _, line := range lines {
				if line[left] != line[right] {
					diff++
				}
			}

			left--
			right++
		}

		if diff == maxDiff {
			return i
		}
	}

	return -1
}

func getSumOfReflections(lines []string, maxDiff int) int {
	chunks := splitByEmptyLine(lines)

	sum := 0
	for _, chunk := range chunks {
		hReflection := findHorizontalReflection(chunk, maxDiff)
		vReflection := findVerticalReflection(chunk, maxDiff)

		if hReflection != -1 {
			sum += (hReflection + 1) * 100
		}

		if vReflection != -1 {
			sum += vReflection + 1
		}
	}

	return sum
}

func part1() int {
	return getSumOfReflections(getInputLines(), 0)
}

func part2() int {
	return getSumOfReflections(getInputLines(), 1)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
