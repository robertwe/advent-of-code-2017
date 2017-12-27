package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(loadFromInput().traverse(), should.Equal, 788))
	fmt.Println(assert.So(lowestSafeDelay(loadFromInput), should.Equal, 0))
}

func loadFromInput() Firewall {
	return loadFirewall(util.InputScanner())
}

func lowestSafeDelay(loader func() Firewall) int {
	for x := 0; ; x++ {
		firewall := loader()
		firewall.delay(x)
		if cost := firewall.traverse(); cost == 0 {
			return x
		}
	}
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
