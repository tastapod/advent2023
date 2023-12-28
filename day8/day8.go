package day8

import (
	"fmt"
	"strings"
)

type Key string

func (k Key) IsStart() bool {
	return k[2] == 'A'
}

func (k Key) IsEnd() bool {
	return k[2] == 'Z'
}

type Node struct {
	Left, Right Key
}

type NodeMap struct {
	Directions []byte
	Nodes      map[Key]Node
}

func (m *NodeMap) Turn(key Key, direction byte) Key {
	if direction == 'L' {
		return m.Nodes[key].Left
	}
	return m.Nodes[key].Right
}

func (m *NodeMap) TurnRight(key Key) Key {
	return m.Nodes[key].Right
}

func (m *NodeMap) CountStepsToZZZ(start Key) (steps int) {
	current := start
	for i := 0; current != "ZZZ"; i = (i + 1) % len(m.Directions) {
		current = m.Turn(current, m.Directions[i])
		steps++
	}
	return
}

func (m *NodeMap) CountStepsToAnyZ(start Key) (steps int) {
	current := start
	for i := 0; current[2] != 'Z'; i = (i + 1) % len(m.Directions) {
		current = m.Turn(current, m.Directions[i])
		steps++
	}
	return
}

func (m *NodeMap) CountGhostSteps() (steps int) {
	startNodes := make([]Key, 0)

	// set up range
	for key := range m.Nodes {
		if key.IsStart() {
			startNodes = append(startNodes, key)
		}
	}
	println("Starting with", len(startNodes), "nodes")

	// Identify cycle for each ghost node
	loopCounts := make([]int, len(startNodes))
	for i, key := range startNodes {
		loopCounts[i] = m.CountStepsToAnyZ(key)
	}
	fmt.Printf("Counts: %v\n", loopCounts)

	// this only works because each ghost loops in a cycle
	return LcmMany(loopCounts)
}

func NewNodeMap(input string) (result NodeMap) {
	parts := strings.Split(input, "\n\n")
	nodeLines := strings.Split(parts[1], "\n")

	result = NodeMap{
		Directions: []byte(strings.TrimSpace(parts[0])),
		Nodes:      make(map[Key]Node),
	}

	for _, line := range nodeLines {
		// xxx = (aaa, bbb)
		// 0....5....0....5

		key := Key(line[0:3])
		left := Key(line[7:10])
		right := Key(line[12:15])
		result.Nodes[key] = Node{Left: left, Right: right}
	}
	return
}

func GcdMany(values []int) (result int) {
	result = values[0]
	for i := 1; i < len(values); i++ {
		result = Gcd(result, values[i])
	}
	return
}

func Gcd(v1, v2 int) int {
	big := max(v1, v2)
	small := min(v1, v2)
	for big != small {
		diff := big - small
		if diff < small {
			big = small
			small = diff
		} else {
			big = diff
		}
	}
	return big // or small
}

func Lcm(v1, v2 int) int {
	return (v1 / Gcd(v1, v2)) * v2
}

func LcmMany(values []int) (result int) {
	result = values[0]
	for i := 1; i < len(values); i++ {
		result = Lcm(result, values[i])
	}
	return
}
