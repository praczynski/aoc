package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

// RemoveIndex removes the element at the specified index from a slice
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	content := utils.ReadFileToString("input.txt")

	mulPattern := `mul\((\d+),(\d+)\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	mulRe, err := regexp.Compile(mulPattern)
	if err != nil {
		fmt.Println("Error compiling mul regex:", err)
		return
	}
	doRe, err := regexp.Compile(doPattern)
	if err != nil {
		fmt.Println("Error compiling do regex:", err)
		return
	}
	dontRe, err := regexp.Compile(dontPattern)
	if err != nil {
		fmt.Println("Error compiling don't regex:", err)
		return
	}

	mulEnabled := true
	sum := 0

	matches := mulRe.FindAllStringSubmatchIndex(content, -1)
	doMatches := doRe.FindAllStringIndex(content, -1)
	dontMatches := dontRe.FindAllStringIndex(content, -1)

	type match struct {
		start, end int
		kind       string
		submatches []int
	}
	allMatches := []match{}
	for _, m := range matches {
		allMatches = append(allMatches, match{m[0], m[1], "mul", m})
	}
	for _, m := range doMatches {
		allMatches = append(allMatches, match{m[0], m[1], "do", nil})
	}
	for _, m := range dontMatches {
		allMatches = append(allMatches, match{m[0], m[1], "don't", nil})
	}

	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i].start < allMatches[j].start
	})

	for _, m := range allMatches {
		switch m.kind {
		case "mul":
			if mulEnabled {
				x, err1 := strconv.Atoi(content[m.submatches[2]:m.submatches[3]])
				y, err2 := strconv.Atoi(content[m.submatches[4]:m.submatches[5]])
				if err1 == nil && err2 == nil {
					sum += x * y
				} else {
					fmt.Println("Error converting captured groups to integers")
				}
			}
		case "do":
			mulEnabled = true
		case "don't":
			mulEnabled = false
		}
	}

	fmt.Println("Sum of results:", sum)
}
