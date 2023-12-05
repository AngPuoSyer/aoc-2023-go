package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	buf, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := string(buf)
	part1(input)
	part2(input)
}

func part1(input string) int {
	total := int(0)

	delimiterRegex := regexp.MustCompile("[:|]")

	for _, line := range strings.Split(input, "\n") {
		split := delimiterRegex.Split(line, -1)

		numberRegexStr := "\\b(?:"
		for _, num := range strings.Fields(strings.TrimSpace(split[1])) {
			numberRegexStr += strings.TrimSpace(num) + "|"
		}

		numberRegexStr = numberRegexStr[:len(numberRegexStr)-1] + ")\\b"

		numberRegex := regexp.MustCompile(numberRegexStr)

		found := numberRegex.FindAllString(split[2], -1)
		foundLen := len(found)

		if foundLen > 0 {
			total += int(math.Pow(2, float64(foundLen-1)))
		}
	}
	fmt.Printf("Part 1: %d\n", total)
	return total
}

func part2(input string) int {
	total := int(0)
	cards := make(map[int]int)

	delimiterRegex := regexp.MustCompile("[:|]")

	for cardNo, line := range strings.Split(input, "\n") {
		_, ok := cards[cardNo]
		if !ok {
			cards[cardNo] = 0
		}
		cards[cardNo] += 1

		split := delimiterRegex.Split(line, -1)

		numberRegexStr := "\\b(?:"
		for _, num := range strings.Fields(strings.TrimSpace(split[1])) {
			numberRegexStr += strings.TrimSpace(num) + "|"
		}

		numberRegexStr = numberRegexStr[:len(numberRegexStr)-1] + ")\\b"

		numberRegex := regexp.MustCompile(numberRegexStr)

		found := numberRegex.FindAllString(split[2], -1)

		foundLen := len(found)

		if foundLen > 0 {
			count := cards[cardNo]
			for j := cardNo + 1; j <= cardNo+foundLen; j++ {
				cards[j] += count
			}

		}
	}
	for _, val := range cards {
		total += val
	}
	fmt.Printf("Part 2: %d\n", total)
	return total
}
