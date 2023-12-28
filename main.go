package main

import (
	"fmt"
	"github.com/tastapod/advent2023/day1"
	"github.com/tastapod/advent2023/day2"
	"github.com/tastapod/advent2023/day3"
	"github.com/tastapod/advent2023/day4"
	"github.com/tastapod/advent2023/day5"
	"github.com/tastapod/advent2023/day6"
	"github.com/tastapod/advent2023/day7"
	"github.com/tastapod/advent2023/day8"
	"github.com/tastapod/advent2023/input"
)

func main() {
	solveDay1()
	solveDay2()
	solveDay3()
	solveDay4()
	solveDay5()
	solveDay6()
	solveDay7()
	solveDay8()
}

func solveDay1() {
	lines := input.LinesForDay(1)
	fmt.Println("Answer for day 1 part 1: ", day1.SumEnds(lines))          // 56465
	fmt.Println("Answer for day 2 part 1: ", day1.SumEndsWithWords(lines)) // 55902
}

func solveDay2() {
	lines := input.LinesForDay(2)
	fmt.Println("Answer for day 2 part 1: ", day2.SumPossibleIds(lines, day2.GameLimits)) // 2600
	fmt.Println("Answer for day 2 part 2: ", day2.SumPowersOfGames(lines))                // 86036
}

func solveDay3() {
	lines := input.LinesForDay(3)
	fmt.Println("Answer for day 3 part 1: ", day3.NewPartNumberFinder(lines).SumPartNumbers()) // 530495
	fmt.Println("Answer for day 3 part 2: ", day3.NewPartNumberFinder(lines).SumGears())       // 80253814
}

func solveDay4() {
	lines := input.LinesForDay(4)
	fmt.Println("Answer for day 4 part 1: ", day4.SumScratchCards(lines))  // 19855
	fmt.Println("Answer for day 4 part 2: ", day4.MoarScratchCards(lines)) // 10378710
}

func solveDay5() {
	almanac := input.ForDay(5)
	fmt.Println("Answer for day 5 part 1: ", day5.FindSmallestLocation(almanac))       // 600279879
	fmt.Println("Answer for day 5 part 2: ", day5.FindSmallestMappedLocation(almanac)) // 20191102
}

func solveDay6() {
	data := input.ForDay(6)
	fmt.Println("Answer for day 6 part 1: ", day6.NewRaceData(data).CalculateProductOfWins()) // 114400

	var bigRace = day6.Race{Time: 35937366, DistanceRecord: 212206012011044}
	fmt.Println("Answer for day 6 part 2: ", bigRace.CountWaysToWin()) // 21039729
}

func solveDay7() {
	lines := input.LinesForDay(7)
	fmt.Println("Answer for day 7 part 1: ", day7.NewSimpleRound(lines).Score())
	fmt.Println("Answer for day 7 part 2: ", day7.NewJokerRound(lines).Score())
}

func solveDay8() {
	mapInput := input.ForDay(8)
	nodeMap := day8.NewNodeMap(mapInput)
	fmt.Println("Answer for day 8 part 1: ", nodeMap.CountStepsToZZZ("AAA"))
	fmt.Println("Answer for day 8 part 2: ", nodeMap.CountGhostSteps())
}
