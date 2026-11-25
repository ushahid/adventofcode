package day21

import (
	"strconv"
	"strings"
)

type Position struct {
	row, col int
}

type Code struct {
	Raw string
	Num int
}

var numKeypad = map[byte]Position{
	'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
	'0': {3, 1}, 'A': {3, 2},
}

var dirKeypad = map[byte]Position{
	'^': {0, 1}, 'A': {0, 2},
	'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
}

func parseInput(input string) []Code {
	lines := strings.Split(input, "\n")
	codes := make([]Code, 0, len(lines))

	for _, line := range lines {
		digits := strings.TrimRight(line, "Aa")
		num, err := strconv.Atoi(digits)
		if err != nil {
			continue
		}
		codes = append(codes, Code{
			Raw: line,
			Num: num,
		})
	}

	return codes
}

func stepPaths(pad map[byte]Position, gap Position, from, to byte) []string {
	start := pad[from]
	end := pad[to]

	dr := end.row - start.row
	dc := end.col - start.col

	var vMove strings.Builder
	if dr > 0 {
		vMove.Grow(dr)
		for i := 0; i < dr; i++ {
			vMove.WriteByte('v')
		}
	} else if dr < 0 {
		vMove.Grow(-dr)
		for i := 0; i < -dr; i++ {
			vMove.WriteByte('^')
		}
	}

	var hMove strings.Builder
	if dc > 0 {
		hMove.Grow(dc)
		for i := 0; i < dc; i++ {
			hMove.WriteByte('>')
		}
	} else if dc < 0 {
		hMove.Grow(-dc)
		for i := 0; i < -dc; i++ {
			hMove.WriteByte('<')
		}
	}

	vStr := vMove.String()
	hStr := hMove.String()

	if vStr == "" {
		return []string{hStr + "A"}
	}
	if hStr == "" {
		return []string{vStr + "A"}
	}

	var paths []string
	if (Position{row: start.row, col: end.col}) != gap {
		paths = append(paths, hStr+vStr+"A")
	}

	if (Position{row: end.row, col: start.col}) != gap {
		paths = append(paths, vStr+hStr+"A")
	}

	return paths
}

func expandSequence(seq string, pad map[byte]Position, gap Position) []string {
	results := []string{""}
	curr := byte('A')

	for i := 0; i < len(seq); i++ {
		next := seq[i]
		options := stepPaths(pad, gap, curr, next)
		var nextResults []string
		for _, prefix := range results {
			for _, opt := range options {
				nextResults = append(nextResults, prefix+opt)
			}
		}
		results = nextResults
		curr = next
	}

	return results
}

type CacheKey struct {
	seq   string
	depth int
}

func minPresses(seq string, depth int, memo map[CacheKey]int) int {
	if depth == 0 {
		return len(seq)
	}

	key := CacheKey{seq: seq, depth: depth}
	if val, ok := memo[key]; ok {
		return val
	}

	total := 0
	start := 0
	for i := 0; i < len(seq); i++ {
		if seq[i] == 'A' {
			chunk := seq[start : i+1]
			start = i + 1

			candidates := expandSequence(chunk, dirKeypad, Position{row: 0, col: 0})
			best := -1
			for _, cand := range candidates {
				cost := minPresses(cand, depth-1, memo)
				if best == -1 || cost < best {
					best = cost
				}
			}
			total += best
		}
	}

	memo[key] = total
	return total
}

func totalComplexity(codes []Code, dirRobots int) int {
	memo := make(map[CacheKey]int)
	numGap := Position{row: 3, col: 0}
	sum := 0

	for _, code := range codes {
		numSeqs := expandSequence(code.Raw, numKeypad, numGap)
		bestLen := -1
		for _, seq := range numSeqs {
			length := minPresses(seq, dirRobots, memo)
			if bestLen == -1 || length < bestLen {
				bestLen = length
			}
		}
		sum += bestLen * code.Num
	}

	return sum
}

func SolvePartOne(input string) int {
	codes := parseInput(input)
	return totalComplexity(codes, 2)
}

func SolvePartTwo(input string) int {
	codes := parseInput(input)
	return totalComplexity(codes, 25)
}
