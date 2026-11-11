package day02

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	var rows [][]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		intRow := make([]int, len(row))
		for i, s := range row {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error parsing string %q at index %d: %v\n", s, i, err)
			}
			intRow[i] = num
		}
		rows = append(rows, intRow)
	}
	return rows
}

func SolvePartOne(input string) int {
	vals := parseInput(input)
	for idx, row := range vals {

	}
	return 0
}

func SolvePartTwo(input string) int {
	return 0
}
