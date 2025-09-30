package day17

import (
	"2018/utils"
	"fmt"
	"math"
)

type point struct {
	y int
	x int
}

type block string

const (
	clay    block = "#"
	water   block = "~"
	sand    block = "."
	source  block = "+"
	wetSand block = "|"
)

type boundary struct {
	minX int
	maxX int
	minY int
	maxY int
}

func SolveP1() {
	grid, bounds := initGrid()
	dfsFill(grid, point{0, 500}, bounds)
	fmt.Println("found", countWater(grid, bounds)+countWetSand(grid, bounds), "water")
}

func SolveP2() {
	grid, bounds := initGrid()
	dfsFill(grid, point{0, 500}, bounds)
	fmt.Println("found", countWater(grid, bounds), "water")
}

func initGrid() (map[point]block, boundary) {
	lines := utils.ReadFile("day17/input.txt")
	clayCoordinates, bounds := parseClayCoordinates(lines)
	bounds.maxX++
	bounds.minX--

	grid := make(map[point]block)
	for coord := range clayCoordinates {
		grid[coord] = clay
	}
	grid[point{0, 500}] = source
	return grid, bounds
}

func dfsFill(grid map[point]block, from point, bounds boundary) bool {
	if grid[from] == clay {
		return true
	}

	if grid[from] == water || grid[from] == wetSand {
		return grid[from] == water
	}

	if from.y > bounds.maxY {
		return false
	}

	canBeWater := dfsFill(grid, point{from.y + 1, from.x}, bounds)

	if canBeWater {
		canSpread := canSpreadHorizontally(grid, from, left, bounds)
		canSpread = canSpread && canSpreadHorizontally(grid, from, right, bounds)
		if canSpread {
			grid[from] = water
			spreadHorizontally(grid, from, left, bounds)
			spreadHorizontally(grid, from, right, bounds)
		} else {
			grid[from] = wetSand
			wetHorizontally(grid, from, left, bounds)
			wetHorizontally(grid, from, right, bounds)
		}
		return canSpread
	} else {
		grid[from] = wetSand
	}

	return false
}

type direction string

const (
	left  direction = "left"
	right direction = "right"
)

func canSpreadHorizontally(grid map[point]block, from point, dir direction, bounds boundary) bool {
	if from.x < bounds.minX || from.x > bounds.maxX {
		return false
	}

	below := point{from.y + 1, from.x}

	if grid[from] == clay || grid[from] == water {
		return true
	}

	if grid[below] != clay && grid[below] != water {
		return dfsFill(grid, from, bounds)
	}

	next := from

	if dir == left {
		next.x--
	} else {
		next.x++
	}

	if canSpreadHorizontally(grid, next, dir, bounds) {
		return true
	}

	return false

}

func spreadHorizontally(grid map[point]block, from point, dir direction, bounds boundary) {
	if from.x < bounds.minX || from.x > bounds.maxX || grid[from] == clay {
		return
	}

	grid[from] = water

	next := from
	if dir == left {
		next.x--
	} else {
		next.x++
	}
	spreadHorizontally(grid, next, dir, bounds)
}

func wetHorizontally(grid map[point]block, from point, dir direction, bounds boundary) {
	if from.x < bounds.minX || from.x > bounds.maxX || grid[from] == clay {
		return
	}

	below := point{from.y + 1, from.x}
	if grid[below] != clay && grid[below] != water {
		grid[from] = wetSand
		dfsFill(grid, below, bounds)
		return
	}

	grid[from] = wetSand

	next := from
	if dir == left {
		next.x--
	} else {
		next.x++
	}
	wetHorizontally(grid, next, dir, bounds)
}

func parseClayCoordinates(lines []string) (map[point]block, boundary) {
	coordinates := make(map[point]block)
	minX, maxX, minY, maxY := math.MaxInt, 0, math.MaxInt, 0

	for _, line := range lines {
		switch line[0] {
		case 'x':
			var x, y1, y2 int
			fmt.Sscanf(line, "x=%d, y=%d..%d", &x, &y1, &y2)
			for y := y1; y <= y2; y++ {
				coordinates[point{y, x}] = clay
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		case 'y':
			var y, x1, x2 int
			fmt.Sscanf(line, "y=%d, x=%d..%d", &y, &x1, &x2)
			for x := x1; x <= x2; x++ {
				coordinates[point{y, x}] = clay
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	return coordinates, boundary{minX, maxX, minY, maxY}
}

func countBlocks(grid map[point]block, bounds boundary, types ...block) int {
	count := 0
	for y := bounds.minY; y <= bounds.maxY; y++ {
		for x := bounds.minX; x <= bounds.maxX; x++ {
			for _, t := range types {
				if grid[point{y, x}] == t {
					count++
					break
				}
			}
		}
	}
	return count
}

func countWater(grid map[point]block, bounds boundary) int {
	return countBlocks(grid, bounds, water)
}

func countWetSand(grid map[point]block, bounds boundary) int {
	return countBlocks(grid, bounds, wetSand)
}
