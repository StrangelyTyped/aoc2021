package day8

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"strings"
)

var segmentMapping = []uint {
	0b1110111,
	0b0100100,
	0b1011101,
	0b1101101,
	0b0101110,
	0b1101011,
	0b1111011,
	0b0100101,
	0b1111111,
	0b1101111,
}

var segmentBits []int
func init(){
	for _, digit := range segmentMapping {
		segmentBits = append(segmentBits, bits.OnesCount(digit))
	}
}


func partOne(input []inputRow) int {
	count := 0
	for _, row := range input {
		for _, digit := range row.digits {
			digitBits := bits.OnesCount(digit)
			if digitBits == segmentBits[1] || digitBits == segmentBits[4] || digitBits == segmentBits[7] || digitBits == segmentBits[8] {
				count++
			}
		}
	}
	return count
}

func partTwo(input []inputRow) int {
	sum := 0
	for _, row := range input {
		digitMapping := make(map[int]uint)
		var bits5, bits6 []uint
		for _, digit := range row.patterns {
			digitBits := bits.OnesCount(digit)
			if digitBits == segmentBits[1] {
				digitMapping[1] = digit
			} else if digitBits == segmentBits[4] {
				digitMapping[4] = digit
			} else if digitBits == segmentBits[7] {
				digitMapping[7] = digit
			} else if digitBits == segmentBits[8] {
				digitMapping[8] = digit
			} else if digitBits == 5 {
				bits5 = append(bits5, digit)
			} else if digitBits == 6 { 
				bits6 = append(bits6, digit)
			} else {
				panic("Failed to length-map digit")
			}
		}
		fourOrSeven := (digitMapping[4] | digitMapping[7])
		for _, digit := range bits6 {
			if bits.OnesCount(digit & digitMapping[1]) == 1 {
				digitMapping[6] = digit
			} else if bits.OnesCount(digit & digitMapping[4]) == 3 {
				digitMapping[0] = digit
			} else if (digit & fourOrSeven) == fourOrSeven {
				digitMapping[9] = digit
			} else {
				panic("Failed to map 6-bit number")
			}
		}
		for _, digit := range bits5 {
			if (digit & digitMapping[9]) == digit {
				if (digit & digitMapping[1]) == digitMapping[1] {
					digitMapping[3] = digit
				} else {
					digitMapping[5] = digit
				} 
			} else {
				digitMapping[2] = digit
			}
		}

		reading := 0
		for _, digit := range row.digits {
			reading *= 10
			for intVal, segVal := range digitMapping {
				if segVal == digit {
					reading += intVal
					break
				}
			}
		}
		sum += reading
	}
	return sum
}

type inputRow struct {
	patterns []uint
	digits []uint
}

func stringToBitmap(digit string) uint {
	var bitmap uint = 0
	for i, j := 'a', 0; i <= 'g'; i,j=i+1,j+1 {
		if strings.ContainsRune(digit, i) {
			bitmap |= (1 << j)
		}
	}
	return bitmap
}

func readInput(input io.Reader) []inputRow {
	var instructions []inputRow
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		var patterns []uint
		for _, digit := range strings.Split(line[0], " ") {
			patterns = append(patterns, stringToBitmap(digit))
		}

		var digits []uint
		for _, digit := range strings.Split(line[1], " ") {
			digits = append(digits, stringToBitmap(digit))
		}

		instructions = append(instructions, inputRow{
			patterns: patterns,
			digits: digits,
		})
	}
	return instructions
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
