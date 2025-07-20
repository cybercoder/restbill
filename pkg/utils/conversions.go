package utils

import (
	"fmt"
	"strconv"
)

func StringToUint(s string) (uint, error) {
	// Use 64-bit parsing but verify it fits in uint
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	// Explicit check for 32-bit systems
	if uint64(uint(num)) != num {
		return 0, fmt.Errorf("value %d overflows uint on this platform", num)
	}

	return uint(num), nil
}
