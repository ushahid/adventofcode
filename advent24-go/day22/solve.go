package day22

import (
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	strNums := strings.Split(input, "\n")
	nums := make([]int, 0, len(strNums))
	for _, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)
		nums = append(nums, num)
	}
	return nums
}

func mixAndPrune(prev int, res int) int {
	return (prev ^ res) % 16777216
}

func nextRandom(seed int) int {
	num := mixAndPrune(seed, seed*64)
	num = mixAndPrune(num, num/32)
	num = mixAndPrune(num, num*2048)
	return num
}

func SolvePartOne(input string) int {
	seeds := parseInput(input)
	total := 0
	for _, seed := range seeds {
		num := seed
		for range 2000 {
			num = nextRandom(num)
		}
		total += num
	}
	return total
}

func SolvePartTwo(input string) int {
	seeds := parseInput(input)

	const keySpace = 19 * 19 * 19 * 19
	totals := make([]int, keySpace)
	seen := make([]int, keySpace)

	for buyerIdx, seed := range seeds {
		buyerID := buyerIdx + 1
		num := seed
		prevPrice := num % 10
		key := 0

		for step := range 2000 {
			num = nextRandom(num)
			price := num % 10
			delta := price - prevPrice + 9
			prevPrice = price
			key = ((key * 19) + delta) % keySpace

			if step >= 3 && seen[key] != buyerID {
				seen[key] = buyerID
				totals[key] += price
			}
		}
	}

	maxBananas := 0
	for _, total := range totals {
		if total > maxBananas {
			maxBananas = total
		}
	}
	return maxBananas
}
