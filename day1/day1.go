package day1

import "github.com/clefever/aoc2017/utils"

// PartOne generates the solution to part one of day 1 of AOC2017
func PartOne(input string) (result int) {
	nums := utils.NumberStringToElements(input)
	length := len(nums)
	for i, val := range nums {
		if i == length-1 {
			if nums[i] == nums[0] {
				result += val
			}
		} else {
			if nums[i] == nums[i+1] {
				result += val
			}
		}
	}
	return result
}

// PartTwo generates the solution to part two of day 1 of AOC2017
func PartTwo(input string) (result int) {
	nums := utils.NumberStringToElements(input)
	length := len(nums)
	for i, val := range nums {
		if i < length/2 {
			if input[i] == input[i+length/2] {
				result += val
			}
		}
	}
	return result * 2
}
