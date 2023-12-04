package main

import "testing"

func TestPart1(t *testing.T) {
	input := `467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	ans := part1(input)

	if ans != 4361 {
		t.Fatal("Wrong answer")
	}
}
