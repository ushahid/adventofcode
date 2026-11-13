package day04

import (
	"strings"
)

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func SolvePartOne(input string) int {
	lines := parseInput(input)
	num_rows := len(lines)
	num_cols := len(lines[0])
	deltas := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
	target := []byte("MAS")
	total := 0
	for row := range num_rows {
		for col := range num_cols {
			if lines[row][col] == 'X' {
				for _, delta := range deltas {
					curr_row := row
					curr_col := col
					for idx := 0; idx < len(target); idx++ {
						curr_row = curr_row + delta[0]
						curr_col = curr_col + delta[1]
						if curr_row >= 0 && curr_row < num_rows && curr_col >= 0 && curr_col < num_cols && lines[curr_row][curr_col] == target[idx] {
							// fmt.Println("Found", string(target[idx]))
							if idx == len(target)-1 {
								total += 1
							}
						} else {
							break
						}
					}
				}
			}
		}
	}
	return total
}

func SolvePartTwo(input string) int {
	lines := parseInput(input)
	num_rows := len(lines)
	num_cols := len(lines[0])
	total := 0
	valid := map[string]struct{}{
		"MS": {},
		"SM": {},
	}
	for row := 1; row < num_rows-1; row++ {
		for col := 1; col < num_cols-1; col++ {
			if lines[row][col] == 'A' {
				primary_diag := string(lines[row-1][col-1]) + string(lines[row+1][col+1])
				secondary_diag := string(lines[row-1][col+1]) + string(lines[row+1][col-1])
				_, p_ok := valid[primary_diag]
				_, s_ok := valid[secondary_diag]
				if p_ok && s_ok {
					total += 1
				}
			}
		}
	}
	return total
}
