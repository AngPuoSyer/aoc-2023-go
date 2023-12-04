package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type RowNumRange struct {
	x []int
	y int
}

func main() {}

func part1(input string) int {
	rows := strings.Split(input, "\n")
	var numberArray [][][]int
	var symbolsArray []Coordinate

	symbolRegex := regexp.MustCompile(`[^0-9.]`)
	intRegex := regexp.MustCompile(`[0-9]+`)

	for rowIndex, row := range rows {
		str := strings.TrimSpace(row)
		charSlice := strings.Split(str, "")

		intCoors := intRegex.FindAllStringIndex(str, -1)

		numberArray = append(numberArray, intCoors)

		for colIndex, char := range charSlice {
			if symbolRegex.MatchString(char) {
				symbolsArray = append(symbolsArray, Coordinate{
					x: colIndex,
					y: rowIndex,
				})
			}
		}
	}

	total := 0
	fmt.Printf("%v\n", numberArray)
	fmt.Printf("%v\n", symbolsArray)

	for yCoor, num := range numberArray {
		boxes := [][]int{
			{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
		}
		for _, translation := range boxes {
			for _, n := range num {
				for i := n[0]; i < n[1]; i++ {
					tempX := i + translation[0]
					tempY := yCoor + translation[1]

				}
			}
		}
	}

	for _, sym := range symbolsArray {
		boxes := [][]int{
			{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
		}
		for _, translation := range boxes {
			tempX := sym.x + translation[0]
			tempY := sym.y + translation[1]

			fmt.Printf("%d, ", tempX)
			fmt.Printf("%d\n", tempY)

			for _, numCoor := range numberArray[tempY] {
				if numCoor[0] <= tempX && tempX < numCoor[1] {
					num, _ := strconv.Atoi(rows[tempY][numCoor[0]:numCoor[1]])
					fmt.Printf("num: %d\n", num)
					total += num
				}
			}
		}
	}

	fmt.Printf("%d\n", total)
	return total
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
