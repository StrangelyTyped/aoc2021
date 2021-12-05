package utils

import "strconv"

func ParseIntOrPanic(in string) int {
	ret, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return ret
}

func IntTernary(cond bool, a,b int) int {
	if cond { 
		return a
	}
	return b
}

func Signum(x int) int{
	if x == 0 {
		return 0
	}
	if x > 0 {
		return 1
	}
	return -1
}