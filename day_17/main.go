package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func bruteforce(quantity int, used int, containers []int, combination map[int]int) (int, map[int]int) {
	total := 0
	for i, container := range containers {
		if container > quantity {
			continue
		}
		if container == quantity {
			combination[used+1] += 1
			total += 1
		} else {
			value, _ := bruteforce(quantity-container, used+1, containers[i+1:], combination)
			total += value
		}
	}
	return total, combination
}

func solve(containers []int) {
	total, combination := bruteforce(150, 0, containers, make(map[int]int))
	minimum := math.MaxInt
	totalMinimum := 0
	for value, quantity := range combination {
		if value < minimum {
			minimum = value
			totalMinimum = quantity
		}
	}
	fmt.Println("Result of part 1:", total)
	fmt.Println("Result of part 2:", totalMinimum)
}

func parseContent(content string) []int {
	lines := strings.Split(content, "\n")
	containers := make([]int, len(lines))
	for i, line := range lines {
		containers[i], _ = strconv.Atoi(line)
	}
	slices.Sort(containers)
	slices.Reverse(containers)
	return containers
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	containers := parseContent(content)
	solve(containers)
}
