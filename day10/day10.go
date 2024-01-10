package day10

import (
	"fmt"
	"strings"
)

type Direction rune

func (d Direction) String() string {
	return string(d)
}

const (
	N Direction = 'N'
	S Direction = 'S'
	W Direction = 'W'
	E Direction = 'E'
)

type Pipe struct {
	End1, End2 Direction
}

func (p Pipe) String() string {
	return fmt.Sprintf("%s-%s", string(p.End1), string(p.End2))
}

// Route maps an entry point to a corresponding exit point.
//
// Travelling e.g. south means you enter the pipe from N.
func (p Pipe) Route(heading Direction) (result Direction, err error) {
	// map heading to pipe end
	var pipeEnd Direction
	switch heading {
	case N:
		pipeEnd = S
	case S:
		pipeEnd = N
	case W:
		pipeEnd = E
	case E:
		pipeEnd = W
	}
	switch pipeEnd {
	case p.End1:
		result = p.End2
	case p.End2:
		result = p.End1
	default:
		err = fmt.Errorf("unknown source %s for connector %s", pipeEnd, p)
	}
	return
}

// Connects checks whether a Pipe connects with something heading in this direction
func (p Pipe) Connects(heading Direction) bool {
	_, err := p.Route(heading)
	return err == nil
}

type PipeMap map[rune]Pipe

var Pipes = PipeMap{
	'|': {N, S},
	'-': {W, E},
	'L': {N, E},
	'J': {N, W},
	'7': {S, W},
	'F': {S, E},
}

type Point struct {
	X, Y int
}

func (p Point) Move(direction Direction) Point {
	switch direction {
	case N:
		return Point{p.X, p.Y - 1}
	case S:
		return Point{p.X, p.Y + 1}
	case W:
		return Point{p.X - 1, p.Y}
	case E:
		return Point{p.X + 1, p.Y}
	default:
		panic("Unknown direction: " + string(direction))
	}
}

type Sketch struct {
	Grid [][]rune
}

func (s *Sketch) FurthestPoint() int {
	return (len(s.MapPoints()) + 1) / 2
}

type State struct {
	Point
	Tile    rune
	Heading Direction
}

func (s *State) String() string {
	return fmt.Sprintf("(%d,%d): %s heading %s", s.X, s.Y, string(s.Tile), string(s.Heading))
}

func (s *State) Route() (result Direction, err error) {
	if pipe, found := Pipes[s.Tile]; found {
		result, err = pipe.Route(s.Heading)
	} else {
		err = fmt.Errorf("unknown tile: %s", string(s.Tile))
	}
	return
}

func (s *Sketch) MapPoints() (result []Point) {
	start := s.StartPoint()
	result = append(result, start)
	for state := s.FirstMove(start); state.Tile != 'S'; state = s.NextMove(state) {
		result = append(result, state.Point)
	}
	return
}

func (s *Sketch) FirstMove(start Point) State {
	for _, heading := range []Direction{N, S, W, E} {
		next := start.Move(heading)
		tile := s.Tile(next)
		if pipe, found := Pipes[tile]; found {
			if pipe.Connects(heading) {
				return State{Point: next, Tile: tile, Heading: heading}
			}
		}
	}
	panic("No first move found")
}

func (s *Sketch) NextMove(state State) State {
	direction, err := state.Route()
	if err != nil {
		panic(err)
	}
	point := state.Point.Move(direction)
	tile := s.Tile(point)
	return State{
		Point:   point,
		Tile:    tile,
		Heading: direction,
	}
}

func (s *Sketch) Tile(point Point) rune {
	return s.Grid[point.Y][point.X]
}

func (s *Sketch) StartPoint() Point {
	for y, row := range s.Grid {
		if x := strings.Index(string(row), "S"); x != -1 {
			return Point{x, y}
		}
	}
	panic("No start found")
}

func (s *Sketch) CalculateArea() int {
	detSum := 0
	border := s.MapPoints()
	border = append(border, border[0])
	for i := 0; i < len(border)-1; i++ {
		detSum += det(border[i], border[i+1])
	}
	return detSum / 2
}

func det(p1, p2 Point) int {
	return p1.X*p2.Y - p1.Y*p2.X
}

func NewSketch(input []string) (s Sketch) {
	numRows := len(input)
	numCols := len(input[0])
	s = Sketch{
		Grid: make([][]rune, numRows+2),
	}

	// top and bottom
	border := make([]rune, numCols+2)
	for i := 0; i < len(border); i++ {
		border[i] = '.'
	}

	// build the grid
	s.Grid[0] = border
	s.Grid[len(s.Grid)-1] = border
	for y, line := range input {
		gridLine := make([]rune, numCols+2)
		gridLine[0] = '.'
		gridLine[numCols+1] = '.'
		for x, ch := range line {
			gridLine[x+1] = ch
		}
		s.Grid[y+1] = gridLine
	}
	return
}
