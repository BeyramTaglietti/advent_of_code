package day7

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseLine(line string) (int, []int) {
	n1, _ := strconv.Atoi(strings.Split(line, ":")[0])
	n2s := strings.Split(strings.Split(line, ":")[1], " ")

	var n2 []int
	for _, n := range n2s {
		if n != "" {
			n2i, _ := strconv.Atoi(n)
			n2 = append(n2, n2i)
		}
	}

	return n1, n2
}

func SolveP1() {
	lines := utils.ReadFile("./day7/input.txt")

	rules := make(map[int][]int)

	for _, line := range lines {
		n1, n2 := parseLine(line)
		rules[n1] = n2
	}

	sum := 0
	for k, v := range rules {
		combinations := getAllCombinations(v[0:1], v[1:])
		if isValid(combinations, k) {
			sum += k
		}
	}

	fmt.Println("sum of valid:", sum)
}

func getAllCombinations(startingValues, endingValues []int) []int {
	if len(endingValues) == 0 {
		return startingValues
	}

	var newValues []int
	for _, val := range startingValues {
		newValues = append(newValues, val+endingValues[0])
		newValues = append(newValues, val*endingValues[0])
	}

	return getAllCombinations(newValues, endingValues[1:])
}

func isValid(combos []int, target int) bool {
	return slices.Contains(combos, target)
}
