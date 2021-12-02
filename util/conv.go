package util

import "strconv"

func ConvertStringToInt(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return result
}
