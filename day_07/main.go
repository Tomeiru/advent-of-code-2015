package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type operationFunction func(left uint16, right uint16) uint16

type operation struct {
	function operationFunction
	left     string
	right    string
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

func parseOperation(ope string) operation {
	components := strings.Split(ope, " ")
	if len(components) == 1 {
		return operation{
			function: VALUE,
			left:     components[0],
			right:    "",
		}
	}
	if len(components) == 2 {
		return operation{
			function: NOT,
			left:     components[1],
			right:    "",
		}
	}
	functions := map[string]operationFunction{
		"AND":    AND,
		"OR":     OR,
		"LSHIFT": LSHIFT,
		"RSHIFT": RSHIFT,
	}
	return operation{
		function: functions[components[1]],
		left:     components[0],
		right:    components[2],
	}
}

func isDependency(field string) bool {
	if field == "" {
		return false
	}
	_, err := strconv.ParseUint(field, 10, 16)
	if err != nil {
		return true
	}
	return false
}

func addDependencies(instr operation) []string {
	dependsOn := make([]string, 0)
	if isDependency(instr.left) {
		dependsOn = append(dependsOn, instr.left)
	}
	if isDependency(instr.right) {
		dependsOn = append(dependsOn, instr.right)
	}
	return dependsOn
}

func createInstructionOrder(dependencies map[string][]string) []string {
	instructionOrder := make([]string, 0)
	for len(dependencies) != 0 {
		removedWires := make([]string, 0)
		for wire, dependency := range dependencies {
			if len(dependency) == 0 {
				removedWires = append(removedWires, wire)
				delete(dependencies, wire)
			}
		}
		for _, removedWire := range removedWires {
			instructionOrder = append(instructionOrder, removedWire)
			for wire, dependency := range dependencies {
				index := slices.Index(dependency, removedWire)
				if index != -1 {
					dependencies[wire] = append(dependency[:index], dependency[index+1:]...)
				}
			}
		}
	}
	return instructionOrder
}

func parseContent(content string) (map[string]operation, []string) {
	lines := strings.Split(content, "\n")
	instructions := make(map[string]operation)
	dependencies := make(map[string][]string)
	for _, line := range lines {
		instructionParts := strings.Split(line, " -> ")
		op := parseOperation(instructionParts[0])
		instructions[instructionParts[1]] = op
		dependencies[instructionParts[1]] = addDependencies(op)
	}
	return instructions, createInstructionOrder(dependencies)
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

func computeSignals(instructions map[string]operation, instructionOrder []string) map[string]uint16 {
	signals := make(map[string]uint16)
	for _, wire := range instructionOrder {
		left := computeValue(instructions[wire].left, signals)
		right := computeValue(instructions[wire].right, signals)
		signals[wire] = instructions[wire].function(left, right)
	}
	return signals
}

func solve(instructions map[string]operation, instructionOrder []string) {
	firstCircuit := computeSignals(instructions, instructionOrder)
	instructions["b"] = operation{
		instructions["b"].function,
		strconv.Itoa(int(firstCircuit["a"])),
		instructions["b"].right,
	}
	secondCircuit := computeSignals(instructions, instructionOrder)
	fmt.Println("Part 1 result is:", firstCircuit["a"])
	fmt.Println("Part 2 result is:", secondCircuit["a"])
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	instructions, instructionOrder := parseContent(content)
	solve(instructions, instructionOrder)
}
