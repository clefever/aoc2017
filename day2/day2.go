package day2

import (
	"math"
	"strconv"
	"strings"
)

// PartOne generates the solution to part one of day 2 of AOC2017
func PartOne(input string) (result int) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		if l != "" {
			max := math.MinInt32
			min := math.MaxInt32
			numbers := strings.Fields(l)
			for _, n := range numbers {
				number, _ := strconv.Atoi(n)
				if number > max {
					max = number
				}
				if number < min {
					min = number
				}
			}
			result += max - min
		}
	}
	return result
}

// PartTwo generates the solution to part two of day 2 of AOC2017
func PartTwo(input string) (result int) {
	lines := strings.Split(input, "\n")
	for _, r := range lines {
		if r != "" {
			division := 0
			numbers := strings.Fields(r)
			for i, first := range numbers {
				f, _ := strconv.Atoi(first)
				for _, second := range numbers[i+1:] {
					s, _ := strconv.Atoi(second)
					if f%s == 0 {
						division = f / s
					} else if s%f == 0 {
						division = s / f
					}
				}
			}
			result += division
		}
	}
	return result
}
