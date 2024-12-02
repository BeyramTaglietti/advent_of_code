package day2

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("./day2/input.txt")

	safeCounter := 0
	for _, line := range lines {
		if isReportSafe(parseLine(line), false) {
			safeCounter++
		}
	}

	fmt.Printf("Found %d safe reports\n", safeCounter)
}

func SolveP2() {
	lines := utils.ReadFile("./day2/input.txt")

	safeCounter := 0
	for _, line := range lines {
		if isReportSafe(parseLine(line), true) {
			safeCounter++
		}
	}

	fmt.Printf("Found %d safe dampened reports\n", safeCounter)
}

func isReportSafe(report []uint16, allowDampener bool) bool {

	if isIncreasingSafely(report, true, allowDampener) {
		return true
	}

	return isIncreasingSafely(report, false, allowDampener)
}

func isIncreasingSafely(report []uint16, increasing bool, allowDampener bool) bool {

	prevNumber := report[0]

	for idx, level := range report[1:] {

		var difference uint16

		if increasing {
			difference = level - prevNumber
		} else {
			difference = prevNumber - level
		}

		if difference < 1 || difference > 3 {

			if allowDampener {

				firstTry := slices.Clone(report)
				firstTry = slices.Delete(firstTry, idx, idx+1)

				if isIncreasingSafely(firstTry, increasing, false) {
					return true
				}

				secondTry := slices.Clone(report)
				secondTry = slices.Delete(secondTry, idx+1, idx+2)

				return isIncreasingSafely(secondTry, increasing, false)

			} else {
				return false
			}
		}

		prevNumber = level
	}

	return true
}

func parseLine(line string) []uint16 {
	reportStrList := strings.Split(line, " ")

	reportIntList := make([]uint16, len(reportStrList), len(reportStrList))

	for idx, rep := range reportStrList {
		parsedStr, _ := strconv.Atoi(rep)

		reportIntList[idx] = uint16(parsedStr)
	}

	return reportIntList
}
