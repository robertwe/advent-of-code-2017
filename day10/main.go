package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
)

func main() {
	var ring []int
	for x := 0; x < 256; x++ {
		ring = append(ring, x)
	}
	lengths := util.InputInts(",")
	c := NewCircle(ring)
	for _, length := range lengths {
		c.Replace(c.ReverseSlice(length))
		c.Advance(length)
	}
	fmt.Println(ring[0] * ring[1]) // part 1 answer: 6952
}

type Circle struct {
	list []int
	i    int
	skip int
}

func NewCircle(list []int) *Circle {
	return &Circle{list: list}
}

func (this *Circle) ReverseSlice(length int) (slice []int) {
	for x := length - 1; x >= 0; x-- {
		slice = append(slice, this.list[this.offset(x)])
	}
	return slice
}

func (this *Circle) Replace(slice []int) {
	for x := 0; x < len(slice); x++ {
		i := this.offset(x)
		this.list[i] = slice[x]
	}
}

func (this *Circle) Advance(length int) {
	this.i = this.offset(length + this.skip)
	this.skip++
}

func (this *Circle) safe(i int) int {
	return i % len(this.list)
}
func (this *Circle) offset(i int) int {
	return this.safe(this.i + i)
}
