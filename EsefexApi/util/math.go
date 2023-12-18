package util

func ClampInt(v int, min int, max int) int {
	if v < min {
		return min
	} else if v > max {
		return max
	}

	return v
}
