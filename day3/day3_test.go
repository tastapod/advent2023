package day3

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var sampleSchematic = strings.Split(strings.TrimSpace(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`), "\n")

func TestFindsNumbers(t *testing.T) {
	finder := PartNumberFinder{[]string{"...123...456*..."}}
	assert.Equal(t, []int{123, 456}, finder.findNumbers())
}

func TestFindsSymbol(t *testing.T) {
	assert.True(t, isSymbol('@'))
	assert.False(t, isSymbol('.'))
	assert.False(t, isSymbol('3'))
	assert.False(t, isSymbol('9'))
}

func TestFindsPartNumbers(t *testing.T) {
	finder := NewPartNumberFinder([]string{"...123...456*..."})
	assert.Equal(t, []int{456}, finder.findPartNumbers())

	// symbol before
	finder = NewPartNumberFinder([]string{
		"....@...........",
		"...123...456*...",
	})
	assert.Equal(t, []int{123, 456}, finder.findPartNumbers())

	// symbol after
	finder = NewPartNumberFinder([]string{
		"...123...456....",
		"..*.........@...",
	})
	assert.Equal(t, []int{123, 456}, finder.findPartNumbers())
}

func TestSumsPartNumbers(t *testing.T) {
	finder := NewPartNumberFinder(sampleSchematic)
	assert.Equal(t, 4361, finder.SumPartNumbers())
}

func TestFindsGears(t *testing.T) {
	finder := NewPartNumberFinder([]string{
		"...200*300...456....",
		"..*.........@.......",
	})
	assert.Equal(t, 60000, finder.SumGears())
}
