package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var distanceFormat = regexp.MustCompile("(\\w+) to (\\w+) = (\\d+)")

func parseContent(content string) (map[string]map[string]int, error) {
	lines := strings.Split(content, "\n")
	graph := make(map[string]map[string]int)
	for _, line := range lines {
		submatchs := distanceFormat.FindStringSubmatch(line)
		departure := submatchs[1]
		destination := submatchs[2]
		distance64, err := strconv.ParseInt(submatchs[3], 10, 32)
		if err != nil {
			return nil, err
		}
		if graph[destination] == nil {
			graph[destination] = make(map[string]int)
		}
		graph[destination][departure] = int(distance64)
		if graph[departure] == nil {
			graph[departure] = make(map[string]int)
		}
		graph[departure][destination] = int(distance64)
	}
	return graph, nil
}

func rec(graph map[string]map[string]int, location string, distance int, visited []string, totalDistance int, routes map[string]int) {
	destinations := graph[location]
	visited = append(visited, location)
	totalDistance += distance
	if len(graph) == len(visited) {
		routes[strings.Join(visited, "->")] = totalDistance
		return
	}
	for destination, distanceToDestination := range destinations {
		if slices.Contains(visited, destination) {
			continue
		}
		rec(graph, destination, distanceToDestination, append(visited), totalDistance, routes)
	}
}

func solve(graph map[string]map[string]int) {
	routes := make(map[string]int)
	for starting, _ := range graph {
		rec(graph, starting, 0, make([]string, 0), 0, routes)
	}
	shortestDistance := math.MaxInt
	longestDistance := 0
	for _, distance := range routes {
		shortestDistance = min(distance, shortestDistance)
		longestDistance = max(distance, longestDistance)
	}
	fmt.Println("Part 1 result is:", shortestDistance)
	fmt.Println("Part 2 result is:", longestDistance)
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	graph, err := parseContent(content)
	if err != nil {
		panic(err)
	}
	solve(graph)
}
