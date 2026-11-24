package day19

import (
	"strings"
)

func parseInput(input string) ([]string, []string) {
	vals := strings.Split(input, "\n\n")
	return strings.Split(vals[0], ", "), strings.Split(vals[1], "\n")
}

func numArrangements(towels []string, pattern string) int {
	dp := make([]int, len(pattern)+1)
	dp[0] = 1
	for i := 0; i < len(pattern)+1; i++ {
		for _, towel := range towels {
			if i-len(towel) >= 0 {
				if pattern[i-len(towel):i] == towel {
					dp[i] += dp[i-len(towel)]
				}
			}
		}
	}
	return dp[len(pattern)]
}

func SolvePartOne(input string) int {
	towels, patterns := parseInput(input)
	total := 0
	for _, pattern := range patterns {
		if numArrangements(towels, pattern) > 0 {
			total++
		}
	}
	return total
}

func SolvePartTwo(input string) int {
	towels, patterns := parseInput(input)
	total := 0
	for _, pattern := range patterns {
		total += numArrangements(towels, pattern)
	}
	return total
}
