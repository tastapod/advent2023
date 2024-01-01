package day10

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConnectsPipeEnds(t *testing.T) {
	p := Pipe{S, E} // F
	assert.Equal(t, byte(E), p.TravelFrom(N))
	assert.Equal(t, byte(S), p.TravelFrom(W))
	assert.Panics(t, func() { p.TravelFrom(E) })
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
