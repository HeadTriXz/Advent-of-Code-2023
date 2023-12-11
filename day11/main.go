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

func getEmptyCols(lines []string) (result []int) {
	for col := range lines[0] {
		isEmpty := true
		for _, line := range lines {
			if line[col] == '#' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			result = append(result, col)
		}
	}

	return result
}

func getEmptyRows(lines []string) (result []int) {
	for row, line := range lines {
		if !strings.Contains(line, "#") {
			result = append(result, row)
		}
	}

	return result
}

func findGalaxies(lines []string, growthRate int) (galaxies [][]int) {
	emptyRows := getEmptyRows(lines)
	emptyCols := getEmptyCols(lines)

	for row, line := range lines {
		rowOffset := 0
		for _, emptyRow := range emptyRows {
			if row > emptyRow {
				rowOffset += growthRate - 1
			}
		}

		for col, char := range line {
			colOffset := 0
			for _, emptyCol := range emptyCols {
				if col > emptyCol {
					colOffset += growthRate - 1
				}
			}

			if char == '#' {
				galaxy := []int{row + rowOffset, col + colOffset}
				galaxies = append(galaxies, galaxy)
			}
		}
	}

	return galaxies
}

func absDistance(a int, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func getDistance(a []int, b []int) int {
	return absDistance(a[0], b[0]) + absDistance(a[1], b[1])
}

func getDistanceSum(lines []string, growthRate int) int {
	sum := 0
	galaxies := findGalaxies(lines, growthRate)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := getDistance(galaxies[i], galaxies[j])
			sum += dist
		}
	}

	return sum
}

func part1() int {
	return getDistanceSum(getInputLines(), 2)
}

func part2() int {
	return getDistanceSum(getInputLines(), 1000000)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
