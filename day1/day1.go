package day1

import "strconv"

// PartOne generates the solution to part one of day 1 of AOC2017
func PartOne(input string) (result int) {
	for i := range input {
		if i == len(input)-1 {
			if input[i] == input[0] {
				val, _ := strconv.Atoi(string(input[i]))
				result += val
			}
		} else {
			if input[i] == input[i+1] {
				val, _ := strconv.Atoi(string(input[i]))
				result += val
			}
		}
	}
	return result
}

// PartTwo generates the solution to part two of day 1 of AOC2017
func PartTwo(input string) (result int) {
	for i := range input {
		if i < len(input)/2 {
			if input[i] == input[i+(len(input)/2)] {
				val, _ := strconv.Atoi(string(input[i]))
				result += val
			}
		}
	}
	return result * 2
}
