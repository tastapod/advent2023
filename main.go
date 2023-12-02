package main

import (
	"github.com/tastapod/advent2023/day1"
	"github.com/tastapod/advent2023/input"
	"strings"
)

func main() {
	solveDay1()
}

func solveDay1() {
	lines := strings.Split(input.ForDay(1), "\n")
	println("Answer for day 1 part 1: ", day1.SumEnds(lines))
	println("Answer for day 2 part 1: ", day1.SumEndsWithWords(lines))
}
