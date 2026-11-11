package day02

import (
	"bufio"
	"fmt"
	"math"
	"slices"
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

func isSafe(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	increasing := true
	if arr[0] > arr[1] {
		increasing = false
	}
	for i := 0; i < len(arr)-1; i++ {
		if (increasing && arr[i+1] > arr[i]) || (!increasing && arr[i] > arr[i+1]) {
			diff := int(math.Round(math.Abs(float64(arr[i] - arr[i+1]))))
			if !(diff >= 1 && diff <= 3) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func SolvePartOne(input string) int {
	vals := parseInput(input)
	safeCount := 0
	for row := range vals {
		if isSafe(vals[row]) {
			safeCount++
		}
	}
	return safeCount
}

func SolvePartTwo(input string) int {
	vals := parseInput(input)
	safeCount := 0
	for row := range vals {
		data := slices.Clone(vals[row])
		idx := -1
		for !isSafe(data) {
			idx++
			if idx < len(vals[row]) {
				data = slices.Clone(vals[row])
				data = slices.Delete(data, idx, idx+1)
			} else {
				break
			}
		}
		if idx < len(vals[row]) {
			safeCount++
		}
	}
	return safeCount
}
