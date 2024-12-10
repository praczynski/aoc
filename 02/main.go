package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var counter, counter2 int

// RemoveIndex removes the element at the specified index from a slice
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		parsedLine, err := parseLine(line)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}

		if checkLine(parsedLine) {
			counter++
		}

		isAnyModifiedLineSafe := false

		for i := 0; i < len(parsedLine); i++ {
			tmp := make([]int, len(parsedLine))
			copy(tmp, parsedLine)
			modifiedLine := RemoveIndex(tmp, i)

			if checkLine(modifiedLine) {
				isAnyModifiedLineSafe = true
				break
			}
		}

		if isAnyModifiedLineSafe {
			counter2++
		}
	}

	fmt.Println("part1: ", counter)
	fmt.Println("part2: ", counter2)
}

func readLines(filename string) ([]string, error) {
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
	return lines, scanner.Err()
}

func parseLine(line string) ([]int, error) {
	var result []int
	for _, num := range strings.Split(line, " ") {
		val, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

func checkLine(line []int) bool {
	if len(line) < 2 {
		panic("Incorrect input")
	}
	var x, y, z int
	for i := 0; i < len(line)-1; i++ {
		t := line[i+1] - line[i]
		if t == 1 || t == 2 || t == 3 {
			x++
		} else if t == -1 || t == -2 || t == -3 {
			y++
		} else {
			z++
		}
	}

	if x == len(line)-1 || y == len(line)-1 {
		return true
	}

	return false
}
