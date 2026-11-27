package day25

import (
	"fmt"
	"strings"
)

type Item struct {
	heights [5]int
}

func parseInput(input string) ([]Item, []Item) {
	keys := make([]Item, 0, 250)
	locks := make([]Item, 0, 250)
	numRows := 7
	numCols := 5
	for _, itemStr := range strings.Split(input, "\n\n") {
		var i Item
		itemStr := strings.Split(itemStr, "\n")
		for c := range numCols {
			numHash := 0
			for r := range numRows {
				if itemStr[r][c] == '#' {
					numHash++
				}
			}
			i.heights[c] = numHash - 1
		}
		if itemStr[0][0] == '#' {
			locks = append(locks, i)
		} else {
			keys = append(keys, i)
		}
	}
	return keys, locks
}

func getLock(key Item) Item {
	var lock Item
	for idx, h := range key.heights {
		lock.heights[idx] = 5 - h
	}
	return lock
}

func fits(key Item, lock Item) bool {
	numRows := 5
	numCols := 5
	for c := range numRows {
		if lock.heights[c]+key.heights[c] > numCols {
			return false
		}
	}
	return true
}

func SolvePartOne(input string) int {
	keys, locks := parseInput(input)
	fmt.Println(len(keys), len(locks))
	total := 0
	for _, key := range keys {
		for _, lock := range locks {
			if fits(key, lock) {
				total++
			}
		}
	}
	return total
}

func SolvePartTwo(input string) int {
	return 0
}
