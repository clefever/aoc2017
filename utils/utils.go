package utils

import (
	"strconv"
	"strings"
)

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

// NumberStringToElements parses a string into a list of single character numbers.
func NumberStringToElements(input string) (numbers []int) {
	for _, s := range input {
		num, _ := strconv.Atoi(string(s))
		numbers = append(numbers, num)
	}
	return numbers
}

// StringToElements parses a string into a list of single characters.
func StringToElements(input string) (strings []string) {
	for _, s := range input {
		strings = append(strings, string(s))
	}
	return strings
}

// NumberListToSlice parses a single line string into a list of numbers.
func NumberListToSlice(input string) (numbers []int) {
	strs := strings.Fields(input)
	for _, s := range strs {
		if s != "" {
			num, _ := strconv.Atoi(s)
			numbers = append(numbers, num)
		}
	}
	return numbers
}

// StringListToSlice parses a single line string into a list of strings.
func StringListToSlice(input string) (strs []string) {
	fields := strings.Fields(input)
	for _, f := range fields {
		if f != "" {
			strs = append(strs, f)
		}
	}
	return strs
}

// NumbersByLineToSlice parses a string into a list of lines of numbers.
func NumbersByLineToSlice(input string) (numbers [][]int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line != "" {
			nums := NumberListToSlice(line)
			numbers = append(numbers, nums)
		}
	}
	return numbers
}

// StringsByLineToSlice parses a string into a list of lines of strings.
func StringsByLineToSlice(input string) (strs [][]string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line != "" {
			s := StringListToSlice(line)
			strs = append(strs, s)
		}
	}
	return strs
}

// SlicesEqual determines whether slices are equal based on their elements.
func SlicesEqual(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
