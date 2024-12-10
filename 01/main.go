package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var distance = 0
var left []int
var right []int
var counter = 0

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var touple []string
	for fileScanner.Scan() {
		touple = strings.Split(fileScanner.Text(), "   ")
		x, err := strconv.Atoi(touple[0])
		if err != nil {
			fmt.Println(err)
		}
		left = append(left, x)
		y, err := strconv.Atoi(touple[1])
		if err != nil {
			fmt.Println(err)
		}
		right = append(right, y)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		counter = int(math.Abs(float64(right[i]) - float64(left[i])))
		distance += counter
	}
	fmt.Println(distance)

	var similarityCounter int

	similarity := make(map[int]int)

	for _, value := range right {
		similarity[value] += 1
	}

	for i := 0; i < len(left); i++ {
		similarityCounter += left[i] * similarity[left[i]]
	}

	fmt.Println(similarityCounter)
}
