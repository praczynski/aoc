package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Moves struct {
	x, y int
}
type Point struct {
	x, y int
}

type Record struct {
	ButtonA Moves
	ButtonB Moves
	Prize   Point
}

func parseLine(line string) (Moves, Moves, Point, error) {
	var buttonA, buttonB Moves
	var prize Point
	var err error

	if strings.HasPrefix(line, "Button A:") {
		_, err = fmt.Sscanf(line, "Button A: X+%d, Y+%d", &buttonA.x, &buttonA.y)
		if err != nil {
			return buttonA, buttonB, prize, err
		}
	} else if strings.HasPrefix(line, "Button B:") {
		_, err = fmt.Sscanf(line, "Button B: X+%d, Y+%d", &buttonB.x, &buttonB.y)
		if err != nil {
			return buttonA, buttonB, prize, err
		}
	} else if strings.HasPrefix(line, "Prize:") {
		_, err = fmt.Sscanf(line, "Prize: X=%d, Y=%d", &prize.x, &prize.y)
		if err != nil {
			return buttonA, buttonB, prize, err
		}
	}

	return buttonA, buttonB, prize, nil
}

func calculateLowestCost(buttonA, buttonB Moves, prize Point) (int, error) {

	minCost := 401

	for a := range 100 {
		for b := range 100 {
			ca, cb := a+1, b+1
			costOfCurrentMoves := 0
			if ca*buttonA.x+cb*buttonB.x == prize.x && ca*buttonA.y+cb*buttonB.y == prize.y {
				costOfCurrentMoves = ca*3 + cb
				if costOfCurrentMoves < minCost {
					minCost = costOfCurrentMoves
				}
			}
		}
	}
	if minCost <= 400 {
		return minCost, nil
	}
	return 0, fmt.Errorf("no solution")
}

func betterCalculateLowestCost(buttonA, buttonB Moves, prize Point) (int, error) {
	calcButtonAPresses := float64(prize.x*buttonB.y-prize.y*buttonB.x) / float64(buttonA.x*buttonB.y-buttonB.x*buttonA.y)

	if math.Ceil(calcButtonAPresses) != calcButtonAPresses {
		return 0, fmt.Errorf("no solution")
	}
	buttonAPresses := int(calcButtonAPresses)
	calcButtonBPresses := float64(prize.x-buttonAPresses*buttonA.x) / float64(buttonB.x)
	if math.Ceil(calcButtonBPresses) != float64(calcButtonBPresses) {
		return 0, fmt.Errorf("no solution")
	}
	return buttonAPresses*3 + int(calcButtonBPresses), nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var records []Record
	var currentRecord Record
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Skip empty lines
			continue
		}

		buttonA, buttonB, prize, err := parseLine(line)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			return
		}

		if strings.HasPrefix(line, "Button A:") {
			currentRecord.ButtonA = buttonA
			lineCount++
		} else if strings.HasPrefix(line, "Button B:") {
			currentRecord.ButtonB = buttonB
			lineCount++
		} else if strings.HasPrefix(line, "Prize:") {
			currentRecord.Prize = prize
			lineCount++
		}

		if lineCount == 3 {
			records = append(records, currentRecord)
			currentRecord = Record{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalCost := 0
	totalCostPart2 := 0
	extended := 10000000000000

	for _, record := range records {
		costOfRecord, _ := betterCalculateLowestCost(record.ButtonA, record.ButtonB, record.Prize)
		totalCost += costOfRecord
		record.Prize.x += extended
		record.Prize.y += extended
		costOfRecord2, _ := betterCalculateLowestCost(record.ButtonA, record.ButtonB, record.Prize)
		totalCostPart2 += costOfRecord2
	}

	fmt.Println("Cost:", totalCost)
	fmt.Println("Cost Part2:", totalCostPart2)
}
