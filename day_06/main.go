package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coordinates struct {
	x, y int
}

type instruction struct {
	instruction string
	start       coordinates
	end         coordinates
}

type operation func(previousValue int) int

func parseContent(content string) ([]instruction, error) {
	lines := strings.Split(content, "\n")
	instructions := make([]instruction, len(lines))
	regex, err := regexp.Compile("(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)")
	if err != nil {
		return nil, err
	}
	for i, line := range lines {
		groups := regex.FindStringSubmatch(line)
		if groups == nil {
			return nil, fmt.Errorf("formatting error on line %d of the input file", i+1)
		}
		convertedPositions := make([]int, 4)
		for ii := 2; ii < 6; ii++ {
			value, err := strconv.Atoi(groups[ii])
			if err != nil {
				return nil, fmt.Errorf("formatting error on line %d of the input file", i+1)
			}
			convertedPositions[ii-2] = value
		}
		instructions[i] = instruction{
			instruction: groups[1],
			start: coordinates{
				x: convertedPositions[0],
				y: convertedPositions[1],
			},
			end: coordinates{
				x: convertedPositions[2],
				y: convertedPositions[3],
			},
		}
	}
	return instructions, nil
}

func evaluatePlot(plot [1000][1000]int) int {
	total := 0
	for _, line := range plot {
		for _, value := range line {
			total += value
		}
	}
	return total
}

func generatePlot(instructions []instruction, operations map[string]operation) [1000][1000]int {
	var plot [1000][1000]int
	for _, instr := range instructions {
		for i := instr.start.y; i <= instr.end.y; i++ {
			for ii := instr.start.x; ii <= instr.end.x; ii++ {
				plot[i][ii] = operations[instr.instruction](plot[i][ii])
			}
		}
	}
	return plot
}

func solve(instructions []instruction) {
	binaryOperations := map[string]operation{
		"turn on": func(previousValue int) int {
			return 1
		},
		"turn off": func(previousValue int) int {
			return 0
		},
		"toggle": func(previousValue int) int {
			if previousValue == 1 {
				return 0
			}
			return 1
		},
	}
	intensityOperations := map[string]operation{
		"turn on": func(previousValue int) int {
			return previousValue + 1
		},
		"turn off": func(previousValue int) int {
			if previousValue == 0 {
				return 0
			}
			return previousValue - 1
		},
		"toggle": func(previousValue int) int {
			return previousValue + 2
		},
	}
	fmt.Println("Part 1 result is:", evaluatePlot(generatePlot(instructions, binaryOperations)))
	fmt.Println("Part 2 result is:", evaluatePlot(generatePlot(instructions, intensityOperations)))
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	instructions, err := parseContent(content)
	if err != nil {
		panic(err)
	}
	solve(instructions)
}
