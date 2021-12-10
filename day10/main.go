package day10

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

var bracketMap = make(map[string]string)
func init(){
	bracketMap["("] = ")"
	bracketMap["["] = "]"
	bracketMap["{"] = "}"
	bracketMap["<"] = ">"
}

func partOne(input []string) int {
	var scoreMap = make(map[string]int)
	scoreMap[")"] = 3
	scoreMap["]"] = 57
	scoreMap["}"] = 1197
	scoreMap[">"] = 25137

	score := 0
	for _, line := range input {
		_, errChr := parse(line)
		if errChr != nil {
			score += scoreMap[*errChr]
		}
	}
	return score
}

func parse(line string) ([]string, *string){
	var stack []string

	for _, chr := range strings.Split(line, "") {
		close, isOpen := bracketMap[chr]
		if isOpen {
			stack = append(stack, close)
		} else if chr == stack[len(stack) - 1] {
			stack = stack[:len(stack)-1]
		} else {
			return stack, &chr
		}
	}
	return stack, nil
}

func partTwo(input []string) int {
	var scoreMap2 = make(map[string]int)
	scoreMap2[")"] = 1
	scoreMap2["]"] = 2
	scoreMap2["}"] = 3
	scoreMap2[">"] = 4

	var scores []int
	for _, line := range input {
		stack, errChr := parse(line)
		if errChr == nil {
			thisScore := 0
			for i := len(stack) - 1; i >= 0; i-- {
				thisScore = (thisScore * 5) + scoreMap2[stack[i]]
			}
			scores = append(scores, thisScore)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func readInput(input io.Reader) []string {
	var instructions []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	return instructions
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
