package main

import (
	"advent-of-code/utils"
	"fmt"
)

type Point struct {
	x, y int
}

func floodFill(grid [][]rune, x, y int, visited [][]bool, plantType rune) (int, int, int) {
	stack := []Point{{x, y}}
	area, perimeter, sides := 0, 0, 0
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cx, cy := point.x, point.y

		if visited[cx][cy] {
			continue
		}
		visited[cx][cy] = true
		area++
		localPerimeter := 4
		localSides := 0

		for _, d := range directions {
			nx, ny := cx+d.x, cy+d.y
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
				if grid[nx][ny] == plantType {
					if !visited[nx][ny] {
						stack = append(stack, Point{nx, ny})
					}
					localPerimeter--
				}
			} else {
				localSides++
			}
		}

		perimeter += localPerimeter
		sides += localSides
	}
	return area, perimeter, sides
}

func calculateTotalCost(grid [][]rune) (int, int) {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	totalCostPart1, totalCostPart2 := 0, 0

	for x := range grid {
		for y := range grid[0] {
			if !visited[x][y] {
				plantType := grid[x][y]
				area, perimeter, sides := floodFill(grid, x, y, visited, plantType)
				costPart1 := area * perimeter
				costPart2 := area * sides
				totalCostPart1 += costPart1
				totalCostPart2 += costPart2
			}
		}
	}
	return totalCostPart1, totalCostPart2
}

func main() {
	input := utils.ReadFileToRuneGrid("input.txt")
	totalCostPart1, totalCostPart2 := calculateTotalCost(input)

	fmt.Println("Part 1:", totalCostPart1)
	fmt.Println("Part 2:", totalCostPart2)
}
