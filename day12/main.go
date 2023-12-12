package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var cache = map[string]int{}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getInput() (allSprings []string, allDamaged [][]int) {
	lines := getInputLines()

	for _, line := range lines {
		fields := strings.Fields(line)
		allSprings = append(allSprings, fields[0])

		row := []int{}
		for _, field := range strings.Split(fields[1], ",") {
			num, _ := strconv.Atoi(field)
			row = append(row, num)
		}

		allDamaged = append(allDamaged, row)
	}

	return allSprings, allDamaged
}

func getKey(springs string, damaged []int, contiguous int) string {
	key := springs + "-"
	for _, d := range damaged {
		key += strconv.Itoa(d) + ","
	}

	return key + "-" + strconv.Itoa(contiguous)
}

func getWithDamaged(springs string, damaged []int, contiguous int) int {
	if len(damaged) == 0 {
		return 0
	}

	key := getKey(springs, damaged, contiguous)
	if val, ok := cache[key]; ok {
		return val
	}

	contiguous++
	if contiguous == damaged[0] {
		if len(springs) == 1 {
			cache[key] = getCombinations(springs[1:], damaged[1:], 0)
			return cache[key]
		}

		if springs[1] != '#' {
			cache[key] = getCombinations(springs[2:], damaged[1:], 0)
			return cache[key]
		}

		return 0
	}

	if contiguous > damaged[0] {
		return 0
	}

	cache[key] = getCombinations(springs[1:], damaged, contiguous)
	return cache[key]
}

func getWithoutDamaged(springs string, damaged []int, contiguous int) int {
	if contiguous > 0 {
		return 0
	}

	return getCombinations(springs[1:], damaged, 0)
}

func getCombinations(springs string, damaged []int, contiguous int) int {
	if len(springs) == 0 {
		if len(damaged) == 0 {
			return 1
		}

		return 0
	}

	switch springs[0] {
	case '#':
		return getWithDamaged(springs, damaged, contiguous)
	case '.':
		return getWithoutDamaged(springs, damaged, contiguous)
	case '?':
		return getWithDamaged(springs, damaged, contiguous) + getWithoutDamaged(springs, damaged, contiguous)
	}

	return 0
}

func unfold(springs string, damaged []int, n int) (newSprings string, newDamaged []int) {
	repeated := []string{}
	for i := 0; i < n; i++ {
		repeated = append(repeated, springs)
		newDamaged = append(newDamaged, damaged...)
	}

	newSprings = strings.Join(repeated, "?")
	return newSprings, newDamaged
}

func part1() int {
	allSprings, allDamaged := getInput()

	sum := 0
	for i := 0; i < len(allSprings); i++ {
		springs := allSprings[i]
		damaged := allDamaged[i]

		sum += getCombinations(springs, damaged, 0)
	}

	return sum
}

func part2() int {
	allSprings, allDamaged := getInput()

	sum := 0
	for i := 0; i < len(allSprings); i++ {
		springs, damaged := unfold(allSprings[i], allDamaged[i], 5)

		sum += getCombinations(springs, damaged, 0)
	}

	return sum
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
