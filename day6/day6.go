package day6

import (
	"math"

	"github.com/clefever/aoc2017/utils"
)

// PartOne generates the solution to part one of day 6 of AOC2017
func PartOne(input string) (result int) {
	line := utils.NumberListToSlice(input)
	seen := [][]int{}
	seen = appendSliceCopy(seen, line)
	wasSeen := false
	cycles := 0
	for !wasSeen {
		index := getLargestIndex(line)
		redist := line[index]
		line[index] = 0
		index = (index + 1) % len(line)
		redistribute(line, redist, index)
		cycles++
		wasSeen = containsSlice(seen, line)
		seen = appendSliceCopy(seen, line)
	}
	return cycles
}

// PartTwo generates the solution to part two of day 6 of AOC2017
func PartTwo(input string) (result int) {
	line := utils.NumberListToSlice(input)
	seen := [][]int{}
	seen = appendSliceCopy(seen, line)
	seenCount, cycles := 0, 0
	for seenCount < 2 {
		index := getLargestIndex(line)
		redist := line[index]
		line[index] = 0
		index = (index + 1) % len(line)
		redistribute(line, redist, index)
		cycles++
		if containsSlice(seen, line) {
			seenCount++
			seen = nil
			if seenCount < 2 {
				cycles = 0
			}
		}
		seen = appendSliceCopy(seen, line)
	}
	return cycles
}

func appendSliceCopy(slices [][]int, slice []int) [][]int {
	temp := make([]int, len(slice))
	copy(temp, slice)
	return append(slices, temp)
}

func getLargestIndex(slice []int) int {
	max, index := math.MinInt32, 0
	for i, val := range slice {
		if val > max {
			max = val
			index = i
		}
	}
	return index
}

func redistribute(slice []int, amount int, startIndex int) {
	loc := startIndex
	for i := 0; i < amount; i++ {
		slice[loc]++
		loc = (loc + 1) % len(slice)
	}
}

func containsSlice(slices [][]int, slice []int) bool {
	for _, s := range slices {
		if utils.SlicesEqual(s, slice) {
			return true
		}
	}
	return false
}
