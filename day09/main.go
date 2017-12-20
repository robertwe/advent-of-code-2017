package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println("Group score:        ", assert.So(groupScore(util.InputString()), should.Equal, 14212))
	fmt.Println("Garbage characters: ", assert.So(garbageScore(util.InputString()), should.Equal, 6569))
}
