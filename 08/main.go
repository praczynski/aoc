package main

import (
	"fmt"
)

type Position struct {
	x, y int
}

func findAntennas(mapData []string) map[rune][]Position {
	antennas := make(map[rune][]Position)
	for y, row := range mapData {
		for x, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], Position{x, y})
			}
		}
	}
	return antennas
}

func calculateAntinodes(antennas map[rune][]Position, width, height int) map[Position]bool {
	antinodes := make(map[Position]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				x1, y1 := positions[i].x, positions[i].y
				x2, y2 := positions[j].x, positions[j].y

				// Check if they are in line horizontally
				if y1 == y2 {
					dist := abs(x2 - x1)
					if dist%2 == 0 {
						d := dist / 2
						antinode1 := Position{x1 - d, y1}
						antinode2 := Position{x2 + d, y2}
						if antinode1.x >= 0 && antinode1.x < width {
							antinodes[antinode1] = true
						}
						if antinode2.x >= 0 && antinode2.x < width {
							antinodes[antinode2] = true
						}
					}
				}

				// Check if they are in line vertically
				if x1 == x2 {
					dist := abs(y2 - y1)
					if dist%2 == 0 {
						d := dist / 2
						antinode1 := Position{x1, y1 - d}
						antinode2 := Position{x2, y2 + d}
						if antinode1.y >= 0 && antinode1.y < height {
							antinodes[antinode1] = true
						}
						if antinode2.y >= 0 && antinode2.y < height {
							antinodes[antinode2] = true
						}
					}
				}

				// Check if they are in line diagonally
				if abs(x2-x1) == abs(y2-y1) {
					dist := abs(x2 - x1)
					if dist%2 == 0 {
						d := dist / 2
						antinode1 := Position{x1 - d, y1 - d}
						antinode2 := Position{x2 + d, y2 + d}
						antinode3 := Position{x1 - d, y1 + d}
						antinode4 := Position{x2 + d, y2 - d}
						if antinode1.x >= 0 && antinode1.x < width && antinode1.y >= 0 && antinode1.y < height {
							antinodes[antinode1] = true
						}
						if antinode2.x >= 0 && antinode2.x < width && antinode2.y >= 0 && antinode2.y < height {
							antinodes[antinode2] = true
						}
						if antinode3.x >= 0 && antinode3.x < width && antinode3.y >= 0 && antinode3.y < height {
							antinodes[antinode3] = true
						}
						if antinode4.x >= 0 && antinode4.x < width && antinode4.y >= 0 && antinode4.y < height {
							antinodes[antinode4] = true
						}
					}
				}
			}
		}
	}
	return antinodes
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	mapData := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	width := len(mapData[0])
	height := len(mapData)

	antennas := findAntennas(mapData)
	antinodes := calculateAntinodes(antennas, width, height)

	fmt.Println(len(antinodes)) // Output the number of unique antinode locations
}
