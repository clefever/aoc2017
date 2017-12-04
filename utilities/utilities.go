package utilities

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

// NumberStringToElements parses a string into a list of single character numbers
func NumberStringToElements(input string) (numbers []int) {
	for _, s := range input {
		num, _ := strconv.Atoi(string(s))
		numbers = append(numbers, num)
	}
	return numbers
}

// NumberListToSlice parses a single line string into a list of numbers
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

// StringListToSlice parses a single line string into a list of strings
func StringListToSlice(input string) (strs []string) {
	fields := strings.Fields(input)
	for _, f := range fields {
		if f != "" {
			strs = append(strs, f)
		}
	}
	return strs
}

// NumbersByLineToSlice parses a string into a list of lines of numbers
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

// StringsByLineToSlice parses a string into a list of lines of strings
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
