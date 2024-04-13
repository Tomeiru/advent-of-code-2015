package main

import (
	"fmt"
	"maps"
	"os"
	"strings"
)

type molecule string

func parseReplacements(lines []string) ([]molecule, map[molecule][]molecule) {
	replacements := make(map[molecule][]molecule)
	bases := make([]molecule, 0)
	for _, line := range lines[0 : len(lines)-2] {
		molecules := strings.Split(line, " => ")
		base := molecule(molecules[0])
		replacement := molecule(molecules[1])
		if replacements[base] == nil {
			replacements[base] = make([]molecule, 0)
			bases = append(bases, base)
		}
		replacements[base] = append(replacements[base], replacement)
	}
	return bases, replacements
}

func parseMedicineMolecule(line string, bases []molecule) []molecule {
	medicine := make([]molecule, 0)
	for i := 0; i < len(line); {
		fmt.Println(i, line[i:])
		for _, base := range bases {
			if strings.HasPrefix(line[i:], string(base)) {
				medicine = append(medicine, base)
				i += len(base)
				break
			}
		}
	}
	return medicine
}

func parseContent(content string) (string, map[molecule][]molecule) {
	lines := strings.Split(content, "\n")
	_, replacements := parseReplacements(lines)
	//medicine := parseMedicineMolecule(lines[len(lines)-1], bases)
	return lines[len(lines)-1], replacements
}

func generateReplacementPossibilities(medicine string, basesToReplacements map[molecule][]molecule) map[string]bool {
	possibilities := make(map[string]bool)
	for i := 0; i < len(medicine); i++ {
		for base, replacements := range basesToReplacements {
			if strings.HasPrefix(medicine[i:], string(base)) {
				for _, replacement := range replacements {
					before := medicine[0:i]
					after := medicine[i+len(base):]
					possibilities[before+string(replacement)+after] = true
				}
				i += len(base) - 1
				break
			}
		}
	}
	return possibilities
}

func solve(medicine string, basesToReplacements map[molecule][]molecule) {
	generated := generateReplacementPossibilities(medicine, basesToReplacements)
	step := 0
	possibilities := make(map[string]bool)
	possibilities["e"] = true
	for possibilities[medicine] != true {
		newPossibilities := make(map[string]bool)
		for possibility, _ := range possibilities {
			maps.Copy(newPossibilities, generateReplacementPossibilities(possibility, basesToReplacements))
		}
		possibilities = newPossibilities
		step += 1
	}
	fmt.Println("Result of part 1:", len(generated))
	fmt.Println("Result of part 2:", step)
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	medicine, replacements := parseContent(content)
	solve(medicine, replacements)
}
