package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sue map[string]int

func createIntPointerFromValue(value int) int {
	return value
}

func generateTargetSue() sue {
	target := make(sue)
	target["children"] = createIntPointerFromValue(3)
	target["cats"] = createIntPointerFromValue(7)
	target["samoyeds"] = createIntPointerFromValue(2)
	target["pomeranians"] = createIntPointerFromValue(3)
	target["akitas"] = createIntPointerFromValue(0)
	target["vizslas"] = createIntPointerFromValue(0)
	target["goldfish"] = createIntPointerFromValue(5)
	target["trees"] = createIntPointerFromValue(3)
	target["cars"] = createIntPointerFromValue(2)
	target["perfumes"] = createIntPointerFromValue(1)
	return target
}

func parseContent(content string) []sue {
	lines := strings.Split(content, "\n")
	sues := make([]sue, len(lines))
	for i, line := range lines {
		sues[i] = make(sue)
		_, information, _ := strings.Cut(line, ": ")
		stats := strings.Split(information, ", ")
		for _, stat := range stats {
			splitStat := strings.Split(stat, ": ")
			value, _ := strconv.Atoi(splitStat[1])
			sues[i][splitStat[0]] = value
		}
	}
	return sues
}

func checkFirstSue(target sue, checked sue) bool {
	for name, value := range checked {
		if target[name] != value {
			return false
		}
	}
	return true
}

func checkSecondSue(target sue, checked sue) bool {
	for name, value := range checked {
		if name == "tree" || name == "cat" {
			if checked[name] < target[name] {
				return false
			}
			continue
		}
		if name == "pomeranians" || name == "goldfish" {
			if checked[name] > target[name] {
				return false
			}
			continue
		}
		if target[name] != value {
			return false
		}
	}
	return true
}

func findSueNumber(target sue, sues []sue, checker func(sue, sue) bool) int {
	for i, checked := range sues {
		if checker(target, checked) {
			return i + 1
		}
	}
	return 0
}

func solve(target sue, sues []sue) {
	fmt.Println("Result of part 1:", findSueNumber(target, sues, checkFirstSue))
	fmt.Println("Result of part 2:", findSueNumber(target, sues, checkSecondSue))
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	target := generateTargetSue()
	sues := parseContent(content)
	solve(target, sues)
}
