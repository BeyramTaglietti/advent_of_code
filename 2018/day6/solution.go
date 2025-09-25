package day6

import (
	"2018/utils"
	"fmt"
	"math"
)

type point struct {
	y, x int
}

func SolveP1() {
	lines := utils.ReadFile("day6/input.txt")

	minX, maxX, minY, maxY := math.MaxInt, 0, math.MaxInt, 0

	gridPointsGrid := make(map[point]int, len(lines))
	gridPoints := make([]point, len(lines))

	for _, line := range lines {
		y, x := parseLine(line)
		newPoint := point{y, x}
		gridPointsGrid[newPoint] = 0
		gridPoints = append(gridPoints, newPoint)

		minX = int(math.Min(float64(minX), float64(x)))
		minY = int(math.Min(float64(minY), float64(y)))

		maxX = int(math.Max(float64(maxX), float64(x)))
		maxY = int(math.Max(float64(maxY), float64(y)))
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {

			closestPoints := findClosestPoints(point{y, x}, gridPoints)
			if (y == minY || y == maxY || x == minX || x == maxX) && len(closestPoints) == 1 {
				for _, cp := range closestPoints {
					delete(gridPointsGrid, cp)
				}
			} else {
				if len(closestPoints) == 1 {
					if _, ok := gridPointsGrid[closestPoints[0]]; ok {
						gridPointsGrid[closestPoints[0]]++
					}
				}
			}

		}
	}

	maxArea := 0

	for _, area := range gridPointsGrid {
		maxArea = int(math.Max(float64(maxArea), float64(area)))
	}

	fmt.Println("max area", maxArea)
}

func SolveP2() {
	lines := utils.ReadFile("day6/input.txt")

	minX, maxX, minY, maxY := math.MaxInt, 0, math.MaxInt, 0

	gridPoints := make(map[point]int, len(lines))

	for _, line := range lines {
		y, x := parseLine(line)
		newPoint := point{y, x}
		gridPoints[newPoint] = 0

		minX = int(math.Min(float64(minX), float64(x)))
		minY = int(math.Min(float64(minY), float64(y)))

		maxX = int(math.Max(float64(maxX), float64(x)))
		maxY = int(math.Max(float64(maxY), float64(y)))
	}

	totalArea := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			totalDistance := 0
			for p := range gridPoints {
				totalDistance += manhattanDistance(p.x, p.y, x, y)
			}
			if totalDistance < 10000 {
				totalArea++
			}
		}
	}

	fmt.Println("total area", totalArea)
}

func findClosestPoints(p point, grid []point) []point {
	minDistance := math.MaxInt
	distances := make(map[point]int)

	result := []point{}

	for _, gridPoint := range grid {
		distance := manhattanDistance(gridPoint.x, gridPoint.y, p.x, p.y)
		minDistance = int(math.Min(float64(minDistance), float64(distance)))

		distances[gridPoint] = distance
	}

	for point, distance := range distances {
		if distance == minDistance {
			result = append(result, point)
		}
	}

	return result
}

func parseLine(line string) (y int, x int) {
	fmt.Sscanf(line, "%d, %d", &x, &y)
	return y, x
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}
