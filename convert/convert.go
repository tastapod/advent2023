package convert

import "strconv"

func ToInt(num string) int {
	value, _ := strconv.Atoi(num)
	return value
}

func ToInt64(num string) int64 {
	value, _ := strconv.ParseInt(num, 10, 64)
	return value
}
