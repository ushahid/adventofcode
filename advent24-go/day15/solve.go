package day15

import (
	"strings"
)

type Position struct {
	row, col int
}

func parseInput(input string) ([][]rune, []Position, Position) {
	grid := make([][]rune, 0)
	robotPos := Position{0, 0}
	vals := strings.Split(input, "\n\n")
	for rIdx, line := range strings.Split(vals[0], "\n") {
		row := make([]rune, len(line))
		for cIdx, char := range line {
			if char == '@' {
				robotPos.row = rIdx
				robotPos.col = cIdx
			}
			row[cIdx] = char
		}
		grid = append(grid, row)
	}

	moves := make([]Position, 0, len(vals[1]))
	for _, r := range vals[1] {
		switch r {
		case '^':
			moves = append(moves, Position{-1, 0})
		case '>':
			moves = append(moves, Position{0, 1})
		case '<':
			moves = append(moves, Position{0, -1})
		case 'v':
			moves = append(moves, Position{1, 0})
		}
	}

	return grid, moves, robotPos
}

func canMove(pos Position, dir Position, grid [][]rune) bool {
	next := Position{
		row: pos.row + dir.row,
		col: pos.col + dir.col,
	}

	switch grid[next.row][next.col] {
	case '#':
		return false
	case '.':
		return true
	case 'O':
		return canMove(next, dir, grid)
	default:
		return false
	}
}

func movePiece(pos Position, dir Position, grid [][]rune) Position {
	next := Position{
		row: pos.row + dir.row,
		col: pos.col + dir.col,
	}

	if grid[next.row][next.col] == 'O' {
		movePiece(next, dir, grid)
	}

	grid[next.row][next.col] = grid[pos.row][pos.col]
	grid[pos.row][pos.col] = '.'

	return next
}

func SolvePartOne(input string) int {
	grid, moves, robotPos := parseInput(input)
	for _, move := range moves {
		if canMove(robotPos, move, grid) {
			robotPos = movePiece(robotPos, move, grid)
		}
	}
	total := 0
	for r, row := range grid {
		for c, cell := range row {
			if cell == 'O' {
				total += 100*r + c
			}
		}
	}
	return total
}

func expandGrid(grid [][]rune) ([][]rune, Position) {
	expanded := make([][]rune, len(grid))
	var robotPos Position

	for r, row := range grid {
		newRow := make([]rune, 0, len(row)*2)
		for c, cell := range row {
			switch cell {
			case '#':
				newRow = append(newRow, '#', '#')
			case 'O':
				newRow = append(newRow, '[', ']')
			case '.':
				newRow = append(newRow, '.', '.')
			case '@':
				robotPos = Position{row: r, col: c * 2}
				newRow = append(newRow, '@', '.')
			}
		}
		expanded[r] = newRow
	}

	return expanded, robotPos
}

// boxLeft returns the Position of '[' for any box cell ('[' or ']').
func boxLeft(pos Position, grid [][]rune) Position {
	if grid[pos.row][pos.col] == ']' {
		return Position{row: pos.row, col: pos.col - 1}
	}
	return pos
}

func canMovePartTwo(pos Position, dir Position, grid [][]rune) bool {
	next := Position{row: pos.row + dir.row, col: pos.col + dir.col}
	cell := grid[next.row][next.col]

	switch cell {
	case '#':
		return false
	case '.':
		return true
	case '[', ']':
		if dir.row == 0 {
			beyond := Position{row: next.row, col: next.col + dir.col}
			return canMovePartTwo(beyond, dir, grid)
		}
		left := boxLeft(next, grid)
		right := Position{row: left.row, col: left.col + 1}
		return canMovePartTwo(left, dir, grid) && canMovePartTwo(right, dir, grid)
	default:
		return false
	}
}

func movePiecePartTwo(pos Position, dir Position, grid [][]rune) Position {
	next := Position{row: pos.row + dir.row, col: pos.col + dir.col}
	cell := grid[next.row][next.col]

	if cell == '[' || cell == ']' {
		if dir.row == 0 {
			beyond := Position{row: next.row, col: next.col + dir.col}
			movePiecePartTwo(beyond, dir, grid)
			grid[beyond.row][beyond.col] = grid[next.row][next.col]
			grid[next.row][next.col] = grid[pos.row][pos.col]
			grid[pos.row][pos.col] = '.'
			return next
		}

		left := boxLeft(next, grid)
		right := Position{row: left.row, col: left.col + 1}

		pushVerticalBox(left, dir, grid)
		pushVerticalBox(right, dir, grid)

		grid[left.row+dir.row][left.col] = '['
		grid[left.row][left.col] = '.'
		grid[right.row+dir.row][right.col] = ']'
		grid[right.row][right.col] = '.'
	}

	grid[next.row][next.col] = grid[pos.row][pos.col]
	grid[pos.row][pos.col] = '.'
	return next
}

// pushVerticalBox recursively clears the space directly above/below a single
// half of a box ('[' or ']') before moving that half one row in dir.
func pushVerticalBox(half Position, dir Position, grid [][]rune) {
	targetRow := half.row + dir.row
	aboveOrBelow := grid[targetRow][half.col]

	if aboveOrBelow == '[' || aboveOrBelow == ']' {
		otherLeft := boxLeft(Position{row: targetRow, col: half.col}, grid)
		otherRight := Position{row: otherLeft.row, col: otherLeft.col + 1}

		// Recursively push the next box(es) in the chain first.
		pushVerticalBox(otherLeft, dir, grid)
		if otherRight.col != otherLeft.col {
			pushVerticalBox(otherRight, dir, grid)
		}

		grid[otherLeft.row+dir.row][otherLeft.col] = '['
		grid[otherLeft.row][otherLeft.col] = '.'
		grid[otherRight.row+dir.row][otherRight.col] = ']'
		grid[otherRight.row][otherRight.col] = '.'
	}
}

func SolvePartTwo(input string) int {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	rawGrid, moves, _ := parseInput(input)
	grid, robotPos := expandGrid(rawGrid)

	for _, move := range moves {
		if canMovePartTwo(robotPos, move, grid) {
			robotPos = movePiecePartTwo(robotPos, move, grid)
		}
	}

	total := 0
	for r, row := range grid {
		for c, cell := range row {
			if cell == '[' {
				total += 100*r + c
			}
		}
	}
	return total
}
