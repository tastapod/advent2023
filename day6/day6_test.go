package day6

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var sampleRaceData = strings.TrimSpace(`
Time:      7  15   30
DistanceRecord:  9  40  200
`)

func TestParsesRaceData(t *testing.T) {
	raceData := NewRaceData(sampleRaceData)
	assert.Equal(t, 3, len(raceData.Races))
}

func TestCalculatesRaceTime(t *testing.T) {
	r := Race{7, 9}
	assert.Equal(t, r.DistanceRun(0), int64(0))
	assert.Equal(t, r.DistanceRun(1), int64(6))
	assert.Equal(t, r.DistanceRun(2), int64(10))
	assert.Equal(t, r.DistanceRun(3), int64(12))
	assert.Equal(t, r.DistanceRun(4), int64(12))
	assert.Equal(t, r.DistanceRun(5), int64(10))
	assert.Equal(t, r.DistanceRun(6), int64(6))
	assert.Equal(t, r.DistanceRun(7), int64(0))
}

func TestCountsWaysToWin(t *testing.T) {
	r := Race{7, 9}
	assert.Equal(t, int64(4), r.CountWaysToWin())
}

func TestCalculatesProductOfWins(t *testing.T) {
	raceData := NewRaceData(sampleRaceData)
	assert.Equal(t, int64(288), raceData.CalculateProductOfWins())
}
