package day4

import (
	"strconv"
	"strings"
)

type ScratchCard struct {
	line string
}

func (c *ScratchCard) evaluate() (result int) {
	parts := strings.Split(c.line, " | ")   // ["Card n: a b c", "d e f"]
	parts2 := strings.Split(parts[0], ": ") // ["Card n", "a b c"]

	// parse winning numbers
	winningNumbers := map[int]bool{} // use map as a set
	for _, num := range strings.Fields(parts2[1]) {
		winningNumbers[toInt(num)] = true
	}

	// count winners
	for _, num := range strings.Fields(parts[1]) {
		if winningNumbers[toInt(num)] {
			if result == 0 {
				result = 1
			} else {
				result *= 2
			}
		}
	}
	return
}

func toInt(num string) int {
	value, _ := strconv.Atoi(num)
	return value
}

func SumScratchCards(scratchCards []string) (total int) {
	for _, line := range scratchCards {
		scratchCard := ScratchCard{line}
		total += scratchCard.evaluate()
	}
	return
}
