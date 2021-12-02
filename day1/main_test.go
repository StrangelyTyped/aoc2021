package day1

import (
	"testing"
)

var testData string = `199
200
208
210
200
207
240
269
260
263`

func TestPartOne(t *testing.T) {
	depths, err := readInts([]byte(testData))
	if err != nil {
		panic(err)
	}

	result := partOne(depths)

	if result != 7 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	depths, err := readInts([]byte(testData))
	if err != nil {
		panic(err)
	}

	result := partTwo(depths)

	if result != 5 {
		t.Errorf("Unexpected result: got %d", result)
	}
}