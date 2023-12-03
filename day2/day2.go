package day2

import (
	"strconv"
	"strings"
)

var GameLimits = map[string]int{"red": 12, "green": 13, "blue": 14}

func checkGames(gameLines []string, limits map[string]int) (possible []int) {
	possible = make([]int, 0, len(gameLines))
	for _, gameLine := range gameLines {
		if isPossible, game := checkGame(gameLine, limits); isPossible {
			possible = append(possible, game)
		}
	}
	return
}

func checkGame(gameLine string, limits map[string]int) (isPossible bool, game int) {
	isPossible = true
	parts := strings.Split(gameLine, ": ") // ["Game n", turns]
	game, _ = strconv.Atoi(parts[0][5:])
	turns := strings.Split(parts[1], "; ") // ["3 red, 1 blue", "2 blue", ...]
	for _, turn := range turns {
		counts := strings.Split(turn, ", ") // ["3", "red"]
		for _, count := range counts {
			num, colour := parseCount(count)
			if num > limits[colour] {
				isPossible = false
				return
			}
		}
	}
	return
}

/* parse "3 red" */
func parseCount(count string) (num int, colour string) {
	parts := strings.Split(count, " ")
	num, _ = strconv.Atoi(parts[0])
	colour = parts[1]
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

	parts := strings.Split(gameLine, ": ") // ["Game n", turns]
	turns := strings.Split(parts[1], "; ") // ["3 red, 1 blue", "2 blue", ...]
	for _, turn := range turns {
		counts := strings.Split(turn, ", ") // ["3", "red"]
		for _, count := range counts {
			num, colour := parseCount(count)
			if num > mins[colour] {
				mins[colour] = num
			}
		}
	}
	return mins["red"] * mins["green"] * mins["blue"]
}

func SumPowersOfGames(gameLines []string) (total int) {
	for _, gameLine := range gameLines {
		total += powerOfGame(gameLine)
	}
	return
}
