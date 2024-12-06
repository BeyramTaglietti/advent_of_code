package day5

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("./day5/input.txt")

	rulesParsed := false

	rulesMap := make(map[int][]int)

	sum := 0

	for _, line := range lines {
		if line == "" {
			rulesParsed = true
			continue
		}

		if !rulesParsed {
			n1, n2 := parseOrderRules(line)

			rulesMap[n2] = append(rulesMap[n2], n1)
		} else {
			numbers := parseUpdateOrder(line)

			isValid := true

			for i := 0; i < len(numbers)-1; i++ {
				containsInvalidValues, _ := listContainsInvalidValues(rulesMap[numbers[i]], numbers[i+1:])

				if containsInvalidValues {
					isValid = false
					break
				}
			}

			if isValid {
				middleValue := math.Floor(float64(len(numbers) / 2))

				sum += numbers[int(middleValue)]
			}

		}
	}

	fmt.Println("sum resulted in", sum)

}

func SolveP2() {
	lines := utils.ReadFile("./day5/input.txt")

	rulesParsed := false

	rulesMap := make(map[int][]int)

	sum := 0

	for _, line := range lines {
		if line == "" {
			rulesParsed = true
			continue
		}

		if !rulesParsed {
			n1, n2 := parseOrderRules(line)

			rulesMap[n2] = append(rulesMap[n2], n1)
		} else {
			numbers := parseUpdateOrder(line)

			valid := true

			for i := 0; i < len(numbers)-1; i++ {
				valuesToCheck := numbers[i+1:]
				containsInvalidValues, invalidValueIndex := listContainsInvalidValues(rulesMap[numbers[i]], valuesToCheck)

				if containsInvalidValues {
					utils.MoveListElement(numbers, slices.Index(numbers, valuesToCheck[invalidValueIndex]), i)
					i = -1 // start from the beginning
					valid = false
				}

			}

			if !valid {
				middleValue := math.Floor(float64(len(numbers) / 2))
				sum += numbers[int(middleValue)]
			}

		}
	}

	fmt.Println("sum of the fixed rows:", sum)
}

func parseOrderRules(line string) (int, int) {
	var n1, n2 int

	fmt.Sscanf(line, "%d|%d", &n1, &n2)

	return n1, n2
}

func parseUpdateOrder(line string) []int {
	strNumbers := strings.Split(line, ",")

	numbers := make([]int, len(strNumbers), len(strNumbers))

	for idx, number := range strNumbers {
		numericValue, _ := strconv.Atoi(number)
		numbers[idx] = numericValue
	}

	return numbers
}

func listContainsInvalidValues(rules []int, list []int) (containsInvalidValue bool, invalidValueIdx int) {
	for idx, number := range list {
		if slices.Contains(rules, number) {
			return true, idx
		}
	}

	return false, 0
}
