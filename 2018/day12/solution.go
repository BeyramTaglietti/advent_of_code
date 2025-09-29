package day12

import (
	"2018/utils"
	"fmt"
	"strings"
)

type rule struct {
	current bool
	left    string
	right   string
}

func evolve(state string, rules map[rule]bool) string {
	b := make([]byte, len(state))
	for i := range b {
		b[i] = '.'
	}
	for i := 2; i <= len(state)-3; i++ {
		k := rule{left: state[i-2 : i], current: state[i] == '#', right: state[i+1 : i+3]}
		if rules[k] {
			b[i] = '#'
		}
	}
	return string(b)
}

func SolveP1() {
	lines := utils.ReadFile("day12/input.txt")
	initialState, rules := parseInput(lines)

	generations := 20
	pad := strings.Repeat(".", 4*generations+4)
	state := pad + initialState + pad

	for g := 0; g < generations; g++ {
		state = evolve(state, rules)
	}

	origin := len(pad)
	sum := 0
	for i := 0; i < len(state); i++ {
		if state[i] == '#' {
			sum += i - origin
		}
	}
	fmt.Println(sum)

}

func SolveP2() {

}

func parseInput(lines []string) (initialState string, rules map[rule]bool) {
	fmt.Sscanf(lines[0], "initial state: %s", &initialState)

	parseRule := func(line string) (string, bool) {
		var parsedRule string
		var parsedResult string
		fmt.Sscanf(line, "%s => %s", &parsedRule, &parsedResult)

		return parsedRule, parsedResult == "#"
	}

	rules = make(map[rule]bool)

	for _, line := range lines[2:] {
		parsedRule, parsedResult := parseRule(line)

		newRule := rule{
			current: string(parsedRule[2]) == "#",
			left:    parsedRule[0:2],
			right:   parsedRule[3:],
		}

		rules[newRule] = parsedResult
	}

	return initialState, rules
}
