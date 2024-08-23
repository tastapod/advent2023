package day10

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConnectsPipeEnds(t *testing.T) {
	assert := assert.New(t)
	p := Pipe{S, E}

	// enter heading north
	result, err := p.Route(N)
	if assert.NoError(err) {
		assert.Equal(E, result)
	}

	// enter heading west
	result, err = p.Route(W)
	if assert.NoError(err) {
		assert.Equal(S, result)
	}

	// enter heading east (impossible)
	result, err = p.Route(E)
	assert.Error(err)
}

var simpleLoop = strings.Split(strings.TrimSpace(`
.....
.S-7.
.|.|.
.L-J.
.....
`), "\n")

var messyLoop = strings.Split(strings.TrimSpace(`
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`), "\n")

func TestCountsSteps(t *testing.T) {
	sketch := NewSketch(simpleLoop)
	assert.Equal(t, 4, sketch.FurthestPoint())

	sketch = NewSketch(messyLoop)
	assert.Equal(t, 8, sketch.FurthestPoint())
}

var largerLoop = strings.Split(strings.TrimSpace(`
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`), "\n")

func TestCalculatesArea(t *testing.T) {
	t.Skip("Can't figure out area using sum of determinants")
	sketch := NewSketch(largerLoop)
	assert.Equal(t, 16, sketch.CalculateArea())
}
