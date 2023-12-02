package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/HeadTriXz/Advent-of-Code-2023/utils"
)

//go:embed input.txt
var input string

var humanNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getFirstInt(str string) (int, string) {
	for i, char := range str {
		if unicode.IsDigit(char) {
			return i, string(char)
		}
	}

	return -1, ""
}

func getFirstInt2(str string) (int, string) {
	var minIndex int = len(str)
	var minChar string = ""

	for k, v := range humanNumbers {
		i := strings.Index(str, k)
		if i != -1 && i < minIndex {
			minIndex = i
			minChar = strconv.Itoa(v)
		}
	}

	for i, char := range str {
		if i >= minIndex {
			break
		}

		if unicode.IsDigit(char) {
			return i, string(char)
		}
	}

	return minIndex, minChar
}

func getLastInt(str string) (int, string) {
	for i := len(str) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(str[i])) {
			return i, string(str[i])
		}
	}

	return -1, ""
}

func getLastInt2(str string) (int, string) {
	var maxIndex int = -1
	var maxChar string = ""

	for k, v := range humanNumbers {
		i := strings.LastIndex(str, k)
		if i != -1 && i > maxIndex {
			maxIndex = i
			maxChar = strconv.Itoa(v)
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if i <= maxIndex {
			break
		}

		if unicode.IsDigit(rune(str[i])) {
			return i, string(str[i])
		}
	}

	return maxIndex, maxChar
}

func part1() int {
	values := []int{}
	lines := getInputLines()

	for _, line := range lines {
		i1, v1 := getFirstInt(line)
		i2, v2 := getLastInt(line)

		if i1 == -1 || i2 == -1 {
			continue
		}

		s, err := strconv.Atoi(v1 + v2)
		if err != nil {
			log.Fatal("Cannot convert string to integer")
		}

		values = append(values, s)
	}

	sum := utils.Sum(values)
	return sum
}

func part2() int {
	values := []int{}
	lines := getInputLines()

	for _, line := range lines {
		i1, v1 := getFirstInt2(line)
		i2, v2 := getLastInt2(line)

		if i1 == -1 || i2 == -1 {
			continue
		}

		s, err := strconv.Atoi(v1 + v2)
		if err != nil {
			log.Fatal("Cannot convert string to integer")
		}

		values = append(values, s)
	}

	sum := utils.Sum(values)
	return sum
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
