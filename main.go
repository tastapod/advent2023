package main

import (
	"fmt"
	"github.com/tastapod/advent2023/day1"
	"github.com/tastapod/advent2023/day10"
	"github.com/tastapod/advent2023/day2"
	"github.com/tastapod/advent2023/day3"
	"github.com/tastapod/advent2023/day4"
	"github.com/tastapod/advent2023/day5"
	"github.com/tastapod/advent2023/day6"
	"github.com/tastapod/advent2023/day7"
	"github.com/tastapod/advent2023/day8"
	"github.com/tastapod/advent2023/day9"
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
	solveDay9()
	solveDay10()
}

func solveDay1() {
	lines := input.ReadAndSplitDay(1)
	fmt.Println("Answer for day 1 part 1: ", day1.SumEnds(lines))          // 56465
	fmt.Println("Answer for day 2 part 1: ", day1.SumEndsWithWords(lines)) // 55902
}

func solveDay2() {
	lines := input.ReadAndSplitDay(2)
	fmt.Println("Answer for day 2 part 1: ", day2.SumPossibleIds(lines, day2.GameLimits)) // 2600
	fmt.Println("Answer for day 2 part 2: ", day2.SumPowersOfGames(lines))                // 86036
}

func solveDay3() {
	lines := input.ReadAndSplitDay(3)
	fmt.Println("Answer for day 3 part 1: ", day3.NewPartNumberFinder(lines).SumPartNumbers()) // 530495
	fmt.Println("Answer for day 3 part 2: ", day3.NewPartNumberFinder(lines).SumGears())       // 80253814
}

func solveDay4() {
	lines := input.ReadAndSplitDay(4)
	fmt.Println("Answer for day 4 part 1: ", day4.SumScratchCards(lines))  // 19855
	fmt.Println("Answer for day 4 part 2: ", day4.MoarScratchCards(lines)) // 10378710
}

func solveDay5() {
	almanac := input.ReadDay(5)
	fmt.Println("Answer for day 5 part 1: ", day5.FindSmallestLocation(almanac))       // 600279879
	fmt.Println("Answer for day 5 part 2: ", day5.FindSmallestMappedLocation(almanac)) // 20191102
}

func solveDay6() {
	data := input.ReadDay(6)
	fmt.Println("Answer for day 6 part 1: ", day6.NewRaceData(data).CalculateProductOfWins()) // 114400

	var bigRace = day6.Race{Time: 35937366, DistanceRecord: 212206012011044}
	fmt.Println("Answer for day 6 part 2: ", bigRace.CountWaysToWin()) // 21039729
}

func solveDay7() {
	lines := input.ReadAndSplitDay(7)
	fmt.Println("Answer for day 7 part 1: ", day7.NewSimpleRound(lines).Score()) // 255048101
	fmt.Println("Answer for day 7 part 2: ", day7.NewJokerRound(lines).Score())  // 253718286
}

func solveDay8() {
	mapInput := input.ReadDay(8)
	nodeMap := day8.NewNodeMap(mapInput)
	fmt.Println("Answer for day 8 part 1: ", nodeMap.CountStepsToZZZ("AAA")) // 14681
	fmt.Println("Answer for day 8 part 2: ", nodeMap.CountGhostSteps())      // 14321394058031
}

func solveDay9() {
	lines := input.ReadAndSplitDay(9)
	fmt.Println("Answer for day 9 part 1: ", day9.SumNextValues(lines))
	fmt.Println("Answer for day 9 part 2: ", day9.SumPreviousValues(lines))
}

func solveDay10() {
	rows := input.ReadAndSplitDay(10)
	sketch := day10.NewSketch(rows)
	fmt.Println("Answer for day 10 part 1: ", sketch.FurthestPoint())
}
