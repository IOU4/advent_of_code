package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input")
	input := string(file)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	myNumbersCount := getMyNumbersCount(lines[0])
	// part 1
	part1Answer := 0
	for _, line := range lines {
		winningNumers, myNumbers := getNumbers(line, myNumbersCount)
		points := 0
		factor := 1
		for _, myN := range myNumbers {
			for _, winN := range winningNumers {
				if myN == winN {
					if points == 0 {
						points = 1
					}
					points *= factor
					factor++
					factor = min(factor, 2)
					break
				}
			}
		}
		part1Answer += points
	}
	fmt.Println("part 1:", part1Answer)

	// part 2
	part2Answer := 0
	copies := make([]int, len(lines))
	for i, line := range lines {
		winningNumers, myNumbers := getNumbers(line, myNumbersCount)
		nextInedx := i
		copies[i]++
		for _, myN := range myNumbers {
			for _, winN := range winningNumers {
				if myN == winN {
					nextInedx++
					copies[nextInedx] += copies[i]
					break
				}
			}
		}
		part2Answer += copies[i]
	}
	fmt.Println("part 2:", part2Answer)
}

func getNumbers(line string, myNumbersCount int) (winningNumbers, myNumbers []int) {
	expr := regexp.MustCompile(`(?:Card +\d: )|(\d+)`)
	subMatches := expr.FindAllStringSubmatch(line, -1)[1:]
	for i, match := range subMatches {
		num, _ := strconv.Atoi(match[0])
		if i < myNumbersCount {
			myNumbers = append(myNumbers, num)
			continue
		}
		winningNumbers = append(winningNumbers, num)
	}
	return winningNumbers, myNumbers
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMyNumbersCount(line string) int {
	count := 0
	expr := regexp.MustCompile(`Card +\d+: ((?: *\d+)+) \|.*`)
	subMatches := expr.FindStringSubmatch(line)
	for _, v := range strings.Split(subMatches[1], " ") {
		if v != "" {
			count++
		}
	}
	return count
}
