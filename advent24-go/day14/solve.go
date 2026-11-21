package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Velocity struct {
	row, col int
}

type Position struct {
	row, col int
}

type Robot struct {
	pos Position
	vel Velocity
}

func parseInput(input string) []Robot {
	robots := make([]Robot, 0, 500)
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`p=([-+]?\d+),([-+]?\d+) v=([-+]?\d+),([-+]?\d+)`)
	for _, line := range lines {
		vals := re.FindStringSubmatch(line)
		px, _ := strconv.Atoi(vals[1])
		py, _ := strconv.Atoi(vals[2])
		vx, _ := strconv.Atoi(vals[3])
		vy, _ := strconv.Atoi(vals[4])
		robot := Robot{
			vel: Velocity{vy, vx},
			pos: Position{py, px},
		}
		robots = append(robots, robot)
	}
	return robots
}

func fixDim(val int, limit int) int {
	for val < 0 {
		val += limit
	}
	return val % limit
}

func advance(robots []Robot, numRows int, numCols int) {
	for idx := range robots {
		newPos := Position{
			robots[idx].pos.row + robots[idx].vel.row,
			robots[idx].pos.col + robots[idx].vel.col,
		}
		newPos.row = fixDim(newPos.row, numRows)
		newPos.col = fixDim(newPos.col, numCols)
		robots[idx].pos = newPos
	}
}

func countQuad(robots []Robot, rowStart int, rowEnd int, colStart int, colEnd int) int {
	count := 0
	for _, robot := range robots {
		if robot.pos.row >= rowStart && robot.pos.row < rowEnd && robot.pos.col >= colStart && robot.pos.col < colEnd {
			count++
		}
	}
	return count
}

func SolvePartOne(input string) int {
	robots := parseInput(input)
	numCols := 101
	numRows := 103
	for range 100 {
		advance(robots, numRows, numCols)
	}
	count_q1 := countQuad(robots, 0, numRows/2, numCols/2+1, numCols)
	count_q2 := countQuad(robots, 0, numRows/2, 0, numCols/2)
	count_q3 := countQuad(robots, numRows/2+1, numRows, 0, numCols/2)
	count_q4 := countQuad(robots, numRows/2+1, numRows, numCols/2+1, numCols)
	return count_q1 * count_q2 * count_q3 * count_q4
}

func makeGrid(robots []Robot, numRows int, numCols int) [][]byte {
	grid := make([][]byte, numRows)
	for r := range numRows {
		row := make([]byte, numCols)
		for c := range numCols {
			row[c] = '.'
		}
		grid[r] = row
	}
	for _, robot := range robots {
		grid[robot.pos.row][robot.pos.col] = '*'
	}
	return grid
}

func countSolo(robots []Robot, numRows int, numCols int) int {
	occupied := make([][]bool, numRows)
	for r := range numRows {
		occupied[r] = make([]bool, numCols)
	}

	for _, robot := range robots {
		r, c := robot.pos.row, robot.pos.col
		occupied[r][c] = true
	}

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	soloCount := 0
	for _, robot := range robots {
		r, c := robot.pos.row, robot.pos.col
		if r < 0 || r >= numRows || c < 0 || c >= numCols {
			continue
		}

		hasNeighbor := false
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < numRows && nc >= 0 && nc < numCols && occupied[nr][nc] {
				hasNeighbor = true
				break
			}
		}

		if !hasNeighbor {
			soloCount++
		}
	}

	return soloCount
}

func SolvePartTwo(input string) int {
	robots := parseInput(input)
	numCols := 101
	numRows := 103
	idx := 0
	for true {
		numSolo := countSolo(robots, numRows, numCols)
		if numSolo < 150 {
			grid := makeGrid(robots, numRows, numCols)
			for r := range numRows {
				fmt.Println(string(grid[r]))
			}
			fmt.Println(numSolo)
			break
		} else {
			advance(robots, numRows, numCols)
			idx++
		}

	}
	return idx
}
