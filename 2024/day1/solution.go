package day1

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"slices"
)

func SolveP1() {

	lines := utils.ReadFile("day1/input.txt")

	valuesLen := len(lines)

	lefts := make([]int, valuesLen, valuesLen)
	rights := make([]int, valuesLen, valuesLen)
	for idx, line := range lines {
		l, r := parseLine(line)

		lefts[idx] = l
		rights[idx] = r
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	totalDistance := 0

	for i := 0; i < valuesLen; i++ {
		totalDistance += int(math.Abs((float64(lefts[i]) - float64(rights[i]))))
	}

	fmt.Println("total distance", totalDistance)
}

func SolveP2() {

	lines := utils.ReadFile("day1/input.txt")

	valuesLen := len(lines)

	lefts := make([]int, valuesLen, valuesLen)
	rights := make([]int, valuesLen, valuesLen)
	for idx, line := range lines {
		l, r := parseLine(line)

		lefts[idx] = l
		rights[idx] = r
	}

	similarityScore := 0
	for _, lValue := range lefts {
		counter := 0
		for _, rValue := range rights {
			if rValue == lValue {
				counter++
			}
		}

		similarityScore += counter * lValue
	}

	fmt.Println("Similarity score:", similarityScore)

}

func parseLine(line string) (l, r int) {
	var left, right int
	fmt.Sscanf(line, "%d   %d", &left, &right)

	return left, right
}
