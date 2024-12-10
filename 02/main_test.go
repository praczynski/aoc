package main

import "testing"

func TestCheckLine(t *testing.T) {
	tests := []struct {
		line     []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
		{[]int{3, 3, 6, 7, 9}, false},
		{[]int{58, 55, 52, 52, 52}, false},
		{[]int{80, 79, 81, 81, 83, 84, 83}, false},
		{[]int{41, 41, 41, 44, 45, 47, 54}, false},
		{[]int{40, 42, 43, 44, 45, 47, 47}, false},
		{[]int{45, 45, 46}, false},
		{[]int{45, 45, 43}, false},
		{[]int{45, 44, 43}, true},
	}

	for _, test := range tests {
		result := checkLine(test.line)
		if result != test.expected {
			t.Errorf("For line %v, expected %v, but got %v", test.line, test.expected, result)
		}
	}
}
