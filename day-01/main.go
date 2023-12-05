package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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
	totalSum := 0

	var numbersInLine []int

	for _, line := range input {

		for _, i := range line {
			num, err := strconv.Atoi(string(i))

			if err != nil {
				continue
			} else {
				numbersInLine = append(numbersInLine, num)
			}
		}
		firstDigit := numbersInLine[0]
		lastDigit := numbersInLine[len(numbersInLine)-1]

		doubleDigit := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		num, _ := strconv.Atoi(doubleDigit)
		totalSum += num
	}

	fmt.Printf("Part 1: %d\n", totalSum)
}

func part2(input []string) {
	totalSum := 0
	re := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for _, line := range input {
		var numbersInLine []string
		for i := range line {
			found := re.FindString(line[i:])
			if found != "" {
				numbersInLine = append(numbersInLine, found)
			}
		}

		for index, word := range numbersInLine {
			_, err := strconv.Atoi(word)
			if err != nil && word != "" {
				numbersInLine[index] = wordToNum(word)
			}
		}

		firstDigit := numbersInLine[0]
		lastDigit := numbersInLine[len(numbersInLine)-1]

		doubleDigit := fmt.Sprintf("%s%s", firstDigit, lastDigit)
		num, _ := strconv.Atoi(doubleDigit)
		totalSum += num
	}
	fmt.Printf("Part 2: %d\n", totalSum)
}

func wordToNum(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	default:
		return ""
	}
}
