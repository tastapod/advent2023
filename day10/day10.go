package day10

import (
	"fmt"
	"strings"
)

const (
	N = 'N'
	S = 'S'
	W = 'W'
	E = 'E'
)

type Pipe struct {
	End1, End2 byte
}

// TravelFrom maps the entry point to a corresponding exit point.
//
// Travelling e.g. S means you enter the pipe from the N,
// so the Pipe switches the direction so that the travel makes sense.
func (c Pipe) TravelFrom(direction byte) byte {
	var end byte
	switch direction {
	case N:
		end = S
	case S:
		end = N
	case W:
		end = E
	case E:
		end = W
	}

	switch end {
	case c.End1:
		return c.End2
	case c.End2:
		return c.End1
	default:
		panic(fmt.Sprintf("Unknown end %d for connector %v", direction, c))
	}
}

var Pipes = map[byte]Pipe{
	'|': {N, S},
	'-': {W, E},
	'L': {N, E},
	'J': {N, W},
	'7': {S, W},
	'F': {S, E},
}

type Sketch struct {
	Grid [][]byte
}

func (s *Sketch) FurthestPoint() int {
	row, col := s.StartPos()
	numSteps := 0

	direction := s.FirstMove(row, col)

	for row, col = Move(row, col, direction); s.Grid[row][col] != 'S'; row, col = Move(row, col, direction) {
		direction = Pipes[s.Grid[row][col]].TravelFrom(direction)
		numSteps++
	}
	return (numSteps + 1) / 2
}

func (s *Sketch) FirstMove(row, col int) byte {

	for _, direction := range []byte{N, S, W, E} {
		if r, c := Move(row, col, direction); s.Grid[r][c] != '.' {
			return direction
		}
	}
	panic("No pipes found around start position")
}

func Move(row, col int, direction byte) (int, int) {
	switch direction {
	case N:
		return row - 1, col
	case S:
		return row + 1, col
	case W:
		return row, col - 1
	case E:
		return row, col + 1
	default:
		panic("Unknown direction: " + string(direction))
	}
}

func (s *Sketch) StartPos() (int, int) {
	for r, row := range s.Grid {
		if c := strings.Index(string(row), "S"); c != -1 {
			return r, c
		}
	}
	panic("No start found")
}

func NewSketch(input []string) (s Sketch) {
	numRows := len(input)
	numCols := len(input[0])
	s = Sketch{
		Grid: make([][]byte, numRows+2),
	}

	// top and bottom
	border := make([]byte, numCols+2)
	for i := 0; i < len(border); i++ {
		border[i] = '.'
	}

	// build the grid
	s.Grid[0] = border
	s.Grid[len(s.Grid)-1] = border
	for i, line := range input {
		gridLine := make([]byte, numCols+2)
		gridLine[0] = '.'
		gridLine[numCols+1] = '.'
		copy(gridLine[1:], line)
		s.Grid[i+1] = gridLine
	}
	return
}
