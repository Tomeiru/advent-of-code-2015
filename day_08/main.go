package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(lines []string) error {
	totalPart1 := 0
	totalPart2 := 0
	for _, line := range lines {
		interpreted, err := strconv.Unquote(line)
		if err != nil {
			return err
		}
		literal := fmt.Sprintf("%#v", line)
		totalPart1 += len(line) - len(interpreted)
		totalPart2 += len(literal) - len(line)
	}
	fmt.Println("Part 1 result is:", totalPart1)
	fmt.Println("Part 2 result is:", totalPart2)
	return nil
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	lines := strings.Split(content, "\n")
	err = solve(lines)
	if err != nil {
		panic(err)
	}
}
