package day7

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var roundInput = strings.Split(strings.TrimSpace(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`), "\n")

func TestFindsHandType(t *testing.T) {
	checkHandType(t, HighCard, "32T8K")
	checkHandType(t, OnePair, "32T3K")
	checkHandType(t, TwoPairs, "32T3T")
	checkHandType(t, ThreeOfAKind, "T2T3T")
	checkHandType(t, FullHouse, "T2T2T")
	checkHandType(t, FourOfAKind, "T2222")
	checkHandType(t, FiveOfAKind, "22222")
}

func checkHandType(t *testing.T, expected int, cards string) {
	assert.Equal(t, expected, NewSimpleHand(cards).Type())
}

func TestParsesRound(t *testing.T) {
	round := NewSimpleRound(roundInput)
	assert.Equal(t, NewSimpleHand("KK677 28"), round.Hands[2])
}

func TestRanksHands(t *testing.T) {
	ranked := NewSimpleRound(roundInput).Ranked()
	assert.Equal(t, "32T3K", ranked.Hands[0].Cards())
	assert.Equal(t, "KK677", ranked.Hands[2].Cards())
}

func TestScoresRound(t *testing.T) {
	score := NewSimpleRound(roundInput).Score()
	assert.Equal(t, 6440, score)
}

func TestFindsHandTypeWithJokers(t *testing.T) {
	checkJokerHandType(t, FourOfAKind, "T55J5")
}

func TestScoresCards(t *testing.T) {
	hand := JokerHand{CommonHand{"89TJK", 0}}
	assert.Equal(t, 8, hand.CardValue(0))
	assert.Equal(t, 9, hand.CardValue(1))
	assert.Equal(t, 10, hand.CardValue(2))
	assert.Equal(t, 1, hand.CardValue(3))
	assert.Equal(t, 13, hand.CardValue(4))
}

func checkJokerHandType(t *testing.T, expected int, cards string) {
	assert.Equal(t, expected, NewJokerHand(cards).Type())
}
