package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic("can't read file")
	}
	input := string(file)
	sum := 0
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	letters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		firstIndex := len(s)
		lIndex := -1
		for i, v := range append(digits, letters...) {
			tmpIndex := strings.Index(s, v)
			if tmpIndex > -1 && tmpIndex < firstIndex {
				firstIndex = tmpIndex
				lIndex = i
			}
		}
		firstNum := 0
		if s[firstIndex] <= '9' {
			firstNum = int(s[firstIndex]) - 48
		} else {
			firstNum = int(digits[lIndex-9][0]) - 48
		}
		lastIndex := -1
		for i, v := range append(digits, letters...) {
			tmpIndex := strings.LastIndex(s, v)
			if tmpIndex > -1 && tmpIndex > lastIndex {
				lastIndex = tmpIndex
				lIndex = i
			}
		}
		lastNum := 0
		if s[lastIndex] <= '9' {
			lastNum = int(s[lastIndex]) - 48
		} else {
			lastNum = int(digits[lIndex-9][0]) - 48
		}
		sum += firstNum*10 + lastNum
		fmt.Println(firstNum*10 + lastNum)
	}
	fmt.Println(sum)
}
