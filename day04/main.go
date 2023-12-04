package main

import (
	_ "embed"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getCardID(input string) int {
	regex := regexp.MustCompile(`Card\s*(\d+):`)
	match := regex.FindStringSubmatch(input)[1]

	id, _ := strconv.Atoi(match)
	return id
}

func getWinningNumbers(input string) []string {
	result := []string{}
	input = input[strings.Index(input, ":")+1:]

	arr := strings.Split(input, "|")
	winningNums := strings.Fields(strings.Trim(arr[0], " "))
	winningMap := map[string]bool{}
	for _, num := range winningNums {
		winningMap[num] = true
	}

	ourNums := strings.Fields(strings.Trim(arr[1], " "))
	for _, num := range ourNums {
		if winningMap[num] {
			result = append(result, num)
		}
	}

	return result
}

func part1() int {
	lines := getInputLines()

	sum := 0
	for _, line := range lines {
		nums := getWinningNumbers(line)

		if len(nums) > 0 {
			sum += int(math.Pow(2, float64(len(nums))-1))
		}
	}

	return sum
}

func countCards(cardID int, cards map[int]string) int {
	nums := getWinningNumbers(cards[cardID])
	if len(nums) == 0 {
		return 1
	}

	count := 1
	for i := cardID + 1; i <= cardID+len(nums); i++ {
		count += countCards(i, cards)
	}

	return count
}

func part2() int {
	lines := getInputLines()
	count := 0

	mappedCards := map[int]string{}
	for _, line := range lines {
		mappedCards[getCardID(line)] = line
	}

	for cardID := range mappedCards {
		count += countCards(cardID, mappedCards)
	}

	return count
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
