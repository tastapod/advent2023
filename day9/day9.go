package day9

import (
	"github.com/tastapod/advent2023/convert"
	"github.com/tastapod/advent2023/seq"
	"strings"
)

type Predictor struct {
	Values []int
}

func (p Predictor) NextValue() (result int) {
	current := make([]int, len(p.Values))
	copy(current, p.Values)

	// keep going until we get a row of zeros
	for seq.Last(current) != 0 {
		result += seq.Last(current)

		// reduce the row above
		next := make([]int, len(current)-1)
		for i := 0; i < len(next); i++ {
			next[i] = current[i+1] - current[i]
		}
		current = next
	}
	return
}

func NewPredictor(seq string) (p Predictor) {
	nums := strings.Fields(seq)
	p = Predictor{
		Values: make([]int, len(nums)),
	}
	for i, num := range nums {
		p.Values[i] = convert.ToInt(num)
	}
	return
}

func SumNextValues(lines []string) (result int) {
	for _, line := range lines {
		result += NewPredictor(line).NextValue()
	}
	return
}
