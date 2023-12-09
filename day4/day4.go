package day4

import (
	"github.com/tastapod/advent2023/convert"
	"math"
	"strings"
)

type ScratchCard struct {
	number         int
	winningNumbers map[int]bool // use map as a set
	numWinners     int
}

func NewScratchCard(line string) (s *ScratchCard) {
	parts := strings.Split(line, " | ")     // ["Card n: a b s", "d e f"]
	parts2 := strings.Split(parts[0], ": ") // ["Card n", "a b s"]
	s = &ScratchCard{winningNumbers: make(map[int]bool)}

	// parse card number
	s.number = convert.ToInt(strings.Fields(parts2[0])[1])

	// parse winning numbers
	for _, num := range strings.Fields(parts2[1]) {
		s.winningNumbers[convert.ToInt(num)] = true
	}

	// count winners
	for _, num := range strings.Fields(parts[1]) {
		if s.winningNumbers[convert.ToInt(num)] {
			s.numWinners++
		}
	}
	return
}

func (s *ScratchCard) score() int {
	return int(math.Pow(2, float64(s.numWinners-1)))
}

func SumScratchCards(scratchCards []string) (total int) {
	for _, line := range scratchCards {
		scratchCard := NewScratchCard(line)
		total += scratchCard.score()
	}
	return
}

func MoarScratchCards(lines []string) (result int) {
	counts := make([]int, len(lines)+1) // use card number as index
	for i := range counts {
		counts[i] = 1 // we have at least one of each card
	}
	for _, line := range lines {
		card := NewScratchCard(line)

		// count additional cards
		for i := 1; i <= card.numWinners; i++ {
			counts[card.number+i] += counts[card.number]
		}

		// tally cards so far
		result += counts[card.number]
	}
	return
}
