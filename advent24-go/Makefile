# Change this to match your go.mod module path if it differs
MODULE_PREFIX ?= github.com/ushahid/adventofcode/advent24-go


# Usage: make new DAY=01
.PHONY: new
new:
ifndef DAY
	$(error DAY is not set. Usage: make new DAY=01)
endif
	@mkdir -p day$(DAY)
	@if [ ! -f day$(DAY)/solve.go ]; then \
		echo "package day$(DAY)" > day$(DAY)/solve.go; \
		echo "" >> day$(DAY)/solve.go; \
		echo "func SolvePartOne(input string) int {" >> day$(DAY)/solve.go; \
		echo "    return 0" >> day$(DAY)/solve.go; \
		echo "}" >> day$(DAY)/solve.go; \
		echo "" >> day$(DAY)/solve.go; \
		echo "func SolvePartTwo(input string) int {" >> day$(DAY)/solve.go; \
		echo "    return 0" >> day$(DAY)/solve.go; \
		echo "}" >> day$(DAY)/solve.go; \
		echo "Created day$(DAY)/solve.go"; \
	fi
	@if [ ! -f day$(DAY)/solve_test.go ]; then \
		echo "package day$(DAY)_test" > day$(DAY)/solve_test.go; \
		echo "" >> day$(DAY)/solve_test.go; \
		echo "import (" >> day$(DAY)/solve_test.go; \
		echo "    \"testing\"" >> day$(DAY)/solve_test.go; \
		echo "    \"time\"" >> day$(DAY)/solve_test.go; \
		echo "" >> day$(DAY)/solve_test.go; \
		echo "    \"$(MODULE_PREFIX)/day$(DAY)\"" >> day$(DAY)/solve_test.go; \
		echo "    \"$(MODULE_PREFIX)/utils\"" >> day$(DAY)/solve_test.go; \
		echo ")" >> day$(DAY)/solve_test.go; \
		echo "" >> day$(DAY)/solve_test.go; \
		echo "func TestSolvePartOne(t *testing.T) {" >> day$(DAY)/solve_test.go; \
		echo "    input := utils.ReadInput(t)" >> day$(DAY)/solve_test.go; \
		echo "    start := time.Now()" >> day$(DAY)/solve_test.go; \
		echo "    result := day$(DAY).SolvePartOne(input)" >> day$(DAY)/solve_test.go; \
		echo "    elapsed := time.Since(start)" >> day$(DAY)/solve_test.go; \
		echo "    utils.LogTime(t, elapsed)" >> day$(DAY)/solve_test.go; \
		echo "    utils.LogResult(t, result)" >> day$(DAY)/solve_test.go; \
		echo "}" >> day$(DAY)/solve_test.go; \
		echo "" >> day$(DAY)/solve_test.go; \
		echo "func TestSolvePartTwo(t *testing.T) {" >> day$(DAY)/solve_test.go; \
		echo "    input := utils.ReadInput(t)" >> day$(DAY)/solve_test.go; \
		echo "    start := time.Now()" >> day$(DAY)/solve_test.go; \
		echo "    result := day$(DAY).SolvePartTwo(input)" >> day$(DAY)/solve_test.go; \
		echo "    elapsed := time.Since(start)" >> day$(DAY)/solve_test.go; \
		echo "    utils.LogTime(t, elapsed)" >> day$(DAY)/solve_test.go; \
		echo "    utils.LogResult(t, result)" >> day$(DAY)/solve_test.go; \
		echo "}" >> day$(DAY)/solve_test.go; \
		echo "Created day$(DAY)/solve_test.go"; \
	fi
	@if [ ! -f day$(DAY)/input.txt ]; then \
		touch day$(DAY)/input.txt; \
		echo "Created day$(DAY)/input.txt"; \
	fi


.PHONY: test
test:
ifndef DAY
	$(error DAY is not set. Usage: make test DAY=01)
endif
	@go test ./day$(DAY) -v