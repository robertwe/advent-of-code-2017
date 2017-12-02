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
		min, max := util.MinMax(ints...)
		diff := max - min
		checksum += diff
	}
	return checksum
}
