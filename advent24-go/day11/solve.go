package day11

import (
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	vals := strings.Fields(input)
	stones := make([]int, 0, len(vals))
	for idx := range vals {
		stone, _ := strconv.Atoi(vals[idx])
		stones = append(stones, stone)
	}
	return stones
}

func SolvePartOne(input string) int {
	// This is the old brute force solution that works for Part 1 but not Part 2
	// stones := parseInput(input)
	// for range 25 {
	// 	i := 0
	// 	for i < len(stones) {
	// 		str := strconv.Itoa(stones[i])
	// 		if stones[i] == 0 {
	// 			stones[i] = 1
	// 		} else if len(str)%2 == 0 {
	// 			stones[i], _ = strconv.Atoi(str[:len(str)/2])
	// 			next, _ := strconv.Atoi(str[len(str)/2:])
	// 			stones = slices.Insert(stones, i+1, next)
	// 			i++
	// 		} else {
	// 			stones[i] *= 2024
	// 		}
	// 		i++
	// 	}
	// }
	// return len(stones)

	stones := parseInput(input)
	totalStones := 0
	seen := make(map[StoneData]int)
	for i := range stones {
		totalStones += countStones(stones[i], 25, seen)
	}
	return totalStones
}

type StoneData struct {
	seed int
	iter int
}

func countStones(seed int, remIter int, seen map[StoneData]int) int {
	seen[StoneData{seed, 0}] = 1
	if count, exists := seen[StoneData{seed, remIter}]; exists {
		return count
	}

	str := strconv.Itoa(seed)
	if seed == 0 {
		next := 1
		c := countStones(next, remIter-1, seen)
		seen[StoneData{next, remIter - 1}] = c
		return c
	} else if len(str)%2 == 0 {
		nextFst, _ := strconv.Atoi(str[:len(str)/2])
		nextSec, _ := strconv.Atoi(str[len(str)/2:])
		c1 := countStones(nextFst, remIter-1, seen)
		c2 := countStones(nextSec, remIter-1, seen)
		seen[StoneData{nextFst, remIter - 1}] = c1
		seen[StoneData{nextSec, remIter - 1}] = c2
		return c1 + c2
	} else {
		next := seed * 2024
		c := countStones(next, remIter-1, seen)
		seen[StoneData{next, remIter - 1}] = c
		return c
	}
}

func SolvePartTwo(input string) int {
	stones := parseInput(input)
	totalStones := 0
	seen := make(map[StoneData]int)
	for i := range stones {
		totalStones += countStones(stones[i], 75, seen)
	}
	return totalStones
}
