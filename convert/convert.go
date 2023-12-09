package convert

import "strconv"

func ToInt(num string) int {
	value, _ := strconv.Atoi(num)
	return value
}
