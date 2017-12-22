package main

import (
	"encoding/hex"
	"fmt"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(), should.Equal, 6952))
	fmt.Println(assert.So(part2(), should.Equal, "28e7c4360520718a5dc811d3942cf1fd"))
	// fmt.Println(assert.So(KnotHashString(""), should.Equal, "a2582a3a0e66e6e86e3812dcb672a272"))
	// fmt.Println(assert.So(KnotHashString("AoC 2017"), should.Equal, "33efeb34ea91902bb2f59c9920caa6cd"))
	// fmt.Println(assert.So(KnotHashString("1,2,3"), should.Equal, "3efbe78a8d82f29979031a4aa0b16a9d"))
	// fmt.Println(assert.So(KnotHashString("1,2,4"), should.Equal, "63960835bcdc130f0b66d7ff4f6a5a8e"))
}

func part1() int {
	lengths := util.InputInts(",")
	c := NewCircle()
	c.ExecuteRound(lengths)
	return int(c.list[0]) * int(c.list[1])
}

func part2() string {
	return KnotHashString(util.InputString())
}

func KnotHashString(input string) string {
	return hex.EncodeToString(KnotHash(toInts(input)))
}

func toInts(input string) (ints []int) {
	for _, c := range input {
		ints = append(ints, int(c))
	}
	return ints
}

func KnotHash(input []int) []byte {
	var salt = []int{17, 31, 73, 47, 23}
	lengths := append(input, salt...)
	c := NewCircle()
	for x := 0; x < 64; x++ {
		c.ExecuteRound(lengths)
	}
	return c.Digest()
}

type Circle struct {
	list []byte
	i    int
	skip int
}

func NewCircle() *Circle {
	var ring []byte
	for x := 0; x < 256; x++ {
		ring = append(ring, byte(x))
	}
	return &Circle{list: ring}
}

func (this *Circle) ExecuteRound(lengths []int) {
	for _, length := range lengths {
		this.Replace(this.ReverseSlice(length))
		this.Advance(length)
	}
}

func (this *Circle) ReverseSlice(length int) (slice []byte) {
	for x := length - 1; x >= 0; x-- {
		slice = append(slice, this.list[this.offset(x)])
	}
	return slice
}

func (this *Circle) Replace(slice []byte) {
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

func (this *Circle) Digest() (digest []byte) {
	for x := 0; x < 16; x++ {
		var sum byte
		for y := 0; y < 16; y++ {
			sum ^= this.list[x*16+y]
		}
		digest = append(digest, sum)
	}
	return digest
}
