package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/util"
)

const test = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

type Node struct {
	name     string
	weight   int
	children []*Node
}

func (this *Node) Weight() int {
	weight := this.weight
	for _, child := range this.children {
		weight += child.Weight()
	}
	return weight
}

func (this *Node) AppendChild(child *Node) {
	this.children = append(this.children, child)
}

func main() {
	bottom := part1()
	part2(bottom)
}
func part2(bottom string) {
	programs := map[string]*Node{}
	scanner := util.InputScanner()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		name := fields[0]
		weight := util.ParseInt(strings.Trim(fields[1], "()"))
		programs[name] = &Node{
			name:   name,
			weight: weight,
		}
	}

	scanner = util.InputScanner()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		name := fields[0]

		if len(fields) <= 2 {
			continue
		}
		for _, child := range strings.Split(strings.Split(line, " -> ")[1], ", ") {
			programs[name].AppendChild(programs[child])
		}
	}

	current := programs[bottom]
	for {
		if len(current.children) == 0 {
			break
		}
		weights := map[int][]*Node{}
		for _, child := range current.children {
			weight := child.Weight()
			weights[weight] = append(weights[weight], child)
		}
		for _, nodes := range weights {
			if len(nodes) == 1 {
				current = nodes[0]
				fmt.Println(current.name, current.Weight())
				break
			}
		}
	}

}
func part1() string {
	programs := map[string]int{}
	deps := map[string]int{}
	scanner := util.InputScanner()
	//scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		program := fields[0]
		programs[program]++

		if len(fields) <= 2 {
			continue
		}
		allDeps := strings.Split(strings.Split(line, " -> ")[1], ", ")
		for _, field := range allDeps {
			deps[field]++
		}
	}

	for p := range programs {
		if deps[p] == 0 {
			fmt.Println("Part 1:", p)
			return p
		}
	}
	panic("NOT FOUND")
}
