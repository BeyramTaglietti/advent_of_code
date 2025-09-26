package day10

import (
	"2018/utils"
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func SolveP1() {
	lines := utils.ReadFile("day10/input.txt")

	lights := make(map[point]point, len(lines))

	for _, line := range lines {
		position, velocity := parseLine(line)
		lights[position] = velocity
	}

	prevArea := math.MaxInt
	prevLights := lights
	for {
		minX, maxX, minY, maxY := bounds(lights)
		area := (maxX - minX + 1) * (maxY - minY + 1)
		if area > prevArea {
			printMap(prevLights)
			return
		}
		prevArea = area
		prevLights = lights
		lights = step(lights)
	}
}

func SolveP2() {
	lines := utils.ReadFile("day10/input.txt")

	lights := make(map[point]point, len(lines))

	for _, line := range lines {
		position, velocity := parseLine(line)
		lights[position] = velocity
	}

	seconds := 0
	prevArea := math.MaxInt
	prevLights := lights
	for {
		minX, maxX, minY, maxY := bounds(lights)
		area := (maxX - minX + 1) * (maxY - minY + 1)
		if area > prevArea {
			printMap(prevLights)
			fmt.Println("seconds", seconds-1)
			return
		}
		prevArea = area
		prevLights = lights
		lights = step(lights)
		seconds++
	}
}

func parseLine(line string) (position point, velocity point) {
	fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &position.x, &position.y, &velocity.x, &velocity.y)
	return position, velocity
}

func step(lights map[point]point) map[point]point {
	newLights := make(map[point]point, len(lights))
	for pos, vel := range lights {
		newPos := point{pos.x + vel.x, pos.y + vel.y}
		newLights[newPos] = vel
	}
	return newLights
}

func bounds(lights map[point]point) (int, int, int, int) {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for pos := range lights {
		if pos.x < minX {
			minX = pos.x
		}
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y < minY {
			minY = pos.y
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	return minX, maxX, minY, maxY
}

func printMap(lights map[point]point) {
	minX, maxX, minY, maxY := bounds(lights)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := lights[point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
