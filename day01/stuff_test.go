package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
}

func (this *Stuff) TestNeighborSums() {
	this.So(sumOfIdenticalNeighborDigits("1122"), should.Equal, 3)
	this.So(sumOfIdenticalNeighborDigits("1111"), should.Equal, 4)
	this.So(sumOfIdenticalNeighborDigits("1234"), should.Equal, 0)
	this.So(sumOfIdenticalNeighborDigits("91212129"), should.Equal, 9)
}

func (this *Stuff) TestOppositeSums() {
	this.So(sumOfIdenticalOppositeDigits("1212"), should.Equal, 6)
	this.So(sumOfIdenticalOppositeDigits("1221"), should.Equal, 0)
	this.So(sumOfIdenticalOppositeDigits("123425"), should.Equal, 4)
	this.So(sumOfIdenticalOppositeDigits("123123"), should.Equal, 12)
	this.So(sumOfIdenticalOppositeDigits("12131415"), should.Equal, 4)
}
