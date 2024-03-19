package main

import (
	"fmt"
	"os"
	"strings"
)

func isForbiddenCombination(precedent int32, actual int32) bool {
	index := strings.IndexRune("acpx", precedent)
	if index == -1 {
		return false
	}
	return index == strings.IndexRune("bdqy", actual)
}

func isStringNicePart1(line string) bool {
	vowels := 0
	twice := false
	var precedentChar int32 = 0

	for _, char := range line {
		if strings.ContainsRune("aeiou", char) {
			vowels += 1
		}
		if precedentChar == char {
			twice = true
		}
		if isForbiddenCombination(precedentChar, char) {
			return false
		}
		precedentChar = char
	}
	return twice && vowels >= 3
}

func hasOneThreeRep(line string) bool {
	has := false
	for i := 2; i < len(line); i++ {
		if line[i-2] == line[i] {
			has = true
		}
	}
	return has
}

func hasPairRep(line string) bool {
	set := make(map[string]bool)
	has := false
	precedent := ""
	for i := 0; i < len(line)-1; i++ {
		pair := string(line[i]) + string(line[i+1])
		if precedent == pair {
			precedent = ""
			continue
		}
		if set[pair] {
			has = true
		}
		set[pair] = true
		precedent = pair
	}
	return has
}

func isStringNicePart2(line string) bool {
	return hasOneThreeRep(line) && hasPairRep(line)
}

func solve(lines []string) {
	totalPart1 := 0
	totalPart2 := 0
	for _, line := range lines {
		if isStringNicePart1(line) {
			totalPart1++
		}
		if isStringNicePart2(line) {
			totalPart2++
		}
	}
	fmt.Println("Part 1 result is:", totalPart1)
	fmt.Println("Part 2 result is:", totalPart2)
	return
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	lines := strings.Split(content, "\n")
	solve(lines)
}
