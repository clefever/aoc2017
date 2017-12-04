package day2

import (
	"math"

	"github.com/clefever/aoc2017/utils"
)

// PartOne generates the solution to part one of day 2 of AOC2017
func PartOne(input string) (result int) {
	lines := utils.NumbersByLineToSlice(input)
	for _, line := range lines {
		max := math.MinInt32
		min := math.MaxInt32
		for _, num := range line {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		result += max - min
	}
	return result
}

// PartTwo generates the solution to part two of day 2 of AOC2017
func PartTwo(input string) (result int) {
	lines := utils.NumbersByLineToSlice(input)
	for _, line := range lines {
		division := 0
		for i, first := range line {
			for _, second := range line[i+1:] {
				if first%second == 0 {
					division = first / second
				} else if second%first == 0 {
					division = second / first
				}
			}
		}
		result += division
	}
	return result
}
