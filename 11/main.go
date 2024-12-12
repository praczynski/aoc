package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type Count struct {
	stone, blinks int
}

var memo = make(map[Count]int)

func length(stone int) int {
	return len(strconv.Itoa(stone))
}

func split(stone int) (int, int) {
	stringInt := strconv.Itoa(stone)
	split := len(stringInt) / 2
	x, _ := strconv.Atoi(stringInt[:split])
	y, _ := strconv.Atoi(stringInt[split:])
	return x, y
}

func count(stone int, blinks int) int {
	if result, found := memo[Count{stone, blinks}]; found {
		return result
	}
	result := 0
	if blinks == 0 {
		result = 1
	} else if stone == 0 {
		result = count(1, blinks-1)
	} else if length(stone)%2 == 0 {
		x, y := split(stone)
		result = count(x, blinks-1) + count(y, blinks-1)

	} else {
		result = count(stone*2024, blinks-1)
	}

	memo[Count{stone, blinks}] = result

	return result
}

func countStonesAfterNBlinks(input []int, blinks int) int {
	result := 0

	for _, stone := range input {
		result += count(stone, blinks)
	}

	return result
}

func main() {
	input := utils.ReadFileToString("input.txt")
	split := strings.Split(input, " ")
	var tt []int
	for _, v := range split {
		k, _ := strconv.Atoi(v)
		tt = append(tt, k)
	}

	fmt.Println("Count after 25:", countStonesAfterNBlinks(tt, 25))
	fmt.Println("Count after 75:", countStonesAfterNBlinks(tt, 75))
}
