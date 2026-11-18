package day08

import (
	"strings"
)

type Position struct {
	row int
	col int
}

func parseInput(input string) (map[rune][]Position, int, int) {
	towers := make(map[rune][]Position)
	num_rows := 0
	num_cols := 0
	for row, line := range strings.Split(input, "\n") {
		num_cols = len(line)
		for col, r := range line {
			if r != '.' {
				towers[r] = append(towers[r], Position{row, col})
			}
		}
		num_rows++
	}
	return towers, num_rows, num_cols
}

func validPos(p Position, num_rows int, num_cols int) bool {
	return p.row >= 0 && p.row < num_rows && p.col >= 0 && p.col < num_cols
}

func SolvePartOne(input string) int {
	valid := 0
	allTowers, num_rows, num_cols := parseInput(input)
	occupiedPos := make(map[Position]bool)

	for _, towers := range allTowers {
		for i := 0; i < len(towers); i++ {
			for j := i + 1; j < len(towers); j++ {
				rowDiff := towers[i].row - towers[j].row
				colDiff := towers[i].col - towers[j].col
				iPos := Position{
					towers[i].row + rowDiff,
					towers[i].col + colDiff,
				}
				if validPos(iPos, num_rows, num_cols) && !occupiedPos[iPos] {
					valid++
					occupiedPos[iPos] = true
				}
				jPos := Position{
					towers[j].row + (rowDiff * -1),
					towers[j].col + (colDiff * -1),
				}
				if validPos(jPos, num_rows, num_cols) && !occupiedPos[jPos] {
					valid++
					occupiedPos[jPos] = true
				}
			}
		}
	}
	return valid
}

func SolvePartTwo(input string) int {
	valid := 0
	allTowers, num_rows, num_cols := parseInput(input)
	occupiedPos := make(map[Position]bool)

	for _, towers := range allTowers {
		for i := 0; i < len(towers); i++ {
			for j := i + 1; j < len(towers); j++ {
				rowDiff := towers[i].row - towers[j].row
				colDiff := towers[i].col - towers[j].col

				deltaRow := []int{rowDiff, -1 * rowDiff}
				deltaCol := []int{colDiff, -1 * colDiff}
				currPos := []Position{towers[i], towers[j]}

				for k := range 2 {
					pos := currPos[k]
					for mul := 1; validPos(pos, num_rows, num_cols); mul++ {
						if !occupiedPos[pos] {
							valid++
							occupiedPos[pos] = true
						}
						pos = Position{
							currPos[k].row + (deltaRow[k] * mul),
							currPos[k].col + (deltaCol[k] * mul),
						}
					}

				}
			}
		}
	}
	return valid
}
