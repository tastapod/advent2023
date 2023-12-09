package day5

import (
	"github.com/tastapod/advent2023/convert"
	"strings"
)

type Mapping struct {
	dst, src, len int
}

func (m *Mapping) Contains(value int) bool {
	return value >= m.src && value < m.src+m.len
}

func (m *Mapping) MapValue(value int) (int, bool) {
	if m.Contains(value) {
		return value - m.src + m.dst, true
	} else {
		return value, false
	}
}

func NewMapping(mappingChunk string) *Mapping {
	fields := strings.Fields(mappingChunk)
	mapping := &Mapping{
		dst: convert.ToInt(fields[0]),
		src: convert.ToInt(fields[1]),
		len: convert.ToInt(fields[2]),
	}
	return mapping
}

type RangeMap struct {
	srcName, dstName string
	mappings         []*Mapping
}

func (m *RangeMap) MapValue(value int) int {
	for _, mapping := range m.mappings {
		if result, found := mapping.MapValue(value); found {
			return result
		}
	}
	return value
}

func NewRangeMap(mapChunk string) (r *RangeMap) {
	lines := strings.Split(strings.TrimSpace(mapChunk), "\n")
	name := strings.Fields(lines[0])[0]
	nameParts := strings.Split(name, "-")
	mappings := lines[1:]

	r = &RangeMap{
		srcName:  nameParts[0],
		dstName:  nameParts[2],
		mappings: make([]*Mapping, len(mappings)),
	}

	for i, mappingChunk := range mappings {
		r.mappings[i] = NewMapping(mappingChunk)
	}
	return
}

type Almanac struct {
	seeds []int
	maps  map[string]*RangeMap
}

func (a *Almanac) FindLocation(seed int) (result int) {
	result = seed
	for mapSrc := "seed"; mapSrc != "location"; {
		rangeMap := a.maps[mapSrc]
		result = rangeMap.MapValue(result)
		mapSrc = rangeMap.dstName
	}
	return
}

func (a *Almanac) SmallestLocation() (result int) {
	for _, seed := range a.seeds {
		location := a.FindLocation(seed)
		if result == 0 || location < result {
			result = location
		}
	}
	return
}

func NewAlmanac(input string) (a *Almanac) {
	chunks := strings.Split(input, "\n\n")
	seeds := strings.Fields(strings.Split(chunks[0], ": ")[1])
	mapChunks := chunks[1:]

	// create empty Almanac
	a = &Almanac{
		seeds: make([]int, len(seeds)),
		maps:  make(map[string]*RangeMap, len(mapChunks)),
	}

	// parse seeds
	for i, seed := range seeds {
		a.seeds[i] = convert.ToInt(seed)
	}

	// parse each range map
	for _, mapChunk := range mapChunks {
		rangeMap := NewRangeMap(mapChunk)
		a.maps[rangeMap.srcName] = rangeMap
	}
	return
}

func FindSmallestLocation(almanacSrc string) int {
	return NewAlmanac(almanacSrc).SmallestLocation()
}
