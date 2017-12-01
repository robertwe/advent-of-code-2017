package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(), should.Equal, 1251))
	fmt.Println(assert.So(part2(), should.Equal, 1244))
}

func part1() (sum int) {
	input := util.InputString()
	input = input + input[len(input)-1:] // wrap-around

	for x := 0; x < len(input)-1; x++ {
		c1 := input[x : x+1]
		c2 := input[x+1 : x+2]

		if c1 == c2 {
			sum += util.ParseInt(c2)
		}
	}

	return sum
}

func part2() (sum int) {
	input := util.InputString()
	length := len(input)
	half := length / 2

	for x := 0; x < len(input)-1; x++ {
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
