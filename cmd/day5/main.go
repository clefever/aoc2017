package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/clefever/aoc2017/day5"
)

func main() {
	buf, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(buf)
	fmt.Printf("Answer to part one: %v\n", day5.PartOne(str))
	fmt.Printf("Answer to part two: %v\n", day5.PartTwo(str))
}
