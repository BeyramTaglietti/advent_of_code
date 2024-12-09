package day9

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
)

const space = -1

func SolveP1() {
	lines := utils.ReadFile("./day9/input.txt")

	puzzle := lines[0]
	fmt.Println("puzzle input: ", puzzle)

	hardDisk := mapHardDisk(puzzle)
	fmt.Println("hardDisk before: ", hardDisk)

	hardDisk = emptySpace(hardDisk)
	fmt.Println("hardDisk after: ", hardDisk)

	checksum := calculateFileSystemChecksum(hardDisk)
	fmt.Println("checksum: ", checksum)
}

func mapHardDisk(puzzle string) []int {
	hardDisk := []int{}

	for idx, value := range puzzle {
		numValue, _ := strconv.Atoi(string(value))

		isSpace := idx%2 != 0

		var valToAppend int
		if isSpace {
			valToAppend = space
		} else {
			valToAppend = idx / 2
		}

		for i := 0; i < numValue; i++ {
			hardDisk = append(hardDisk, valToAppend)
		}
	}

	return hardDisk
}

func getLastNonEmptyValue(disk []int) (value, idxOfvalue int) {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != space {
			return disk[i], i
		}
	}

	return -1, -1
}

func emptySpace(disk []int) []int {

	newDiskLen := 0
	for _, value := range disk {
		if value != space {
			newDiskLen++
		}
	}

	emptiedDisk := []int{}
	for _, value := range disk {

		if newDiskLen == len(emptiedDisk) {
			break
		}

		if value == -1 {
			lastValue, indexOfLastValue := getLastNonEmptyValue(disk)

			// remove last value from original disk
			disk = disk[:indexOfLastValue]

			emptiedDisk = append(emptiedDisk, lastValue)
		} else {
			emptiedDisk = append(emptiedDisk, value)
		}
	}
	return emptiedDisk
}

func calculateFileSystemChecksum(disk []int) int {
	sum := 0
	for idx, value := range disk {
		sum += value * idx
	}

	return sum
}
