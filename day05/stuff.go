package main

type Program struct {
	part2   bool
	counter int
	steps   int
	ints    []int
}

func (this *Program) Execute() {
	for this.inRange() {
		c := this.counter

		this.steps++

		current := this.ints[this.counter]
		this.counter += current

		if c >= len(this.ints) || c < 0 {
			break
		}
		this.increment(c)
	}
}

func (this *Program) inRange() bool {
	return this.counter < len(this.ints) && this.counter >= 0
}

func (this *Program) increment(c int) {
	if this.part2 && this.ints[c] >= 3 {
		this.ints[c]--
	} else {
		this.ints[c]++
	}
}
