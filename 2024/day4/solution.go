package day4

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type grid map[coordinate]rune

func SolveP1() {

	lines := utils.ReadFile("./day4/input.txt")

	fullGrid := make(map[coordinate]rune)

	for y, line := range lines {
		for x, char := range line {
			fullGrid[coordinate{
				x: x,
				y: y,
			}] = char
		}
	}

	correctXmasFound := 0

	for y, line := range lines {
		for x, char := range line {
			if char == 'X' {

				position := coordinate{
					x: x,
					y: y,
				}

				correctXmasFound += findHorizontally(fullGrid, position)
				correctXmasFound += findVertically(fullGrid, position)
				correctXmasFound += findDiagonally(fullGrid, position)

			}
		}
	}

	fmt.Println("Correct xmas found:", correctXmasFound)
}

func SolveP2() {

	lines := utils.ReadFile("./day4/input.txt")

	fullGrid := make(map[coordinate]rune)

	for y, line := range lines {
		for x, char := range line {
			fullGrid[coordinate{
				x: x,
				y: y,
			}] = char
		}
	}

	correctXmasFound := 0

	for y, line := range lines {
		for x, char := range line {
			if char == 'M' {

				position := coordinate{
					x: x,
					y: y,
				}

				correctXmasFound += findXMas(fullGrid, position)

			}
		}
	}

	fmt.Println("Correct X-MAS found:", correctXmasFound)
}

func findHorizontally(fullGrid grid, startingPosition coordinate) int {

	count := 0

	xPos := startingPosition.x
	mPos := startingPosition.x + 1
	aPos := startingPosition.x + 2
	sPos := startingPosition.x + 3

	// check to the right
	if fullGrid[coordinate{
		x: xPos,
		y: startingPosition.y,
	}] == 'X' &&
		fullGrid[coordinate{
			x: mPos,
			y: startingPosition.y,
		}] == 'M' &&
		fullGrid[coordinate{
			x: aPos,
			y: startingPosition.y,
		}] == 'A' &&
		fullGrid[coordinate{
			x: sPos,
			y: startingPosition.y,
		}] == 'S' {

		count++
	}

	xNeg := startingPosition.x
	mNeg := startingPosition.x - 1
	aNeg := startingPosition.x - 2
	sNeg := startingPosition.x - 3

	// check to the left
	if fullGrid[coordinate{
		x: xNeg,
		y: startingPosition.y,
	}] == 'X' &&
		fullGrid[coordinate{
			x: mNeg,
			y: startingPosition.y,
		}] == 'M' &&
		fullGrid[coordinate{
			x: aNeg,
			y: startingPosition.y,
		}] == 'A' &&
		fullGrid[coordinate{
			x: sNeg,
			y: startingPosition.y,
		}] == 'S' {

		count++
	}

	return count
}

func findVertically(fullGrid grid, startingPosition coordinate) int {

	count := 0

	xPos := startingPosition.y
	mPos := startingPosition.y + 1
	aPos := startingPosition.y + 2
	sPos := startingPosition.y + 3

	if fullGrid[coordinate{
		x: startingPosition.x,
		y: xPos,
	}] == 'X' &&
		fullGrid[coordinate{
			x: startingPosition.x,
			y: mPos,
		}] == 'M' &&
		fullGrid[coordinate{
			x: startingPosition.x,
			y: aPos,
		}] == 'A' &&
		fullGrid[coordinate{
			x: startingPosition.x,
			y: sPos,
		}] == 'S' {

		count++
	}

	xNeg := startingPosition.y
	mNeg := startingPosition.y - 1
	aNeg := startingPosition.y - 2
	sNeg := startingPosition.y - 3

	if fullGrid[coordinate{
		x: startingPosition.x,
		y: xNeg,
	}] == 'X' &&
		fullGrid[coordinate{
			x: startingPosition.x,
			y: mNeg,
		}] == 'M' &&
		fullGrid[coordinate{
			x: startingPosition.x,
			y: aNeg,
		}] == 'A' &&
		fullGrid[coordinate{
			x: startingPosition.x,
			y: sNeg,
		}] == 'S' {

		count++
	}

	return count
}

