package day12_test

import (
    "testing"
    "time"

    "github.com/ushahid/adventofcode/advent24-go/day12"
    "github.com/ushahid/adventofcode/advent24-go/utils"
)

func TestSolvePartOne(t *testing.T) {
    input := utils.ReadInput(t)
    start := time.Now()
    result := day12.SolvePartOne(input)
    elapsed := time.Since(start)
    utils.LogTime(t, elapsed)
    utils.LogResult(t, result)
}

func TestSolvePartTwo(t *testing.T) {
    input := utils.ReadInput(t)
    start := time.Now()
    result := day12.SolvePartTwo(input)
    elapsed := time.Since(start)
    utils.LogTime(t, elapsed)
    utils.LogResult(t, result)
}
