package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/grid"
	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(), should.Equal, 438))
	fmt.Println(assert.So(part2(), should.Equal, 266330))
}

func part1() int {
	input := util.InputInt()
	turtle := Start()
	for {
		turtle.Advance()
		if turtle.id == input {
			at := grid.NewPoint(float64(turtle.at.x), float64(turtle.at.y))
			return int(grid.CityBlockDistance(at, grid.NewPoint(0, 0)))
		}
	}
	return 438 // Solved using pen and paper! (TODO: solve using the turtle and manhattan distance)
}

func part2() int {
	turtle := Start()
	for {
		turtle.Advance()
		written := turtle.Sum()
		if written > util.InputInt() {
			return written
		}
	}
	panic(":(")
}

type Turtle struct {
	at Point
	id int

	distance  int
	traveled  int
	direction Direction

	visited map[Point]int
}

func Start() *Turtle {
	return &Turtle{
		at: Point{0, 0},
		id: 1,

		distance:  2,
		traveled:  0,
		direction: Right,

		visited: map[Point]int{},
	}
}

func (this *Turtle) Sum() (sum int) {
	for _, d := range Neighbors8 {
		sum += this.visited[this.at.Move(d)]
	}
	if sum == 0 {
		sum++
	}
	return sum
}

func (this *Turtle) Advance() {
	at := this.at
	sum := this.Sum()
	this.visited[at] = sum

	if this.traveled == this.distance/2 {
		this.direction = travels[this.direction]
	} else if this.traveled == this.distance {
		this.traveled = 0
		this.distance += 2
		this.direction = travels[this.direction]
	}
	this.traveled++

	next := at.Move(this.direction)
	//log.Printf("Advancing from %v to %v (sum: %d)", at, next, sum)

	this.at = next
	this.id++
}

var travels = map[Direction]Direction{
	Down:  Right,
	Right: Up,
	Up:    Left,
	Left:  Down,
}

/////////////////////////////////////////

// TODO: extract all of this to grid package

type Point struct{ x, y int }

func (this Point) Move(d Direction) Point {
	return Point{
		x: this.x + d.dx,
		y: this.y + d.dy,
	}
}

type Direction struct{ dx, dy int }

var (
	Right = Direction{dx: 1, dy: 0}
	Left  = Direction{dx: -1, dy: 0}
	Up    = Direction{dx: 0, dy: 1}
	Down  = Direction{dx: 0, dy: -1}

	TopRight    = Direction{1, 1}
	TopLeft     = Direction{-1, 1}
	BottomRight = Direction{1, -1}
	BottomLeft  = Direction{-1, -1}

	Neighbors4 = []Direction{Right, Left, Up, Down}
	Neighbors8 = append(Neighbors4, []Direction{TopRight, TopLeft, BottomRight, BottomLeft}...)
)
