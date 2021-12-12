package day12

import (
	"strings"
	"testing"
)

var testData1 string = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var testData2 string = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var testData3 string = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestPartOne1(t *testing.T) {
	nums := readInput(strings.NewReader(testData1))

	result := partOne(nums)

	if result != 10 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartOne2(t *testing.T) {
	nums := readInput(strings.NewReader(testData2))

	result := partOne(nums)

	if result != 19 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartOne3(t *testing.T) {
	nums := readInput(strings.NewReader(testData3))

	result := partOne(nums)

	if result != 226 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo1(t *testing.T) {
	nums := readInput(strings.NewReader(testData1))

	result := partTwo(nums)

	if result != 36 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo2(t *testing.T) {
	nums := readInput(strings.NewReader(testData2))

	result := partTwo(nums)

	if result != 103 {
		t.Errorf("Unexpected result: got %d", result)
	}
}

func TestPartTwo3(t *testing.T) {
	nums := readInput(strings.NewReader(testData3))

	result := partTwo(nums)

	if result != 3509 {
		t.Errorf("Unexpected result: got %d", result)
	}
}
