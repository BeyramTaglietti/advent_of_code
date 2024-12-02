package day12

import (
	"aoc2015/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func SolveP1() {
	puzzle := utils.ReadFile("./day12/input.txt")[0]

	counterMap := make(map[int]int)

	for i := 0; i < len(puzzle); i++ {
		character := rune(puzzle[i])

		if unicode.IsDigit(character) {
			numericValue, valueLength := getFullNumber(puzzle[i:])

			if puzzle[i-1] == '-' {
				counterMap[-numericValue]++
			} else {
				counterMap[numericValue]++
			}

			i += valueLength - 1
		}
	}

	fmt.Println("final sum:", sumMap(counterMap))
}

func sumMap(summedMap map[int]int) int {
	tot := 0

	for key, counter := range summedMap {
		tot += key * counter
	}

	return tot
}

func getFullNumber(line string) (numValue, strLength int) {
	var lineBuilder strings.Builder

	for _, val := range line {
		if unicode.IsDigit(val) {
			lineBuilder.WriteRune(val)
		} else {
			break
		}
	}

	strVal := lineBuilder.String()

	numericValue, _ := strconv.Atoi(strVal)

	return numericValue, len(strVal)
}
