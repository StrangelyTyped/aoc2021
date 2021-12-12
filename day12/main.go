package day12

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func partOne(input map[string][]string) int {
	return len(enumeratePaths(input, []string{"start"}, true))
}

func enumeratePaths(input map[string][]string, path []string, isPartOne bool) [][]string {
	current := path[len(path)-1]
	if current == "end" {
		return [][]string{path}
	}

	var results [][]string
	targets := input[current]
	for _, target := range targets {
		if !isValidTarget(path, target, isPartOne) {
			continue
		}
		var newPath []string
		newPath = append(newPath, path...)
		newPath = append(newPath, target)
		results = append(results, enumeratePaths(input, newPath, isPartOne)...)
	}
	return results
}

func isValidTarget(path []string, target string, isPartOne bool) bool {
	if strings.ToLower(target) != target {
		return true
	}
	if target == "start" {
		return false
	}
	for _, t := range path {
		if t == target {
			if isPartOne {
				return false
			} else if pathHasDuplicateSmall(path) {
				return false
			}
		}
	}
	return true
}

func pathHasDuplicateSmall(path []string) bool {
	occurrences := make(map[string]bool)
	for _, entry := range path {
		if strings.ToLower(entry) == entry {
			_, ok := occurrences[entry]
			if ok {
				return true
			}
			occurrences[entry] = true
		}
	}
	return false
}

func partTwo(input map[string][]string) int {
	return len(enumeratePaths(input, []string{"start"}, false))
}

func readInput(input io.Reader) map[string][]string {
	instructions := make(map[string][]string)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "-") 
		instructions[row[0]] = append(instructions[row[0]], row[1])
		instructions[row[1]] = append(instructions[row[1]], row[0])
	}
	return instructions
}

func Main(input io.Reader) {
	numbers := readInput(input)
	fmt.Printf("Part 1: %d\n", partOne(numbers))
	fmt.Printf("Part 2: %d\n", partTwo(numbers))
}
