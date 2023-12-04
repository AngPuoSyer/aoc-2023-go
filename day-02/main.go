package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CubeCombination struct {
	red   int
	green int
	blue  int
}

func main() {
	input_file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer input_file.Close()
	sc := bufio.NewScanner(input_file)

	var input []string

	for sc.Scan() {
		input = append(input, sc.Text())
	}

	part1(input)
	part2(input)
}

func part1(input []string) {
	validGames := 0

	maxCubeConfig := CubeCombination{
		red:   12,
		green: 13,
		blue:  14,
	}
	for _, line := range input {
		game, prompt := splitGame(line)
		cubes := getMaxCubes(prompt)

		if isValidGame(cubes, maxCubeConfig) {
			validGames += game
		}
	}

	fmt.Printf("Part 1: %d\n", validGames)
}

func part2(input []string) {
	power := int64(0)

	for _, line := range input {
		_, prompt := splitGame(line)
		cubes := getMaxCubes(prompt)

		power += getPower(cubes)

	}

	fmt.Printf("Part 2: %d\n", power)
}

func splitGame(line string) (game int, prompt string) {
	promptSlice := strings.Split(line, ":")
	gameString := promptSlice[0]
	prompt = promptSlice[1]

	gameNum, _ := strconv.Atoi(strings.Replace(gameString, "Game ", "", 1))
	return gameNum, prompt
}

func getMaxCubes(prompt string) CubeCombination {
	cubeMaxTuple := []int{0, 0, 0}

	for _, set := range strings.Split(prompt, ";") {
		for _, colourCountPair := range strings.Split(set, ",") {
			colourCountPair = strings.TrimSpace(colourCountPair)

			tempStr := strings.Split(colourCountPair, " ")
			currMax, _ := strconv.Atoi(tempStr[0])
			colr := tempStr[1]
			idx := colourToIndex(colr)
			if currMax > cubeMaxTuple[idx] {
				cubeMaxTuple[idx] = currMax
			}
		}
	}
	return CubeCombination{
		red:   cubeMaxTuple[0],
		green: cubeMaxTuple[1],
		blue:  cubeMaxTuple[2],
	}
}

func colourToIndex(colour string) (idx int) {
	switch colour {
	case "red":
		return 0
	case "green":
		return 1
	case "blue":
		return 2
	default:
		return -1
	}
}

func isValidGame(cubes CubeCombination, maxCubesConfig CubeCombination) bool {
	if cubes.red > maxCubesConfig.red {
		return false
	}
	if cubes.green > maxCubesConfig.green {
		return false
	}
	if cubes.blue > maxCubesConfig.blue {
		return false
	}

	return true
}

func getPower(cubes CubeCombination) int64 {
	return int64(cubes.blue) * int64(cubes.red) * int64(cubes.green)
}
