package day5

import (
	"github.com/strangelytyped/aoc2021/utils"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)
type terrainMap map[int64]int


func expandAndMap(terrain terrainMap, row []int) {
	x,y := row[0],row[1]
	xInc, yInc := utils.Signum(row[2] - row[0]), utils.Signum(row[3] - row[1])
	for x != row[2] || y != row[3] {
		terrain[int64(x << 32) | int64(y)]++
		x += xInc
		y += yInc
	}
	terrain[int64(x << 32) | int64(y)]++
}

func run(input [][]int, filter bool) int {
	terrain := make(terrainMap)

	for _, row := range input {
		if !filter || row[0] == row[2] || row[1] == row[3] {
			expandAndMap(terrain, row)
		}
	}

	count := 0
	for _, x := range terrain {
		if x > 1 {
			count++
		}
	}

	return count
}

func partOne(input [][]int) int {
	return run(input, true)
}


func partTwo(input [][]int) int {
	return run(input, false)
}

func readInput(input io.Reader) [][]int {
	var result [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.FieldsFunc(line, func(r rune) bool { return !unicode.IsNumber(r) })
		var nums []int
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		result = append(result, nums)
	}
	return result
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
