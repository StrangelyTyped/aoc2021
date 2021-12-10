package day9

import (
	"strings"
	"testing"
)

var testData string = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestPartOne(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partOne(nums)

	if result != 15 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partTwo(nums)

	if result != 1134 {
		t.Errorf("Unexpected result: got %d", result)
	}
}
