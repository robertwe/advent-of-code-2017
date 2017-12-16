package main

import (
	"fmt"
)

func main() {
	generateMatches(40000000, NewGenerator(1, 16807, 65), NewGenerator(1, 48271, 8921)) // example
	generateMatches(40000000, NewGenerator(1, 16807, 277), NewGenerator(1, 48271, 349)) // part1
	generateMatches(5000000, NewGenerator(4, 16807, 277), NewGenerator(8, 48271, 349))  // part2
}
func generateMatches(iterations int, a, b *Generator) {
	count := 0
	for x := 0; x < iterations; x++ {
		a.Next()
		b.Next()
		if a.Mask() == b.Mask() {
			count++
		}
	}
	fmt.Println(count)
}

type Generator struct {
	divisor int
	factor  int
	Current int
}

func NewGenerator(divisor, factor, start int) *Generator {
	return &Generator{divisor: divisor, factor: factor, Current: start}
}

func (this *Generator) Next() {
	this.generate()
	for this.Current%this.divisor != 0 {
		this.generate()
	}
}

func (this *Generator) generate() {
	this.Current *= this.factor
	this.Current %= 2147483647
}

func (this *Generator) Mask() int {
	return this.Current &^ 0xFFFF0000
}
