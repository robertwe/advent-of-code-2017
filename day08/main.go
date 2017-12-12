package main

import (
	"fmt"
	"log"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

var maxes []int
var registers = make(map[string]int)

func main() {
	scanner := util.InputScanner()
	for scanner.Scan() {
		fields := scanner.Fields()
		if len(fields) == 0 {
			continue
		}

		var test func(register string, value int) bool
		switch fields[5] {
		case "==":
			test = equal
		case "!=":
			test = notEqual
		case "<":
			test = less
		case ">":
			test = more
		case "<=":
			test = lessEqual
		case ">=":
			test = moreEqual
		default:
			log.Panic("Test not found:", fields[5])
		}

		if !test(fields[4], util.ParseInt(fields[6])) {
			continue
		}

		var op func(string, int)
		switch fields[1] {
		case "inc":
			op = increment
		case "dec":
			op = decrement
		default:
			log.Panic("Operation not found:", fields[1])
		}
		op(fields[0], util.ParseInt(fields[2]))
		maxes = append(maxes, max())
	}

	fmt.Println(assert.So(max(), should.Equal, 4902))
	fmt.Println(assert.So(util.Max(maxes...), should.Equal, 7037))
}

func max() (m int) {
	for _, value := range registers {
		if value > m {
			m = value
		}
	}
	return m
}

func increment(register string, value int) { registers[register] += value }
func decrement(register string, value int) { registers[register] -= value }

func less(register string, value int) bool      { return registers[register] < value }
func more(register string, value int) bool      { return registers[register] > value }
func equal(register string, value int) bool     { return !less(register, value) && !more(register, value) }
func notEqual(register string, value int) bool  { return !equal(register, value) }
func lessEqual(register string, value int) bool { return !more(register, value) }
func moreEqual(register string, value int) bool { return !less(register, value) }
