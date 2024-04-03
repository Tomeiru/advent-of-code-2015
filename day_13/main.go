package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var lineFormat = regexp.MustCompile("(\\w+) would (gain|lose) (\\d+) happiness units by sitting next to (\\w+)\\.")

func calculateChangeInHappiness(graph map[string]map[string]int, arrangement []string) int {
	change := 0
	for i, person := range arrangement {
		previous := i - 1
		if previous == -1 {
			previous = len(arrangement) - 1
		}
		change += graph[person][arrangement[previous]]
		change += graph[person][arrangement[(i+1)%len(arrangement)]]
	}
	return change
}

func fillAllArrangements(names []string, graph map[string]map[string]int, arrangements map[string]int, current []string) {
	if len(current) == len(graph) {
		arrangements[strings.Join(current, ",")] = calculateChangeInHappiness(graph, current)
		return
	}
	for _, name := range names {
		if slices.Contains(current, name) {
			continue
		}
		next := make([]string, len(current)+1)
		_ = copy(next, current)
		next[len(current)] = name
		fillAllArrangements(names, graph, arrangements, next)
	}
}

func parseContent(content string) map[string]map[string]int {
	lines := strings.Split(content, "\n")
	graph := make(map[string]map[string]int)
	for _, line := range lines {
		submatchs := lineFormat.FindStringSubmatch(line)
		affected := submatchs[1]
		if graph[affected] == nil {
			graph[affected] = make(map[string]int)
		}
		causing := submatchs[4]
		positive := submatchs[2] == "gain"
		amount, _ := strconv.ParseInt(submatchs[3], 10, 32)
		if !positive {
			amount *= -1
		}
		graph[affected][causing] = int(amount)
	}
	return graph
}

func generateArrangements(graph map[string]map[string]int) map[string]int {
	names := make([]string, len(graph))
	i := 0
	for name, _ := range graph {
		names[i] = name
		i++
	}
	arrangements := make(map[string]int)
	fillAllArrangements(names, graph, arrangements, make([]string, 0))
	return arrangements
}

func findOptimalHappinessChange(arrangements map[string]int) int {
	answer := 0
	for _, value := range arrangements {
		answer = max(answer, value)
	}
	return answer
}

func solve(graph map[string]map[string]int) {
	arrangements := generateArrangements(graph)
	mine := make(map[string]int)
	for name, links := range graph {
		mine[name] = 0
		links["Mathieu"] = 0
	}
	graph["Mathieu"] = mine
	arrangementsIncludingMe := generateArrangements(graph)
	fmt.Println("Result of part 1:", findOptimalHappinessChange(arrangements))
	fmt.Println("Result of part 2:", findOptimalHappinessChange(arrangementsIncludingMe))
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	graph := parseContent(content)
	solve(graph)
}
