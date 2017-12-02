package main

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
)

func part2(lines []string) int {
	checksum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		checksum += divider(fields)
	}
	return checksum
}

func divider(i []string) int {
	var nums []int
	for _, x := range i {
		nums = append(nums, util.ParseInt(x))
	}

	for _, a := range nums {
		for _, b := range nums {
			if a == b {
				continue
			}
			if a%b == 0 && a/b > 0 {
				return a / b
			}
		}
	}
	panic("BAD ROW")
}
