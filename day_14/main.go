package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lineFormat = regexp.MustCompile("(\\w+) can fly (\\d+) km\\/s for (\\d+) seconds, but then must rest for (\\d+) seconds\\.")

type reindeer struct {
	name                                string
	speed, flightDuration, restDuration int
}

func (r *reindeer) calculateDistanceTraveledAtSecond(second int) int {
	cycleDuration := r.flightDuration + r.restDuration
	cycleSpeed := r.flightDuration * r.speed
	cycles := second / cycleDuration
	remaining := second % cycleDuration
	if remaining > r.flightDuration {
		return cycleSpeed * (cycles + 1)
	}
	return cycleSpeed*cycles + remaining*r.speed
}

func parseContent(content string) []reindeer {
	lines := strings.Split(content, "\n")
	reindeers := make([]reindeer, len(lines))
	for i, line := range lines {
		submatchs := lineFormat.FindStringSubmatch(line)
		speed, _ := strconv.ParseInt(submatchs[2], 10, 32)
		flightDuration, _ := strconv.ParseInt(submatchs[3], 10, 32)
		restDuration, _ := strconv.ParseInt(submatchs[4], 10, 32)
		reindeers[i] = reindeer{
			name:           submatchs[1],
			speed:          int(speed),
			flightDuration: int(flightDuration),
			restDuration:   int(restDuration),
		}
	}
	return reindeers
}

func getLeaderNameAndDistanceAtSecond(reindeers []reindeer, second int) (string, int) {
	leader := ""
	distance := 0
	for _, r := range reindeers {
		current := r.calculateDistanceTraveledAtSecond(second)
		if current > distance {
			distance = current
			leader = r.name
		}
	}
	return leader, distance

}

func solve(reindeers []reindeer) {
	raceDuration := 2503
	_, firstWinnerDistance := getLeaderNameAndDistanceAtSecond(reindeers, raceDuration)
	reindeerToPoints := make(map[string]int, len(reindeers))
	for second := 1; second <= raceDuration; second++ {
		leaderName, _ := getLeaderNameAndDistanceAtSecond(reindeers, second)
		reindeerToPoints[leaderName] += 1
	}
	secondWinnerPoints := 0
	for _, points := range reindeerToPoints {
		secondWinnerPoints = max(secondWinnerPoints, points)
	}
	fmt.Println("Result of part 1:", firstWinnerDistance)
	fmt.Println("Result of part 2:", secondWinnerPoints)
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(dat)
	reindeers := parseContent(content)
	solve(reindeers)
}
