package main

import (
	"advent-of-code/utils"
	"fmt"
)

var directions = [4][2]int{
	{1, 0},  // Down
	{0, 1},  // Right
	{-1, 0}, // Up
	{0, -1}, // Left
}

func isValid(x, y int, inputMap [][]rune) bool {
	return x >= 0 && x < len(inputMap[0]) && y >= 0 && y < len(inputMap)
}

func countReachableNines(inputMap [][]rune, startX, startY int) int {
	visited := make(map[[2]int]bool)
	var dfs func(x, y, height int)
	reachableNines := 0

	dfs = func(x, y, height int) {
		if !isValid(x, y, inputMap) || visited[[2]int{x, y}] || int(inputMap[y][x]-'0') != height {
			return
		}
		visited[[2]int{x, y}] = true
		if height == 9 {
			reachableNines++
			return
		}
		for _, dir := range directions {
			nextX := x + dir[0]
			nextY := y + dir[1]
			dfs(nextX, nextY, height+1)
		}
	}

	dfs(startX, startY, 0)
	return reachableNines
}

func countDistinctTrails(inputMap [][]rune, points []rune, point int, row, col int, directions [4][2]int) int {
	// Base case: if point is 9, we have found a valid trail
	if point == 9 {
		return 1
	}

	result := 0

	// Iterate over all directions
	for _, v := range directions {
		nextX := col + v[0]
		nextY := row + v[1]

		// Check if nextX and nextY are within bounds
		if !isValid(nextX, nextY, inputMap) {
			continue
		}

		// Check if the next cell matches the required point
		if int(inputMap[nextY][nextX]-'0') == point+1 {
			// Recursively check the next point in the sequence
			result += countDistinctTrails(inputMap, points, point+1, nextY, nextX, directions)
		}
	}

	return result
}

func main() {
	inputMap := utils.ReadFileToRuneGrid("input.txt")
	totalScore := 0
	totalRating := 0

	for y := range inputMap {
		for x := range inputMap[y] {
			if inputMap[y][x] == '0' {
				totalScore += countReachableNines(inputMap, x, y)
				totalRating += countDistinctTrails(inputMap, []rune("0123456789"), 0, y, x, directions)
			}
		}
	}

	fmt.Println("Total Score:", totalScore)
	fmt.Println("Total Rating:", totalRating)
}
