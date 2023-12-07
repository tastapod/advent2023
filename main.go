package main

import (
	"github.com/tastapod/advent2023/day1"
	"github.com/tastapod/advent2023/day2"
	"github.com/tastapod/advent2023/input"
	"strings"
)

func main() {
	solveDay1()
	solveDay2()
}

func solveDay1() {
	lines := strings.Split(input.ForDay(1), "\n")
	println("Answer for day 1 part 1: ", day1.SumEnds(lines))          // 56465
	println("Answer for day 2 part 1: ", day1.SumEndsWithWords(lines)) // 55902
}

func solveDay2() {
	lines := strings.Split(input.ForDay(2), "\n")
	println("Answer for day 2 part 1: ", day2.SumPossibleIds(lines, day2.GameLimits)) // 2600
	println("Answer for day 2 part 2: ", day2.SumPowersOfGames(lines))                // 86036
}
