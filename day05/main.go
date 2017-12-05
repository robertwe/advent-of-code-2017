package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(), should.Equal, 351282))
	fmt.Println(assert.So(part2(), should.Equal, 24568703))
}

func part1() int {
	return NewProgram(util.ParseInts(util.InputLines())).Execute()
}

func part2() int {
	return NewProgram(util.ParseInts(util.InputLines())).Part2().Execute()
}
