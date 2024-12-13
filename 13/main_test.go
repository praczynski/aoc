package main

import (
	"testing"
)

/*
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400
80 * A = 240
40 * B = 40
lowestCost = 280
*/
func TestCalculateLowestCost(t *testing.T) {
	buttonA := Moves{94, 34}
	buttonB := Moves{22, 67}
	prizePosition := Point{8400, 5400}
	expected := 280
	lowestCost, err := calculateLowestCost(buttonA, buttonB, prizePosition)
	if err != nil {
		t.Error("Not expected Errror")
	}
	if lowestCost != expected {
		t.Errorf("Expected %v, but got %v", expected, lowestCost)
	}
}

/*
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400
80 * A = 240
40 * B = 40
lowestCost = 280
*/
func TestCalculateLowestCostPartAlternate(t *testing.T) {
	buttonA := Moves{94, 34}
	buttonB := Moves{22, 67}
	prizePosition := Point{8400, 5400}
	expected := 280

	lowestCost, err := betterCalculateLowestCost(buttonA, buttonB, prizePosition)
	if err != nil {
		t.Error(err)
	}
	if lowestCost != expected {
		t.Errorf("Expected %v, but got %v", expected, lowestCost)
	}
}
func TestCalculateLowestCostPartAlternateWithError(t *testing.T) {
	buttonA := Moves{94, 34}
	buttonB := Moves{22, 67}
	prizePosition := Point{8401, 5400}

	_, err := betterCalculateLowestCost(buttonA, buttonB, prizePosition)
	if err == nil {
		t.Error("Should raise error")
	}
}

/*
Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=10000000012748, Y=10000000012176
*/
func TestCalculateLowestCostPartAlternatePart2(t *testing.T) {
	buttonA := Moves{26, 66}
	buttonB := Moves{67, 21}
	prizePosition := Point{10000000012748, 10000000012176}

	_, err := betterCalculateLowestCost(buttonA, buttonB, prizePosition)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkCalculateLowestCost(b *testing.B) {
	buttonA := Moves{94, 34}
	buttonB := Moves{22, 67}
	prizePosition := Point{8400, 5400}
	for i := 0; i < b.N; i++ {
		calculateLowestCost(buttonA, buttonB, prizePosition)
	}
}

func BenchmarkCalculateLowestCostAlternate(b *testing.B) {
	buttonA := Moves{94, 34}
	buttonB := Moves{22, 67}
	prizePosition := Point{8400, 5400}
	for i := 0; i < b.N; i++ {
		betterCalculateLowestCost(buttonA, buttonB, prizePosition)
	}
}
