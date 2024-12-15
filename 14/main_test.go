package main

import (
	"fmt"
	"math"
	"testing"
)

/*
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3

1.12.......
...........
...........
......11.11
1.1........
.........1.
.......1...

......2..1.
...........
1..........
.11........
.....1.....
...12......
.1....1....

expected: 12
*/
func TestCalculateSafetyFactor(t *testing.T) {

	input := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}

	expected := 12
	var moves, cols, rows int = 100, 11, 7
	result := calculateSafetyFactor(input, moves, cols, rows)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsUpperLeft(t *testing.T) {
	pos := Position{0, 0}
	cols := 11
	rows := 7
	expected := true
	fmt.Println(math.Floor(float64(cols) / 2))
	result := isUpperLeft(pos, cols, rows)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func BenchmarkCalculateLowestCost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
