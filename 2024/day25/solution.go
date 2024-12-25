package day25

import (
	"aoc2024/utils"
	"fmt"
)

func SolveP1() {

	lines := utils.ReadFile("./day25/input.txt")

	keys := make([][]int, 0)
	locks := make([][]int, 0)

	for i := 0; i < len(lines); i += 8 {
		is_key, grid := parseGrid(lines[i : i+7])

		if is_key {
			keys = append(keys, grid)
		} else {
			locks = append(locks, grid)
		}
	}

	valid_pairs := 0

	for _, key := range keys {
		for _, lock := range locks {
			if isValidPair(key, lock) {
				valid_pairs++
			}
		}
	}

	fmt.Println("Valid pairs:", valid_pairs)
}

func parseGrid(
	lines []string,
) (isKey bool, grid []int) {

	isKey = string(lines[0]) == "....."

	if !isKey {
		for x := 0; x < len(lines[0]); x++ {
			grid = append(grid, 0)
			for y := 1; y < len(lines); y++ {
				if lines[y][x] == '#' {
					grid[x]++
				}
			}
		}
	} else {
		for x := 0; x < len(lines[0]); x++ {
			grid = append(grid, 0)
			for y := 0; y < len(lines)-1; y++ {
				if lines[y][x] == '#' {
					grid[x]++
				}
			}
		}
	}

	return isKey, grid
}

func isValidPair(
	key []int,
	lock []int,
) bool {

	const available_space = 5

	for i := 0; i < len(key); i++ {
		if lock[i]+key[i] > available_space {
			return false
		}
	}

	return true
}
