package day16

import (
	"container/heap"
	"strings"
)

type Position struct {
	row, col int
}

type Node struct {
	pos  Position
	dir  Position
	cost int
}

type StateKey struct {
	pos Position
	dir Position
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func findRunePos(target rune, grid []string) Position {
	for rIdx, line := range grid {
		for cIdx, r := range line {
			if r == target {
				return Position{rIdx, cIdx}
			}
		}
	}
	return Position{-1, -1}
}

func parseInput(input string) (grid []string, start Position, end Position) {
	grid = strings.Split(input, "\n")
	start = findRunePos('S', grid)
	end = findRunePos('E', grid)
	return grid, start, end
}

func minEndCost(end Position, dist map[StateKey]int) int {
	dirs := [4]Position{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	best := -1
	for _, d := range dirs {
		if c, ok := dist[StateKey{pos: end, dir: d}]; ok {
			if best == -1 || c < best {
				best = c
			}
		}
	}
	return best
}

func dijkstraAllShortest(
	grid []string,
	start, end, start_dir Position,
) (dist map[StateKey]int, prev map[StateKey][]StateKey) {
	rows := len(grid)
	cols := len(grid[0])

	startKey := StateKey{pos: start, dir: start_dir}
	dist = make(map[StateKey]int)
	prev = make(map[StateKey][]StateKey)

	pq := &MinHeap{}
	heap.Init(pq)

	dist[startKey] = 0
	heap.Push(pq, Node{pos: start, dir: start_dir, cost: 0})

	bestEndCost := -1
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Node)

		if bestEndCost != -1 && curr.cost > bestEndCost {
			break
		}

		currKey := StateKey{pos: curr.pos, dir: curr.dir}
		if best, ok := dist[currKey]; ok && curr.cost > best {
			continue
		}

		if curr.pos == end {
			if bestEndCost == -1 || curr.cost < bestEndCost {
				bestEndCost = curr.cost
			}
			continue
		}

		relax := func(nextPos, nextDir Position, stepCost int) {
			nextCost := curr.cost + stepCost
			nextKey := StateKey{pos: nextPos, dir: nextDir}
			prevBest, seen := dist[nextKey]
			if !seen || nextCost < prevBest {
				dist[nextKey] = nextCost
				prev[nextKey] = []StateKey{currKey}
				heap.Push(pq, Node{pos: nextPos, dir: nextDir, cost: nextCost})
			} else if nextCost == prevBest {
				prev[nextKey] = append(prev[nextKey], currKey)
			}
		}

		// 1. Move forward (+1)
		nextPos := Position{row: curr.pos.row + curr.dir.row, col: curr.pos.col + curr.dir.col}
		if nextPos.row >= 0 && nextPos.row < rows && nextPos.col >= 0 && nextPos.col < cols {
			if grid[nextPos.row][nextPos.col] != '#' {
				relax(nextPos, curr.dir, 1)
			}
		}

		// 2. Turn 90 degrees clockwise or counter-clockwise (+1000)
		turnDirs := [2]Position{
			{row: curr.dir.col, col: -curr.dir.row}, // right
			{row: -curr.dir.col, col: curr.dir.row}, // left
		}
		for _, nd := range turnDirs {
			relax(curr.pos, nd, 1000)
		}
	}

	return dist, prev
}

func countOptimalTiles(end Position, dist map[StateKey]int, prev map[StateKey][]StateKey) int {
	bestEndCost := minEndCost(end, dist)
	if bestEndCost == -1 {
		return 0
	}

	dirs := [4]Position{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	visitedStates := make(map[StateKey]bool)
	bestTiles := make(map[Position]bool)
	queue := make([]StateKey, 0)

	for _, d := range dirs {
		k := StateKey{pos: end, dir: d}
		if c, ok := dist[k]; ok && c == bestEndCost {
			visitedStates[k] = true
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		bestTiles[curr.pos] = true

		for _, p := range prev[curr] {
			if !visitedStates[p] {
				visitedStates[p] = true
				queue = append(queue, p)
			}
		}
	}

	return len(bestTiles)
}

func lowestScore(grid []string, start, end, start_dir Position) int {
	if len(grid) == 0 || start.row < 0 || end.row < 0 {
		return 0
	}
	dist, _ := dijkstraAllShortest(grid, start, end, start_dir)
	best := minEndCost(end, dist)
	if best == -1 {
		return -1
	}
	return best
}

func SolvePartOne(input string) int {
	grid, start, end := parseInput(input)
	start_dir := Position{0, 1}
	return lowestScore(grid, start, end, start_dir)
}

func SolvePartTwo(input string) int {
	grid, start, end := parseInput(input)
	start_dir := Position{0, 1}
	dist, prev := dijkstraAllShortest(grid, start, end, start_dir)
	return countOptimalTiles(end, dist, prev)
}
