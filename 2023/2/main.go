package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic("can't read file")
	}
	input := string(file)
	// 	input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	fmt.Println(getTotal(input))
}

func getTotal(input string) int {
	sum := 0
	for _, round := range strings.Split(strings.TrimSpace(input), "\n") {
		sum += dealWithRound(round)
	}
	return sum
}

func dealWithRound(round string) int {
	expression := regexp.MustCompile(`^(?:Game \d+): (.*)`)
	subMatches := expression.FindStringSubmatch(round)
	maxRed, maxGreen, maxBlue := 0, 0, 0
	for _, game := range strings.Split(subMatches[1], ";") {
		for _, finding := range strings.Split(game, ",") {
			expression = regexp.MustCompile(`(\d+) (red|green|blue)`)
			matches := expression.FindStringSubmatch(finding)
			number, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0
			}
			switch color := matches[2]; color {
			case "red":
				if number > maxRed {
					maxRed = number
				}
			case "green":
				if number > maxGreen {
					maxGreen = number
				}
			case "blue":
				if number > maxBlue {
					maxBlue = number
				}
			}
		}
	}
	return maxBlue * maxGreen * maxRed
}
