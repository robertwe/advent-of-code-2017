package main

import (
	"github.com/mdwhatcott/advent-of-code/util"
)

func sumOfIdenticalNeighborDigits(input string) (sum int) {
	input = input + input[0:1] // wrap-around

	for x := 0; x < len(input)-1; x++ {
		c1 := input[x : x+1]
		c2 := input[x+1 : x+2]

		if c1 == c2 {
			sum += util.ParseInt(c2)
		}
	}

	return sum
}

func sumOfIdenticalOppositeDigits(input string) (sum int) {
	length := len(input)
	half := length / 2

	for x := 0; x < len(input); x++ {
		c1 := input[x : x+1]
		x2 := x + half
		if x2 >= length {
			x2 -= length // wrap-around
		}
		c2 := input[x2 : x2+1]

		if c1 == c2 {
			sum += util.ParseInt(c2)
		}
	}

	return sum
}
