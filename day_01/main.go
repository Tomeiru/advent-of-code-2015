package main

import (
	"fmt"
	"os"
)

func solve(content string) {
	floor := 0
	basementPos := -1
	for i := 0; i < len(content); i++ {
		if content[i] == '(' {
			floor += 1
		} else if content[i] == ')' {
			floor -= 1
		}
		if basementPos == -1 && floor == -1 {
			basementPos = i + 1
		}
	}
	fmt.Println("Part 1 result is:", floor)
	fmt.Println("Part 2 result is:", basementPos)
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	solve(content)
}
