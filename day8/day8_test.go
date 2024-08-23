package day8

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var sampleInput = strings.TrimSpace(`
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`)

func TestParsesInput(t *testing.T) {
	assert := assert.New(t)

	nodeMap := NewNodeMap(sampleInput)
	assert.Equal(3, len(nodeMap.Nodes))
	assert.Equal(Key("AAA"), nodeMap.Turn("BBB", 'L'))
	assert.Equal(Key("ZZZ"), nodeMap.Turn("ZZZ", 'R'))
}

func TestFollowsDirections(t *testing.T) {
	nodeMap := NewNodeMap(sampleInput)
	assert.Equal(t, 6, nodeMap.CountStepsToZZZ("AAA"))
}

var ghostInput = strings.TrimSpace(`
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`)

func TestFollowsGhostPath(t *testing.T) {
	//t.Skip("Not yet")
	nodeMap := NewNodeMap(ghostInput)
	assert.Equal(t, 6, nodeMap.CountGhostSteps())
}

func TestCalculatesGCD(t *testing.T) {
	assert.Equal(t, 5, Gcd(25, 35))
}

func TestCalculatesLCM(t *testing.T) {
	assert.Equal(t, 175, Lcm(25, 35))
	assert.Equal(t, 156, LcmReduce([]int{26, 39, 52}))
}