func findDiagonally(fullGrid grid, startingPosition coordinate) int {

	count := 0

	var line strings.Builder

	// top

	// top right
	for i := 0; i <= 3; i++ {
		line.WriteRune(fullGrid[coordinate{
			x: startingPosition.x + i,
			y: startingPosition.y - i,
		}])
	}

	if line.String() == "XMAS" {
		count++
	}

	line.Reset()

	// top left
	for i := 0; i <= 3; i++ {
		line.WriteRune(fullGrid[coordinate{
			x: startingPosition.x - i,
			y: startingPosition.y - i,
		}])
	}

	if line.String() == "XMAS" {
		count++
	}

	// bottom
	line.Reset()

	// bottom right
	for i := 0; i <= 3; i++ {
		line.WriteRune(fullGrid[coordinate{
			x: startingPosition.x + i,
			y: startingPosition.y + i,
		}])
	}

	if line.String() == "XMAS" {
		count++
	}

	line.Reset()

	// bottom left
	for i := 0; i <= 3; i++ {
		line.WriteRune(fullGrid[coordinate{
			x: startingPosition.x - i,
			y: startingPosition.y + i,
		}])
	}

	if line.String() == "XMAS" {
		count++
	}

	return count
}

func findXMas(fullGrid grid, startingPosition coordinate) int {

	count := 0

	firstMPosition := startingPosition
	var secondMPosition coordinate
	var aPos, s1Pos, s2Pos coordinate

	// second M is on the right
	if fullGrid[coordinate{x: firstMPosition.x + 2, y: firstMPosition.y}] == 'M' {

		secondMPosition = coordinate{x: firstMPosition.x + 2, y: firstMPosition.y}

		// regular X
		aPos = coordinate{x: firstMPosition.x + 1, y: firstMPosition.y + 1}
		s1Pos, s2Pos = coordinate{x: firstMPosition.x, y: firstMPosition.y + 2}, coordinate{x: secondMPosition.x, y: secondMPosition.y + 2}

		if fullGrid[aPos] == 'A' && fullGrid[s1Pos] == 'S' && fullGrid[s2Pos] == 'S' {
			count++
		}

		// inverted X
		aPos = coordinate{x: firstMPosition.x + 1, y: firstMPosition.y - 1}
		s1Pos, s2Pos = coordinate{x: firstMPosition.x, y: firstMPosition.y - 2}, coordinate{x: secondMPosition.x, y: secondMPosition.y - 2}

		if fullGrid[aPos] == 'A' && fullGrid[s1Pos] == 'S' && fullGrid[s2Pos] == 'S' {
			count++
		}
	}

	// second M is below
	if fullGrid[coordinate{x: firstMPosition.x, y: firstMPosition.y + 2}] == 'M' {
		secondMPosition = coordinate{x: firstMPosition.x, y: firstMPosition.y + 2}

		// regular X
		aPos = coordinate{x: firstMPosition.x + 1, y: firstMPosition.y + 1}
		s1Pos, s2Pos = coordinate{x: firstMPosition.x + 2, y: firstMPosition.y}, coordinate{x: secondMPosition.x + 2, y: secondMPosition.y}

		if fullGrid[aPos] == 'A' && fullGrid[s1Pos] == 'S' && fullGrid[s2Pos] == 'S' {
			count++
		}

		// inverted X
		aPos = coordinate{x: firstMPosition.x - 1, y: firstMPosition.y + 1}
		s1Pos, s2Pos = coordinate{x: firstMPosition.x - 2, y: firstMPosition.y}, coordinate{x: secondMPosition.x - 2, y: secondMPosition.y}

		if fullGrid[aPos] == 'A' && fullGrid[s1Pos] == 'S' && fullGrid[s2Pos] == 'S' {
			count++
		}
	}

	return count
}
