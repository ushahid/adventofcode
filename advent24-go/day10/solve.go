package day10

import (
	"slices"
	"strings"
)

type Position struct {
	row, col int
}

func parseInput(input string) []string {
	topoMap := strings.Split(input, "\n")
	return topoMap
}

func SolvePartOne(input string) int {
	topoMap := parseInput(input)
	numRows := len(topoMap)
	numCols := len(topoMap[0])
	starts := make([]Position, 0, 100)
	for r := range numRows {
		for c := range numCols {
			if topoMap[r][c] == '0' {
				starts = append(starts, Position{r, c})
			}
		}
	}
	total := 0
	for _, start := range starts {
		horizon := make([]Position, 1, 100)
		horizon[0] = start
		seen := make(map[Position]struct{})
		for len(horizon) > 0 {
			curr := horizon[len(horizon)-1]
			horizon = horizon[:len(horizon)-1]
			for _, delta := range []Position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				next := Position{curr.row + delta.row, curr.col + delta.col}
				_, exists := seen[next]
				if !exists && next.row >= 0 && next.row < numRows && next.col >= 0 && next.col < numCols &&
					int(topoMap[next.row][next.col])-int(topoMap[curr.row][curr.col]) == 1 {
					seen[next] = struct{}{}
					if topoMap[next.row][next.col] == '9' {
						total += 1
					} else {
						horizon = append(horizon, next)
					}
				}
			}
		}

	}
	return total
}

func SolvePartTwo(input string) int {
	topoMap := parseInput(input)
	numRows := len(topoMap)
	numCols := len(topoMap[0])
	starts := make([]Position, 0, 100)
	for r := range numRows {
		for c := range numCols {
			if topoMap[r][c] == '0' {
				starts = append(starts, Position{r, c})
			}
		}
	}
	total := 0
	horizon := slices.Clone(starts)
	for len(horizon) > 0 {
		curr := horizon[len(horizon)-1]
		horizon = horizon[:len(horizon)-1]
		for _, delta := range []Position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Position{curr.row + delta.row, curr.col + delta.col}
			if next.row >= 0 && next.row < numRows && next.col >= 0 && next.col < numCols &&
				int(topoMap[next.row][next.col])-int(topoMap[curr.row][curr.col]) == 1 {
				if topoMap[next.row][next.col] == '9' {
					total += 1
				} else {
					horizon = append(horizon, next)
				}
			}
		}
	}
	return total
}
