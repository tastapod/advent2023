package day5

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParsesMaps(t *testing.T) {
	assert := assert.New(t)
	input := strings.TrimSpace(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48
`)
	almanac := NewPointAlmanac(input)
	assert.Equal(1, len(almanac.maps))
	assert.Equal(2, len(almanac.maps["seed"].mappings))
	assert.Equal([]int{79, 14, 55, 13}, almanac.seeds)
}

func TestMappingMapsValues(t *testing.T) {
	assert := assert.New(t)
	mapping := NewMapping("52 50 48")

	assert.False(mapping.Contains(49))
	assert.True(mapping.Contains(50))
	assert.True(mapping.Contains(97))
	assert.False(mapping.Contains(98))

	result, found := mapping.MapValue(50)
	assert.Equal(52, result)
	assert.True(found)

	result, found = mapping.MapValue(97)
	assert.Equal(99, result)
	assert.True(found)

	result, found = mapping.MapValue(100)
	assert.Equal(100, result)
	assert.False(found)
}

func TestRangeMapMapsValues(t *testing.T) {
	assert := assert.New(t)
	m := NewResourceMap(strings.TrimSpace(`
seed-to-soil map:
50 98 2
52 50 48
`))
	// from puzzle text
	assert.Equal(81, m.MapValue(79))
	assert.Equal(14, m.MapValue(14))
	assert.Equal(57, m.MapValue(55))
	assert.Equal(13, m.MapValue(13))
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
	assert := assert.New(t)
	almanac := NewPointAlmanac(sampleAlmanac)
	assert.Equal(82, almanac.FindLocation(79))
	assert.Equal(43, almanac.FindLocation(14))
	assert.Equal(86, almanac.FindLocation(55))
	assert.Equal(35, almanac.FindLocation(13))
}

func TestFindsSmallestLocation(t *testing.T) {
	assert.Equal(t, 35, FindSmallestLocation(sampleAlmanac))
}

func TestSplitsRangeBasedOnMapping(t *testing.T) {
	assert := assert.New(t)
	var mapping = NewMapping("200 100 10") // 100 -> 200
	var result RangeGroup

	// below range
	result = mapping.apply([]Range{{5, 8}})
	assert.Equal([]Range{{5, 8}}, result.unmapped)
	assert.Empty(result.mapped)

	// above range
	result = mapping.apply([]Range{{150, 152}})
	assert.Equal([]Range{{150, 152}}, result.unmapped)
	assert.Empty(result.mapped)

	// inside range
	result = mapping.apply([]Range{{101, 103}})
	assert.Empty(result.unmapped)
	assert.Equal([]Range{{201, 203}}, result.mapped)

	// overlap below
	result = mapping.apply([]Range{{95, 104}})
	assert.Equal([]Range{{95, 99}}, result.unmapped)
	assert.Equal([]Range{{200, 204}}, result.mapped)

	// overlap above
	result = mapping.apply([]Range{{108, 120}})
	assert.Equal([]Range{{110, 120}}, result.unmapped)
	assert.Equal([]Range{{208, 209}}, result.mapped)
}

func TestSplitsMultipleRanges(t *testing.T) {
	assert := assert.New(t)
	mapping := NewMapping("49 53 8")
	result := mapping.apply([]Range{{57, 69}})
	assert.Equal([]Range{{61, 69}}, result.unmapped)
	assert.Equal([]Range{{53, 56}}, result.mapped)
}

func TestFindsSmallestLocationForAnySeed(t *testing.T) {
	assert.Equal(t, 46, FindSmallestMappedLocation(sampleAlmanac))
}
