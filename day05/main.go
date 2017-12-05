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
	program := &Program{ints: util.ParseInts(util.InputLines())}
	program.Execute()
	return program.steps
}

func part2() int {
	program := &Program{ints: util.ParseInts(util.InputLines())}
	program.part2 = true
	program.Execute()
	return program.steps
}
