package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type present struct {
	length, width, height int
}

func (p *present) getSmallestSideDimension() [2]int {
	if p.length >= p.height && p.length >= p.width {
		return [2]int{p.height, p.width}
	}
	if p.width >= p.height && p.width >= p.length {
		return [2]int{p.height, p.length}
	}
	return [2]int{p.length, p.width}
}

func (p *present) calculateSmallestSideArea() int {
	dimensions := p.getSmallestSideDimension()
	return dimensions[0] * dimensions[1]
}

func (p *present) calculateNeededWrappingPaper() int {
	return 2*p.length*p.width + 2*p.width*p.height + 2*p.height*p.length + p.calculateSmallestSideArea()
}

func (p *present) calculateSmallestSidePerimeter() int {
	dimensions := p.getSmallestSideDimension()
	return dimensions[0]*2 + dimensions[1]*2
}

func (p *present) calculateCubicFeet() int {
	return p.length * p.height * p.width
}

func (p *present) calculateNeededRibbon() int {
	return p.calculateSmallestSidePerimeter() + p.calculateCubicFeet()
}

func solve(presents []present) {
	neededWrappingPaper := 0
	neededRibbon := 0
	for _, p := range presents {
		neededWrappingPaper += p.calculateNeededWrappingPaper()
		neededRibbon += p.calculateNeededRibbon()
	}
	fmt.Println("Part 1 result is:", neededWrappingPaper)
	fmt.Println("Part 2 result is:", neededRibbon)
}

func parseContent(content string) []present {
	lines := strings.Split(content, "\n")
	presents := make([]present, len(lines))
	for i := 0; i < len(lines); i++ {
		dimensions := strings.Split(lines[i], "x")
		convertedDimensions := make([]int, len(dimensions))
		for ii := 0; ii < len(dimensions); ii++ {
			value, err := strconv.Atoi(dimensions[ii])
			if err != nil {
				panic("should never happen")
			}
			convertedDimensions[ii] = value
		}
		presents[i] = present{
			length: convertedDimensions[0],
			width:  convertedDimensions[1],
			height: convertedDimensions[2],
		}
	}
	return presents
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	presents := parseContent(content)
	solve(presents)
}
