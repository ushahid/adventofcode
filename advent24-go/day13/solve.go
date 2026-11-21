package day13

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Machine struct {
	da    Position
	db    Position
	prize Position
}

func parseInput(input string) []Machine {
	var machines []Machine
	a_regex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	b_regex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	p_regex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		// Read first
		line := scanner.Text()
		vals := a_regex.FindStringSubmatch(line)
		ax, _ := strconv.Atoi(vals[1])
		ay, _ := strconv.Atoi(vals[2])

		// Read second
		scanner.Scan()
		line = scanner.Text()
		vals = b_regex.FindStringSubmatch(line)
		bx, _ := strconv.Atoi(vals[1])
		by, _ := strconv.Atoi(vals[2])

		// Read prize
		scanner.Scan()
		line = scanner.Text()
		vals = p_regex.FindStringSubmatch(line)
		px, _ := strconv.Atoi(vals[1])
		py, _ := strconv.Atoi(vals[2])
		scanner.Scan()

		machines = append(machines, Machine{da: Position{ax, ay}, db: Position{bx, by}, prize: Position{px, py}})
	}
	return machines
}

func solveEq(m Machine) (int, int, error) {
	det := m.da.x*m.db.y - m.da.y*m.db.x
	if det == 0 {
		return 0, 0, fmt.Errorf("No solution")
	}
	a_press := float64(m.prize.x*m.db.y-m.prize.y*m.db.x) / float64(det)
	b_press := float64(m.da.x*m.prize.y-m.da.y*m.prize.x) / float64(det)
	if math.Trunc(a_press) == a_press && math.Trunc(b_press) == b_press {
		return int(a_press), int(b_press), nil
	}
	return 0, 0, fmt.Errorf("No integer solution")
}

func SolvePartOne(input string) int {
	machines := parseInput(input)
	total := 0
	for _, m := range machines {
		if a_press, b_press, err := solveEq(m); err == nil {
			total += 3*a_press + b_press
		}

	}
	return total
}

func SolvePartTwo(input string) int {
	machines := parseInput(input)
	for idx := range machines {
		machines[idx].prize.x += 10000000000000
		machines[idx].prize.y += 10000000000000
	}
	total := 0
	for _, m := range machines {
		if a_press, b_press, err := solveEq(m); err == nil {
			total += 3*a_press + b_press
		}

	}
	return total
}
