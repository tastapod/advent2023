package day7

import (
	"github.com/tastapod/advent2023/convert"
	"slices"
	"strings"
)

const (
	HighCard = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand interface {
	Cards() string
	Bid() int
	CardValue(i int) int
	Type() int
}

type CommonHand struct {
	cards string
	bid   int
}

func (h CommonHand) Cards() string {
	return h.cards
}

func (h CommonHand) Bid() int {
	return h.bid
}

func (h CommonHand) CardValue(i int) int {
	card := h.cards[i]

	switch card {
	case '2', '3', '4', '5', '6', '7', '8', '9':
		return convert.ToInt(string(card))
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		panic("Unknown card: " + string(card))
	}
}

func (h CommonHand) Type() int {
	// count cards
	cardCounts := map[rune]int{}
	for _, card := range h.cards {
		cardCounts[card]++
	}

	// count counts
	countCounts := make([]int, 6)
	for _, count := range cardCounts {
		countCounts[count]++
	}

	// figure out type
	switch {
	case countCounts[5] == 1:
		return FiveOfAKind
	case countCounts[4] == 1:
		return FourOfAKind
	case countCounts[3] == 1:
		if countCounts[2] == 1 {
			return FullHouse
		}
		return ThreeOfAKind
	case countCounts[2] == 2:
		return TwoPairs
	case countCounts[2] == 1:
		return OnePair
	default:
		return HighCard
	}
}

func NewCommonHand(line string) CommonHand {
	parts := strings.Fields(line)

	var bid int
	if len(parts) == 1 {
		bid = 0
	} else {
		bid = convert.ToInt(parts[1])
	}

	return CommonHand{
		cards: parts[0],
		bid:   bid,
	}
}

type SimpleHand struct {
	CommonHand
}

func NewSimpleHand(line string) Hand {
	return SimpleHand{NewCommonHand(line)}
}

type Round struct {
	Hands []Hand
}

// Ranked ranks hands
func (r Round) Ranked() (ranked Round) {
	ranked = r
	slices.SortFunc(ranked.Hands, func(a, b Hand) int {
		if diff := a.Type() - b.Type(); diff != 0 {
			return diff
		}
		for i := 0; i < len(a.Cards()); i++ {
			if diff := a.CardValue(i) - b.CardValue(i); diff != 0 {
				return diff
			}
		}
		return 0
	})
	return
}

func (r Round) Score() (total int) {
	ranked := r.Ranked()
	for i, hand := range ranked.Hands {
		total += hand.Bid() * (i + 1)
	}
	return
}

func newRound(lines []string, newHand func(string) Hand) (r Round) {
	r = Round{
		Hands: make([]Hand, len(lines)),
	}
	for i, line := range lines {
		r.Hands[i] = newHand(line)
	}
	return
}

func NewSimpleRound(lines []string) Round {
	return newRound(lines, NewSimpleHand)
}

type JokerHand struct {
	CommonHand
}

func (h JokerHand) CardValue(i int) int {
	card := h.cards[i]
	if card == 'J' {
		return 1 // lowest value
	}
	return h.CommonHand.CardValue(i)
}

func (h JokerHand) Type() int {
	// strip all the jokers
	cards := strings.Replace(h.cards, "J", "", -1)

	numJokers := 5 - len(cards)

	// count cards
	cardCounts := map[rune]int{}
	for _, card := range cards {
		cardCounts[card]++
	}

	// count counts
	countCounts := make([]int, 6)
	for _, count := range cardCounts {
		countCounts[count]++
	}

	// figure out type
	switch {
	case countCounts[5] == 1:
		return FiveOfAKind
	case countCounts[4] == 1:
		switch numJokers {
		case 1:
			return FiveOfAKind
		default:
			return FourOfAKind
		}
	case countCounts[3] == 1:
		switch numJokers {
		case 2:
			return FiveOfAKind
		case 1:
			return FourOfAKind
		default:
			if countCounts[2] == 1 {
				return FullHouse
			}
			return ThreeOfAKind
		}
	case countCounts[2] == 2:
		switch numJokers {
		case 1:
			return FullHouse
		default:
			return TwoPairs
		}
	case countCounts[2] == 1:
		switch numJokers {
		case 3:
			return FiveOfAKind
		case 2:
			return FourOfAKind
		case 1:
			return ThreeOfAKind
		default:
			return OnePair
		}
	default:
		switch numJokers {
		case 5, 4:
			return FiveOfAKind
		case 3:
			return FourOfAKind
		case 2:
			return ThreeOfAKind
		case 1:
			return OnePair
		default:
			return HighCard
		}
	}
}

func NewJokerHand(line string) Hand {
	return JokerHand{NewCommonHand(line)}
}

func NewJokerRound(lines []string) Round {
	return newRound(lines, NewJokerHand)
}
