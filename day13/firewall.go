package main

import (
	"bytes"
)

type Firewall []*Layer

func (this Firewall) Reset() {
	for _, layer := range this {
		layer.Reset()
	}
}

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

func (this Firewall) traverseUntilCaught() int {
	for at := 0; at < len(this); at++ {
		if this[at].Alarm() {
			return at
		}
		this.Scan()
	}
	return -1
}
func (this Firewall) traverse() (cost int) {
	caught := false
	for at := 0; at < len(this); at++ {
		cost += this[at].Severity(at)
		caught = caught || this[at].Alarm()
		//fmt.Println("Cumulative cost at", at, cost)
		//fmt.Println(this.Draw(at))
		this.Scan()
		//fmt.Print("<ENTER> to continue...")
		//fmt.Scanln()
		//fmt.Println()
	}
	if cost == 0 && caught {
		cost++
	}
	return cost
}

func (this Firewall) Draw(at int) string {
	buffer := new(bytes.Buffer)
	for x := 0; x < len(this); x++ {
		if this[x] == nil {
			if x == at {
				buffer.WriteString("( )")
			}
			buffer.WriteString("\n")
			continue
		}
		for y := 0; y < this[x].Depth; y++ {
			if this[x].position == y {
				if x == at && y == 0 {
					buffer.WriteString("(x)")
				} else {
					buffer.WriteString(" x ")
				}
			} else {
				if x == at && y == 0 {
					buffer.WriteString("(.)")
				} else {
					buffer.WriteString(" . ")
				}
			}
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("----------------------------------------------")
	return buffer.String()
}
