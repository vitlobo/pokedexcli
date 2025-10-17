package util

// Clamp ensures val stays within [min, max].
func Clamp[T ~float64 | ~int](val, min, max T) T {
	if val < min { return min }
	if val > max { return max }
	return val
}

func ClampInt(val, min, max int) int {
	if val < min { return min }
	if val > max { return max }
	return val
}