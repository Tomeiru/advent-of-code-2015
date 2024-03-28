package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func incrementLetter(letter int) int {
	if letter == 'z' {
		return 'a'
	}
	return letter + 1
}

func incrementPassword(password []int) []int {
	password[len(password)-1] = incrementLetter(password[len(password)-1])
	for i := len(password) - 1; i > 0 && password[i] == 'a'; i-- {
		password[i-1] = incrementLetter(password[i-1])
	}
	return password
}

func checkPasswordRequirements(password []int) bool {
	forbiddenLetter := []int{'i', 'o', 'l'}
	maxIncreasing := 0
	increasingCount := 1
	precedent := 0
	pairs := make(map[int]bool)
	for _, letter := range password {
		if slices.Contains(forbiddenLetter, letter) {
			return false
		}
		if precedent+1 == letter {
			increasingCount += 1
		} else {
			maxIncreasing = max(increasingCount, maxIncreasing)
			increasingCount = 1
		}
		if precedent == letter {
			pairs[letter] = true
		}
		precedent = letter
	}
	maxIncreasing = max(increasingCount, maxIncreasing)
	return maxIncreasing >= 3 && len(pairs) >= 2
}

func convertPassword(password []int) string {
	chars := make([]string, len(password))
	for i, char := range password {
		chars[i] = string(rune(char))
	}
	return strings.Join(chars, "")
}

func mutateToNextCompliantPassword(password []int) []int {
	for i := 0; !checkPasswordRequirements(password); i++ {
		password = incrementPassword(password)
	}
	return password
}

func solve(password []int) {
	nextCompliantPassword := mutateToNextCompliantPassword(password)
	fmt.Println("Part 1 result is:", convertPassword(nextCompliantPassword))
	nextNextCompliantPassword := mutateToNextCompliantPassword(incrementPassword(nextCompliantPassword))
	fmt.Println("Part 2 result is:", convertPassword(nextNextCompliantPassword))
}

func parseContent(content string) []int {
	sequence := make([]int, len(content))
	for i, num := range content {
		sequence[i] = int(num)
	}
	return sequence
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	password := parseContent(content)
	solve(password)
}
