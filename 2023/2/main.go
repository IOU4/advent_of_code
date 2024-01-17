package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic("can't read file")
	}
	input := string(file)
	input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`
	fmt.Println(getTotal(input))

}

func getTotal(input string) int {
	sum := 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	for _, round := range strings.Split(input, "\n") {
		expression := regexp.MustCompile(`^(?:Game )(\d): (.*)`)
		subMatches := expression.FindStringSubmatch(round)
		gameId := int(rune(subMatches[1][0] - '0'))
		for _, game := range strings.Split(subMatches[2], ";") {
			for _, finding := range strings.Split(game, ",") {
				expression = regexp.MustCompile(`(\d) (red|green|blue)`)
				matches := expression.FindStringSubmatch(finding)

				switch color := matches[2]; color {
				case "red":
					number := int(rune(matches[1][0] - '0'))
					if number > 12 {
						panic("red")
					}
				}

			}
		}
		sum += gameId
	}
	return sum
}
