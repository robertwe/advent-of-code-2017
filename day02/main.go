package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	input := util.InputLines()

	fmt.Println(assert.So(part1(input), should.Equal, 34925))
	fmt.Println(assert.So(part2(input), should.Equal, 221))
}
