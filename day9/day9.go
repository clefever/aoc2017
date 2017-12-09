package day9

import "github.com/clefever/aoc2017/utils"

// PartOne generates the solution to part one of day 9 of AOC2017
func PartOne(input string) (result int) {
	parse := utils.StringToElements(input)
	score, garbageLayer, groupLayer := 0, 0, 0
	for i := 0; i < len(parse); i++ {
		if parse[i] == "!" && garbageLayer > 0 {
			i++
		} else if parse[i] == "<" && garbageLayer < 1 {
			garbageLayer++
		} else if parse[i] == ">" && garbageLayer > 0 {
			garbageLayer--
		} else if parse[i] == "{" && garbageLayer < 1 {
			groupLayer++
		} else if parse[i] == "}" && groupLayer > 0 && garbageLayer < 1 {
			groupLayer--
		}
		if parse[i] == "{" && groupLayer > 0 && garbageLayer < 1 {
			score += groupLayer
		}
	}
	return score
}

// PartTwo generates the solution to part two of day 9 of AOC2017
func PartTwo(input string) (result int) {
	parse := utils.StringToElements(input)
	garbageLayer, groupLayer, garbageCount := 0, 0, 0
	for i := 0; i < len(parse); i++ {
		if parse[i] == "!" && garbageLayer > 0 {
			i++
		} else if parse[i] == "<" && garbageLayer < 1 {
			garbageLayer++
		} else if parse[i] == ">" && garbageLayer > 0 {
			garbageLayer--
		} else if parse[i] == "{" && garbageLayer < 1 {
			groupLayer++
		} else if parse[i] == "}" && groupLayer > 0 && garbageLayer < 1 {
			groupLayer--
		} else if garbageLayer > 0 {
			garbageCount++
		}
	}
	return garbageCount
}
