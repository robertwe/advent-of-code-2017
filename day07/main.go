package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

const test = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func main() {
	tower := NewTower()
	scanner := util.InputScanner()
	for scanner.Scan() {
		tower.AddProgram(scanner.Text())
	}
	bottom := tower.FindBottom()
	node := tower.listing[bottom]
	diff, value := node.FindUnbalance()

	fmt.Println(assert.So(bottom, should.Equal, "dgoocsw"))
	fmt.Println(assert.So(diff+value, should.Equal, 1275))
	fmt.Println(util.ElapsedTime())
}
