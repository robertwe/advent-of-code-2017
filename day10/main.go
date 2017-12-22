package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(), should.Equal, 6952))
	fmt.Println(assert.So(part2(), should.Equal, "28e7c4360520718a5dc811d3942cf1fd"))
}

func part1() int {
	circle := NewLoop()
	circle.TwistAll(toBytes(util.InputInts(",")))
	return int(circle.list[0]) * int(circle.list[1])
}

func part2() string {
	return KnotHashString(util.InputString())
}

func toBytes(input []int) (bytes []byte) {
	for _, l := range input {
		bytes = append(bytes, byte(l))
	}
	return bytes
}
