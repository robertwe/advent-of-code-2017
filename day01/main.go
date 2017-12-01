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
	return sumOfIdenticalNeighborDigits(util.InputCharacters())
}

func part2() (sum int) {
	return sumOfIdenticalOppositeDigits(util.InputCharacters())
}
