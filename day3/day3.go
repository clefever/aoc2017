package day3

import (
	"github.com/clefever/aoc2017/utilities"
)

type cell struct {
	X int
	Y int
}

func (c cell) getLength() int {
	return utilities.Max(utilities.Abs(c.X), utilities.Abs(c.Y))
}

// PartOne generates the solution to part one of day 3 of AOC2017
func PartOne(input int) (result int) {
	currCell := cell{0, 0}
	m := map[cell]int{currCell: 1}
	for i := 2; i <= input; i++ {
		currCell = getNextCell(currCell)
		m[currCell] = i
	}
	return utilities.Abs(currCell.X) + utilities.Abs(currCell.Y)
}

// PartTwo generates the solution to part two of day 3 of AOC2017
func PartTwo(input int) (result int) {
	currCell := cell{0, 0}
	m := map[cell]int{currCell: 1}
	for i := 2; i <= input; i++ {
		currCell = getNextCell(currCell)
		cellSum := sumAdjacentCells(m, currCell)
		m[currCell] = cellSum
		if cellSum > input {
			return cellSum
		}
	}
	return -1
}

func getNextCell(c cell) cell {
	if c.X == 0 && c.Y == 0 {
		return cell{1, 0}
	}
	diff := c.X - c.Y
	length := c.getLength()
	if diff > 0 && diff < length*2 && c.X == length {
		return cell{c.X, c.Y + 1}
	}
	if diff <= 0 && c.X != -length {
		return cell{c.X - 1, c.Y}
	}
	if diff < 0 && c.X == -length {
		return cell{c.X, c.Y - 1}
	}
	if diff >= 0 && c.Y == -length {
		return cell{c.X + 1, c.Y}
	}
	return c
}

func sumAdjacentCells(m map[cell]int, c cell) (result int) {
	for x := c.X - 1; x <= c.X+1; x++ {
		for y := c.Y - 1; y <= c.Y+1; y++ {
			if !(x == c.X && y == c.Y) {
				result += m[cell{x, y}]
			}
		}
	}
	return result
}
