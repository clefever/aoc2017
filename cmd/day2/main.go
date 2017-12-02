package main

import (
	"aoc2017/day2"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	buf, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(buf)
	fmt.Printf("Answer to part one: %v\n", day2.PartOne(str))
	fmt.Printf("Answer to part two: %v\n", day2.PartTwo(str))
}
