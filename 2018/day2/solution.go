package day2

import (
	"2018/utils"
	"fmt"
	"math"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("day2/input.txt")

	twices := 0
	thrices := 0

	for _, line := range lines {
		counter := make(map[rune]int)

		for _, letter := range line {
			counter[letter] += 1
		}

		foundTwice := false
		foundThrice := false

		for _, value := range counter {

			if foundThrice && foundTwice {
				break
			}

			if !foundTwice && value == 2 {
				foundTwice = true
				twices += 1
			} else if !foundThrice && value == 3 {
				foundThrice = true
				thrices += 1
			}
		}

	}

	fmt.Println("twices:", twices, "thrices:", thrices, "result:", twices*thrices)
}

func SolveP2() {
	lines := utils.ReadFile("day2/input.txt")

	var closest string
	closestDifference := math.MaxInt

	for index, line := range lines {
		for _, comparingLine := range lines[index+1:] {
			difference := 0
			common := strings.Builder{}

			for idx := range line {
				if line[idx] != comparingLine[idx] {
					difference += 1
				} else {
					common.WriteByte(line[idx])
				}
			}

			if difference < closestDifference {
				closest = common.String()
				closestDifference = difference
			}
		}
	}

	fmt.Println("result:", closest)
}
