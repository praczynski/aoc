package main

import (
	"testing"
)

func TestFindAntennas(t *testing.T) {
	input := "2333133121414131402"
	checksum := 1928

	antennas := findAntennas(mapData)

	if len(antennas) != len(expected) {
		t.Fatalf("expected %d antenna types, got %d", len(expected), len(antennas))
	}

	for key, positions := range expected {
		if len(antennas[key]) != len(positions) {
			t.Fatalf("expected %d positions for frequency %c, got %d", len(positions), key, len(antennas[key]))
		}
		for i, pos := range positions {
			if antennas[key][i] != pos {
				t.Fatalf("expected position %v for frequency %c, got %v", pos, key, antennas[key][i])
			}
		}
	}
}

func TestCalculateAntinodes(t *testing.T) {
	mapData := []string{
		"............",
		"............",
		"............",
		"....a.......",
		"............",
		"......a.....",
		"............",
		"............",
		"............",
		"............",
	}

	width := len(mapData[0])
	height := len(mapData)

	antennas := findAntennas(mapData)
	antinodes := calculateAntinodes(antennas, width, height)

	expectedAntinodes := map[Position]bool{
		{2, 1}: true, {8, 7}: true,
	}

	if len(antinodes) != len(expectedAntinodes) {
		t.Fatalf("expected %d unique antinodes, got %d", len(expectedAntinodes), len(antinodes))
	}

	for pos := range expectedAntinodes {
		if !antinodes[pos] {
			t.Fatalf("expected antinode at position %v", pos)
		}
	}
}
