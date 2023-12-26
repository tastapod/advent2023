package day5

import (
	"fmt"
	"github.com/tastapod/advent2023/convert"
	"math"
	"strings"
)

type Mapping struct {
	dstStart, srcStart, length int
}

func (m Mapping) Contains(value int) bool {
	return value >= m.srcStart && value < m.srcStart+m.length
}

func (m Mapping) MapValue(value int) (int, bool) {
	if m.Contains(value) {
		return value - m.srcStart + m.dstStart, true
	} else {
		return value, false
	}
}

func (m Mapping) String() string {
	return fmt.Sprintf("mapping: [%d..%d] -> [%d..%d]",
		m.srcStart, m.srcStart+m.length-1,
		m.dstStart, m.dstStart+m.length-1)

}

type Range struct {
	first, last int
}

func (r Range) length() int {
	return r.last + r.first + 1
}

type RangeGroup struct {
	unmapped, mapped []Range
}

func (m Mapping) apply(inputs []Range) (result RangeGroup) {
	result = RangeGroup{}

	for _, input := range inputs {
		if input.last < m.srcStart || input.first > m.sourceLast() {
			// no overlap
			result.unmapped = append(result.unmapped, input)
			continue
		}

		// overlapping

		// check for unmapped range at the beginning
		if input.first < m.srcStart {
			result.unmapped = append(result.unmapped, Range{input.first, min(input.last, m.srcStart-1)})
		}

		// overlap
		first := max(input.first, m.srcStart)
		last := min(input.last, m.sourceLast())
		shift := m.dstStart - m.srcStart
		result.mapped = append(result.mapped, Range{first + shift, last + shift})

		// check for unmapped range at the end
		if input.last > m.sourceLast() {
			result.unmapped = append(result.unmapped, Range{
				first: m.sourceLast() + 1,
				last:  input.last})
		}
	}
	return
}

// sourceLast returns the highest inclusive value in the source range
func (m Mapping) sourceLast() int {
	return m.srcStart + m.length - 1
}

func NewMapping(mappingChunk string) Mapping {
	fields := strings.Fields(mappingChunk)
	return Mapping{
		dstStart: convert.ToInt(fields[0]),
		srcStart: convert.ToInt(fields[1]),
		length:   convert.ToInt(fields[2]),
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
		first := convert.ToInt(seedLine[i*2])
		length := convert.ToInt(seedLine[i*2+1])
		a.seedRanges[i] = Range{first, first + length - 1}
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
