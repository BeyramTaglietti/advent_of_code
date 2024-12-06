package day6

import (
	"aoc2024/utils"
	"fmt"
)

type coordinate struct {
	x int
	y int
}

type cell struct {
	visited   bool
	obstructs bool
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

// . nothing
// # obstruction
// ^ guard

func SolveP1() {
	lines := utils.ReadFile("./day6/input.txt")

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
				currentMap[position] = cell{visited: false, obstructs: false}
			case '#':
				currentMap[position] = cell{visited: false, obstructs: true}
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
				currentMap[position] = cell{visited: true, obstructs: false}
			}
		}
	}

	moveGuard(guard, &currentMap)

	visitedCells := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			pos := coordinate{x: j, y: i}
			cell := currentMap[pos]
			if cell.obstructs {
				fmt.Printf("#")
			} else if cell.visited {
				fmt.Printf("X")
				visitedCells++
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println()
	}

	fmt.Println("visited cells:", visitedCells)

}

func moveGuard(g guard, currentMap *map[coordinate]cell) {

	getNextCell := func(nextPosition coordinate) (obstructed, exit bool) {
		cell, isWithingTheBorder := (*currentMap)[nextPosition]

		if !isWithingTheBorder {
			return false, true
		}

		return cell.obstructs, false
	}

	moveNextCell := func(nextCell coordinate) {
		(*currentMap)[nextCell] = cell{visited: true, obstructs: false}
		g.coordinate = nextCell
	}

	var nextCell coordinate

	switch g.facingDirection {
	case up:
		nextCell = coordinate{x: g.x, y: g.y - 1}
		isObstructed, isExit := getNextCell(nextCell)

		if isExit {
			return
		}

		if isObstructed {
			g.facingDirection = right
		} else {
			moveNextCell(nextCell)
		}

	case down:
		nextCell = coordinate{x: g.x, y: g.y + 1}
		isObstructed, isExit := getNextCell(nextCell)
		if isExit {
			return
		}

		if isObstructed {
			g.facingDirection = left
		} else {
			moveNextCell(nextCell)
		}

	case left:
		nextCell = coordinate{x: g.x - 1, y: g.y}
		isObstructed, isExit := getNextCell(nextCell)
		if isExit {
			return
		}

		if isObstructed {
			g.facingDirection = up
		} else {
			moveNextCell(nextCell)
		}

	case right:
		nextCell = coordinate{x: g.x + 1, y: g.y}
		isObstructed, isExit := getNextCell(nextCell)
		if isExit {
			return
		}

		if isObstructed {
			g.facingDirection = down
		} else {
			moveNextCell(nextCell)
		}

	}
	moveGuard(g, currentMap)
}
