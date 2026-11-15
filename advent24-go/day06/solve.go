package day06

import (
	"strings"
)

const (
	NORTH = 0
	SOUTH = 1
	EAST  = 2
	WEST  = 3
)

type Guard struct {
	row, col, direction int
}

const (
	EMPTY     = 0
	FOOTPRINT = 1
	WALL      = 2
)

func parseInput(input string) ([][]int, Guard) {
	var grid [][]int
	var guard = Guard{0, 0, NORTH}
	row_idx := 0
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		row := make([]int, len(line))
		for col_idx, r := range line {
			switch r {
			case '.':
				row[col_idx] = EMPTY
			case '#':
				row[col_idx] = WALL
			case '^':
				guard = Guard{row_idx, col_idx, NORTH}
			}
		}
		grid = append(grid, row)
		row_idx++
	}
	return grid, guard
}

func nextStep(guard Guard) (int, int) {
	switch guard.direction {
	case NORTH:
		return guard.row - 1, guard.col
	case SOUTH:
		return guard.row + 1, guard.col
	case EAST:
		return guard.row, guard.col + 1
	default: // WEST
		return guard.row, guard.col - 1
	}
}

func rotateDirection(direction int) int {
	switch direction {
	case NORTH:
		return EAST
	case SOUTH:
		return WEST
	case EAST:
		return SOUTH
	default: // WEST
		return NORTH
	}
}

func guardStep(grid [][]int, guard Guard) (Guard, bool) {
	new_row, new_col := nextStep(guard)
	num_rows := len(grid)
	num_cols := len(grid[0])
	movedOut := !(new_row >= 0 && new_row < num_rows && new_col >= 0 && new_col < num_cols)
	if movedOut {
		guard.row = new_row
		guard.col = new_col
		return guard, true
	} else {
		switch grid[new_row][new_col] {
		case EMPTY:
			{
				guard.row = new_row
				guard.col = new_col
				return guard, false
			}
		case FOOTPRINT:
			{
				guard.row = new_row
				guard.col = new_col
				return guard, false
			}
		default: // WALL
			{
				guard.direction = rotateDirection(guard.direction)
				return guard, false
			}
		}
	}
}

func SolvePartOne(input string) int {
	grid, guard := parseInput(input)
	movedOut := false
	for !movedOut {
		grid[guard.row][guard.col] = FOOTPRINT
		guard, movedOut = guardStep(grid, guard)
	}
	marked := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == FOOTPRINT {
				marked++
			}
		}
	}
	return marked
}

func checkLoop(grid [][]int, guard Guard) bool {
	num_rows := len(grid)
	num_cols := len(grid[0])

	visited := make([][][]bool, num_rows)
	for r := range visited {
		visited[r] = make([][]bool, num_cols)
		for c := range visited[r] {
			visited[r][c] = make([]bool, 4)
		}
	}

	movedOut := false
	for !movedOut {
		if visited[guard.row][guard.col][guard.direction] {
			return true
		}
		visited[guard.row][guard.col][guard.direction] = true
		guard, movedOut = guardStep(grid, guard)
	}
	return false
}

func SolvePartTwo(input string) int {
	grid, startGuard := parseInput(input)
	movedOut := false

	visitedCells := make(map[[2]int]bool)
	guard := startGuard
	for !movedOut {
		guard, movedOut = guardStep(grid, guard)
		if !movedOut {
			visitedCells[[2]int{guard.row, guard.col}] = true
		}
	}

	numObstacles := 0
	for cell := range visitedCells {
		r, c := cell[0], cell[1]
		duplicate := make([][]int, len(grid))
		for i := range grid {
			duplicate[i] = make([]int, len(grid[i]))
			copy(duplicate[i], grid[i])
		}
		duplicate[r][c] = WALL
		if checkLoop(duplicate, startGuard) {
			numObstacles++
		}
	}

	return numObstacles
}
