package main

import (
	"fmt"
	"io"
	"os"

	"github.com/strangelytyped/aoc2021/day1"
	"github.com/strangelytyped/aoc2021/day2"
)

func openOrPanic(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func main(){
	fmt.Println("Day 1")
	day1.Main(openOrPanic("inputs/day1.txt"))
	fmt.Println("Day 2")
	day2.Main(openOrPanic("inputs/day2.txt"))
}