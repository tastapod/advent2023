package day1

import (
	"strings"
)

const DIGITS = "0123456789"

type FindInt = func(string) int

var DigitWords = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func findEndDigits(input string, findInt FindInt) int {
	var first, last = -1, -1
	for i := 0; i < len(input); i++ {
		if i := findInt(input[i:]); i != -1 {
			first = i
			break
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if i := findInt(input[i:]); i != -1 {
			last = i
			break
		}
	}
	return first*10 + last
}

func findInt(input string) int {
	return strings.Index(DIGITS, string(input[0]))
}

func findIntOrName(input string) int {
	if result := findInt(input); result != -1 {
		return result
	}
	for pos, word := range DigitWords {
		if strings.HasPrefix(input, word) {
			return pos
		}
	}
	return -1
}

func sumEnds(lines []string, findInt FindInt) int {
	var total = 0
	for _, line := range lines {
		total += findEndDigits(line, findInt)
	}
	return total
}

func SumEnds(lines []string) int {
	return sumEnds(lines, findInt)
}

func SumEndsWithWords(lines []string) int {
	return sumEnds(lines, findIntOrName)
}
