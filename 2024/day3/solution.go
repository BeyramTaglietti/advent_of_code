package day3

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("./day3/input.txt")

	pattern := `mul\(d+,d+\)`

	total := 0
	for _, line := range lines {
		re, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatalf("Error compiling regex: %s", err.Error())
		}

		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			total += multiply(match)
		}

	}

	fmt.Println("total sum:", total)
}

func SolveP2() {
	lines := utils.ReadFile("./day3/input.txt")

	var fullPuzzle strings.Builder

	for _, line := range lines {
		fullPuzzle.WriteString(line)
	}

	line := fullPuzzle.String()

	matchesMap := make(matchMap)
	matchesMap.findMultiplications(line)
	matchesMap.findCmds(line)

	fmt.Println("total sum with rules:", multiplyWithRules(matchesMap))
}

type command string

const (
	do   command = "do"
	dont command = "don't"
)

type matchFound struct {
	isCmd                bool
	cmd                  command
	multiplicationResult int
}

type matchMap map[int]matchFound

func (m *matchMap) findMultiplications(line string) {
	pattern := `mul\(\d+,\d+\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Error compiling regex: %s", err.Error())

	}

	matches := re.FindAllString(line, -1)
	matchesIdxs := re.FindAllStringIndex(line, -1)

	for idx, match := range matches {
		(*m)[matchesIdxs[idx][0]] = matchFound{
			isCmd:                false,
			multiplicationResult: multiply(match),
		}
	}

}

func (m *matchMap) findCmds(line string) {
	pattern := `do`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Error compiling regex: %s", err.Error())
	}

	matches := re.FindAllString(line, -1)
	matchesIdxs := re.FindAllStringIndex(line, -1)

	for idx := range matches {
		matchesAtIdx := matchesIdxs[idx][0]
		fullWord := line[matchesAtIdx : matchesAtIdx+5]

		if fullWord == string(dont) {
			(*m)[matchesAtIdx] = matchFound{
				isCmd: true,
				cmd:   dont,
			}
		} else {
			(*m)[matchesAtIdx] = matchFound{
				isCmd: true,
				cmd:   do,
			}
		}
	}

}

func multiplyWithRules(matchesMap matchMap) (result int) {

	keys := make([]int, len(matchesMap), len(matchesMap))

	i := 0
	for k := range matchesMap {
		keys[i] = k
		i++
	}

	slices.Sort(keys)

	sum := 0
	can := true
	for _, key := range keys {

		match := matchesMap[key]

		switch match.isCmd {
		case true:
			switch match.cmd {
			case do:
				can = true
			case dont:
				can = false
			}

		case false:
			if can {
				sum += match.multiplicationResult
			}
		}

	}

	return sum
}

func multiply(op string) int {
	var x1, x2 int

	fmt.Sscanf(op, "mul(%d,%d)", &x1, &x2)

	return x1 * x2
}
