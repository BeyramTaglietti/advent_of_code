package day10

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("day10/input.txt")
	puzzle := lines[0]

	for j := 0; j < 40; j++ {
		puzzle = buildString(puzzle)
	}

	fmt.Println("Final result's length:", len(puzzle))
}

func SolveP2() {
	lines := utils.ReadFile("day10/input.txt")
	puzzle := lines[0]

	for j := 0; j < 50; j++ {
		puzzle = buildString(puzzle)
	}

	fmt.Println("Final result's length:", len(puzzle))
}

func buildString(puzzle string) string {
	var builder strings.Builder
	for i := 0; i < len(puzzle); i++ {
		letter := rune(puzzle[i])
		repeats := countRepeats(letter, puzzle[i:])
		builder.WriteString(strconv.Itoa(repeats) + string(letter))
		i += repeats - 1
	}
	return builder.String()
}

func countRepeats(letter rune, line string) int {
	counter := 0

	for _, x := range line {
		if letter == x {
			counter++
			continue
		}

		break
	}

	return counter
}
