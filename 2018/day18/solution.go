package day18

import (
	"2018/utils"
	"fmt"
)

type point struct {
	y int
	x int
}

type block string

const (
	open       = "."
	tree       = "|"
	lumberyard = "#"
)

func SolveP1() {
	lines := utils.ReadFile("day18/input.txt")

	grid, m, n := parseGrid(lines)

	printGrid(grid, m, n)
	minutes := 10

	for range minutes {
		grid = nextCycle(grid, m, n)
	}

	printGrid(grid, m, n)
	fmt.Println("result", calculateResult(grid))
}

func SolveP2() {
	lines := utils.ReadFile("day18/input.txt")

	grid, m, n := parseGrid(lines)

	printGrid(grid, m, n)
	minutes := 1000000000

	seenGrids := make(map[string]int)

	for currentCycle := 0; currentCycle < minutes; currentCycle++ {
		gridStr := gridToString(grid, m, n)

		if prevValue, ok := seenGrids[gridStr]; ok {
			cycleLength := currentCycle - prevValue
			remainingCycles := minutes - currentCycle
			remainder := remainingCycles % cycleLength

			for i := 0; i < remainder; i++ {
				grid = nextCycle(grid, m, n)
			}
			break
		}

		seenGrids[gridStr] = currentCycle

		grid = nextCycle(grid, m, n)
	}

	printGrid(grid, m, n)
	fmt.Println("result", calculateResult(grid))
}

func nextCycle(grid map[point]block, m, n int) map[point]block {
	newGrid := make(map[point]block)
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			newGrid[point{y, x}] = evolve(point{y, x}, grid)
		}
	}
	return newGrid
}

func parseGrid(lines []string) (map[point]block, int, int) {
	grid := make(map[point]block)
	m := len(lines)
	n := len(lines[0])

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			var newBlock block
			switch lines[y][x] {
			case '#':
				newBlock = lumberyard
			case '|':
				newBlock = tree
			case '.':
				newBlock = open
			}
			grid[point{y, x}] = newBlock
		}
	}

	return grid, m, n
}

func evolve(p point, grid map[point]block) block {

	adjacentBlocks := []point{
		{p.y - 1, p.x - 1},
		{p.y - 1, p.x},
		{p.y - 1, p.x + 1},
		{p.y, p.x - 1},
		{p.y, p.x + 1},
		{p.y + 1, p.x - 1},
		{p.y + 1, p.x},
		{p.y + 1, p.x + 1},
	}

	countNeighbours := func(ofType block) int {
		counter := 0
		for _, adjBlock := range adjacentBlocks {
			if _, ok := grid[adjBlock]; ok {
				if grid[adjBlock] == ofType {
					counter++
				}
			}
		}

		return counter
	}

	switch grid[p] {
	case open:
		if countNeighbours(tree) >= 3 {
			return tree
		}
	case tree:
		if countNeighbours(lumberyard) >= 3 {
			return lumberyard
		}
	case lumberyard:
		if countNeighbours(lumberyard) >= 1 && countNeighbours(tree) >= 1 {
			return lumberyard
		} else {
			return open
		}
	}

	return grid[p]
}

func printGrid(grid map[point]block, m, n int) {
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			fmt.Print(grid[point{y, x}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func calculateResult(grid map[point]block) int {
	lumberyards := 0
	woodAcres := 0
	for _, v := range grid {
		switch v {
		case lumberyard:
			lumberyards++
		case tree:
			woodAcres++
		}
	}

	return woodAcres * lumberyards
}

func gridToString(grid map[point]block, m, n int) string {
	var result string
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			result += string(grid[point{y, x}])
		}
	}
	return result
}
