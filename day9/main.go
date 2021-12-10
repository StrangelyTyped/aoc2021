package day9

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/strangelytyped/aoc2021/utils"
)

func getAdjacent(input [][]int, x int, y int) [][]int {
	var squares [][]int
	if  x > 0 {
		squares = append(squares, []int{x-1, y, input[x-1][y]})
	}
	if y > 0 {
		squares = append(squares, []int{x, y-1, input[x][y-1]})
	}
	if x < len(input) - 1 {
		squares = append(squares, []int{x+1, y, input[x+1][y]})
	}
	if y < len(input[x]) - 1 {
		squares = append(squares, []int{x, y+1, input[x][y+1]})
	}
	return squares
}

func findLowPoints(input [][]int) [][]int {
	var lowPoints [][]int
	for x, row := range input {
		for y, chr := range row {
			isLow := true
			for _, sq := range getAdjacent(input, x, y) {
				if sq[2] <= chr {
					isLow = false
					break
				}
			}
			if isLow {
				lowPoints = append(lowPoints, []int{x, y, chr})
			}
		}
	}
	return lowPoints
}

func partOne(input [][]int) int {
	score := 0
	for _, point := range findLowPoints(input){
		score += point[2] + 1
	}
	return score
}


func partTwo(input [][]int) int {
	var basinSizes []int
	for _, basin := range findLowPoints(input) {
		basinSize := 0
		searchQueue := [][]int{basin}
		checked := make(map[int]bool)
		for len(searchQueue) > 0 {
			basinSize++
			thisSpace := searchQueue[0]
			searchQueue = searchQueue[1:]
			adjacent := getAdjacent(input, thisSpace[0], thisSpace[1])
			for _, adjSpace := range adjacent {
				if adjSpace[2] != 9 && adjSpace[2] > thisSpace[2] && !checked[(adjSpace[0] << 32) | adjSpace[1]] {
					searchQueue = append(searchQueue, adjSpace)
					checked[(adjSpace[0] << 32) | adjSpace[1]] = true
				}
			}
		}
		basinSizes = append(basinSizes, basinSize)
	}
	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes) - 3] * basinSizes[len(basinSizes) - 2] * basinSizes[len(basinSizes) - 1]
}

func readInput(input io.Reader) [][]int {
	var instructions [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, c := range strings.Split(line, "") {
			row = append(row, utils.ParseIntOrPanic(c))
		}
		instructions = append(instructions, row)
	}
	return instructions
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
