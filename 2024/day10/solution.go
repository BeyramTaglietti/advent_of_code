package day10

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
)

type point struct {
	x int
	y int
}

type trail struct {
	startingPoint point
	endingPoint   point
}

func SolveP1() {
	lines := utils.ReadFile("./day10/input.txt")

	grid := createGrid(lines)

	trailsFound := []trail{}
	for x, line := range lines {
		for y, char := range line {
			if char == '0' {
				fmt.Println("found 0")
				foundTrailHeads := createTrail([]point{{x: x, y: y}}, grid)
				for _, head := range foundTrailHeads {
					trailsFound = append(trailsFound, trail{startingPoint: point{x: x, y: y}, endingPoint: head})
				}

				fmt.Println("trailsFound:", trailsFound)
			}
		}
	}

	fmt.Println("unique trails found:", findUniqueTrails(trailsFound))
}

func createGrid(lines []string) map[point]int {
	grid := make(map[point]int)

	for x, line := range lines {
		for y, char := range line {
			numVal, _ := strconv.Atoi(string(char))
			grid[point{x: x, y: y}] = numVal
		}
	}

	return grid
}

func findNextPoints(p point, grid map[point]int) []point {

	possibleNextPoints := []point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}

	nextPoints := []point{}

	for _, nextPoint := range possibleNextPoints {
		if _, ok := grid[nextPoint]; ok {
			if grid[nextPoint] == grid[p]+1 {
				nextPoints = append(nextPoints, nextPoint)
			}
		}
	}

	return nextPoints
}

func createTrail(points []point, grid map[point]int) []point {

	fmt.Println("creating trail with points:", points)

	firstPoint := points[0]

	if grid[firstPoint] == 9 {
		return points
	}

	nextPoints := []point{}

	for _, p := range points {
		nextPoints = append(nextPoints, findNextPoints(p, grid)...)
	}

	return createTrail(nextPoints, grid)
}

func findUniqueTrails(trails []trail) int {
	uniqueTrails := []trail{}

	for _, p := range trails {
		if !slices.Contains(uniqueTrails, p) {
			uniqueTrails = append(uniqueTrails, p)
		}
	}

	return len(uniqueTrails)
}
