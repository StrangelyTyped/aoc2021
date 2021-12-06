package day6

import (
	"strings"
	"testing"
)

var testData string = `3,4,3,1,2`

func TestPartOne(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partOne(nums)

	if result != 5934 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partTwo(nums)

	if result != 26984457539 {
		t.Errorf("Unexpected result: got %d", result)
	}
}