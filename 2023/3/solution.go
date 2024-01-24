package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumberStruct struct {
	line  int
	start int
	end   int
	value int
}

type SymbolStruct struct {
	line int
	pos  int
}

func main() {
	file, _ := os.ReadFile("input")
	input := string(file)
	numbers, symbols, gears := getNumbersSymbolsGears(input)
	// part 1
	part1Answer := 0
	partNumbers := []NumberStruct{}
	for _, number := range numbers {
		if isPartNumber(number, symbols) {
			part1Answer += number.value
			partNumbers = append(partNumbers, number)
		}
	}
	fmt.Println("part 1 Answer:", part1Answer)

	// part 2
	part2Answer := 0
	for _, gear := range gears {
		gearRatio := 0
		parts := 0
		for _, number := range partNumbers {
			if (gear.line == number.line-1 || gear.line == number.line || gear.line == number.line+1) && isAdjacent(gear.pos, number.start, number.end) {
				parts++
				if parts == 1 {
					gearRatio = number.value
				}
				if parts == 2 {
					gearRatio *= number.value
				}
				if parts > 2 {
					break
				}
			}
		}
		if parts == 2 {
			part2Answer += gearRatio
		}
	}

	fmt.Println("part 2 answer:", part2Answer)
}

func getNumbersSymbolsGears(input string) (numbers []NumberStruct, symbols []SymbolStruct, gears []SymbolStruct) {
	for lineIndex, line := range strings.Split(input, "\n") {
		numberString := ""
		start := -1
		for i, v := range line {
			if v >= '0' && v <= '9' {
				if start == -1 {
					start = i
				}
				numberString += string(v)
			} else {
				if start != -1 {
					convertedNumber, _ := strconv.Atoi(numberString)
					numbers = append(numbers, NumberStruct{lineIndex, start, i - 1, convertedNumber})
				}
				start = -1
				numberString = ""
				if v != '.' {
					symbols = append(symbols, SymbolStruct{lineIndex, i})
				}
				if v == '*' {
					gears = append(gears, SymbolStruct{lineIndex, i})
				}
			}

			if start != -1 && i == len(line)-1 {
				convertedNumber, _ := strconv.Atoi(numberString)
				numbers = append(numbers, NumberStruct{lineIndex, start, i - 1, convertedNumber})
			}
		}
	}
	return numbers, symbols, gears
}

func isPartNumber(number NumberStruct, symbols []SymbolStruct) bool {
	for _, symbol := range symbols {
		if (symbol.line == number.line+1 || symbol.line == number.line-1 || symbol.line == number.line) && isAdjacent(symbol.pos, number.start, number.end) {
			return true
		}
	}
	return false
}

func isAdjacent(pos, start, end int) bool {
	return pos >= start-1 && pos <= end+1
}
