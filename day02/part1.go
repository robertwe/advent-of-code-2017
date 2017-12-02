package main

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
)

func part1(lines []string) int {
	checksum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		max := maxOfAll(fields)
		min := minOfAll(fields)
		diff := max - min
		checksum += diff
	}
	return checksum
}
func minOfAll(i []string) int {
	min := 0xffffff
	for _, x := range i {
		y := util.ParseInt(x)
		if y < min {
			min = y
		}
	}
	return min
}
func maxOfAll(i []string) int {
	max := 0
	for _, x := range i {
		y := util.ParseInt(x)
		if y > max {
			max = y
		}
	}
	return max
}
