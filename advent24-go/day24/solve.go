package day24

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	Left  string
	Op    string
	Right string
	Out   string
}

type Circuit struct {
	Wires map[string]int
	Gates []Gate
}

func parseInput(input string) Circuit {
	c := Circuit{
		Wires: make(map[string]int),
	}

	parts := strings.Split(input, "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		wireParts := strings.Split(line, ":")
		wire := wireParts[0]
		val, err := strconv.Atoi(strings.TrimSpace(wireParts[1]))
		if err == nil {
			c.Wires[wire] = val
		}
	}
	for _, line := range strings.Split(parts[1], "\n") {
		tokens := strings.Fields(line)
		c.Gates = append(c.Gates, Gate{
			Left:  tokens[0],
			Op:    tokens[1],
			Right: tokens[2],
			Out:   tokens[4],
		})
	}
	return c
}

func evalCircuit(c Circuit) map[string]int {
	vals := make(map[string]int, len(c.Wires)+len(c.Gates))
	for k, v := range c.Wires {
		vals[k] = v
	}

	var eval func(wire string) int
	eval = func(wire string) int {
		if v, ok := vals[wire]; ok {
			return v
		}
		for _, g := range c.Gates {
			if g.Out == wire {
				a := eval(g.Left)
				b := eval(g.Right)
				var res int
				switch g.Op {
				case "AND":
					if a == 1 && b == 1 {
						res = 1
					}
				case "OR":
					if a == 1 || b == 1 {
						res = 1
					}
				case "XOR":
					if a != b {
						res = 1
					}
				}
				vals[wire] = res
				return res
			}
		}
		return 0
	}

	for _, g := range c.Gates {
		eval(g.Out)
	}

	return vals
}

func SolvePartOne(input string) int {
	c := parseInput(input)
	vals := evalCircuit(c)

	var zWires []string
	for wire := range vals {
		if strings.HasPrefix(wire, "z") {
			zWires = append(zWires, wire)
		}
	}
	sort.Strings(zWires)
	result := 0
	for i, wire := range zWires {
		if vals[wire] == 1 {
			result |= 1 << i
		}
	}
	return result
}

func isXYWire(w string) bool {
	return strings.HasPrefix(w, "x") || strings.HasPrefix(w, "y")
}

func isXYBit00(g Gate) bool {
	return (g.Left == "x00" && g.Right == "y00") || (g.Left == "y00" && g.Right == "x00")
}

func FindSwappedWires(input string) string {
	c := parseInput(input)

	var maxZ string
	for _, g := range c.Gates {
		if strings.HasPrefix(g.Out, "z") && g.Out > maxZ {
			maxZ = g.Out
		}
	}

	bad := make(map[string]bool)

	for _, g := range c.Gates {
		leftXY := isXYWire(g.Left)
		rightXY := isXYWire(g.Right)
		bothXY := leftXY && rightXY
		outIsZ := strings.HasPrefix(g.Out, "z")

		if outIsZ && g.Out != maxZ && g.Op != "XOR" {
			bad[g.Out] = true
		}

		if g.Op == "XOR" && !outIsZ && !bothXY {
			bad[g.Out] = true
		}

		if g.Op == "XOR" && bothXY && !isXYBit00(g) {
			feedsXOR := false
			for _, other := range c.Gates {
				if other.Op == "XOR" && (other.Left == g.Out || other.Right == g.Out) {
					feedsXOR = true
					break
				}
			}
			if !feedsXOR {
				bad[g.Out] = true
			}
		}

		if g.Op == "AND" && !isXYBit00(g) {
			feedsOR := false
			for _, other := range c.Gates {
				if other.Op == "OR" && (other.Left == g.Out || other.Right == g.Out) {
					feedsOR = true
					break
				}
			}
			if !feedsOR {
				bad[g.Out] = true
			}
		}
	}

	result := make([]string, 0, len(bad))
	for w := range bad {
		result = append(result, w)
	}
	sort.Strings(result)
	return strings.Join(result, ",")
}

func SolvePartTwo(input string) int {
	swapped := FindSwappedWires(input)
	fmt.Println("    Wires:", swapped)
	if swapped == "" {
		return 0
	}
	return len(strings.Split(swapped, ","))
}
