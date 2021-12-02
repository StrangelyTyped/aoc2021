package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type position struct {
	depth int
	horiz int
}

func partOne(input []instruction) int {
	pos := position{}

	for _, instr := range input {
		switch(instr.direction){
		case "forward":
			pos.horiz += instr.distance
		case "down":
			pos.depth += instr.distance
		case "up":
			pos.depth -= instr.distance
		}
	}

	return pos.depth * pos.horiz
}

func partTwo(input []instruction) int {
	pos := position{}
	aim := 0

	for _, instr := range input {
		switch(instr.direction){
		case "forward":
			pos.horiz += instr.distance
			pos.depth += aim * instr.distance
		case "down":
			aim += instr.distance
		case "up":
			aim -= instr.distance
		}
	}

	return pos.depth * pos.horiz
}

type instruction struct {
	direction string
	distance int
}

func readInstructions(input io.Reader) []instruction {
	var instructions []instruction
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		dir := scanner.Text()
		scanner.Scan()
		dist, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{
			direction: dir,
			distance: dist,
		})
	}
	return instructions
}

func Main(input io.Reader){
	instructions := readInstructions(input)
	fmt.Printf("Part 1: %d\n", partOne(instructions))
	fmt.Printf("Part 2: %d\n", partTwo(instructions))
}