package main

import (
	_ "embed"
	"log"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func getInput() (string, map[string][]string) {
	lines := getInputLines()
	nodes := map[string][]string{}

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		regex := regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)
		matches := regex.FindStringSubmatch(line)

		if len(matches) != 4 {
			log.Fatalf("Invalid line: %s", line)
			continue
		}

		nodes[matches[1]] = []string{matches[2], matches[3]}
	}

	return lines[0], nodes
}

func rotate(str *string) rune {
	char := (*str)[0]
	*str = (*str)[1:] + (*str)[:1]

	return rune(char)
}

func hasFoundAll(nodes []string) bool {
	for _, node := range nodes {
		if node[len(node)-1] != 'Z' {
			return false
		}
	}

	return true
}

func gdc(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gdc(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func part1() int {
	instructions, nodes := getInput()

	steps := 0
	currentElement := "AAA"
	for currentElement != "ZZZ" {
		if rotate(&instructions) == 'L' {
			currentElement = nodes[currentElement][0]
		} else {
			currentElement = nodes[currentElement][1]
		}

		steps++
	}

	return steps
}

func part2() int {
	instructions, nodes := getInput()

	startNodes := []string{}
	for node := range nodes {
		if node[len(node)-1] == 'A' {
			startNodes = append(startNodes, node)
		}
	}

	steps := 0
	for !hasFoundAll(startNodes) {
		lr := rotate(&instructions)
		for i, node := range startNodes {
			if lr == 'L' {
				startNodes[i] = nodes[node][0]
			} else {
				startNodes[i] = nodes[node][1]
			}
		}

		steps++
	}

	return steps
}

func part2LCM() int {
	instructions, nodes := getInput()

	stepCounts := []int{}
	for node := range nodes {
		if node[len(node)-1] != 'A' {
			continue
		}

		steps := 0
		currentElement := node
		for currentElement[len(currentElement)-1] != 'Z' {
			if rotate(&instructions) == 'L' {
				currentElement = nodes[currentElement][0]
			} else {
				currentElement = nodes[currentElement][1]
			}

			steps++
		}

		stepCounts = append(stepCounts, steps)
	}

	return lcm(stepCounts[0], stepCounts[1], stepCounts...)
}

func main() {
	part1Result := part1()
	part2Result := part2LCM()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
