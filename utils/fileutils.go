package utils

import (
	"bufio"
	"os"
)

// readLines reads lines from a file and returns them as a slice of strings
func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic("")
	}
	return lines
}

func ReadFileToString(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found!")
	}
	return string(content)
}

func ReadFileToRuneGrid(filename string) [][]rune {
	lines, err := ReadLines(filename)
	if err != nil {
		panic("File not found!")
	}

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}
