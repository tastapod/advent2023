package main

import (
	"github.com/tastapod/advent2023/day1"
	"github.com/tastapod/advent2023/day2"
	"github.com/tastapod/advent2023/day3"
	"github.com/tastapod/advent2023/input"
)

func main() {
	solveDay1()
	solveDay2()
	solveDay3()
}

func solveDay1() {
	lines := input.LinesForDay(1)
	println("Answer for day 1 part 1: ", day1.SumEnds(lines))          // 56465
	println("Answer for day 2 part 1: ", day1.SumEndsWithWords(lines)) // 55902
}

func solveDay2() {
	lines := input.LinesForDay(2)
	println("Answer for day 2 part 1: ", day2.SumPossibleIds(lines, day2.GameLimits)) // 2600
	println("Answer for day 2 part 2: ", day2.SumPowersOfGames(lines))                // 86036
}

func solveDay3() {
	lines := input.LinesForDay(3)
	println("Answer for day 3 part 1: ", day3.NewPartNumberFinder(lines).SumPartNumbers()) // 530495
	println("Answer for day 3 part 2: ", day3.NewPartNumberFinder(lines).SumGears())       // 80253814
}
