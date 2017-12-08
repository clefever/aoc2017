package day8

import (
	"math"
	"strconv"

	"github.com/clefever/aoc2017/utils"
)

// PartOne generates the solution to part one of day 8 of AOC2017
func PartOne(input string) (result int) {
	lines := utils.StringsByLineToSlice(input)
	registers := map[string]int{}
	for _, line := range lines {
		if checkCondition(registers, line) {
			writeRegister(registers, line)
		}
	}
	max := math.MinInt32
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return max
}

// PartTwo generates the solution to part two of day 8 of AOC2017
func PartTwo(input string) (result int) {
	lines := utils.StringsByLineToSlice(input)
	registers := map[string]int{}
	maxEver := math.MinInt32
	for _, line := range lines {
		if checkCondition(registers, line) {
			writeRegister(registers, line)
			for _, v := range registers {
				if v > maxEver {
					maxEver = v
				}
			}
		}
	}
	return maxEver
}

func checkCondition(m map[string]int, line []string) bool {
	num, _ := strconv.Atoi(line[6])
	switch line[5] {
	case ">":
		return m[line[4]] > num
	case "<":
		return m[line[4]] < num
	case ">=":
		return m[line[4]] >= num
	case "<=":
		return m[line[4]] <= num
	case "==":
		return m[line[4]] == num
	case "!=":
		return m[line[4]] != num
	}
	return false
}

func writeRegister(m map[string]int, line []string) {
	num, _ := strconv.Atoi(line[2])
	switch line[1] {
	case "inc":
		m[line[0]] += num
	case "dec":
		m[line[0]] -= num
	}
}
