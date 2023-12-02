package day1

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindsFirstAndLastDigits(t *testing.T) {
	assert.Equal(t, 12, findEndDigits("1abc2", findInt))
}

var sampleLines = strings.Split(strings.TrimSpace(`
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`), "\n")

func TestAddsLines(t *testing.T) {
	assert.Equal(t, 142, SumEnds(sampleLines))
}

func TestFindsSpelledNumbers(t *testing.T) {
	assert.Equal(t, 83, findEndDigits("eightwothree", findIntOrName))
	assert.Equal(t, 76, findEndDigits("7pqrstsixteen", findIntOrName))
}

var sampleLines2 = strings.Split(strings.TrimSpace(`
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`), "\n")

func TestAddsLinesWithWords(t *testing.T) {
	assert.Equal(t, 281, SumEndsWithWords(sampleLines2))
}
