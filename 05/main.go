package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) ([][2]int, [][]int) {
	sections := strings.Split(input, "\n\n")
	rulesSection := sections[0]
	updatesSection := sections[1]

	var rules [][2]int
	for _, line := range strings.Split(rulesSection, "\n") {
		parts := strings.Split(line, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		rules = append(rules, [2]int{x, y})
	}

	var updates [][]int
	for _, line := range strings.Split(updatesSection, "\n") {
		var update []int
		for _, part := range strings.Split(line, ",") {
			page, _ := strconv.Atoi(part)
			update = append(update, page)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func createGraph(rules [][2]int) map[int][]int {
	graph := make(map[int][]int)
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		graph[x] = append(graph[x], y)
	}
	return graph
}

func dfs(graph map[int][]int, start int, visited map[int]bool, result *[]int) {
	visited[start] = true

	*result = append(*result, start)

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited, result)
		}
	}
}

func checkUpdate(sorted []int, update []int) bool {
	updatePointer := 0
	for i := 0; i < len(sorted); i++ {
		if update[updatePointer] == sorted[i] {
			updatePointer++
		}
		if updatePointer == len(update) {
			return true
		}
	}
	return false
}

func main() {
	input := utils.ReadFileToString("0.txt")

	rules, updates := parseInput(input)

	graph := createGraph(rules)

	visited := make(map[int]bool)

	var result []int

	for node := range graph {
		if !visited[node] {
			dfs(graph, node, visited, &result)
		}
	}

	var part1 [][]int

	for _, v := range updates {
		visited := make(map[int]bool)

		var result []int

		dfs(graph, v[0], visited, &result)
		for node := range graph {
			if !visited[node] {
				dfs(graph, node, visited, &result)
			}
		}
		if checkUpdate(result, v) {
			part1 = append(part1, v)
		}
	}

	fmt.Println(graph, result, part1)
}
