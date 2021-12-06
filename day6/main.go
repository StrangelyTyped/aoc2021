package day6

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/strangelytyped/aoc2021/utils"
)

func run(input []int, days int) int {
	fish := make([]int, 10)
	for _, age := range input {
		fish[age]++
	}

	for i := 0; i < days; i++ {
		spawning := fish[0]
		fish = append(fish[1:], 0)
		fish[6] += spawning
		fish[8] = spawning
	}

	sum := 0
	for _, c := range fish {
		sum += c
	}
	return sum
}

func partOne(input []int) int {
	return run(input, 80)
}

func partTwo(input []int) int {
	return run(input, 256)
}

func readInput(input io.Reader) []int {
	var result []int
	inputTxt, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}
	for _, digit := range strings.Split(string(inputTxt), ",") {
		result = append(result, utils.ParseIntOrPanic(digit))
	}
	return result
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
