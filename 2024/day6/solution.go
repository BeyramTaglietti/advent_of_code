package day6

import (
	"aoc2024/utils"
	"fmt"
	"sync"
)

type coordinate struct {
	x int
	y int
}

type cell struct {
	visited        int
	guardWasFacing direction
	obstructs      bool
}

type direction string

const (
	up    direction = "up"
	down  direction = "down"
	left  direction = "left"
	right direction = "right"
)

type guard struct {
	facingDirection direction
	coordinate
}

// '.' nothing
//
// '#' obstruction
//
// '^' guard
func createMap(lines []string) (map[coordinate]cell, guard) {

	currentMap := make(map[coordinate]cell)
	var guard guard

	for y, line := range lines {
		for x, char := range line {

			position := coordinate{
				x: x,
				y: y,
			}

			switch char {
			case '.':
				currentMap[position] = cell{visited: 0, obstructs: false}
			case '#':
				currentMap[position] = cell{visited: 0, obstructs: true}
			default:
				switch char {
				case '<':
					guard.facingDirection = left
				case '>':
					guard.facingDirection = right
				case '^':
					guard.facingDirection = up
				case 'v':
					guard.facingDirection = down
				}
				guard.coordinate = position
				currentMap[position] = cell{visited: 1, obstructs: false, guardWasFacing: guard.facingDirection}
			}
		}
	}

	return currentMap, guard
}

func SolveP1() {
	lines := utils.ReadFile("./day6/input.txt")

	currentMap, guard := createMap(lines)

	moveGuard(guard, &currentMap)

	visitedCells := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			pos := coordinate{x: j, y: i}
			cell := currentMap[pos]

			if cell.visited > 0 {
				visitedCells++
			}
		}
	}

	printMap(&currentMap)

	fmt.Println("visited cells:", visitedCells)
}

func SolveP2() {
	lines := utils.ReadFile("./day6/input.txt")

	currentMap, guard := createMap(lines)

	doesLoop := moveGuard(guard, &currentMap)

	var visitedCells []coordinate

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			pos := coordinate{x: j, y: i}
			cell := currentMap[pos]

			if cell.visited > 0 {
				visitedCells = append(visitedCells, pos)
			}
		}
	}

	numberOfLoops := 0

	wg := sync.WaitGroup{}
	wg.Add(len(visitedCells))
	mutex := sync.Mutex{}

	for _, visitedCell := range visitedCells {
		go func() {
			defer wg.Done()
			copyMap := make(map[coordinate]cell)
			copyMap, _ = createMap(lines)

			if visitedCell == guard.coordinate {
				return
			}

			copyMap[visitedCell] = cell{visited: 0, obstructs: true}

			doesLoop = moveGuard(guard, &copyMap)

			if doesLoop {
				mutex.Lock()
				numberOfLoops++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println("number of loops:", numberOfLoops)
}

func printMap(currentMap *map[coordinate]cell) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			pos := coordinate{x: j, y: i}
			cell := (*currentMap)[pos]

			if cell.obstructs {
				fmt.Printf("#")
			} else if cell.visited >= 1 {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getNextCell(currentPosition coordinate, facingDirection direction, currentMap *map[coordinate]cell) (obstructed, exit bool, nextCell coordinate) {

	var nextPosition coordinate

	switch facingDirection {
	case up:
		nextPosition = coordinate{x: currentPosition.x, y: currentPosition.y - 1}
	case down:
		nextPosition = coordinate{x: currentPosition.x, y: currentPosition.y + 1}
	case left:
		nextPosition = coordinate{x: currentPosition.x - 1, y: currentPosition.y}
	case right:
		nextPosition = coordinate{x: currentPosition.x + 1, y: currentPosition.y}
	}

	cell, isWithingTheBorder := (*currentMap)[nextPosition]

	if !isWithingTheBorder {
		return false, true, nextPosition
	}

	return cell.obstructs, false, nextPosition
}

func rotateGuard(g guard) direction {
	switch g.facingDirection {
	case up:
		return right
	case down:
		return left
	case left:
		return up
	case right:
		return down
	}
	return ""
}

func isGuardLooping(currentMap *map[coordinate]cell, nextCell coordinate, g guard) bool {

	if (*currentMap)[nextCell].visited > 1 {
		if (*currentMap)[nextCell].guardWasFacing == g.facingDirection {
			return true
		}
	}

	return false
}

func moveGuard(g guard, currentMap *map[coordinate]cell) (loop bool) {

	moveToNextCell := func(nextCell coordinate) {
		(*currentMap)[nextCell] = cell{visited: (*currentMap)[nextCell].visited + 1, obstructs: false, guardWasFacing: g.facingDirection}
		g.coordinate = nextCell
	}

	isObstructed, isExit, nextCell := getNextCell(g.coordinate, g.facingDirection, currentMap)
	if isExit {
		return false
	}

	if isObstructed {
		g.facingDirection = rotateGuard(g)
	} else {
		if isGuardLooping(currentMap, nextCell, g) {
			return true
		}

		moveToNextCell(nextCell)
	}

	return moveGuard(g, currentMap)
}
