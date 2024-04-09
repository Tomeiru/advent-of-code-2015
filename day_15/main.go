package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lineFormat = regexp.MustCompile("(\\w+): capacity (-?\\d+), durability (-?\\d+), flavor (-?\\d+), texture (-?\\d+), calories (-?\\d+)")

type category int

const (
	capacity category = iota
	durability
	flavor
	texture
	calories
)

type ingredient struct {
	name       string
	properties [5]int
}

func parseContent(content string) []ingredient {
	lines := strings.Split(content, "\n")
	ingredients := make([]ingredient, len(lines))
	for i, line := range lines {
		submatchs := lineFormat.FindStringSubmatch(line)
		capacity, _ := strconv.ParseInt(submatchs[2], 10, 32)
		durability, _ := strconv.ParseInt(submatchs[3], 10, 32)
		flavor, _ := strconv.ParseInt(submatchs[4], 10, 32)
		texture, _ := strconv.ParseInt(submatchs[5], 10, 32)
		calories, _ := strconv.ParseInt(submatchs[6], 10, 32)
		ingredients[i] = ingredient{
			name:       submatchs[0],
			properties: [5]int{int(capacity), int(durability), int(flavor), int(texture), int(calories)},
		}
	}
	return ingredients
}

func generateNextGeneration(nbIngredients int, precedentGeneration [][]int) [][]int {
	nextGeneration := make([][]int, 0)
	set := make(map[string]bool)
	for _, combination := range precedentGeneration {
		for i := 1; i < nbIngredients; i++ {
			nextCombination := make([]int, len(combination))
			_ = copy(nextCombination, combination)
			nextCombination[0] -= 1
			nextCombination[i] += 1
			key := fmt.Sprint(nextCombination)
			if set[key] != true {
				nextGeneration = append(nextGeneration, nextCombination)
				set[key] = true
			}
		}
	}
	return nextGeneration
}

func generatePossibleMeasures(nbIngredients int) [][]int {
	measure := make([]int, nbIngredients)
	measure[0] = 100
	possibilities := make([][]int, 1)
	possibilities[0] = measure
	precedentGeneration := make([][]int, 1)
	precedentGeneration[0] = measure
	for i := 0; i < 100; i++ {
		nextGeneration := generateNextGeneration(nbIngredients, precedentGeneration)
		possibilities = append(possibilities, nextGeneration...)
		precedentGeneration = nextGeneration
	}
	return possibilities
}

func calculateCategoryTotal(ingredients []ingredient, measures []int, category category) int {
	total := 0
	for i, ingr := range ingredients {
		total += ingr.properties[category] * measures[i]
	}
	if total <= 0 {
		return 0
	}
	return total
}

func calculateScore(ingredients []ingredient, measures []int) int {
	totals := [4]int{0, 0, 0, 0}
	for cat := capacity; cat < calories; cat++ {
		totals[cat] = calculateCategoryTotal(ingredients, measures, cat)
	}
	return totals[0] * totals[1] * totals[2] * totals[3]
}

func solve(ingredients []ingredient, possibilities [][]int) {
	bestCookie := 0
	best500kcalCookie := 0
	for _, possibility := range possibilities {
		kcals := calculateCategoryTotal(ingredients, possibility, calories)
		bestCookie = max(bestCookie, calculateScore(ingredients, possibility))
		if kcals == 500 {
			best500kcalCookie = max(best500kcalCookie, calculateScore(ingredients, possibility))
		}
	}
	fmt.Println("Result of part 1:", bestCookie)
	fmt.Println("Result of part 2:", best500kcalCookie)
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	ingredients := parseContent(content)
	possibilities := generatePossibleMeasures(len(ingredients))
	solve(ingredients, possibilities)
}
