package day7

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var roundInput = strings.Split(strings.TrimSpace(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`), "\n")

func TestFindsHandType(t *testing.T) {
	checkHandType := func(expected int, cards string) {
		assert.Equal(t, expected, NewSimpleHand(cards).Type())
	}
	checkHandType(HighCard, "32T8K")
	checkHandType(OnePair, "32T3K")
	checkHandType(TwoPairs, "32T3T")
	checkHandType(ThreeOfAKind, "T2T3T")
	checkHandType(FullHouse, "T2T2T")
	checkHandType(FourOfAKind, "T2222")
	checkHandType(FiveOfAKind, "22222")
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
	assert := assert.New(t)

	hand := JokerHand{CommonHand{"89TJK", 0}}
	assert.Equal(8, hand.CardValue(0))
	assert.Equal(9, hand.CardValue(1))
	assert.Equal(10, hand.CardValue(2))
	assert.Equal(1, hand.CardValue(3))
	assert.Equal(13, hand.CardValue(4))
}

func checkJokerHandType(t *testing.T, expected int, cards string) {
	assert.Equal(t, expected, NewJokerHand(cards).Type())
}
