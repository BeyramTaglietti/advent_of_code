package day18

import (
	"aoc2015/utils"
	"fmt"
)

type point struct {
	x int
	y int
}

func SolveP1() {
	lines := utils.ReadFile("./day18/input.txt")

	limitX, limitY := len(lines[0]), len(lines)

	grid := createGrid(lines)

	const steps int = 100

	fmt.Println("grid before")
	printGrid(grid, limitX, limitY)

	for step := range steps {
		fmt.Println("grid at step", step+1, ":")
		grid = nextState(grid, false, limitX, limitY)
		printGrid(grid, limitX, limitY)
		fmt.Println()
	}

	fmt.Println("found", countOnLights(grid), "on lights")
}

func SolveP2() {
	lines := utils.ReadFile("./day18/input.txt")

	limitX, limitY := len(lines[0])-1, len(lines)-1

	grid := createGrid(lines)

	const steps int = 100

	for range steps {
		grid = nextState(grid, true, limitX, limitY)
	}

	fmt.Println("found", countOnLights(grid), "on lights")
}

func createGrid(lines []string) map[point]bool {
	grid := make(map[point]bool)
	for y, line := range lines {
		for x, char := range line {
			p := point{x, y}

			if char == '.' {
				grid[p] = false
			} else {
				grid[p] = true
			}
		}
	}

	return grid
}

func printGrid(grid map[point]bool, limitX, limitY int) {
	for y := 0; y <= limitY; y++ {
		for x := 0; x <= limitX; x++ {
			if isOn, ok := grid[point{x, y}]; ok {
				if isOn {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findNeighbours(p point, grid map[point]bool) []point {

	adjacentPoints := []point{
		{p.x + 1, p.y},
		{p.x + 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y + 1},
		{p.x - 1, p.y},
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
	}

	onNeighbours := []point{}
	for _, n := range adjacentPoints {
		if isOn, exists := grid[n]; exists {
			if isOn {
				onNeighbours = append(onNeighbours, n)
			}
		}
	}

	return onNeighbours
}

func nextState(grid map[point]bool, cornersStuckOn bool, limitX, limitY int) map[point]bool {

	newGrid := make(map[point]bool)

	corners := []point{
		{0, 0},
		{limitX, 0},
		{0, limitY},
		{limitX, limitY},
	}

	for p, lightIsOn := range grid {
		isCorner := func(c point) bool {

			for _, corner := range corners {
				if corner.x == c.x && corner.y == c.y {
					return true
				}
			}

			return false
		}

		if cornersStuckOn && isCorner(p) {
			newGrid[p] = true
			continue
		}

		onNeighbours := findNeighbours(p, grid)

		switch lightIsOn {
		case true:
			if len(onNeighbours) != 2 && len(onNeighbours) != 3 {
				newGrid[p] = false
			} else {
				newGrid[p] = lightIsOn
			}
		case false:
			if len(onNeighbours) == 3 {
				newGrid[p] = true
			} else {
				newGrid[p] = lightIsOn
			}
		}

	}

	return newGrid

}

func countOnLights(grid map[point]bool) int {
	sum := 0

	for _, v := range grid {
		if v {
			sum++
		}
	}

	return sum
}
