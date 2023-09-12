package utils

import "strconv"

func ParseUint(numStr string) (uint64, error) {
	return strconv.ParseUint(numStr, 10, 64)
}
