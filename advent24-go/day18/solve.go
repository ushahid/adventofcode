package day18

import (
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	row, col int
}

type State struct {
	pos  Position
	cost int
}

func parseInput(input string) []Position {
	lines := strings.Split(input, "\n")
	corruptList := make([]Position, 0, len(lines))
	for _, line := range lines {
		vals := strings.Split(line, ",")
		row, _ := strconv.Atoi(vals[0])
		col, _ := strconv.Atoi(vals[1])
		corruptList = append(corruptList, Position{row, col})
	}
	return corruptList
}

func makeGrid(pos []Position, numRows int, numCols int) [][]byte {
	grid := make([][]byte, 0, numRows)
	for range numRows {
		row := make([]byte, numCols)
		for c := range numCols {
			row[c] = '.'
		}
		grid = append(grid, row)
	}
	for _, p := range pos {
		grid[p.row][p.col] = '#'
	}
	return grid
}

func findShortest(grid [][]byte) int {
	numRows := len(grid)
	numCols := len(grid[0])
	target := Position{numRows - 1, numCols - 1}
	horizon := []State{{Position{0, 0}, 0}}
	seen := make(map[Position]bool)
	for len(horizon) > 0 {
		curr := horizon[0]
		if curr.pos.row == target.row && curr.pos.col == target.col {
			return curr.cost
		}
		horizon = slices.Delete(horizon, 0, 1)
		delta := []Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range delta {
			next := Position{curr.pos.row + d.row, curr.pos.col + d.col}
			if !seen[next] && next.row >= 0 && next.row < numRows && next.col >= 0 && next.col < numCols && grid[next.row][next.col] == '.' {
				horizon = append(horizon, State{pos: next, cost: curr.cost + 1})
				seen[next] = true
			}
		}
	}
	return -1
}

func SolvePartOne(input string) int {
	corruptPos := parseInput(input)
	grid := makeGrid(corruptPos[:1024], 71, 71)
	return findShortest(grid)
}

func SolvePartTwo(input string) int {
	corruptPos := parseInput(input)
	grid := makeGrid(corruptPos[:1024], 71, 71)
	for i := 1024; i < len(corruptPos); i++ {
		grid[corruptPos[i].row][corruptPos[i].col] = '#'
		shortest := findShortest(grid)
		if shortest == -1 {
			return corruptPos[i].row*100 + corruptPos[i].col
		}
	}
	return 0
}
