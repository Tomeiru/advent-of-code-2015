package main

import (
	"fmt"
	"os"
)

func generateLookAndSay(input []int) []int {
	previousNumber := 0
	runLength := 0
	result := make([]int, 0)
	for _, digit := range input {
		if previousNumber == digit {
			runLength++
			continue
		}
		if runLength != 0 {
			result = append(result, runLength, previousNumber)
		}
		runLength = 1
		previousNumber = digit
	}
	result = append(result, runLength, previousNumber)
	return result
}

func solve(sequence []int) {
	solution := sequence
	resultPart1 := 0
	for i := 0; i < 50; i++ {
		solution = generateLookAndSay(solution)
		if i == 39 {
			resultPart1 = len(solution)
		}
	}
	fmt.Println("Part 1 result is:", resultPart1)
	fmt.Println("Part 2 result is:", len(solution))
}

func parseContent(content string) []int {
	sequence := make([]int, len(content))
	for i, num := range content {
		sequence[i] = int(num) - 48
	}
	return sequence
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	sequence := parseContent(content)
	solve(sequence)
}
