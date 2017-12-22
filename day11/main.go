package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	var maxDistance int
	var start Hex
	var position Hex

	path := strings.Split(util.InputString(), ",")

	for _, direction := range path {
		position = position.Offset(Offsets[direction])
		if distance := position.DistanceTo(start); distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println(assert.So(position.DistanceTo(start), should.Equal, 707))
	fmt.Println(assert.So(maxDistance, should.Equal, 1490))
}

var Offsets = map[string]Hex{
	"n":  North,
	"ne": NorthEast,
	"nw": NorthWest,
	"s":  South,
	"se": SouthEast,
	"sw": SouthWest,
}
