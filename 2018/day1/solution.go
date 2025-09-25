package day1

import (
	"2018/utils"
	"fmt"
	"strconv"
)

func SolveP1() {
	lines := utils.ReadFile("day1/input.txt")

	currentFrequency := 0

	for _, line := range lines {
		value, _ := strconv.Atoi(line)

		currentFrequency += value
	}

	fmt.Println("ending frequency", currentFrequency)
}

func SolveP2() {

	lines := utils.ReadFile("day1/input.txt")

	currentFrequency := 0
	visitedFrequencies := make(map[int]bool)

outer:
	for {

		for _, line := range lines {
			value, _ := strconv.Atoi(line)

			currentFrequency += value

			if visitedFrequencies[currentFrequency] {
				fmt.Println("found:", currentFrequency)
				break outer
			}

			visitedFrequencies[currentFrequency] = true
		}

	}
}
