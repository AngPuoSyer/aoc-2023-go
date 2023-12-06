package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/alphadose/itogami"
)

type GardeningConfigMap struct {
	source      int64
	destination int64
	validRange  int64
}

type SeedConfig struct {
	seed       int64
	soil       int64
	fertilizer int64
	water      int64
	light      int64
	humidity   int64
	location   int64
}

func main() {
	buf, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := string(buf)
	part1(input)
	part2(input)
}

func part1(input string) int64 {
	minLocation := int64(math.MaxInt64)

	seeds,
		seedToSoilMap,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemp,
		tempToHumid,
		humidToLocation := parseMaps(input)

	for _, seed := range seeds {
		soil := getConversation(seedToSoilMap, seed)
		fertilizer := getConversation(soilToFertilizer, soil)
		water := getConversation(fertilizerToWater, fertilizer)
		light := getConversation(waterToLight, water)
		temp := getConversation(lightToTemp, light)
		humid := getConversation(tempToHumid, temp)
		location := getConversation(humidToLocation, humid)

		if location < minLocation {
			minLocation = location
		}
	}
	fmt.Printf("Part 1: %d\n", minLocation)

	return minLocation
}

func part2(input string) int64 {
	runtime.GOMAXPROCS(runtime.NumCPU())
	minLocation := int64(math.MaxInt64)
	var lock sync.Mutex
	var wg sync.WaitGroup

	pool := itogami.NewPool(100)

	seeds,
		seedToSoilMap,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemp,
		tempToHumid,
		humidToLocation := parseMaps(input)

	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			wg.Add(1)

			tempFunc := func(idx int64) func() {
				return func() {
					defer wg.Done()
					fmt.Printf("Seed[%d]\n", idx)
					soil := getConversation(seedToSoilMap, idx)
					fertilizer := getConversation(soilToFertilizer, soil)
					water := getConversation(fertilizerToWater, fertilizer)
					light := getConversation(waterToLight, water)
					temp := getConversation(lightToTemp, light)
					humid := getConversation(tempToHumid, temp)
					location := getConversation(humidToLocation, humid)
					lock.Lock()
					if location < minLocation {
						minLocation = location
					}
					lock.Unlock()
				}
			}(j)

			pool.Submit(tempFunc)
		}
	}

	wg.Wait()
	fmt.Printf("Part 2: %d\n", minLocation)

	return minLocation
}

func parseMaps(input string) (
	seeds []int64,
	seedToSoilMap []GardeningConfigMap,
	soilToFertilizer []GardeningConfigMap,
	fertilizerToWater []GardeningConfigMap,
	waterToLight []GardeningConfigMap,
	lightToTemp []GardeningConfigMap,
	tempToHumid []GardeningConfigMap,
	humidToLocation []GardeningConfigMap,
) {
	strArr := strings.Split(input, "\n\n")
	seedArr := strings.Split(strings.TrimSpace(strArr[0][7:]), " ")

	for _, seed := range seedArr {
		seedNum, _ := strconv.Atoi(seed)
		seeds = append(seeds, int64(seedNum))
	}

	for _, chunk := range strArr[1:] {
		trimmedChunk := strings.TrimSpace(chunk)
		if strings.HasPrefix(trimmedChunk, "seed-to-soil") {
			seedToSoilMap = parseMapValue(trimmedChunk)
		} else if strings.HasPrefix(trimmedChunk, "soil-to-fertilizer") {
			soilToFertilizer = parseMapValue(trimmedChunk)
		} else if strings.HasPrefix(trimmedChunk, "fertilizer-to-water") {
			fertilizerToWater = parseMapValue(trimmedChunk)
		} else if strings.HasPrefix(trimmedChunk, "water-to-light") {
			waterToLight = parseMapValue(trimmedChunk)
		} else if strings.HasPrefix(trimmedChunk, "light-to-temperature") {
			lightToTemp = parseMapValue(trimmedChunk)
		} else if strings.HasPrefix(trimmedChunk, "temperature-to-humidity") {
			tempToHumid = parseMapValue(trimmedChunk)
		} else if strings.HasPrefix(trimmedChunk, "humidity-to-location") {
			humidToLocation = parseMapValue(trimmedChunk)
		}
	}

	return seeds, seedToSoilMap, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumid, humidToLocation
}

func parseMapValue(chunk string) []GardeningConfigMap {
	var configMaps []GardeningConfigMap
	for _, line := range strings.Split(chunk, "\n") {
		if strings.Contains(line, "map") {
			continue
		}

		var configParamTuple []int64

		for _, val := range strings.Fields(line) {
			num, _ := strconv.Atoi(val)
			configParamTuple = append(configParamTuple, int64(num))
		}

		configMap := GardeningConfigMap{
			source:      configParamTuple[1],
			destination: configParamTuple[0],
			validRange:  configParamTuple[2],
		}
		configMaps = append(configMaps, configMap)
	}
	return configMaps
}

func getConversation(configs []GardeningConfigMap, query int64) int64 {
	for _, config := range configs {
		if query >= config.source && query < config.source+config.validRange {
			return config.destination + query - config.source
		}
	}

	return query
}

func processPart2Seed(configs []GardeningConfigMap, query int64) int64 {
	for _, config := range configs {
		if query >= config.source && query < config.source+config.validRange {
			return config.destination + query - config.source
		}
	}

	return query
}
