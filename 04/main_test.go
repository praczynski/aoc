package main

import (
	"fmt"
	"testing"
)

func TestCountXMAS(t *testing.T) {
	// grid := [][]rune{
	// 	{'S', '.', 'M', '.', 'D'},
	// 	{'.', 'A', '.', '.', 'D'},
	// 	{'S', '.', 'M', '.', 'D'},
	// }
	grid := [][]rune{
		{'.', 'M', '.', 'S', '.', '.', '.', '.', '.', '.'},
		{'.', '.', 'A', '.', '.', 'M', 'S', 'M', 'S', '.'},
		{'.', 'M', '.', 'S', '.', 'M', 'A', 'A', '.', '.'},
		{'.', '.', 'A', '.', 'A', 'S', 'M', 'S', 'M', '.'},
		{'.', 'M', '.', 'S', '.', 'M', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', '.'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', '.'},
		{'M', '.', 'M', '.', 'M', '.', 'M', '.', 'M', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	expectedCountXMAS := 0
	fmt.Println(len(grid), len(grid[0]))
	expectedChecks := (len(grid) - 2) * (len(grid[0]) - 2) * 6
	expectedCountXMASPatterns := 9

	count := countXMAS(grid)

	if count != expectedCountXMAS {
		t.Errorf("countXMAS() = %d; want %d", count, expectedCountXMAS)
	}

	count, checks := countXMASPatterns(grid)

	fmt.Println(expectedChecks, checks)
	if count != expectedCountXMASPatterns {
		t.Errorf("countXMASPatterns() = %d; want %d", count, expectedCountXMASPatterns)
	}
	if checks != expectedChecks {
		t.Errorf("checks = %d; want %d", checks, expectedChecks)
	}

}
