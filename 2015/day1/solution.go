package day1

import (
	"aoc2015/utils"
	"fmt"
)

func SolveP1() {
	lines := utils.ReadFile("day1/input.txt")

	currentFloor := 0
	for _, r := range lines[0] {
		switch r {
		case '(':
			currentFloor++
		case ')':
			currentFloor--
		}
	}

	fmt.Printf("Final floor: %d\n", currentFloor)
}

func SolveP2() {
	lines := utils.ReadFile("day1/input.txt")

	currentFloor := 0
	var firstNegativeFloor *int = nil
	for idx, r := range lines[0] {
		switch r {
		case '(':
			currentFloor++
		case ')':
			currentFloor--
		}

		if currentFloor == -1 && firstNegativeFloor == nil {
			firstNegativeFloor = &idx
		}
	}

	fmt.Printf("First time in the basement: %d\n", *firstNegativeFloor+1)
}
