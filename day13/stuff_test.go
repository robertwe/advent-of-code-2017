package main

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code/util"
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
	this.l = NewLayer(3, 0)
	this.assertOscillations(0, 1, 2, 1, 0, 1, 2, 1, 0)
}

func (this *Stuff) assertOscillations(all ...int) {
	for _, expected := range all {
		this.So(this.l.Scan(), should.Equal, expected)
	}
}

func (this *Stuff) TestSeverityOfScan() {
	this.l = NewLayer(3, 0)
	this.So(this.l.Severity(3), should.Equal, 9)
	this.l.Scan()
	this.So(this.l.Severity(3), should.Equal, 0)
}

func (this *Stuff) TestSample() {
	const input = `0: 3
1: 2
4: 4
6: 4
`
	scanner := util.NewScanner(strings.NewReader(input))
	this.So(loadFirewall(scanner).traverse(), should.Equal, 24)
}

func (this *Stuff) TestSampleWithDelay() {
	const input = `0: 3
1: 2
4: 4
6: 4
`

	loader := func() Firewall {
		scanner := util.NewScanner(strings.NewReader(input))
		return loadFirewall(scanner)
	}
	this.So(lowestSafeDelay(loader), should.Equal, 10)
}
