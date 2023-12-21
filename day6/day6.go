package day6

import (
	"github.com/tastapod/advent2023/convert"
	"strings"
)

type Race struct {
	Time, DistanceRecord int64
}

func (race Race) CountWaysToWin() (result int64) {
	for i := int64(0); i < race.Time; i++ {
		if distanceRun := race.DistanceRun(i); distanceRun > race.DistanceRecord {
			result++
		}
	}
	return
}

func (race Race) DistanceRun(chargeTime int64) int64 {
	return chargeTime * (race.Time - chargeTime)
}

type RaceData struct {
	Races []Race
}

func (r RaceData) CalculateProductOfWins() (result int64) {
	result = 1
	for _, race := range r.Races {
		result *= race.CountWaysToWin()
	}
	return
}

func NewRaceData(input string) RaceData {
	parts := strings.Split(input, "\n")
	times := strings.Fields(parts[0])[1:]
	distances := strings.Fields(parts[1])[1:]

	result := RaceData{make([]Race, len(times))}

	for i := 0; i < len(times); i++ {
		result.Races[i] = Race{Time: convert.ToInt64(times[i]), DistanceRecord: convert.ToInt64(distances[i])}
	}
	return result
}
