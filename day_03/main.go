package main

import (
	"fmt"
	"maps"
	"os"
)

type position struct {
	x, y int
}

func (p *position) moveDown() {
	p.y += 1
}

func (p *position) moveRight() {
	p.x += 1
}

func (p *position) moveUp() {
	p.y -= 1
}

func (p *position) moveLeft() {
	p.x -= 1
}

func (p *position) generateString() string {
	return fmt.Sprintf("%d;%d", p.x, p.y)
}

func distributePresents(instruction string) map[string]int {
	initialPosition := position{x: 0, y: 0}
	encounteredPositions := make(map[string]int)
	for _, direction := range instruction {
		encounteredPositions[initialPosition.generateString()] += 1
		switch direction {
		case 'v':
			initialPosition.moveDown()
		case '<':
			initialPosition.moveRight()
		case '>':
			initialPosition.moveLeft()
		case '^':
			initialPosition.moveUp()
		default:
		}
	}
	encounteredPositions[initialPosition.generateString()] += 1
	return encounteredPositions
}

func solve(content string) {
	ogSantaDelivery := distributePresents(content)
	fmt.Println("Part 1 result is:", len(ogSantaDelivery))
	newSantaInstruction := ""
	robotSantaInstruction := ""
	for i, direction := range content {
		if i%2 == 0 {
			newSantaInstruction += string(direction)
		} else {
			robotSantaInstruction += string(direction)
		}
	}
	newSantaDelivery := distributePresents(newSantaInstruction)
	robotSantaDelivery := distributePresents(robotSantaInstruction)
	maps.Copy(newSantaDelivery, robotSantaDelivery)
	fmt.Println("Part 2 result is:", len(newSantaDelivery))
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	solve(content)
}
