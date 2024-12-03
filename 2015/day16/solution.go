package day16

import (
	"aoc2015/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolveP1() {

	lines := utils.ReadFile("./day16/input.txt")

	auntSue := []string{
		"children: 3",
		"cats: 7",
		"samoyeds: 2",
		"pomeranians: 3",
		"akitas: 0",
		"vizslas: 0",
		"goldfish: 5",
		"trees: 3",
		"cars: 2",
		"perfumes: 1",
	}

	for _, line := range lines {

		points := 0
		for _, charateristic := range auntSue {
			if strings.Contains(line, charateristic) {
				points++
			}
		}

		if points == 3 {
			fmt.Println(line)
		}

	}
}

func SolveP2() {

	lines := utils.ReadFile("./day16/input.txt")

	auntSue := []string{
		"children: 3",
		"samoyeds: 2",
		"akitas: 0",
		"vizslas: 0",
		"cars: 2",
		"perfumes: 1",
	}

	for _, line := range lines {

		points := 0

		hasRangeCharateristic, respectsRangeCharateristic := respectsRangeCharateristics(line)
		if hasRangeCharateristic {
			if respectsRangeCharateristic {
				points++
			}
		}

		for _, charateristic := range auntSue {
			if strings.Contains(line, charateristic) {
				points++
			}
		}

		if points == 3 {
			fmt.Println(line)
		}

	}
}

func respectsRangeCharateristics(line string) (contains, respects bool) {

	minRanges := map[string]int{
		"cats":  7,
		"trees": 3,
	}

	maxRanges := map[string]int{
		"pomeranians": 3,
		"goldfish":    5,
	}

	for k := range minRanges {
		if strings.Contains(line, k) {
			keyIndex := strings.Index(line, k)
			commaIndex := strings.Index(line[keyIndex:], ",")

			wordLen := len(k)

			var value string

			if commaIndex == -1 {
				value = line[keyIndex+wordLen+2:]
			} else {
				value = line[keyIndex+wordLen+2 : commaIndex+keyIndex]
			}

			numericValue, _ := strconv.Atoi(value)

			if numericValue > minRanges[k] {
				return true, true
			} else {
				return true, false
			}
		}
	}

	for k := range maxRanges {
		if strings.Contains(line, k) {
			keyIndex := strings.Index(line, k)
			commaIndex := strings.Index(line[keyIndex:], ",")
			wordLen := len(k)

			var value string

			if commaIndex == -1 {
				value = line[keyIndex+wordLen+2:]
			} else {
				value = line[keyIndex+wordLen+2 : commaIndex+keyIndex]
			}

			numericValue, _ := strconv.Atoi(value)

			if numericValue < maxRanges[k] {
				return true, true
			} else {
				return true, false
			}
		}
	}

	return false, false
}
