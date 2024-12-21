package day20

import (
	"aoc2024/utils"
	"fmt"
	"maps"
	"slices"
	"sync"
)

type point struct {
	x int
	y int
}

type path struct {
	current_position point
	visited_points   map[point]int
	current_steps    int
	cheated          bool
}

type gameMap map[point]bool

func SolveP1() {
	lines := utils.ReadFile("./day20/input.txt")

	game_map, entrance, exit := parseMap(lines)

	paths := []path{{current_position: entrance, visited_points: map[point]int{entrance: 0}, current_steps: 0}}

	_, validPath := findShortestPath(game_map, paths, exit, []path{})

	// game_map.printMap(len(lines[0]), len(lines))
	// validPath.printPath(len(lines[0]), len(lines))

	fmt.Println("Path length", validPath.visited_points[exit])

	savings := validPath.tryCheatingPath(game_map, validPath.visited_points[exit])

	fmt.Println("Total savings that save at least 100 picoseconds", countSavings(savings, 100))
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
		if _, ok := m[p]; ok && current_path.visited_points[p] == 0 {
			if !m[p] {
				next_points = append(next_points, p)
			}
		}
	}

	return next_points
}

func findShortestPath(m map[point]bool, currentPaths []path, exitPoint point, validPaths []path) (bool, path) {

	if len(currentPaths) == 0 {
		return false, path{}
	}

	nextPaths := []path{}

	wg := sync.WaitGroup{}
	wg.Add(len(currentPaths))

	mu := sync.Mutex{}

	for _, p := range currentPaths {

		go func(x path, exit point) {
			defer wg.Done()

			nextPoints := findNextPoints(m, p)

			if slices.Contains(nextPoints, exitPoint) {
				visited_ps := maps.Clone(x.visited_points)
				visited_ps[exitPoint] = x.current_steps + 1
				validP := path{current_position: exitPoint, visited_points: visited_ps, current_steps: x.current_steps + 1}
				mu.Lock()
				validPaths = append(validPaths, validP)
				mu.Unlock()
			}

			for _, np := range nextPoints {

				newVisitedPoints := maps.Clone(p.visited_points)
				newVisitedPoints[np] = p.current_steps + 1

				mu.Lock()
				nextPaths = append(nextPaths, path{current_position: np, visited_points: newVisitedPoints, current_steps: p.current_steps + 1})
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

func parseMap(lines []string) (game_map gameMap, entrance point, exit point) {

	game_map = make(map[point]bool)

	for y, line := range lines {
		for x, char := range line {

			switch char {
			case '#':
				game_map[point{x, y}] = true
			default:
				switch char {
				case 'S':
					entrance = point{x, y}
				case 'E':
					exit = point{x, y}
				}
				game_map[point{x, y}] = false
			}
		}
	}

	return game_map, entrance, exit
}

func (g gameMap) printMap(limitX, limitY int) {
	for y := 0; y < limitY; y++ {
		for x := 0; x < limitX; x++ {
			if g[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (p path) printPath(limitX, limitY int) {
	for y := 0; y < limitY; y++ {
		for x := 0; x < limitX; x++ {
			if p.visited_points[point{x, y}] != 0 {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type direction string

const (
	up    direction = "up"
	down  direction = "down"
	left  direction = "left"
	right direction = "right"
)

type cheatingPoint struct {
	point
	dir direction
}

func (final_path path) tryCheatingPath(m map[point]bool, currentRecord int) map[int]int {
	savings := make(map[int]int)

	isPartOfOriginalPath := func(p point) (bool, int) {
		return final_path.visited_points[p] != 0, final_path.visited_points[p]
	}

	checkCheatedPoint := func(p cheatingPoint, steps int) {
		partOfPath, ogStepCount := isPartOfOriginalPath(p.point)

		if partOfPath {
			saved := ogStepCount - steps

			if currentRecord-saved < currentRecord {
				savings[saved-2]++
			}
		}
	}

	for p, steps := range final_path.visited_points {
		blockedAdjacentPoints := []cheatingPoint{}

		possible_next_points := []cheatingPoint{
			{point: point{p.x, p.y - 1}, dir: up},
			{point: point{p.x, p.y + 1}, dir: down},
			{point: point{p.x - 1, p.y}, dir: left},
			{point: point{p.x + 1, p.y}, dir: right},
		}

		for _, p := range possible_next_points {
			if _, ok := m[p.point]; ok && m[p.point] {
				blockedAdjacentPoints = append(blockedAdjacentPoints, p)
			}
		}

		for _, bp := range blockedAdjacentPoints {
			tryPoint := cheatingPoint{point: bp.point, dir: bp.dir}
			switch bp.dir {
			case up:
				tryPoint.point.y--
			case down:
				tryPoint.point.y++
			case left:
				tryPoint.point.x--
			case right:
				tryPoint.point.x++
			}
			checkCheatedPoint(tryPoint, steps)
		}
	}

	return savings
}

func countSavings(m map[int]int, min int) int {
	total := 0

	for k, v := range m {
		if k >= min {
			total += v
		}
	}

	return total
}
