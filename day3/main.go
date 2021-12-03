package day3

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func mostCommonBits(input []string) []int {
	refs := make([]int, len(input[0]))
	for _, line := range input {
		for idx, chr := range line {
			if chr == '1' {
				refs[idx]++
			} else {
				refs[idx]--
			}
		}
	}
	return refs
}

func partOne(input []string) int {
	gamma := 0
	epsilon := 0
	refs := mostCommonBits(input)
	for i := 0; i < len(refs); i++ {
		gamma *= 2
		epsilon *= 2
		if refs[i] > 0 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	return gamma * epsilon
}



func partTwo(input []string) int {
	o2 := runFilter(input, true)
	co2 := runFilter(input, false)
	return o2 * co2
}

func runFilter(input []string, mostCommon bool) int {
	bitIdx := 0
	for len(input) > 1 {
		refs := mostCommonBits(input)
		mostCommonInSlot := refs[bitIdx]
		desiredChr := '0'
		if (mostCommon && mostCommonInSlot >= 0) || (!mostCommon && mostCommonInSlot < 0){
			desiredChr = '1'
		}
		input = filter(input, bitIdx, byte(desiredChr))

		bitIdx++

	}
	result, err := strconv.ParseInt(input[0], 2, 32)
	if err != nil {
		panic(err)
	}
	return int(result)
}

func filter(input []string, bitIdx int, desiredChr byte) []string {
	var result []string
	for _, line := range input {
		if line[bitIdx] == byte(desiredChr) {
			result = append(result, line)
		}
	}
	return result
}

func readInstructions(input io.Reader) []string {
	var instructions []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	return instructions
}

func Main(input io.Reader) {
	instructions := readInstructions(input)
	fmt.Printf("Part 1: %d\n", partOne(instructions))
	fmt.Printf("Part 2: %d\n", partTwo(instructions))
}
