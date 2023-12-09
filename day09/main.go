package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getDifference(arr []int) (result []int) {
	for i := 1; i < len(arr); i++ {
		result = append(result, arr[i]-arr[i-1])
	}

	return result
}

func isAllZero(arr []int) bool {
	for _, value := range arr {
		if value != 0 {
			return false
		}
	}

	return true
}

func extrapolateForward(arr []int) []int {
	diff := getDifference(arr)
	if isAllZero(diff) {
		return append(arr, arr[0])
	}

	ex := extrapolateForward(diff)
	value := arr[len(arr)-1] + ex[len(ex)-1]

	return append(arr, value)
}

func extrapolateBackwards(arr []int) []int {
	diff := getDifference(arr)
	if isAllZero(diff) {
		return append(arr, arr[0])
	}

	ex := extrapolateBackwards(diff)
	value := arr[0] - ex[0]

	return append([]int{value}, arr...)
}

func part1() int {
	lines := getInputLines()

	sum := 0
	for _, line := range lines {
		values := strings.Fields(line)

		integers := make([]int, len(values))
		for i, value := range values {
			integers[i], _ = strconv.Atoi(value)
		}

		ex := extrapolateForward(integers)
		sum += ex[len(ex)-1]
	}

	return sum
}

func part2() int {
	lines := getInputLines()

	sum := 0
	for _, line := range lines {
		values := strings.Fields(line)

		integers := make([]int, len(values))
		for i, value := range values {
			integers[i], _ = strconv.Atoi(value)
		}

		ex := extrapolateBackwards(integers)
		sum += ex[0]
	}

	return sum
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
