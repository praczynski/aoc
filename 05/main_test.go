package main

import (
	"fmt"
	"testing"
)

func TestSortGraph(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {3},
		4: {3},
	}

	visited := make(map[int]bool)

	var result []int

	dfs(graph, 12, visited, &result)

	fmt.Println(result)

	// fmt.Println(expectedChecks, checks)
	// if count != expectedCountXMASPatterns {
	// 	t.Errorf("countXMASPatterns() = %d; want %d", count, expectedCountXMASPatterns)
	// }
	// if checks != expectedChecks {
	// 	t.Errorf("checks = %d; want %d", checks, expectedChecks)
	// }

}
