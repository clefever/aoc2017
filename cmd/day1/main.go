package main

import (
	"aoc2017/day1"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	str := strings.TrimSpace(string(buf))
	fmt.Printf("Answer to part one: %v\n", day1.PartOne(str))
	fmt.Printf("Answer to part two: %v\n", day1.PartTwo(str))
}
