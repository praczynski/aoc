package main

import (
	"testing"
)

func TestCountNumbersAfterBlinks(t *testing.T) {
	input := []int{125, 17}
	tests := []struct {
		blinks int
		stones int
	}{
		{1, 3},
		{2, 4},
		{3, 5},
		{4, 9},
		{5, 13},
		{6, 22},
		{25, 55312},
	}

	for _, test := range tests {
		result := countStonesAfterNBlinks(input, test.blinks)
		if result != test.stones {
			t.Errorf("Expected %v, but got %v", test.stones, result)
		}
	}
}

var num = 75

func BenchmarkCountNumbersAfterBlinks(b *testing.B) {
	input := []int{510613, 358, 84, 40702, 4373582, 2, 0, 1584}

	for i := 0; i < b.N; i++ {
		countStonesAfterNBlinks(input, num)
	}
}
