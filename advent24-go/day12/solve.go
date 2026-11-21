package day12

import (
	"strings"
)

type Position struct {
	row, col int
}

func SolvePartOne(input string) int {
	grid := strings.Split(input, "\n")
	numRows := len(grid)
	numCols := len(grid[0])
	seen := make(map[Position]bool)
	total := 0
	for r := range numRows {
		for c := range numCols {
			pos := Position{r, c}
			if !seen[pos] {
				seen[pos] = true
				plant := grid[r][c]
				horizon := make([]Position, 0, numCols*numRows)
				horizon = append(horizon, pos)
				delta := []Position{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
				area := 0
				perimeter := 0
				for len(horizon) > 0 {
					lastPos := horizon[len(horizon)-1]
					horizon = horizon[:len(horizon)-1]
					area++
					for _, d := range delta {
						curr := Position{lastPos.row + d.row, lastPos.col + d.col}
						if curr.row >= 0 && curr.row < numRows &&
							curr.col >= 0 && curr.col < numCols &&
							grid[curr.row][curr.col] == plant {
							if !seen[curr] {
								horizon = append(horizon, curr)
								seen[curr] = true
							}
						} else {
							perimeter++
						}
					}
				}
				total += area * perimeter
			}
		}
	}
	return total
}

func countSides(region map[Position]bool) int {
	dirs := []int{-1, 1}
	sides := 0
	for p := range region {
		for _, vr := range dirs {
			for _, hc := range dirs {
				adjV := region[Position{p.row + vr, p.col}]
				adjH := region[Position{p.row, p.col + hc}]
				diag := region[Position{p.row + vr, p.col + hc}]
				if (!adjV && !adjH) || (adjV && adjH && !diag) {
					sides++
				}
			}
		}
	}

	return sides
}

func SolvePartTwo(input string) int {
	grid := strings.Split(input, "\n")
	numRows := len(grid)
	numCols := len(grid[0])
	seen := make(map[Position]bool)
	total := 0
	for r := range numRows {
		for c := range numCols {
			pos := Position{r, c}
			if !seen[pos] {
				region := make(map[Position]bool)
				region[pos] = true
				seen[pos] = true
				plant := grid[r][c]
				horizon := make([]Position, 0, numCols*numRows)
				horizon = append(horizon, pos)
				delta := []Position{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
				area := 0
				perimeter := 0
				for len(horizon) > 0 {
					lastPos := horizon[len(horizon)-1]
					horizon = horizon[:len(horizon)-1]
					area++
					for _, d := range delta {
						curr := Position{lastPos.row + d.row, lastPos.col + d.col}
						if curr.row >= 0 && curr.row < numRows &&
							curr.col >= 0 && curr.col < numCols &&
							grid[curr.row][curr.col] == plant {
							if !seen[curr] {
								horizon = append(horizon, curr)
								seen[curr] = true
								region[curr] = true
							}
						} else {
							perimeter++
						}
					}
				}
				sides := countSides(region)
				total += area * sides
			}

		}
	}
	return total
}
