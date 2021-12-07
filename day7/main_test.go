package day7

import (
	"strings"
	"testing"
)

var testData string = `16,1,2,0,4,2,7,1,2,14`

func TestPartOne(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partOne(nums)

	if result != 37 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partTwo(nums)

	if result != 168 {
		t.Errorf("Unexpected result: got %d", result)
	}
}