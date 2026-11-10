package day01

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func getDistribution(numbers []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range numbers {
		counts[num]++
	}
	return counts
}

func parseInput(input string) ([]int, []int) {
	var first []int
	var second []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		vals := strings.Fields(scanner.Text())
		left, err := strconv.Atoi(vals[0])
		if err != nil {
			fmt.Println("Could not parse integer:", left)
		}
		right, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("Could not parse integer:", right)
		}
		first = append(first, left)
		second = append(second, right)
	}
	return first, second
}

func SolvePartOne(input string) int {
	left, right := parseInput(input)
	slices.Sort(left[:])
	slices.Sort(right[:])
	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i]) - float64(right[i])))
	}
	return distance
}

func SolvePartTwo(input string) int {
	left, right := parseInput(input)
	counts := getDistribution(right)
	distance := 0
	for _, val := range left {
		// defaults to 0 if key is not found
		distance += val * counts[val]
	}
	return distance
}
