package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	//fmt.Println(assert.So(loadFromInput().traverse(), should.Equal, 788))
	fmt.Println(assert.So(lowestSafeDelay(loadFromInput), should.Equal, 0))
	//fmt.Println(lowestSafeDelay(loadTestInput))
}

func loadTestInput() Firewall {
	scanner := util.NewScanner(strings.NewReader("0: 3 \n 1: 2 \n 4: 4 \n 6: 4"))
	return loadFirewall(scanner)
}

var input = util.InputBytes()

func loadFromInput() Firewall {
	return loadFirewall(util.NewScanner(bytes.NewReader(input)))
}

// TODO: Learn from: https://github.com/lukaszroz/advent-of-code-2017/blob/master/day13.go
func lowestSafeDelay(loader func() Firewall) int {
	firewall := loader()
	for delay := 0; ; delay++ {
		if delay % 10000 == 0 {
			fmt.Println("Delay:", delay)
		}
		if delay > 10000 {
			//break
		}
		if delay%2 > 0 {
			continue
		}
		if (delay-2)%4 == 0 {
			continue
		}
		if delay%8 == 0 {
			continue
		}
		if (delay-20)%24 == 0 {
			continue
		}
		if (delay-28)%24 == 0 {
			continue
		}
		if (delay-60)%120 == 0 {
			continue
		}
		if (delay-84)%120 == 0 {
			continue
		}
		//fmt.Println("delay:", delay)
		firewall.Reset()
		firewall.delay(delay)

		//fmt.Println("About to traverse:")
		//fmt.Println(firewall.Draw(-1))
		//fmt.Print("<ENTER> to continue...")
		//fmt.Scanln()
		//fmt.Println()

		at := firewall.traverseUntilCaught()
		//fmt.Println("cost:", cost)
		if at < 0 {
			return delay
		} else {
			//fmt.Println("Caught at", at, "with delay of", delay)
		}
	}
	return -1
}

func loadFirewall(scanner *util.Scanner) (firewall Firewall) {
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ":", "", 1)
		fields := strings.Fields(line)
		level := util.ParseInt(fields[0])
		depth := util.ParseInt(fields[1])
		for level > len(firewall) {
			firewall = append(firewall, nil)
		}
		firewall = append(firewall, NewLayer(depth, len(firewall)))
	}
	return firewall
}
