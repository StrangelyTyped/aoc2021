package day11

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/strangelytyped/aoc2021/utils"
)

func flash(input [][]int, x int, y int) {
	input[x][y] = -1
	for x2 := x - 1; x2 <= x+1; x2++ {
		if x2 < 0 || x2 >= len(input) {
			continue
		}
		for y2 := y - 1; y2 <= y+1; y2++ {
			if y2 < 0 || y2 >= len(input[x2]) || input[x2][y2] < 0 {
				continue
			}
			input[x2][y2]++
			if input[x2][y2] > 9 {
				flash(input, x2, y2)
			}
		}
	}
}

func step(input [][]int) ([][]int, int) {
	flashCount := 0

	for x, row := range input {
		for y, val := range row {
			input[x][y] = val + 1
		}
	}

	for x, row := range input {
		for y, val := range row {
			if val > 9 {
				flash(input, x, y)
			}
		}
	}

	for x, row := range input {
		for y, val := range row {
			if val < 0 {
				input[x][y] = 0
				flashCount++
			}
		}
	}

	return input, flashCount
}

func partOne(input [][]int) int {
	flashSum := 0
	var flashes int
	for i := 0; i < 100; i++ {
		input, flashes = step(input)
		flashSum += flashes
	}
	return flashSum
}

func partTwo(input [][]int) int {
	stepNo := 0
	var flashes int
	for {
		stepNo++
		input, flashes = step(input)
		if flashes == len(input)*len(input[0]) {
			return stepNo
		}
	}
}

func readInput(input io.Reader) [][]int {
	var instructions [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var row []int
		for _, c := range strings.Split(scanner.Text(), "") {
			row = append(row, utils.ParseIntOrPanic(c))
		}
		instructions = append(instructions, row)
	}
	return instructions
}

func copy(input [][]int) [][]int {
	var result [][]int
	for _, row := range input {
		result = append(result, append(make([]int, 0, len(row)), row...))
	}
	return result
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(copy(numbers)))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
