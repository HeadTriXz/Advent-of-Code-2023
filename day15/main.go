package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getInput() []string {
	return strings.Split(strings.TrimRight(input, "\n"), ",")
}

type Lens struct {
	focalLength int
	label       string
}

func (l *Lens) getHashCode() int {
	return hashString(l.label)
}

func split(char rune) bool {
	return char == '=' || char == '-'
}

func parseLens(l string) (*Lens, bool) {
	parts := strings.FieldsFunc(l, split)

	focalLength := 0
	if len(parts) > 1 {
		focalLength, _ = strconv.Atoi(parts[1])
	}

	lens := &Lens{focalLength, parts[0]}

	return lens, len(parts) > 1
}

func hashString(s string) int {
	hash := 0
	for _, char := range s {
		hash += int(char)
		hash *= 17
		hash %= 256
	}

	return hash
}

func part1() int {
	arr := getInput()

	sum := 0
	for _, s := range arr {
		sum += hashString(s)
	}

	return sum
}

func part2() int {
	arr := getInput()
	boxes := map[int][]Lens{}

	sum := 0
	for _, s := range arr {
		lens, shouldAdd := parseLens(s)
		hash := lens.getHashCode()

		if shouldAdd {
			if items, ok := boxes[hash]; ok {
				alreadyExists := false
				for i, item := range items {
					if item.label == lens.label {
						items[i] = *lens
						alreadyExists = true
						break
					}
				}

				if !alreadyExists {
					boxes[hash] = append(items, *lens)
				}
			} else {
				boxes[hash] = []Lens{*lens}
			}
		} else {
			if items, ok := boxes[hash]; ok {
				for i, item := range items {
					if item.label == lens.label {
						boxes[hash] = append(items[:i], items[i+1:]...)
						break
					}
				}
			}
		}
	}

	for box, items := range boxes {
		for i, item := range items {
			sum += (1 + box) * (1 + i) * item.focalLength
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
