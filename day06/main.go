package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	d := NewDebugger(util.ParseInts(strings.Fields(util.InputString())))
	//d := NewDebugger([]int{0, 2, 7, 0})

	part1 := d.Debug() + 1

	d.states = make(map[string]struct{})
	d.states[fmt.Sprint(d.registers)] = struct{}{}
	part2 := d.Debug()

	fmt.Println(assert.So(part1, should.Equal, 5042))
	fmt.Println(assert.So(part2, should.Equal, 1086))
}

type Debugger struct {
	registers []int
	cursor    int
	states    map[string]struct{}
}

func NewDebugger(registers []int) *Debugger {
	return &Debugger{
		registers: registers,
		states:    make(map[string]struct{}),
	}
}

func (this *Debugger) Debug() int {
	for {
		this.redistributeLargestBlock()
		if this.alreadySeenCurrentRegisterState() {
			break
		}
	}

	return len(this.states)
}
func (this *Debugger) redistributeLargestBlock() {
	largest := this.findLargestRegister()
	value := this.registers[largest]
	this.registers[largest] = 0
	this.cursor = largest
	for ; value > 0; value-- {
		this.cursor++
		if this.cursor >= len(this.registers) {
			this.cursor = 0
		}
		this.registers[this.cursor]++
	}
}
func (this *Debugger) findLargestRegister() int {
	max := -1
	index := -1
	for i, value := range this.registers {
		if value > max {
			index = i
			max = value
		}
	}
	return index
}
func (this *Debugger) alreadySeenCurrentRegisterState() bool {
	state := fmt.Sprint(this.registers)
	_, found := this.states[state]
	if !found {
		this.states[state] = struct{}{}
	}
	return found
}
