package day10

import (
	"strings"
	"testing"
)

var testData string = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestPartOne(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partOne(nums)

	if result != 26397 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	nums := readInput(strings.NewReader(testData))

	result := partTwo(nums)

	if result != 288957 {
		t.Errorf("Unexpected result: got %d", result)
	}
}
