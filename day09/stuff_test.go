package main

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
}

func (this *Stuff) Setup() {
}

func (this *Stuff) Test() {
}
