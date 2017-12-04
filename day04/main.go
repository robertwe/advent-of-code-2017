package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	valid1 := 0
	valid2 := 0
	input := util.InputScanner()

	for input.Scan() {
		line := input.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if Valid1(line) {
			valid1++
		}
		if Valid(line) {
			valid2++
		}
	}

	fmt.Println(assert.So(valid1, should.Equal, 451))
	fmt.Println(assert.So(valid2, should.Equal, 223))
}
