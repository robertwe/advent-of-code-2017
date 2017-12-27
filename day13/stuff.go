package main

type Layer struct {
	Depth int
	increment int
	position int
}

func NewLayer(depth int) *Layer {
	return &Layer{
		Depth: depth,
		increment: 1,
	}
}

func (this *Layer) Scan() int {
	defer this.travel()
	return this.position
}

func (this *Layer) travel() {
	this.position += this.increment
	if this.position == this.Depth - 1 {
		this.increment = -this.increment
	} else if this.position == 0 {
		this.increment = -this.increment
	}
}