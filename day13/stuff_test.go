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
	l *Layer
}

func (this *Stuff) TestOscillatingScanner() {
	this.l = NewLayer(3)
	this.assertOscillations(0,1,2,1,0,1,2,1,0)
}

func (this *Stuff) assertOscillations(all ...int) {
	for _, expected := range all {
		this.So(this.l.Scan(), should.Equal, expected)
	}
}
