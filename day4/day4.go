package day4

import (
	"sort"
	"strings"

	"github.com/clefever/aoc2017/utilities"
)

// PartOne generates the solution to part one of day 4 of AOC2017
func PartOne(input string) (result int) {
	wordSlices := utilities.StringsByLineToSlice(input)
	for _, line := range wordSlices {
		m := map[string]bool{}
		for _, str := range line {
			m[str] = true
		}
		if len(m) == len(line) {
			result++
		}
	}
	return result
}

// PartTwo generates the solution to part two of day 4 of AOC2017
func PartTwo(input string) (result int) {
	wordSlices := utilities.StringsByLineToSlice(input)
	for _, line := range wordSlices {
		m := map[string]bool{}
		for _, str := range line {
			s := strings.Split(str, "")
			sort.Strings(s)
			m[strings.Join(s, "")] = true
		}
		if len(m) == len(line) {
			result++
		}
	}
	return result
}
