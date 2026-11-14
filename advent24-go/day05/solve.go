package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	SMALLER = 1
	GREATER = 2
)

type Pair struct {
	x, y int
}

func parseInput(input string) (map[Pair]int, [][]int) {
	var order = make(map[Pair]int)
	var cases [][]int

	segments := strings.Split(input, "\n\n")
	// Read order
	for line := range strings.Lines(segments[0]) {
		vals := strings.FieldsFunc(line, func(r rune) bool {
			return r == '|'
		})
		before, err_one := strconv.Atoi(strings.TrimSpace(vals[0]))
		after, err_two := strconv.Atoi(strings.TrimSpace(vals[1]))

		if err_one != nil || err_two != nil {
			fmt.Println("Could not parse int (order)")
		}

		order[Pair{before, after}] = SMALLER
		order[Pair{after, before}] = GREATER
	}

	// Read cases
	for line := range strings.Lines(segments[1]) {
		vals := strings.FieldsFunc(line, func(r rune) bool {
			return r == ','
		})
		ints := make([]int, len(vals))
		for i, s := range vals {
			n, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				fmt.Println("Could not parse int (cases)")
			}
			ints[i] = n
		}

		cases = append(cases, ints)
	}
	return order, cases
}

func getCmp(order map[Pair]int) func(int, int) int {
	return func(a, b int) int {
		switch order[Pair{a, b}] {
		case SMALLER:
			return -1
		case GREATER:
			return 1
		default:
			return 0
		}
	}
}

func getMedian(list []int) int {
	mid := len(list) / 2
	if len(list)%2 == 0 {
		return list[mid-1]
	} else {
		return list[mid]
	}
}

func SolvePartOne(input string) int {
	order, cases := parseInput(input)
	sum := 0
	for _, ints := range cases {
		if slices.IsSortedFunc(ints, getCmp(order)) {
			sum += getMedian(ints)
		}
	}
	return sum
}

func SolvePartTwo(input string) int {
	order, cases := parseInput(input)
	sum := 0
	cmp := getCmp(order)
	for _, ints := range cases {
		if !slices.IsSortedFunc(ints, cmp) {
			slices.SortFunc(ints, cmp)
			sum += getMedian(ints)
		}
	}
	return sum
}
