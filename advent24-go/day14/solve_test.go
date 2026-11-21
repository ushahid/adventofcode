package day14_test

import (
    "testing"
    "time"

    "github.com/ushahid/adventofcode/advent24-go/day14"
    "github.com/ushahid/adventofcode/advent24-go/utils"
)

func TestSolvePartOne(t *testing.T) {
    input := utils.ReadInput(t)
    start := time.Now()
    result := day14.SolvePartOne(input)
    elapsed := time.Since(start)
    utils.LogTime(t, elapsed)
    utils.LogResult(t, result)
}

func TestSolvePartTwo(t *testing.T) {
    input := utils.ReadInput(t)
    start := time.Now()
    result := day14.SolvePartTwo(input)
    elapsed := time.Since(start)
    utils.LogTime(t, elapsed)
    utils.LogResult(t, result)
}
