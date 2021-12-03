package day3

import (
	"strings"
	"testing"
)

var testData string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPartOne(t *testing.T) {
	instrs := readInstructions(strings.NewReader(testData))

	result := partOne(instrs)

	if result != 198 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	instrs := readInstructions(strings.NewReader(testData))

	result := partTwo(instrs)

	if result != 230 {
		t.Errorf("Unexpected result: got %d", result)
	}
}