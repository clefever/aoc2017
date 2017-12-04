package utilities

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
