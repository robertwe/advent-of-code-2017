package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
)

func main() {
	d := NewDebugger(util.ParseInts(strings.Fields(util.InputString())))
	//d := NewDebugger([]int{0, 2, 7, 0})
	part1 := d.Debug() + 1
	fmt.Println(part1)
	d.states = make(map[string]struct{})
	d.states[fmt.Sprint(d.registers)] = struct{}{}
	part2 := d.Debug()
	fmt.Println(part2)
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
		state := fmt.Sprint(this.registers)
		_, found := this.states[state]
		if found {
			return len(this.states)
		} else {
			this.states[state] = struct{}{}
		}
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
