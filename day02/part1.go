package main

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
)

func part1(lines []string) int {
	checksum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		ints := util.ParseInts(fields)
		max := maxOfAll(ints)
		min := minOfAll(ints)
		diff := max - min
		checksum += diff
	}
	return checksum
}

func minOfAll(i []int) int {
	min := 0xffffff
	for _, x := range i {
		if x < min {
			min = x
		}
	}
	return min
}

func maxOfAll(i []int) int {
	max := 0
	for _, y := range i {
		if y > max {
			max = y
		}
	}
	return max
}
