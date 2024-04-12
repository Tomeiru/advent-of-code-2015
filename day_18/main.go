package main

import (
	"fmt"
	"os"
	"strings"
)

func countTurnedOnNeighbours(grid [][]bool, x int, y int) int {
	neighbours := 0
	for ii := y - 1; ii <= y+1; ii++ {
		if ii < 0 || ii >= len(grid) {
			continue
		}
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= len(grid[ii]) || (ii == y && i == x) {
				continue
			}
			if grid[ii][i] {
				neighbours++
			}
		}
	}
	return neighbours
}

func copyGrid(src [][]bool) [][]bool {
	dest := make([][]bool, len(src))
	for i, line := range src {
		dest[i] = make([]bool, len(line))
		_ = copy(dest[i], src[i])
	}
	return dest
}

func determineNextState(grid [][]bool, x int, y int, broken bool) bool {
	if broken && (x == 0 || x == len(grid[0])-1) && (y == 0 || y == len(grid)-1) {
		return true
	}
	current := grid[y][x]
	onNeighbours := countTurnedOnNeighbours(grid, x, y)
	if current && onNeighbours != 2 && onNeighbours != 3 {
		return false
	}
	if !current && onNeighbours == 3 {
		return true
	}
	return current
}

func createNextGeneration(old [][]bool, broken bool) [][]bool {
	next := copyGrid(old)
	for y := 0; y < len(old); y++ {
		for x := 0; x < len(old[y]); x++ {
			next[y][x] = determineNextState(old, x, y, broken)
		}
	}
	return next
}

func countTurnedOnLights(grid [][]bool) int {
	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] {
				total += 1
			}
		}
	}
	return total
}

func parseContent(content string) [][]bool {
	lines := strings.Split(content, "\n")
	grid := make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for ii, state := range line {
			grid[i][ii] = state == '#'
		}
	}
	return grid
}

func solve(grid [][]bool) {
	result := copyGrid(grid)
	brokenResult := copyGrid(grid)
	brokenResult[0][0] = true
	brokenResult[0][len(grid[0])-1] = true
	brokenResult[len(grid)-1][0] = true
	brokenResult[len(grid)-1][len(grid[0])-1] = true
	for i := 0; i < 100; i++ {
		result = createNextGeneration(result, false)
		brokenResult = createNextGeneration(brokenResult, true)
	}
	fmt.Println("Result of part 1:", countTurnedOnLights(result))
	fmt.Println("Result of part 2:", countTurnedOnLights(brokenResult))
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	grid := parseContent(content)
	solve(grid)
}
