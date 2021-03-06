package main

import (
	"fmt"
	"io"
	"os"

	"github.com/strangelytyped/aoc2021/day1"
	"github.com/strangelytyped/aoc2021/day2"
	"github.com/strangelytyped/aoc2021/day3"
	"github.com/strangelytyped/aoc2021/day4"
	"github.com/strangelytyped/aoc2021/day5"
	"github.com/strangelytyped/aoc2021/day6"
	"github.com/strangelytyped/aoc2021/day7"
	"github.com/strangelytyped/aoc2021/day8"
	"github.com/strangelytyped/aoc2021/day9"
	"github.com/strangelytyped/aoc2021/day10"
	"github.com/strangelytyped/aoc2021/day11"
	"github.com/strangelytyped/aoc2021/day12"
)

func openOrPanic(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func main() {
	fmt.Println("Day 1")
	day1.Main(openOrPanic("inputs/day1.txt"))
	fmt.Println("Day 2")
	day2.Main(openOrPanic("inputs/day2.txt"))
	fmt.Println("Day 3")
	day3.Main(openOrPanic("inputs/day3.txt"))
	fmt.Println("Day 4")
	day4.Main(openOrPanic("inputs/day4.txt"))
	fmt.Println("Day 5")
	day5.Main(openOrPanic("inputs/day5.txt"))
	fmt.Println("Day 6")
	day6.Main(openOrPanic("inputs/day6.txt"))
	fmt.Println("Day 7")
	day7.Main(openOrPanic("inputs/day7.txt"))
	fmt.Println("Day 8")
	day8.Main(openOrPanic("inputs/day8.txt"))
	fmt.Println("Day 9")
	day9.Main(openOrPanic("inputs/day9.txt"))
	fmt.Println("Day 10")
	day10.Main(openOrPanic("inputs/day10.txt"))
	fmt.Println("Day 11")
	day11.Main(openOrPanic("inputs/day11.txt"))
	fmt.Println("Day 12")
	day12.Main(openOrPanic("inputs/day12.txt"))
}
