package day17

import (
	"fmt"
	"regexp"
	"strconv"
)

type Computer struct {
	a, b, c int
	program []int
}

func parseInput(input string) Computer {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(input, -1)
	if len(matches) < 4 {
		panic("invalid input: expected 3 registers and program instructions")
	}

	nums := make([]int, len(matches))
	for i, m := range matches {
		v, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		nums[i] = v
	}

	prog := make([]int, len(nums)-3)
	for i := 3; i < len(nums); i++ {
		prog[i-3] = nums[i]
	}

	return Computer{
		a:       nums[0],
		b:       nums[1],
		c:       nums[2],
		program: prog,
	}
}

func truncDivPow2(numerator int, shift int) int {
	if shift >= 63 {
		return 0
	}
	if shift <= 0 {
		return numerator
	}
	return numerator / (int(1) << shift)
}

func (c Computer) run() []int {
	aReg, bReg, cReg := c.a, c.b, c.c
	ip := 0
	var out []int

	combo := func(op int) int {
		switch op {
		case 0, 1, 2, 3:
			return op
		case 4:
			return aReg
		case 5:
			return bReg
		case 6:
			return cReg
		default:
			panic(fmt.Sprintf("invalid combo operand: %d", op))
		}
	}

	for ip >= 0 && ip+1 < len(c.program) {
		opcode := c.program[ip]
		operand := c.program[ip+1]

		switch opcode {
		case 0: // adv: A = trunc(A / 2^combo)
			aReg = truncDivPow2(aReg, combo(operand))
			ip += 2
		case 1: // bxl: B = B XOR literal
			bReg ^= operand
			ip += 2
		case 2: // bst: B = combo % 8
			bReg = combo(operand) % 8
			ip += 2
		case 3: // jnz: jump if A != 0
			if aReg != 0 {
				ip = operand
			} else {
				ip += 2
			}
		case 4: // bxc: B = B XOR C
			bReg ^= cReg
			ip += 2
		case 5: // out: output combo % 8
			out = append(out, combo(operand)%8)
			ip += 2
		case 6: // bdv: B = trunc(A / 2^combo)
			bReg = truncDivPow2(aReg, combo(operand))
			ip += 2
		case 7: // cdv: C = trunc(A / 2^combo)
			cReg = truncDivPow2(aReg, combo(operand))
			ip += 2
		default:
			panic(fmt.Sprintf("unknown opcode: %d", opcode))
		}
	}

	return out
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findQuineA(base Computer, idx int, currentA int) (int, bool) {
	if idx < 0 {
		return currentA, true
	}

	targetSuffix := base.program[idx:]
	for bits := 0; bits < 8; bits++ {
		nextA := (currentA << 3) | bits
		trial := Computer{
			a:       nextA,
			b:       base.b,
			c:       base.c,
			program: base.program,
		}
		if equalSlice(trial.run(), targetSuffix) {
			if ans, ok := findQuineA(base, idx-1, nextA); ok {
				return ans, true
			}
		}
	}

	return 0, false
}

func SolvePartOne(input string) int {
	comp := parseInput(input)
	fmt.Println(comp)
	out := comp.run()

	var result int
	for _, digit := range out {
		result = result*10 + digit
	}
	return int(result)
}

func SolvePartTwo(input string) int {
	comp := parseInput(input)
	ans, ok := findQuineA(comp, len(comp.program)-1, 0)
	if !ok {
		panic("no valid initial register A found")
	}
	return ans
}
