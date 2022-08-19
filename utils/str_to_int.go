package utils

import (
	"strconv"
)

// StrToInt converts string number to int
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
