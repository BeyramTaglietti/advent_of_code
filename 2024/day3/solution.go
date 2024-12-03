package day3

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("./day3/input.txt")

	pattern := `mul\(\d+,\d+\)`

	total := 0
	for _, line := range lines {
		re, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatalf("Error compiling regex: %s", err.Error())
		}

		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			total += multiply(match)
		}

	}

	fmt.Println("total sum:", total)
}

func SolveP2() {
	lines := utils.ReadFile("./day3/input.txt")

	var fullPuzzle strings.Builder

	for _, line := range lines {
		fullPuzzle.WriteString(line)
	}

	line := fullPuzzle.String()

	pattern := `mul\((\d+,\d+)\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)
	operations := re.FindAllString(line, -1)

	sum := 0
	enabled := true
	for _, v := range operations {
		if v == "don't()" {
			enabled = false
			continue
		}
		if v == "do()" {
			enabled = true
			continue
		}
		if enabled {
			sum += multiply(v)
		}
	}

	fmt.Println("total sum with rules:", sum)
}

func multiply(op string) int {
	var x1, x2 int

	fmt.Sscanf(op, "mul(%d,%d)", &x1, &x2)

	return x1 * x2
}
