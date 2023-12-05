package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getSeeds() []int {
	input := strings.TrimPrefix(input, "seeds: ")
	row := strings.Split(input, "\n")[0]
	raw := strings.Fields(row)

	seeds := []int{}
	for _, num := range raw {
		seed, _ := strconv.Atoi(num)
		seeds = append(seeds, seed)
	}

	return seeds
}

func getSeedChunks() (chunks [][]int) {
	seeds := getSeeds()

	for i := 0; i < len(seeds); i += 2 {
		chunks = append(chunks, seeds[i:i+2])
	}

	return chunks
}

func getGenericMap(name string) [][]int {
	genericMap := [][]int{}

	target := name + " map:"
	idx := strings.Index(input, target)
	chunks := strings.Split(input[idx+len(target):], "\n\n")
	rows := strings.Split(strings.Trim(chunks[0], "\n"), "\n")

	for _, row := range rows {
		raw := strings.Fields(row)
		destStart, _ := strconv.Atoi(raw[0])
		sourceStart, _ := strconv.Atoi(raw[1])
		rangeLength, _ := strconv.Atoi(raw[2])

		genericMap = append(genericMap, []int{destStart, sourceStart, rangeLength})
	}

	return genericMap
}

func convertValue(value int, mapArr [][]int) int {
	for _, mapRow := range mapArr {
		destStart := mapRow[0]
		sourceStart := mapRow[1]
		rangeLength := mapRow[2]

		sourceEnd := sourceStart + rangeLength

		if value >= sourceStart && value < sourceEnd {
			return destStart + (value - sourceStart)
		}
	}

	return value
}

func convertValueRange(start int, length int, mapArr [][]int) (result [][]int) {
	for _, mapRow := range mapArr {
		destStart := mapRow[0]
		sourceStart := mapRow[1]
		rangeLength := mapRow[2]

		sourceEnd := sourceStart + rangeLength

		if start >= sourceStart && start < sourceEnd {
			end := start + length
			if end > sourceEnd {
				end = sourceEnd
			}

			result = append(result, []int{destStart + (start - sourceStart), end - start})
		} else if sourceStart >= start && sourceStart < start+length {
			end := sourceEnd
			if end > start+length {
				end = start + length
			}

			result = append(result, []int{destStart, end - sourceStart})
		}
	}

	if len(result) == 0 {
		result = append(result, []int{start, length})
	}

	return result
}

func getSeedToSoilMap() [][]int {
	return getGenericMap("seed-to-soil")
}

func getSoilToFertilizerMap() [][]int {
	return getGenericMap("soil-to-fertilizer")
}

func getFertilizerToWaterMap() [][]int {
	return getGenericMap("fertilizer-to-water")
}

func getWaterToLightMap() [][]int {
	return getGenericMap("water-to-light")
}

func getLightToTemperatureMap() [][]int {
	return getGenericMap("light-to-temperature")
}

func getTemperatureToHumidityMap() [][]int {
	return getGenericMap("temperature-to-humidity")
}

func getHumidityToLocationMap() [][]int {
	return getGenericMap("humidity-to-location")
}

func part1() int {
	soilMap := getSeedToSoilMap()
	fertilizerMap := getSoilToFertilizerMap()
	waterMap := getFertilizerToWaterMap()
	lightMap := getWaterToLightMap()
	temperatureMap := getLightToTemperatureMap()
	humidityMap := getTemperatureToHumidityMap()
	locationMap := getHumidityToLocationMap()

	lowestLocation := -1
	for _, seed := range getSeeds() {
		soil := convertValue(seed, soilMap)
		fertilizer := convertValue(soil, fertilizerMap)
		water := convertValue(fertilizer, waterMap)
		light := convertValue(water, lightMap)
		temperature := convertValue(light, temperatureMap)
		humidity := convertValue(temperature, humidityMap)
		location := convertValue(humidity, locationMap)

		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func part2() int {
	soilMap := getSeedToSoilMap()
	fertilizerMap := getSoilToFertilizerMap()
	waterMap := getFertilizerToWaterMap()
	lightMap := getWaterToLightMap()
	temperatureMap := getLightToTemperatureMap()
	humidityMap := getTemperatureToHumidityMap()
	locationMap := getHumidityToLocationMap()

	lowestLocation := -1

	// We shall never speak of this again
	for _, chunk := range getSeedChunks() {
		for _, soil := range convertValueRange(chunk[0], chunk[1], soilMap) {
			for _, fertilizer := range convertValueRange(soil[0], soil[1], fertilizerMap) {
				for _, water := range convertValueRange(fertilizer[0], fertilizer[1], waterMap) {
					for _, light := range convertValueRange(water[0], water[1], lightMap) {
						for _, temperature := range convertValueRange(light[0], light[1], temperatureMap) {
							for _, humidity := range convertValueRange(temperature[0], temperature[1], humidityMap) {
								for _, location := range convertValueRange(humidity[0], humidity[1], locationMap) {
									if lowestLocation == -1 || location[0] < lowestLocation {
										lowestLocation = location[0]
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return lowestLocation
}

func main() {
	part1Result := part1()
	part2Result := part2()

	log.Printf("Part 1: %d", part1Result)
	log.Printf("Part 2: %d", part2Result)
}
