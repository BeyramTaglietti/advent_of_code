package day8

import (
	"aoc2022/utils"
	"fmt"
)

func SolveP1() {
	lines := utils.ReadFile("day8/input.txt")

	totalMemory := 0
	totalChars := 0

	for _, line := range lines {
		totalMemory += len(line)
		totalChars += numberOfChars(line)
	}

	fmt.Println("total number of chars:", totalChars)
	fmt.Println("total number of memory:", totalMemory)
	fmt.Println("difference:", totalMemory-totalChars)
}

func SolveP2() {
	lines := utils.ReadFile("day8/input.txt")

	totalMemory := 0

	for _, line := range lines {
		encLine := encodeLine(line)
		totalMemory += len(encLine) - len(line)
	}

	fmt.Println("difference:", totalMemory)
}

func encodeLine(line string) string {
	encodedLine := "\""
	for i := 0; i < len(line); i++ {
		letter := line[i]

		if letter == '"' {
			encodedLine += "\\\""
			continue
		}

		if letter == '\\' {
			encodedLine += "\\\\"
			continue
		}

		encodedLine += string(letter)
	}

	encodedLine += "\""

	return encodedLine
}

func numberOfChars(line string) int {
	count := 0
	for i := 0; i < len(line); i++ {
		letter := line[i]

		if letter == '"' {
			continue
		}

		if letter == '\\' {
			if line[i+1] == 'x' {
				i += 3
			} else {
				i++
			}

			count++
			continue

		}

		count++
	}

	return count
}
