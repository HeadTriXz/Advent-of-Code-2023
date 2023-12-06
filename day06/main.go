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

func calcDistance(start int, time int) int {
	return start * (time - start)
}

func parseIntArray(str string) (result []int) {
	for _, num := range strings.Fields(str)[1:] {
		n, _ := strconv.Atoi(num)
		result = append(result, n)
	}

	return result
}

func parseInt(str string) int {
	arr := strings.Fields(str)[1:]
	joined := strings.Join(arr, "")
	n, _ := strconv.Atoi(joined)

	return n
}

func part1() int {
	lines := getInputLines()
	timeArr := parseIntArray(lines[0])
	distArr := parseIntArray(lines[1])

	margin := 0
	for i := 0; i < len(timeArr); i++ {
		time := timeArr[i]
		dist := distArr[i]

		count := 0
		for j := 0; j < time; j++ {
			if calcDistance(j, time) >= dist {
				count++
			}
		}

		if margin == 0 {
			margin = count
		} else {
			margin *= count
		}
	}

	return margin
}

func part2() int {
	lines := getInputLines()
	time := parseInt(lines[0])
	dist := parseInt(lines[1])

	count := 0
	for j := 0; j < time; j++ {
		if calcDistance(j, time) >= dist {
			count++
		}
	}

	return count
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
