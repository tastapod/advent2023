package day5

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParsesMaps(t *testing.T) {
	input := strings.TrimSpace(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48
`)
	almanac := NewPointAlmanac(input)
	assert.Equal(t, 1, len(almanac.maps))
	assert.Equal(t, 2, len(almanac.maps["seed"].mappings))
	assert.Equal(t, []int{79, 14, 55, 13}, almanac.seeds)
}

func TestMappingMapsValues(t *testing.T) {
	mapping := NewMapping("52 50 48")

	assert.False(t, mapping.Contains(49))
	assert.True(t, mapping.Contains(50))
	assert.True(t, mapping.Contains(97))
	assert.False(t, mapping.Contains(98))

	result, found := mapping.MapValue(50)
	assert.Equal(t, 52, result)
	assert.True(t, found)

	result, found = mapping.MapValue(97)
	assert.Equal(t, 99, result)
	assert.True(t, found)

	result, found = mapping.MapValue(100)
	assert.Equal(t, 100, result)
	assert.False(t, found)
}

func TestRangeMapMapsValues(t *testing.T) {
	m := NewResourceMap(strings.TrimSpace(`
seed-to-soil map:
50 98 2
52 50 48
`))
	// from puzzle text
	assert.Equal(t, 81, m.MapValue(79))
	assert.Equal(t, 14, m.MapValue(14))
	assert.Equal(t, 57, m.MapValue(55))
	assert.Equal(t, 13, m.MapValue(13))
}

//seeds: 79 14 55 13

var sampleAlmanac = strings.TrimSpace(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`)

func TestFindsLocationForSeed(t *testing.T) {
	almanac := NewPointAlmanac(sampleAlmanac)
	assert.Equal(t, 82, almanac.FindLocation(79))
	assert.Equal(t, 43, almanac.FindLocation(14))
	assert.Equal(t, 86, almanac.FindLocation(55))
	assert.Equal(t, 35, almanac.FindLocation(13))
}

func TestFindsSmallestLocation(t *testing.T) {
	assert.Equal(t, 35, FindSmallestLocation(sampleAlmanac))
}

func TestSplitsRangeBasedOnMapping(t *testing.T) {
	var mapping = NewMapping("200 100 10") // 100 -> 200
	var result RangeGroup

	// below range
	result = mapping.apply([]Range{{5, 3}})
	assert.Equal(t, []Range{{5, 3}}, result.unmapped)
	assert.Empty(t, result.mapped)

	// above range
	result = mapping.apply([]Range{{150, 3}})
	assert.Equal(t, []Range{{150, 3}}, result.unmapped)
	assert.Empty(t, result.mapped)

	// inside range
	result = mapping.apply([]Range{{101, 3}})
	assert.Empty(t, result.unmapped)
	assert.Equal(t, []Range{{201, 3}}, result.mapped)

	// overlap below
	result = mapping.apply([]Range{{95, 10}})
	assert.Equal(t, []Range{{95, 5}}, result.unmapped)
	assert.Equal(t, []Range{{200, 5}}, result.mapped)

	// overlap above
	result = mapping.apply([]Range{{108, 12}})
	assert.Equal(t, []Range{{110, 10}}, result.unmapped)
	assert.Equal(t, []Range{{208, 2}}, result.mapped)
}

func TestSplitsMultipleRanges(t *testing.T) {
	mapping := NewMapping("49 53 8")
	result := mapping.apply([]Range{{57, 13}})
	assert.Equal(t, []Range{{61, 9}}, result.unmapped)
	assert.Equal(t, []Range{{53, 4}}, result.mapped)
}

func TestFindsSmallestLocationForAnySeed(t *testing.T) {
	t.Skip("What am I missing here?")
	assert.Equal(t, 46, FindSmallestMappedLocation(sampleAlmanac))
}
