package main

import (
	_ "embed"
	"log"
	"math/rand"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Component struct {
	name     string
	attached map[string]bool
}

type Group struct {
	names    map[string]bool
	attached map[string]bool
}

func getInputLines() []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func parseInput() []Component {
	lines := getInputLines()
	components := []Component{}

	for _, line := range lines {
		fields := strings.Fields(line)
		name := strings.TrimSuffix(fields[0], ":")
		attached := map[string]bool{}

		for _, field := range fields[1:] {
			attached[field] = true
		}

		components = append(components, Component{name, attached})
	}

	for _, component := range components {
		for attached := range component.attached {
			i := slices.IndexFunc(components, func(c Component) bool {
				return c.name == attached
			})

			if i == -1 {
				components = append(components, Component{attached, map[string]bool{}})
				i = len(components) - 1
			}

			components[i].attached[component.name] = true
		}
	}

	return components
}

func getRandomKey(m map[string]bool) string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys[rand.Intn(len(keys))]
}

func mergeComponents(groups []Group, index1 int, index2 int) []Group {
	group1 := groups[index1]
	group2 := groups[index2]

	names := union(group1.names, group2.names)
	attached := union(group1.attached, group2.attached)
	attached = exclude(attached, names)

	newGroups := []Group{}
	for i, group := range groups {
		if i == index1 || i == index2 {
			continue
		}

		newGroups = append(newGroups, group)
	}

	return append(newGroups, Group{names, attached})
}

func union(a map[string]bool, b map[string]bool) map[string]bool {
	result := map[string]bool{}

	for key := range a {
		result[key] = true
	}

	for key := range b {
		result[key] = true
	}

	return result
}

func exclude(a map[string]bool, b map[string]bool) map[string]bool {
	result := map[string]bool{}

	for key := range a {
		if !b[key] {
			result[key] = true
		}
	}

	return result
}

func kargerMinCut(components []Component) int {
	a := []Group{}
	for _, component := range components {
		a = append(a, Group{
			names:    map[string]bool{component.name: true},
			attached: component.attached,
		})
	}

	for len(a) > 2 {
		index := rand.Intn(len(a))
		component := a[index]

		attached := getRandomKey(component.attached)
		attachedIndex := slices.IndexFunc(a, func(c Group) bool {
			return c.names[attached]
		})

		a = mergeComponents(a, index, attachedIndex)
	}

	if len(a[0].attached) == 3 || len(a[1].attached) == 3 {
		return len(a[0].names) * len(a[1].names)
	}

	return kargerMinCut(components)
}

func part1() int {
	components := parseInput()

	// If you are unlucky, you get the wrong answer. Try again.
	return kargerMinCut(components)
}

func main() {
	log.Printf("Part 1: %d", part1())
}
