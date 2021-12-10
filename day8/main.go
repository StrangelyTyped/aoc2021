package day8

import (
	"bufio"
	"fmt"
	"io"
	"strings"

)

var segmentMapping = []string {
	"abcefg",
	"cf",
	"acdeg",
	"acdfg",
	"bcdf",
	"abdfg",
	"abdefg",
	"acf",
	"abcdefg",
	"abcdfg",
}

func computeLengthBasedCandidates(inString string) []int {
	var result []int
	for digit, segs := range(segmentMapping) {
		if len(segs) == len(inString) {
			result = append(result, digit)
		}
	}
	return result
}

func partOne(input []inputRow) int {
	count := 0
	for _, row := range input {
		for _, digit := range row.digits {
			candidates := computeLengthBasedCandidates(digit)
			if len(candidates) == 1 {
				count++
			}
		}
	}
	return count
}

type rowMapping map[string][]string

func applyRowMapping(digitString string, digit int, rowMapping *rowMapping) {
	validSegStr := segmentMapping[digit]
	validSegs := strings.Split(validSegStr, "")
	for _, sourceChr := range strings.Split(digitString, "") {
		candidateSegs, ok := (*rowMapping)[sourceChr]
		if ok {
			var newCandidateSegs []string
			for _, c := range candidateSegs {
				if strings.Contains(validSegStr, c) {
					newCandidateSegs = append(newCandidateSegs, c)
				}
			}
			if len(newCandidateSegs) == 0 {
				panic("No valid candidates1")
			}
			(*rowMapping)[string(sourceChr)] = newCandidateSegs
		} else {
			(*rowMapping)[string(sourceChr)] = validSegs
		}
	}
	for sourceChr, candidateSegs := range *rowMapping {
		if !strings.Contains(digitString, sourceChr) {
			var newCandidateSegs []string
			for _, c := range candidateSegs {
				if !strings.Contains(validSegStr, c) {
					newCandidateSegs = append(newCandidateSegs, c)
				}
			}
			if len(newCandidateSegs) == 0 {
				fmt.Printf("Source: %s\n", sourceChr)
				panic("No valid candidates2")
			}
			(*rowMapping)[string(sourceChr)] = newCandidateSegs
		}
	}
}

func computeMappedCandidates(inString string, rowMapping rowMapping) []int {
	var result []int
	candidates := computeLengthBasedCandidates(inString)
	if len(candidates) == 1 {
		return candidates
	}
	for _, digit := range(candidates) {
		segs := segmentMapping[digit]

		if digitIsValid(inString, segs, rowMapping) {
			result = append(result, digit)
		}
	}
	return result
}

func digitIsValid(inString, segs string, rowMapping rowMapping) bool {
	var remainingSegMappings [][]string
	for _, x := range strings.Split(inString, "") {
		remainingSegMappings = append(remainingSegMappings, rowMapping[x])
	}
	return canBuildValidPerm(strings.Split(segs, ""), remainingSegMappings)
}

func canBuildValidPerm(segs []string, remainingSegMappings [][]string) bool {
	if len(segs) == 0 {
		return true
	}
	searchStr := segs[0]
	for idx, x := range remainingSegMappings {
		for _, y := range x {
			if y == searchStr {
				if canBuildValidPerm(segs[1:], append(remainingSegMappings[0:idx], remainingSegMappings[idx+1:]...)) {
					return true
				}
			}
		}
	}
	return false
}


func partTwo(input []inputRow) int {
	sum := 0
	for _, row := range input {
		fmt.Printf("Row: %v\n", row)
		rowMapping := make(rowMapping)
		for {
			badDigits := false
			for _, digit := range row.patterns {
				candidates := computeMappedCandidates(digit, rowMapping)
				if len(candidates) == 1 {
					fmt.Printf("Mapping after %s=%d: %v\n", digit, candidates[0], rowMapping)
					applyRowMapping(digit, candidates[0], &rowMapping)
				}else{
					fmt.Printf("Candidates for %s=%v, mapping: %v\n", digit, candidates, rowMapping)
					badDigits = true
				}
			}
			if !badDigits {
				break;
			}
		}

		rowValue := 0
		for _, digit := range row.digits {
			candidates := computeMappedCandidates(digit, rowMapping)
			rowValue = (rowValue * 10) + candidates[0]
		}
		sum += rowValue

	}
	return sum
}

type inputRow struct {
	patterns []string
	digits []string
}

func readInput(input io.Reader) []inputRow {
	var instructions []inputRow
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")

		instructions = append(instructions, inputRow{
			patterns: strings.Split(line[0], " "),
			digits: strings.Split(line[1], " "),
		})
	}
	return instructions
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
