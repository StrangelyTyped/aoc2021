package day2

import (
	"strings"
	"testing"
)

var testData string = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPartOne(t *testing.T) {
	instrs := readInstructions(strings.NewReader(testData))

	result := partOne(instrs)

	if result != 150 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	instrs := readInstructions(strings.NewReader(testData))

	result := partTwo(instrs)

	if result != 900 {
		t.Errorf("Unexpected result: got %d", result)
	}
}