package day7

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseLine(line string) (uint, []uint) {
	n1, _ := strconv.Atoi(strings.Split(line, ":")[0])
	n2s := strings.Split(strings.Split(line, ":")[1], " ")

	var n2 []uint
	for _, n := range n2s {
		if n != "" {
			n2i, _ := strconv.Atoi(n)
			n2 = append(n2, uint(n2i))
		}
	}

	return uint(n1), n2
}

func SolveP1() {
	lines := utils.ReadFile("./day7/input.txt")

	rules := make(map[uint][]uint)

	for _, line := range lines {
		n1, n2 := parseLine(line)
		rules[n1] = n2
	}

	var sum uint
	for k, v := range rules {
		combinations := getAllCombinations(v[0:1], v[1:], false)
		if isValid(combinations, k) {
			sum += k
		}
	}

	fmt.Println("sum of valid lines:", sum)
}

func SolveP2() {
	lines := utils.ReadFile("./day7/input.txt")

	rules := make(map[uint][]uint)

	for _, line := range lines {
		n1, n2 := parseLine(line)
		rules[n1] = n2
	}

	var sum uint
	for k, v := range rules {
		combinations := getAllCombinations(v[0:1], v[1:], true)
		if isValid(combinations, k) {
			sum += k
		}
	}

	fmt.Println("sum of valid lines:", sum)
}

func getAllCombinations(startingValues, endingValues []uint, includeConcatenation bool) []uint {
	if len(endingValues) == 0 {
		return startingValues
	}

	var newValues []uint
	for _, val := range startingValues {
		newValues = append(newValues, val+endingValues[0])
		newValues = append(newValues, val*endingValues[0])

		if includeConcatenation {
			numericVal, _ := strconv.Atoi(strconv.Itoa(int(val)) + strconv.Itoa(int(endingValues[0])))
			newValues = append(newValues, uint(numericVal))
		}
	}

	return getAllCombinations(newValues, endingValues[1:], includeConcatenation)
}

func isValid(combos []uint, target uint) bool {
	return slices.Contains(combos, target)
}
