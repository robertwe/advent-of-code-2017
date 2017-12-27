package main

import "fmt"

type Firewall []*Layer

func (this Firewall) Scan() {
	for _, layer := range this {
		layer.Scan()
	}
}

func (this Firewall) delay(ticks int) {
	for x := 0; x < ticks; x++ {
		this.Scan()
	}
}

func (firewall Firewall) traverse() (cost int) {
	for at := 0; at < len(firewall); at++ {
		cost += firewall[at].Severity(at)
		firewall.Scan()
	}
	return cost
}

type Layer struct {
	Depth     int
	layer     int
	increment int
	position  int
}

func NewLayer(depth int, layer int) *Layer {
	return &Layer{
		Depth:     depth,
		layer:     layer,
		increment: 1,
	}
}

func (this *Layer) Scan() int {
	if this == nil {
		return -1
	}
	defer this.travel()
	return this.position
}

func (this *Layer) travel() {
	this.position += this.increment
	if this.position == this.Depth-1 {
		this.increment = -this.increment
	} else if this.position == 0 {
		this.increment = -this.increment
	}
}

func (this *Layer) Severity(layer int) int {
	if this == nil {
		return 0
	}
	if this.position > 0 {
		return 0
	}
	return this.Depth * layer
}

func (this *Layer) String() string {
	return fmt.Sprintln("layer:", this.layer, "position:", this.position)
}
