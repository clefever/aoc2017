package day7

import (
	"strconv"
	"strings"

	"github.com/clefever/aoc2017/utils"
)

type Program struct {
	Name         string
	Value        int
	ProgramNames []string
	Programs     []*Program
}

func (p Program) getFullWeight() int {
	if len(p.Programs) == 0 {
		return p.Value
	}
	return sumAllSubWeights(&p)
}

// PartOne generates the solution to part one of day 7 of AOC2017
func PartOne(input string) (result string) {
	programs := utils.StringsByLineToSlice(input)
	programSlice := []*Program{}
	for _, program := range programs {
		name := program[0]
		valueStr := program[1]
		valueStr = strings.Replace(valueStr, "(", "", -1)
		valueStr = strings.Replace(valueStr, ")", "", -1)
		value, _ := strconv.Atoi(valueStr)
		names := []string{}
		if len(program) > 3 {
			for _, s := range program[3:] {
				n := strings.Replace(s, ",", "", -1)
				names = append(names, n)
			}
		}
		programSlice = append(programSlice, &Program{name, value, names, nil})
	}
	allPrograms := make([]*Program, len(programSlice))
	copy(allPrograms, programSlice)
	for _, program := range programSlice {
		for _, p := range program.ProgramNames {
			pRef := findProgramByName(programSlice, p)
			allPrograms = removeFromSlice(allPrograms, pRef)
			program.Programs = append(program.Programs, pRef)
		}
	}
	return allPrograms[0].Name
}

// PartTwo generates the solution to part two of day 7 of AOC2017
func PartTwo(input string) (result int) {
	programs := utils.StringsByLineToSlice(input)
	programSlice := []*Program{}
	for _, program := range programs {
		name := program[0]
		valueStr := program[1]
		valueStr = strings.Replace(valueStr, "(", "", -1)
		valueStr = strings.Replace(valueStr, ")", "", -1)
		value, _ := strconv.Atoi(valueStr)
		names := []string{}
		if len(program) > 3 {
			for _, s := range program[3:] {
				n := strings.Replace(s, ",", "", -1)
				names = append(names, n)
			}
		}
		programSlice = append(programSlice, &Program{name, value, names, nil})
	}
	allPrograms := make([]*Program, len(programSlice))
	copy(allPrograms, programSlice)
	for _, program := range programSlice {
		for _, p := range program.ProgramNames {
			pRef := findProgramByName(programSlice, p)
			allPrograms = removeFromSlice(allPrograms, pRef)
			program.Programs = append(program.Programs, pRef)
		}
	}
	wrong := findWrongProgram(allPrograms[0])
	siblings := findAllSiblings(programSlice, wrong)
	for _, prog := range siblings {
		if prog != wrong {
			result = wrong.Value - (wrong.getFullWeight() - prog.getFullWeight())
		}
	}
	return result
}

func findProgramByName(programs []*Program, name string) *Program {
	for _, p := range programs {
		if p.Name == name {
			return p
		}
	}
	return nil
}

func findAllSiblings(programs []*Program, toFind *Program) []*Program {
	for _, p := range programs {
		for _, prog := range p.Programs {
			if prog == toFind {
				return p.Programs
			}
		}
	}
	return nil
}

func removeFromSlice(slice []*Program, program *Program) []*Program {
	for i, p := range slice {
		if p == program {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func sumAllSubWeights(p *Program) int {
	if p == nil {
		return 0
	}
	sum := 0
	for _, prog := range p.Programs {
		sum += sumAllSubWeights(prog)
	}
	return sum + p.Value
}

func findWrongProgram(p *Program) *Program {
	if p == nil {
		return nil
	}
	m := map[int]int{}
	for _, prog := range p.Programs {
		m[prog.getFullWeight()]++
	}
	hasUnequal := false
	for _, prog := range p.Programs {
		if m[prog.getFullWeight()] == 1 {
			hasUnequal = true
			return findWrongProgram(prog)
		}
	}
	if !hasUnequal {
		return p
	}
	return nil
}
