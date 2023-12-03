package day2

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParsesGameLines(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	}

	possible := checkGames(lines, GameLimits)
	assert.Equal(t, []int{1}, possible)
}

var sampleLines = strings.Split(strings.TrimSpace(`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`), "\n")

func TestAddsIdsOfPossibleGames(t *testing.T) {
	assert.Equal(t, 8, SumPossibleIds(sampleLines, GameLimits))
}

// Power of game is product of fewest number of dice to play the game
func TestCalculatesPowerOfGame(t *testing.T) {
	assert.Equal(t, 48, powerOfGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
}

func TestAddsPowersOfGames(t *testing.T) {
	assert.Equal(t, 2286, SumPowersOfGames(sampleLines))
}
