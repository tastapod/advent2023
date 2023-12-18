package day5

import (
	"fmt"
	"github.com/tastapod/advent2023/convert"
	"math"
	"strings"
)

type Mapping struct {
	dest, source, length int
}

func (m Mapping) Contains(value int) bool {
	return value >= m.source && value < m.source+m.length
}

func (m Mapping) MapValue(value int) (int, bool) {
	if m.Contains(value) {
		return value - m.source + m.dest, true
	} else {
		return value, false
	}
}

type Range struct {
	first, length int
}

func (r Range) last() int {
	return r.first + r.length - 1
}

type RangeGroup struct {
	unmapped, mapped []Range
}

func (m Mapping) apply(inputs []Range) (result RangeGroup) {
	result = RangeGroup{}

	for _, input := range inputs {
		if input.last() < m.source || input.first > m.sourceLast() {
			// no overlap
			result.unmapped = inputs
		} else {
			// overlapping

			if input.first < m.source {
				// unmapped range at the beginning
				result.unmapped = append(result.unmapped, Range{input.first, m.source - input.first})
			}

			// overlap
			first := max(input.first, m.source)
			last := min(input.last(), m.sourceLast())
			length := last - first + 1
			dest := m.dest + first - m.source
			result.mapped = append(result.mapped, Range{dest, length})

			if input.last() > m.sourceLast() {
				// unmapped range at the end
				result.unmapped = append(result.unmapped, Range{
					first:  input.first + length,
					length: input.length - length})
			}
		}
	}
	//fmt.Printf("%v on %v\n-> %v\n", m, inputs, result)

	return
}

// sourceLast returns the highest inclusive value in the source range
func (m Mapping) sourceLast() int {
	return m.source + m.length - 1
}

func NewMapping(mappingChunk string) Mapping {
	fields := strings.Fields(mappingChunk)
	return Mapping{
		dest:   convert.ToInt(fields[0]),
		source: convert.ToInt(fields[1]),
		length: convert.ToInt(fields[2]),
	}
}

type ResourceMap struct {
	srcName, destName string
	mappings          []Mapping
}

func (m ResourceMap) MapValue(value int) int {
	for _, mapping := range m.mappings {
		if result, found := mapping.MapValue(value); found {
			return result
		}
	}
	return value
}

func (m ResourceMap) Apply(input []Range) []Range {
	unmapped := input
	mapped := make([]Range, 0, len(input))

	for _, mapping := range m.mappings {
		result := mapping.apply(unmapped)
		unmapped = result.unmapped
		mapped = append(mapped, result.mapped...)
	}
	return append(unmapped, mapped...)
}

func NewResourceMap(mapChunk string) (r ResourceMap) {
	lines := strings.Split(strings.TrimSpace(mapChunk), "\n")
	name := strings.Fields(lines[0])[0]
	nameParts := strings.Split(name, "-")
	mappings := lines[1:]

	r = ResourceMap{
		srcName:  nameParts[0],
		destName: nameParts[2],
		mappings: make([]Mapping, len(mappings)),
	}

	for i, mappingChunk := range mappings {
		r.mappings[i] = NewMapping(mappingChunk)
	}
	return
}

type PointAlmanac struct {
	seeds []int
	maps  map[string]ResourceMap
}

func (a PointAlmanac) FindLocation(seed int) (result int) {
	result = seed
	for resource := "seed"; resource != "location"; {
		m := a.maps[resource]
		result = m.MapValue(result)
		resource = m.destName
	}
	return
}

func (a PointAlmanac) SmallestLocation() (result int) {
	for _, seed := range a.seeds {
		location := a.FindLocation(seed)
		if result == 0 || location < result {
			result = location
		}
	}
	return
}

func NewPointAlmanac(input string) (result PointAlmanac) {
	chunks := strings.Split(input, "\n\n")
	seeds := strings.Fields(strings.Split(chunks[0], ": ")[1])
	mapChunks := chunks[1:]

	// create empty Almanac
	result = PointAlmanac{
		seeds: make([]int, len(seeds)),
		maps:  make(map[string]ResourceMap, len(mapChunks)),
	}

	// parse seeds
	for i, seed := range seeds {
		result.seeds[i] = convert.ToInt(seed)
	}

	// parse each resource map
	for _, mapChunk := range mapChunks {
		r := NewResourceMap(mapChunk)
		result.maps[r.srcName] = r
	}
	return
}

func FindSmallestLocation(almanacSrc string) int {
	return NewPointAlmanac(almanacSrc).SmallestLocation()
}

// -------------------------------------------------------

type RangeAlmanac struct {
	seedRanges []Range
	maps       []ResourceMap
}

func (a RangeAlmanac) findSmallestMappedLocation() (result int) {
	ranges := a.seedRanges
	for _, resourceMap := range a.maps {
		fmt.Printf("%s = %v\n", resourceMap.srcName, ranges)
		ranges = resourceMap.Apply(ranges)
	}

	result = math.MaxInt
	for _, l := range ranges {
		result = min(result, l.first)
	}
	return
}

func NewRangeAlmanac(input string) (a RangeAlmanac) {
	chunks := strings.Split(input, "\n\n")
	seedLine := strings.Fields(strings.Split(chunks[0], ": ")[1])
	mapChunks := chunks[1:]

	a = RangeAlmanac{
		seedRanges: make([]Range, len(seedLine)/2),
		maps:       make([]ResourceMap, len(mapChunks)),
	}

	// parse seed ranges
	for i := 0; i < len(seedLine)/2; i++ {
		start := convert.ToInt(seedLine[i*2])
		length := convert.ToInt(seedLine[i*2+1])
		a.seedRanges[i] = Range{start, length}
	}

	// parse each resource map
	for i, mapChunk := range mapChunks {
		a.maps[i] = NewResourceMap(mapChunk)
	}
	return
}

func FindSmallestMappedLocation(almanac string) int {
	return NewRangeAlmanac(almanac).findSmallestMappedLocation()
}
