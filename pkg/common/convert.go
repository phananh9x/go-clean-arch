package common

import "strconv"

func ConvertStringToInt64(input string) int64 {
	i, _ := strconv.ParseInt(input, 10, 64)
	return i
}

func Int64ToFloat64(in int64) float64 {
	return float64(in)
}
