package main

import (
	"bufio"
	"fmt"
	"os"
)

// Define the patterns for "X-MAS"
var patterns = [][][]rune{
	{{'M', '.', 'S'}, {'.', 'A', '.'}, {'M', '.', 'S'}},
	{{'S', '.', 'M'}, {'.', 'A', '.'}, {'S', '.', 'M'}},
	{{'M', '.', 'M'}, {'.', 'A', '.'}, {'S', '.', 'S'}},
	{{'S', '.', 'S'}, {'.', 'A', '.'}, {'M', '.', 'M'}},
}

// Directions for searching in the grid
var directions = [8][2]int{
	{0, 1},   // Right
	{1, 0},   // Down
	{1, 1},   // Down-Right
	{1, -1},  // Down-Left
	{0, -1},  // Left
	{-1, 0},  // Up
	{-1, -1}, // Up-Left
	{-1, 1},  // Up-Right
}

// Function to check if a word exists in the grid starting from (row, col) in a given direction
func checkWord(grid [][]rune, word string, row, col, dirX, dirY int) bool {
	for i := 0; i < len(word); i++ {
		newRow := row + i*dirX
		newCol := col + i*dirY
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] != rune(word[i]) {
			return false
		}
	}
	return true
}

// Function to count all occurrences of the word "XMAS" in the grid
func countXMAS(grid [][]rune) int {
	word := "XMAS"
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			for _, direction := range directions {
				if checkWord(grid, word, row, col, direction[0], direction[1]) {
					count++
				}
			}
		}
	}

	return count
}

// Function to check if a pattern exists in the grid starting from (row, col)
func checkPattern(grid [][]rune, pattern [][]rune, row, col int) bool {
	// fmt.Println(grid, pattern, row, col)
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			if pattern[i][j] == '.' {
				continue
			}
			if grid[i][j] != pattern[i][j] {
				return false
			}
		}
	}
	return true
}

// Function to count all occurrences of the X-MAS patterns in the grid
func countXMASPatterns(grid [][]rune) (int, int) {
	count := 0
	checks := 0
	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[0])-2; col++ {
			for _, pattern := range patterns {
				toCheck := [][]rune{
					grid[row][col : col+3],
					grid[row+1][col : col+3],
					grid[row+2][col : col+3],
				}

				checks++
				if checkPattern(toCheck, pattern, row, col) {
					count++
				}
			}
		}
	}
	return count, checks
}

// Function to read the grid from a file
func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid, scanner.Err()
}

func main() {
	// Read the grid from the input file
	grid, err := readGrid("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Count all occurrences of "XMAS"
	countPart1 := countXMAS(grid)
	fmt.Println("Total occurrences of XMAS:", countPart1)

	// Count all occurrences of X-MAS patterns
	countPart2, checks := countXMASPatterns(grid)
	fmt.Printf("Total occurrences of X-MAS: %d after %d checks\n", countPart2, checks)
}
