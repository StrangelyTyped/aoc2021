package day11

import (
	"strings"
	"testing"
)

var testData string = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestPartOne(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partOne(nums)

	if result != 1656 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partTwo(nums)

	if result != 195 {
		t.Errorf("Unexpected result: got %d", result)
	}
}
