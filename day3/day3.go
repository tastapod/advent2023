package day3

import (
	"regexp"
	"strconv"
	"strings"
)

var NumberRE = regexp.MustCompile(`\d+`)

type PartNumberFinder struct {
	schematic []string
}

func NewPartNumberFinder(schematic []string) *PartNumberFinder {
	paddedWidth := len(schematic[0]) + 2
	paddedHeight := len(schematic) + 2
	padded := make([]string, paddedHeight)

	// empty top and bottom lines
	padded[0] = strings.Repeat(".", paddedWidth)
	padded[paddedHeight-1] = padded[0]

	// empty first and last column
	for i, line := range schematic {
		padded[i+1] = "." + line + "."
	}
	return &PartNumberFinder{padded}
}

func (f *PartNumberFinder) findNumbers() (result []int) {
	for _, line := range f.schematic {
		idxs := NumberRE.FindAllIndex([]byte(line), -1)
		for _, idx := range idxs {
			value, _ := strconv.Atoi(line[idx[0]:idx[1]])
			result = append(result, value)
		}
	}
	return
}

func (f *PartNumberFinder) findPartNumbers() (result []int) {
	for lineNo, line := range f.schematic {
		idxs := NumberRE.FindAllIndex([]byte(line), -1)
		for _, idx := range idxs {
			if f.hasAdjacentSymbol(lineNo, idx) {
				value, _ := strconv.Atoi(line[idx[0]:idx[1]])
				result = append(result, value)
			}
		}
	}
	return
}

func (f *PartNumberFinder) hasAdjacentSymbol(lineNo int, idx []int) bool {
	start, end := idx[0], idx[1]
	line := f.schematic[lineNo]
	// check ends
	if isSymbol(line[start-1]) || isSymbol(line[end]) {
		return true
	}
	if containsSymbol(idx, f.schematic[lineNo-1]) {
		return true
	}
	if containsSymbol(idx, f.schematic[lineNo+1]) {
		return true
	}
	return false
}

func (f *PartNumberFinder) SumPartNumbers() (total int) {
	for _, num := range f.findPartNumbers() {
		total += num
	}
	return
}

func containsSymbol(idx []int, line string) bool {
	for i := idx[0] - 1; i <= idx[1]; i++ {
		if isSymbol(line[i]) {
			return true
		}
	}
	return false
}

func isSymbol(ch byte) bool {
	return ch != '.' && !strings.Contains("0123456789", string(rune(ch)))
}
