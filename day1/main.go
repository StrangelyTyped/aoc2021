package main

import (
	"fmt"
	"io"
	"os"
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

func partOne(depths []int) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	return count
}

func partTwo(depths []int) int {
	count := 0
	for i := 3; i < len(depths); i++ {
		if sum(depths[i-2:i+1]) > sum(depths[i-3:i]) {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	dataBytes, err := io.ReadAll(file)
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