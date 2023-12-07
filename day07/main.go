package main

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	TypeFiveOfAKind  = 6
	TypeFourOfAKind  = 5
	TypeFullHouse    = 4
	TypeThreeOfAKind = 3
	TypeTwoPair      = 2
	TypeOnePair      = 1
	TypeHighCard     = 0
)

var ranks = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

type Hand struct {
	bid   int
	score int
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getCounts(hand string, withJoker bool) []int {
	counts := []int{}

	for _, rank := range ranks {
		if withJoker && rank == 'J' {
			continue
		}

		count := strings.Count(hand, string(rank))
		if count > 0 {
			counts = append(counts, count)
		}
	}

	if withJoker {
		count := strings.Count(hand, "J")
		if count > 0 {
			maxCount := 0
			maxIndex := 0

			for i, count := range counts {
				if count > maxCount {
					maxCount = count
					maxIndex = i
				}
			}

			if len(counts) == 0 {
				counts = append(counts, count)
			} else {
				counts[maxIndex] += count
			}
		}
	}

	return counts
}

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}

	return false
}

func getType(hand string, withJoker bool) int {
	counts := getCounts(hand, withJoker)

	if contains(counts, 5) {
		return TypeFiveOfAKind
	}

	if contains(counts, 4) {
		return TypeFourOfAKind
	}

	if contains(counts, 3) {
		if contains(counts, 2) {
			return TypeFullHouse
		}

		return TypeThreeOfAKind
	}

	if contains(counts, 2) {
		if len(counts) == 3 {
			return TypeTwoPair
		}

		return TypeOnePair
	}

	return TypeHighCard
}

func getScore(hand string, withJoker bool) int {
	cardType := getType(hand, withJoker)
	score := cardType

	for _, card := range hand {
		score *= 100

		if withJoker && card == 'J' {
			continue
		}

		for i, rank := range ranks {
			if card == rank {
				score += (len(ranks) - i)
				break
			}
		}
	}

	return score
}

func getTotalWinnings(withJoker bool) int {
	lines := getInputLines()
	hands := make([]Hand, len(lines))

	for i, line := range lines {
		arr := strings.Fields(line)

		score := getScore(arr[0], withJoker)
		bid, _ := strconv.Atoi(arr[1])

		hands[i] = Hand{bid, score}
	}

	sort.Slice(hands[:], func(i, j int) bool {
		return hands[i].score < hands[j].score
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	return total
}

func part1() int {
	return getTotalWinnings(false)
}

func part2() int {
	return getTotalWinnings(true)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
