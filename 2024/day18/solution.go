package day18

import (
	"aoc2024/utils"
	"fmt"
	"maps"
	"slices"
	"sync"
)

type point struct {
	x, y int
}

type path struct {
	current_position point
	visited_points   map[point]bool
}

func SolveP1() {

	lines := utils.ReadFile("day18/input.txt")

	limitX, limitY := 70, 70
	gameMap := make(map[point]bool)

	for x := 0; x <= limitX; x++ {
		for y := 0; y <= limitY; y++ {
			gameMap[point{x, y}] = false
		}
	}

	bytesMax := 1024
	index := 0
	for _, line := range lines {

		if index == bytesMax {
			break
		}

		x1, x2 := paseLine(line)
		gameMap[point{x1, x2}] = true

		index++
	}

	// printMap(gameMap, limitX, limitY)

	entrance := point{0, 0}

	paths := []path{{current_position: entrance, visited_points: map[point]bool{entrance: true}}}

	_, validPath := findShortestPath(gameMap, paths, point{limitX, limitY}, []path{})

	// printPath(validPath, limitX, limitY)

	fmt.Println("Path length", len(validPath.visited_points))
}

func SolveP2() {

	lines := utils.ReadFile("day18/input.txt")

	limitX, limitY := 70, 70
	gameMap := make(map[point]bool)

	for x := 0; x <= limitX; x++ {
		for y := 0; y <= limitY; y++ {
			gameMap[point{x, y}] = false
		}
	}

	entrance := point{0, 0}

	for i := len(lines); i > 0; i-- {
		f_lines := lines[:i]

		gameMap = createMap(f_lines, limitX, limitY)

		paths := []path{{current_position: entrance, visited_points: map[point]bool{entrance: true}}}
		foundValidPath, _ := findShortestPath(gameMap, paths, point{limitX, limitY}, []path{})
		if foundValidPath {
			freeing_point_x, freeing_point_y := paseLine(lines[i])
			fmt.Println("first point which blocks the map", point{freeing_point_x, freeing_point_y})
			break
		}
	}

}

func createMap(lines []string, limitX, limitY int) map[point]bool {
	gameMap := make(map[point]bool)

	for x := 0; x <= limitX; x++ {
		for y := 0; y <= limitY; y++ {
			gameMap[point{x, y}] = false
		}
	}

	for _, line := range lines {
		x1, x2 := paseLine(line)
		gameMap[point{x1, x2}] = true
	}

	return gameMap
}

func printPath(p path, limitX, limitY int) {
	for y := 0; y <= limitY; y++ {
		for x := 0; x <= limitX; x++ {
			if p.visited_points[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func paseLine(line string) (int, int) {
	var x1, x2 int

	fmt.Sscanf(line, "%d,%d", &x1, &x2)

	return x1, x2
}

func printMap(gameMap map[point]bool, limitX, limitY int) {
	for y := 0; y <= limitY; y++ {
		for x := 0; x <= limitX; x++ {
			if gameMap[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findShortestPath(m map[point]bool, currentPaths []path, exitPoint point, validPaths []path) (bool, path) {

	if len(currentPaths) == 0 {
		return false, path{}
	}

	currentPaths = clearSimilarPaths(currentPaths)

	nextPaths := []path{}

	wg := sync.WaitGroup{}
	wg.Add(len(currentPaths))

	mu := sync.Mutex{}

	for _, p := range currentPaths {

		go func(x path, exit point) {
			defer wg.Done()

			nextPoints := findNextPoints(m, p)

			if slices.Contains(nextPoints, exitPoint) {
				mu.Lock()
				validPaths = append(validPaths, p)
				mu.Unlock()
			}

			for _, np := range nextPoints {
				newVisitedPoints := maps.Clone(p.visited_points)
				newVisitedPoints[np] = true

				mu.Lock()
				nextPaths = append(nextPaths, path{current_position: np, visited_points: newVisitedPoints})
				mu.Unlock()
			}
		}(p, exitPoint)

	}

	wg.Wait()

	if len(validPaths) > 0 {
		return true, validPaths[0]
	}

	return findShortestPath(m, nextPaths, exitPoint, validPaths)

}

func findNextPoints(m map[point]bool, current_path path) []point {

	var next_points []point

	possible_next_points := []point{
		{current_path.current_position.x, current_path.current_position.y - 1},
		{current_path.current_position.x, current_path.current_position.y + 1},
		{current_path.current_position.x - 1, current_path.current_position.y},
		{current_path.current_position.x + 1, current_path.current_position.y},
	}

	for _, p := range possible_next_points {
		if _, ok := m[p]; ok && !m[p] && !current_path.visited_points[p] {
			next_points = append(next_points, p)
		}
	}

	return next_points
}

func clearSimilarPaths(paths []path) []path {
	uniquePaths := []path{}

	for _, p := range paths {
		if !slices.ContainsFunc(uniquePaths, func(x1 path) bool {
			return x1.current_position == p.current_position && len(x1.visited_points) == len(p.visited_points)
		}) {
			uniquePaths = append(uniquePaths, p)
		}
	}

	return uniquePaths
}
