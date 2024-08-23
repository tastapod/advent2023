package day9

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindsNextNumber(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(18, NewPredictor("0 3 6 9 12 15").NextValue())
	assert.Equal(28, NewPredictor("1 3 6 10 15 21").NextValue())
	assert.Equal(68, NewPredictor("10 13 16 21 30 45").NextValue())
}

var sequences = strings.Split(strings.TrimSpace(`
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`), "\n")

func TestSumsNextValues(t *testing.T) {
	assert.Equal(t, 114, SumNextValues(sequences))
}

func TestFindsPreviousValue(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(-3, NewPredictor("0 3 6 9 12 15").PreviousValue())
	assert.Equal(0, NewPredictor("1 3 6 10 15 21").PreviousValue())
	assert.Equal(5, NewPredictor("10 13 16 21 30 45").PreviousValue())
}

func TestSumsPreviousValues(t *testing.T) {
	assert.Equal(t, 2, SumPreviousValues(sequences))
}
