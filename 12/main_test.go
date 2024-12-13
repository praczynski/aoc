package main

import (
	"testing"
)

func TestExample0(t *testing.T) {
	input := [][]rune{
		{'A', 'B'},
		{'C', 'D'},
	}
	expectedPrice := 16
	expectedSidePrice := 16
	result, result2 := calculateTotalCost(input)
	if result != expectedPrice {
		t.Errorf("Expected perimeter price %v, but got %v", expectedPrice, result)
	}
	if result2 != expectedSidePrice {
		t.Errorf("Expected side price %v, but got %v", expectedSidePrice, result2)
	}
}

func TestExample1(t *testing.T) {
	input := [][]rune{
		{'A', 'A', 'A', 'A'},
		{'B', 'B', 'C', 'D'},
		{'B', 'B', 'C', 'C'},
		{'E', 'E', 'E', 'C'},
	}
	expectedPrice := 140
	expectedSidePrice := 80
	result, result2 := calculateTotalCost(input)
	if result != expectedPrice {
		t.Errorf("Expected perimeter price %v, but got %v", expectedPrice, result)
	}
	if result2 != expectedSidePrice {
		t.Errorf("Expected side price %v, but got %v", expectedSidePrice, result2)
	}
}

func TestExample2(t *testing.T) {
	input := [][]rune{
		{'0', '0', '0', '0', '0'},
		{'0', 'X', '0', 'X', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', 'X', '0', 'X', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expectedPrice := 772
	expectedSidePrice := 436
	result, result2 := calculateTotalCost(input)
	if result != expectedPrice {
		t.Errorf("Expected perimeter price %v, but got %v", expectedPrice, result)
	}
	if result2 != expectedSidePrice {
		t.Errorf("Expected side price %v, but got %v", expectedSidePrice, result2)
	}
}
