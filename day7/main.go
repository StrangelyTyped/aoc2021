package day7

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"sort"
	"strings"

	"github.com/strangelytyped/aoc2021/utils"
)

func run(input []int, costFunc func(int)int) int {
	var minFuel int = math.MaxInt64
	sort.Ints(input)
	for targetPos := input[0]; targetPos < input[len(input)-1]; targetPos++ {
		thisPosCost := 0
		for _, x := range input {
			thisPosCost += costFunc(targetPos - x)
		}
		if thisPosCost > minFuel {
			return minFuel
		}
		minFuel = thisPosCost
	}
	return minFuel
}

func partOne(input []int) int {
	return run(input, func(x int) int { return int(math.Abs(float64(x))) })
}

func partTwo(input []int) int {
	return run(input, func(x int) int { 
		n := int(math.Abs(float64(x)))
		return(n * (n+1)) / 2
	})
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
