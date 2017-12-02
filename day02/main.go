package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(util.InputLines()), should.Equal, 34925))
	fmt.Println(assert.So(part2(util.InputLines()), should.Equal, 221))
}
