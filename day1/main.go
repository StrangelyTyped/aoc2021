package day1

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func readInts(dataBytes []byte) ([]int, error) {
	var ints []int
	for _, line := range strings.Split(strings.TrimSpace(string(dataBytes)), "\n") {
		depth, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
		ints = append(ints, depth)
	}
	return ints, nil
}

func sum(slice []int) int {
	x := 0
	for _, y := range slice {
		x += y
	}
	return x
}

func windowSum(depths []int, windowSize int) int {
	count := 0
	for i := windowSize; i < len(depths); i++ {
		if sum(depths[i-(windowSize-1):i+1]) > sum(depths[i-windowSize:i]) {
			count++
		}
	}
	return count
}

func partOne(depths []int) int {
	return windowSum(depths, 1)
}

func partTwo(depths []int) int {
	return windowSum(depths, 3)
}

func Main(input io.Reader) {
	dataBytes, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	depths, err := readInts(dataBytes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", partOne(depths))
	fmt.Printf("Part 2: %d\n", partTwo(depths))
}
