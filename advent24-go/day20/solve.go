package day20

import "strings"

type Position struct {
	row, col int
}

type Track struct {
	rows int
	cols int
	dist []int      // flattened grid: -1 means wall or unvisited
	path []Position // track tiles ordered by distance from S
}

func (t *Track) idx(row int, col int) int {
	return row*t.cols + col
}

func parseInput(input string) Track {
	lines := strings.Split(input, "\n")
	numRows := len(lines)
	numCols := len(lines[0])

	var start Position
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if lines[r][c] == 'S' {
				start = Position{r, c}
			}
		}
	}

	dist := make([]int, numRows*numCols)
	for i := range dist {
		dist[i] = -1
	}

	path := make([]Position, 0, numRows*numCols)
	curr := start
	dist[curr.row*numCols+curr.col] = 0
	path = append(path, curr)

	dirs := [4]Position{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	moved := true
	for moved {
		moved = false
		for _, d := range dirs {
			nextRow, nextCol := curr.row+d.row, curr.col+d.col
			if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols {
				continue
			}
			nextIdx := nextRow*numCols + nextCol
			if lines[nextRow][nextCol] != '#' && dist[nextIdx] == -1 {
				dist[nextIdx] = dist[curr.row*numCols+curr.col] + 1
				curr = Position{nextRow, nextCol}
				path = append(path, curr)
				moved = true
				break
			}
		}
	}

	return Track{
		rows: numRows,
		cols: numCols,
		dist: dist,
		path: path,
	}
}

func countCheats(t Track, maxCheat int, minSave int) int {
	total := 0

	for _, currPos := range t.path {
		dist := t.dist[t.idx(currPos.row, currPos.col)]
		for deltaRow := -maxCheat; deltaRow <= maxCheat; deltaRow++ {
			testRow := currPos.row + deltaRow
			if testRow < 0 || testRow >= t.rows {
				continue
			}
			absDeltaRow := deltaRow
			if absDeltaRow < 0 {
				absDeltaRow = -absDeltaRow
			}
			remCol := maxCheat - absDeltaRow

			for deltaCol := -remCol; deltaCol <= remCol; deltaCol++ {
				c2 := currPos.col + deltaCol
				if c2 < 0 || c2 >= t.cols {
					continue
				}

				d2 := t.dist[t.idx(testRow, c2)]
				if d2 == -1 {
					continue
				}

				absDeltaCol := deltaCol
				if absDeltaCol < 0 {
					absDeltaCol = -absDeltaCol
				}
				cheatLen := absDeltaRow + absDeltaCol
				saved := (d2 - dist) - cheatLen

				if saved >= minSave {
					total++
				}
			}
		}
	}
	return total
}

func SolvePartOne(input string) int {
	t := parseInput(input)
	return countCheats(t, 2, 100)
}

func SolvePartTwo(input string) int {
	t := parseInput(input)
	return countCheats(t, 20, 100)
}
