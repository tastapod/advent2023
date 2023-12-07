package day2

import (
	"strconv"
	"strings"
)

var GameLimits = map[string]int{"red": 12, "green": 13, "blue": 14}

func checkGames(gameLines []string, limits map[string]int) (possible []int) {
	possible = make([]int, 0, len(gameLines))
	for _, gameLine := range gameLines {
		game := 0
		isPossible := true

		parseGameLine(gameLine, func(thisGame, turn, drawNum int, draw Draw) (finished bool) {
			game = thisGame
			if draw.num > limits[draw.colour] {
				isPossible = false
				finished = true // short-circuit
			}
			return
		})
		if isPossible {
			possible = append(possible, game)
		}
	}
	return
}

type Draw struct {
	num    int
	colour string
}

type DrawOperation func(game, turn, drawNum int, draw Draw) (finished bool)

func parseGameLine(gameLine string, processDraw DrawOperation) {
	parts := strings.Split(gameLine, ": ") // ["Game n", turnStr]
	game, _ := strconv.Atoi(parts[0][5:])
	turnStr := strings.Split(parts[1], "; ") // ["3 red, 1 blue", "2 blue", ...]

	for iTurn, drawsStr := range turnStr {
		drawStr := strings.Split(drawsStr, ", ") // ["3", "red"]
		for iDraw, rawDraw := range drawStr {
			draw := parseDraw(rawDraw)
			if finished := processDraw(game, iTurn, iDraw, draw); finished {
				return // short circuit
			}
		}
	}
	return
}

/* parse "3 red" */
func parseDraw(count string) (draw Draw) {
	parts := strings.Split(count, " ")
	draw.num, _ = strconv.Atoi(parts[0])
	draw.colour = parts[1]
	return
}

func SumPossibleIds(lines []string, limits map[string]int) (total int) {
	possible := checkGames(lines, limits)
	for _, id := range possible {
		total += id
	}
	return
}

func powerOfGame(gameLine string) int {
	mins := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	// scan every draw for max number of that colour
	parseGameLine(gameLine, func(_, turn, drawNum int, draw Draw) (finished bool) {
		if draw.num > mins[draw.colour] {
			mins[draw.colour] = draw.num
		}
		return
	})

	return mins["red"] * mins["green"] * mins["blue"]
}

func SumPowersOfGames(gameLines []string) (total int) {
	for _, gameLine := range gameLines {
		total += powerOfGame(gameLine)
	}
	return
}
