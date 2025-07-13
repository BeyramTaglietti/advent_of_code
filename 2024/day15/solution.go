package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

type Position struct {
	y, x int
}

func readInput(filename string) ([][]rune, []rune, Position) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	moves := []rune{}
	mapFinished := false
	var start Position

	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			mapFinished = true
			continue
		}
		if mapFinished {
			moves = append(moves, []rune(line)...)
		} else {
			row := []rune(line)
			for x, ch := range row {
				if ch == '@' {
					start = Position{y, x}
				}
			}
			grid = append(grid, row)
			y++
		}
	}

	return grid, moves, start
}

func move(grid [][]rune, y, x int, dir Direction) (bool, int, int) {
	if grid[y][x] == '#' {
		return false, 0, 0
	}

	ny, nx := y, x
	switch dir {
	case LEFT:
		nx--
	case RIGHT:
		nx++
	case UP:
		ny--
	case DOWN:
		ny++
	}

	if grid[y][x] == '.' {
		return true, ny, nx
	}

	canMove, _, _ := move(grid, ny, nx, dir)
	if canMove {
		grid[ny][nx] = grid[y][x]
		grid[y][x] = '.'
	}
	return canMove, ny, nx
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, ch := range row {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}

func sumBoxesCoordinates(grid [][]rune) int {
	sum := 0
	for y, row := range grid {
		for x, ch := range row {
			if ch == 'O' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func SolveP1() {
	grid, moves, start := readInput("./day15/input.txt")
	ry, rx := start.y, start.x
	fmt.Println("starting position", ry, rx)

	for _, mv := range moves {
		dir := UP
		switch mv {
		case '<':
			dir = LEFT
		case '^':
			dir = UP
		case '>':
			dir = RIGHT
		case 'v':
			dir = DOWN
		}
		canMove, cy, cx := move(grid, ry, rx, dir)
		if canMove {
			ry, rx = cy, cx
		}
	}

	printGrid(grid)
	fmt.Println("ending position", ry, rx, "with coordinates", 100*ry+rx)
	fmt.Println("the sum of all boxes coordinates is", sumBoxesCoordinates(grid))
}
