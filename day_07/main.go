package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation func(left uint16, right uint16) uint16

type instruction struct {
	operation operation
	left      string
	right     string
	wire      string
}

func AND(left uint16, right uint16) uint16 {
	return left & right
}

func OR(left uint16, right uint16) uint16 {
	return left | right
}

func LSHIFT(left uint16, right uint16) uint16 {
	return left << right
}

func RSHIFT(left uint16, right uint16) uint16 {
	return left >> right
}

func NOT(left uint16, right uint16) uint16 {
	return ^left
}

func VALUE(left uint16, right uint16) uint16 {
	return left
}

func parseInstruction(line string) instruction {
	instructionParts := strings.Split(line, " -> ")
	components := strings.Split(instructionParts[0], " ")
	if len(components) == 1 {
		return instruction{
			operation: VALUE,
			left:      components[0],
			right:     "",
			wire:      instructionParts[1],
		}
	}
	if len(components) == 2 {
		return instruction{
			operation: NOT,
			left:      components[1],
			right:     "",
			wire:      instructionParts[1],
		}
	}
	operations := map[string]operation{
		"AND":    AND,
		"OR":     OR,
		"LSHIFT": LSHIFT,
		"RSHIFT": RSHIFT,
	}
	return instruction{
		operation: operations[components[1]],
		left:      components[0],
		right:     components[2],
		wire:      instructionParts[1],
	}
}

func parseContent(content string) []instruction {
	lines := strings.Split(content, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}
	return instructions
}

func computeValue(field string, signals map[string]uint16) uint16 {
	if field == "" {
		return 0
	}
	value, err := strconv.ParseUint(field, 10, 16)
	if err != nil {
		return signals[field]
	}
	return uint16(value)
}

func solve(instructions []instruction) {
	signals := make(map[string]uint16)
	for _, instr := range instructions {
		left := computeValue(instr.left, signals)
		right := computeValue(instr.right, signals)
		signals[instr.wire] = instr.operation(left, right)
	}
	fmt.Println(signals)
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	instructions := parseContent(content)
	solve(instructions)
}
