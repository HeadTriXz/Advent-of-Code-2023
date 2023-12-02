package main

import (
	_ "embed"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/HeadTriXz/Advent-of-Code-2023/utils"
)

//go:embed input.txt
var input string

type Game struct {
	id    int
	red   int
	green int
	blue  int
}

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func parseID(input string) (int, string) {
	idx := strings.Index(input, ":")
	id, _ := strconv.Atoi(input[:idx])

	return id, input[idx+2:]
}

func parseGameInput(input string) (*Game, error) {
	if !strings.HasPrefix(input, "Game") {
		return nil, errors.New("invalid input")
	}

	input = strings.TrimPrefix(input, "Game ")
	id, input := parseID(input)

	batch := strings.Split(input, "; ")

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, subset := range batch {
		clusters := strings.Split(subset, ", ")

		for _, cluster := range clusters {
			arr := strings.Split(cluster, " ")
			num, err := strconv.Atoi(arr[0])
			if err != nil {
				return nil, errors.New("invalid number of cubes")
			}

			switch arr[1] {
			case "red":
				if num > maxRed {
					maxRed = num
				}
			case "green":
				if num > maxGreen {
					maxGreen = num
				}
			case "blue":
				if num > maxBlue {
					maxBlue = num
				}
			default:
				return nil, errors.New("invalid color")
			}
		}
	}

	return &Game{
		id:    id,
		red:   maxRed,
		green: maxGreen,
		blue:  maxBlue,
	}, nil
}

func part1() int {
	lines := getInputLines()
	possibleGames := []int{}

	for _, line := range lines {
		game, error := parseGameInput(line)
		if error != nil {
			log.Fatalf("Error parsing game input: %s", error)
			continue
		}

		if game.red <= maxRed && game.green <= maxGreen && game.blue <= maxBlue {
			possibleGames = append(possibleGames, game.id)
		}
	}

	return utils.Sum(possibleGames)
}

func part2() int {
	lines := getInputLines()
	totalPower := 0

	for _, line := range lines {
		game, error := parseGameInput(line)
		if error != nil {
			log.Fatalf("Error parsing game input: %s", error)
			continue
		}

		power := game.red * game.green * game.blue
		totalPower += power
	}

	return totalPower
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
