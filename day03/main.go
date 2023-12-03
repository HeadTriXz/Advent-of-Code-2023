package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getNumberIndices(input string) [][]int {
	regex := regexp.MustCompile(`\d+`)
	return regex.FindAllStringIndex(input, -1)
}

func getGearIndices(input string) []int {
	indices := []int{}
	for i, char := range input {
		if char == '*' {
			indices = append(indices, i)
		}
	}

	return indices
}

func part1() int {
	lines := getInputLines()

	sum := 0
	for i, line := range lines {
		minRow := max(i-1, 0)
		maxRow := min(i+2, len(lines))

		indices := getNumberIndices(line)
		for _, indexPair := range indices {
			minCol := max(indexPair[0]-1, 0)
			maxCol := min(indexPair[1]+1, len(line)-1)

			for _, row := range lines[minRow:maxRow] {
				for _, col := range row[minCol:maxCol] {
					if col == '.' || unicode.IsDigit(col) {
						continue
					}

					number, _ := strconv.Atoi(line[indexPair[0]:indexPair[1]])
					sum += number
					break
				}
			}
		}
	}

	return sum
}

func part2() int {
	lines := getInputLines()

	sum := 0
	for i, line := range lines {
		minRow := max(i-1, 0)
		maxRow := min(i+2, len(lines))

		gearIndices := getGearIndices(line)
		if len(gearIndices) == 0 {
			continue
		}

		for _, gearIndex := range gearIndices {
			minCol := max(gearIndex-1, 0)
			maxCol := min(gearIndex+1, len(line)-1)

			var first int
			for _, row := range lines[minRow:maxRow] {
				numIndices := getNumberIndices(row)
				for _, indexPair := range numIndices {
					if indexPair[0] > maxCol || indexPair[1] <= minCol {
						continue
					}

					num, _ := strconv.Atoi(row[indexPair[0]:indexPair[1]])
					if first == 0 {
						first = num
						continue
					}

					sum += num * first
				}
			}
		}
	}

	return sum
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
