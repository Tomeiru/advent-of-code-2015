package day_15

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lineFormat = regexp.MustCompile("(\\w+): capacity (-?\\d+), durability (-?\\d+), flavor (-?\\d+), texture (-?\\d+), calories (-?\\d+)")

type ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
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
			capacity:   int(capacity),
			flavor:     int(flavor),
			texture:    int(texture),
			calories:   int(calories),
			durability: int(durability),
		}
	}
	return ingredients
}

func generateNextGeneration(nbIngredients int, precedentGeneration [][]int) [][]int {
	nextGeneration := make([][]int, 0)
	for i, combination := range precedentGeneration {
		for ing := 1; ing < nbIngredients; i++ {
		}
	}
	return nextGeneration
}

func generatePossibleMeasures(nbIngredients int) {
	measure := make([]int, nbIngredients)
	measure[0] = 100
	possibilities := make([][]int, 1)
	possibilities[0] = measure
	precedentGeneration := make([][]int, 1)
	precedentGeneration[0] = measure
	for i := 0; i < 100; i++ {
		nextGeneration := generateNextGeneration(precedentGeneration)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	ingredients := parseContent(content)
}
