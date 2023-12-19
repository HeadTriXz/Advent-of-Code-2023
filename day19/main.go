package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var operators = map[rune]func(int, int) bool{
	'>': func(a, b int) bool { return a > b },
	'<': func(a, b int) bool { return a < b },
	' ': func(a, b int) bool { return true },
}

type Rule struct {
	category rune
	operator rune
	operand  int
	result   string
}

type Part struct {
	x, m, a, s int
}

type PartRange struct {
	start Part
	end   Part
}

func (p *Part) get(category rune) int {
	switch category {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	}

	return 0
}

func (p *Part) set(category rune, value int) {
	switch category {
	case 'x':
		p.x = value
	case 'm':
		p.m = value
	case 'a':
		p.a = value
	case 's':
		p.s = value
	}
}

func getInput() (map[string][]Rule, []Part) {
	sections := strings.Split(strings.TrimRight(input, "\n"), "\n\n")

	workflows := map[string][]Rule{}
	for _, line := range strings.Split(sections[0], "\n") {
		splits := strings.Split(line, "{")
		name := splits[0]
		rules := strings.Split(strings.TrimSuffix(splits[1], "}"), ",")

		workflows[name] = []Rule{}

		for _, rule := range rules {
			splits := strings.FieldsFunc(rule, func(r rune) bool {
				return r == '>' || r == '<' || r == ':'
			})

			if len(splits) == 3 {
				operator := rune(rule[1])
				category := rune(splits[0][0])
				operand, _ := strconv.Atoi(splits[1])
				result := splits[2]

				workflows[name] = append(workflows[name], Rule{
					category: category,
					operator: operator,
					operand:  operand,
					result:   result,
				})
				continue
			}

			workflows[name] = append(workflows[name], Rule{
				category: ' ',
				operator: ' ',
				operand:  0,
				result:   splits[0],
			})
		}
	}

	parts := []Part{}
	for _, line := range strings.Split(sections[1], "\n") {
		line = line[1 : len(line)-1]
		splits := strings.Split(line, ",")

		part := Part{}
		for _, split := range splits {
			category := rune(split[0])
			operand, _ := strconv.Atoi(split[2:])

			part.set(category, operand)
		}

		parts = append(parts, part)
	}

	return workflows, parts
}

func isAccepted(part Part, workflows map[string][]Rule, name string) bool {
	workflow := workflows[name]
	for _, rule := range workflow {
		isValid := operators[rule.operator](part.get(rule.category), rule.operand)
		if !isValid {
			continue
		}

		if rule.result == "A" {
			return true
		}

		if rule.result == "R" {
			return false
		}

		return isAccepted(part, workflows, rule.result)
	}

	return false
}

func calcCombinations(workflows map[string][]Rule, name string, ranges PartRange) int {
	if name == "R" {
		return 0
	}

	if name == "A" {
		sum := 1

		sum *= ranges.end.x - ranges.start.x + 1
		sum *= ranges.end.m - ranges.start.m + 1
		sum *= ranges.end.a - ranges.start.a + 1
		sum *= ranges.end.s - ranges.start.s + 1

		return sum
	}

	total := 0
	for _, rule := range workflows[name] {
		if rule.operator == ' ' {
			total += calcCombinations(workflows, rule.result, ranges)
			continue
		}

		if ranges.start.get(rule.category) >= rule.operand || rule.operand >= ranges.end.get(rule.category) {
			continue
		}

		newRanges := ranges
		if rule.operator == '>' {
			newRanges.start.set(rule.category, rule.operand+1)
			ranges.end.set(rule.category, rule.operand)
		}

		if rule.operator == '<' {
			newRanges.end.set(rule.category, rule.operand-1)
			ranges.start.set(rule.category, rule.operand)
		}

		total += calcCombinations(workflows, rule.result, newRanges)
	}

	return total
}

func part1() int {
	workflows, parts := getInput()

	sum := 0
	for _, part := range parts {
		if !isAccepted(part, workflows, "in") {
			continue
		}

		sum += part.x
		sum += part.m
		sum += part.a
		sum += part.s
	}

	return sum
}

func part2() int {
	workflows, _ := getInput()
	ranges := PartRange{
		start: Part{x: 1, m: 1, a: 1, s: 1},
		end:   Part{x: 4000, m: 4000, a: 4000, s: 4000},
	}

	return calcCombinations(workflows, "in", ranges)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
