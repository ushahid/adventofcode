package day07

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, [][]int) {
	var targets []int
	var inputs [][]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		vals := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return r == ':' || r == ' '
		})
		target, err := strconv.Atoi(vals[0])
		if err != nil {
			fmt.Println("Failed to parse int")
		}
		targets = append(targets, target)
		numbersRow := make([]int, len(vals)-1)
		for idx := range vals[1:] {
			target, err := strconv.Atoi(vals[idx+1])
			if err != nil {
				fmt.Println("Failed to parse int")
			}
			numbersRow[idx] = target
		}
		inputs = append(inputs, numbersRow)
	}
	return targets, inputs
}

func concat(a int, b int) int {
	val, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return val
}

func possiblyTrueOne(target int, inputs []int, current int) bool {
	if current > target {
		return false
	}
	if len(inputs) == 1 {
		return (inputs[0]+current == target) || (inputs[0]*current == target)
	}
	return possiblyTrueOne(target, inputs[1:], inputs[0]+current) || possiblyTrueOne(target, inputs[1:], inputs[0]*current)
}

func possiblyTrueTwo(target int, inputs []int, current int) bool {
	if current > target {
		return false
	}
	if len(inputs) == 1 {
		return (inputs[0]+current == target) || (inputs[0]*current == target) || (concat(current, inputs[0]) == target)
	}
	return possiblyTrueTwo(target, inputs[1:], inputs[0]+current) ||
		possiblyTrueTwo(target, inputs[1:], inputs[0]*current) ||
		possiblyTrueTwo(target, inputs[1:], concat(current, inputs[0]))
}

func SolvePartOne(input string) int {
	targets, inputsList := parseInput(input)
	totalValid := 0
	for idx := range targets {
		if possiblyTrueOne(targets[idx], inputsList[idx], 0) {
			totalValid += targets[idx]
		}
	}
	return totalValid
}

func SolvePartTwo(input string) int {
	targets, inputsList := parseInput(input)
	totalValid := 0
	for idx := range targets {
		if possiblyTrueTwo(targets[idx], inputsList[idx], 0) {
			totalValid += targets[idx]
		}
	}
	return totalValid
}
