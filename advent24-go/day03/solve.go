package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolvePartOne(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	sum := 0
	for line := range strings.Lines(input) {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, expr := range matches {
			left, err_left := strconv.Atoi(expr[1])
			right, err_right := strconv.Atoi(expr[2])
			if err_left != nil || err_right != nil {
				fmt.Println("Numerical parsing error")
			}
			sum += left * right
		}
	}
	return sum
}

func SolvePartTwo(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)
	sum := 0
	enabled := true
	for line := range strings.Lines(input) {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, expr := range matches {
			if strings.HasPrefix(expr[0], "don't") {
				enabled = false
			} else if strings.HasPrefix(expr[0], "do") {
				enabled = true
			} else if enabled {
				left, err_left := strconv.Atoi(expr[1])
				right, err_right := strconv.Atoi(expr[2])
				if err_left != nil || err_right != nil {
					fmt.Println("Numerical parsing error")
				}
				sum += left * right
			}
		}
	}
	return sum
}
