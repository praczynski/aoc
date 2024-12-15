package main

import (
	"fmt"
	"math"
	"advent-of-code/utils"
)

type Position struct {
	x, y int
}
type Velocity struct {
	x, y int
}

func parseLine(input string) (Position, Velocity) {
	var pos Position
	var vel Velocity

	_, err := fmt.Sscanf(input, "p=%d,%d v=%d,%d", &pos.x, &pos.y, &vel.x, &vel.y)

	if err != nil {
		panic(err)
	}

	return pos, vel
}

func moveGuard(moves, cols, rows int, pos Position, vel Velocity) Position {
	newX := int(math.Remainder(float64(pos.x)+float64(moves)*float64(vel.x), float64(cols)))
	if newX < 0 {
		newX = cols + newX
	}
	newY := int(math.Remainder(float64(pos.y)+float64(moves)*float64(vel.y), float64(rows)))
	if newY < 0 {
		newY = rows + newY
	}
	return Position{newX, newY}
}

func isUpperLeft(pos Position, cols, rows int) bool {
	return pos.x >= 0 && pos.x < int(math.Floor(float64(cols)/2)) && pos.y >= 0 && pos.y <= int(math.Floor(float64(rows)/2))
}

func isUpperRight(pos Position, cols, rows int) bool {
	return pos.x <= cols && pos.x >= int(math.Ceil(float64(cols)/2)) && pos.y >= 0 && pos.y <= int(math.Floor(float64(rows)/2))
}

func isBottomLeft(pos Position, cols, rows int) bool {
	return pos.x >= 0 && pos.x <= int(math.Floor(float64(cols)/2)) && pos.y <= rows && pos.y >= int(math.Ceil(float64(rows)/2))
}

func isBottomRight(pos Position, cols, rows int) bool {
	return pos.x <= cols && pos.x >= int(math.Ceil(float64(cols)/2)) && pos.y <= rows && pos.y >= int(math.Ceil(float64(rows)/2))
}

func calculateSafetyFactor(input []string, moves, cols, rows int) int {
	var ul, ur, bl, br int = 0, 0, 0, 0
	for _, guard := range input {
		pos, val := parseLine(guard)
		newPosistion := moveGuard(moves, cols, rows, pos, val)
		if isUpperLeft(newPosistion, cols, rows) {
			ul++
		}
		if isUpperRight(newPosistion, cols, rows) {
			ur++
		}
		if isBottomLeft(newPosistion, cols, rows) {
			bl++
		}
		if isBottomRight(newPosistion, cols, rows) {
			br++
		}
	}
	return ul * ur * bl * br
}

var moves, cols, rows int = 100, 101, 103

func main() {
	for _, input := range utils.ReadLines("input.txt")

	var input string = "p=0,4 v=3,-3"

	pos, vel := parseLine(input)

	var lines []string
	lines = append(lines, input)
	fmt.Printf("Position: %+v\n", pos)
	fmt.Printf("Velocity: %+v\n", vel)

	safetyFactor := calculateSafetyFactor(lines, moves, cols, rows)
	fmt.Println(safetyFactor)
}
