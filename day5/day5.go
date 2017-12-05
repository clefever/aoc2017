package day5

import "github.com/clefever/aoc2017/utils"

// PartOne generates the solution to part one of day 5 of AOC2017
func PartOne(input string) (result int) {
	nums := utils.NumberListToSlice(input)
	pos, steps := 0, 0
	for pos < len(nums) {
		temp := pos
		pos += nums[pos]
		nums[temp]++
		steps++
	}
	return steps
}

// PartTwo generates the solution to part two of day 5 of AOC2017
func PartTwo(input string) (result int) {
	nums := utils.NumberListToSlice(input)
	pos, steps := 0, 0
	for pos < len(nums) {
		temp := pos
		pos += nums[pos]
		if nums[temp] >= 3 {
			nums[temp]--
		} else {
			nums[temp]++
		}
		steps++
	}
	return steps
}
