package main

import (
	_ "embed"
	"log"
	"slices"
	"strings"

	"github.com/HeadTriXz/Advent-of-Code-2023/utils"
)

//go:embed input.txt
var input string

type ModuleType int

type Module struct {
	dest []string
	name string
	typ  ModuleType
}

type Pulse struct {
	dest string
	src  string
	high bool
}

const (
	Broadcast ModuleType = iota
	Conjunction
	FlipFlop
)

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func parseInput() (map[string]Module, map[string]bool, map[string]map[string]bool) {
	lines := getInputLines()
	modules := make(map[string]Module, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		dest := strings.Split(parts[1], ", ")
		name := parts[0]
		typ := Broadcast

		if strings.Contains(name, "&") {
			name = name[1:]
			typ = Conjunction
		} else if strings.Contains(name, "%") {
			name = name[1:]
			typ = FlipFlop
		}

		modules[name] = Module{
			dest: dest,
			name: name,
			typ:  typ,
		}
	}

	states := map[string]bool{}
	memory := map[string]map[string]bool{}
	for _, mod := range modules {
		if mod.typ == FlipFlop {
			states[mod.name] = false
		}

		if mod.typ == Conjunction {
			memory[mod.name] = map[string]bool{}
		}
	}

	for _, mod := range modules {
		for _, dest := range mod.dest {
			if modules[dest].typ == Conjunction {
				memory[dest][mod.name] = false
			}
		}
	}

	return modules, states, memory
}

func pushButton(modules map[string]Module, states map[string]bool, memory map[string]map[string]bool) (low int, high int, lowRx map[string]bool) {
	queue := []Pulse{{"broadcaster", "button", false}}
	lowRx = map[string]bool{}

	for len(queue) > 0 {
		pulse := queue[0]
		if pulse.high {
			if slices.Contains(modules[pulse.dest].dest, "rx") {
				lowRx[pulse.src] = true
			}

			high++
		} else {
			low++
		}

		queue = queue[1:]
		module := modules[pulse.dest]

		switch module.typ {
		case Broadcast:
			for _, dest := range module.dest {
				queue = append(queue, Pulse{dest, module.name, pulse.high})
			}
		case Conjunction:
			memory[module.name][pulse.src] = pulse.high

			isHighPulse := true
			for _, isHigh := range memory[module.name] {
				if !isHigh {
					isHighPulse = false
					break
				}
			}

			for _, dest := range module.dest {
				queue = append(queue, Pulse{dest, module.name, !isHighPulse})
			}
		case FlipFlop:
			if pulse.high {
				continue
			}

			states[module.name] = !states[module.name]
			for _, dest := range module.dest {
				queue = append(queue, Pulse{dest, module.name, states[module.name]})
			}
		}
	}

	return low, high, lowRx
}

func part1() int {
	modules, states, memory := parseInput()

	lowPulses := 0
	highPulses := 0

	for i := 0; i < 1000; i++ {
		low, high, _ := pushButton(modules, states, memory)
		lowPulses += low
		highPulses += high
	}

	return highPulses * lowPulses
}

func part2() int {
	modules, states, memory := parseInput()

	rxIndices := map[string][]int{}
	for i := 0; i < 100_000; i++ {
		_, _, lowRx := pushButton(modules, states, memory)
		for mod := range lowRx {
			rxIndices[mod] = append(rxIndices[mod], i)
		}
	}

	arr := []int{}
	for _, indices := range rxIndices {
		arr = append(arr, indices[1]-indices[0])
	}

	return utils.LCM(arr[0], arr[1], arr[2:]...)
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
