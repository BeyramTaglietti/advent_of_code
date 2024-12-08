package day8

import (
	"aoc2024/utils"
	"fmt"
	"slices"
)

type point struct {
	x int
	y int
}

type mapLimit struct {
	maxX int
	maxY int
}

func SolveP1() {

	lines := utils.ReadFile("./day8/input.txt")

	antennasMap := make(map[rune][]point)
	grid := make(map[point]rune)

	mapLimits := mapLimit{
		maxX: len(lines[0]),
		maxY: len(lines),
	}

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				antennasMap[char] = append(antennasMap[char], point{
					x: x,
					y: y,
				})
				grid[point{x, y}] = char
			}
		}
	}

	printMap(grid, mapLimits)

	totalAntennas := []point{}
	for _, v := range antennasMap {
		antiNodes := []point{}
		for idx, anten := range v {
			antiNodes = slices.Concat(antiNodes, findAntennasCouples(anten, v[idx+1:], []point{}))
		}
		antiNodes = slices.DeleteFunc(antiNodes, func(p point) bool {

			if slices.Contains(totalAntennas, p) {
				return true
			}

			return !isValid(p, mapLimits)
		})
		totalAntennas = slices.Concat(totalAntennas, antiNodes)
		for _, node := range antiNodes {
			if _, exists := grid[node]; !exists {
				grid[node] = '#'
			}
		}
	}

	fmt.Println("Antinodes", totalAntennas)
	fmt.Println("Antinodes count", len(totalAntennas))
	printMap(grid, mapLimits)
}

func isValid(p point, mapLimits mapLimit) bool {

	return p.x >= 0 && p.y >= 0 && p.x < mapLimits.maxX && p.y < mapLimits.maxY
}

func printMap(grid map[point]rune, mapLimits mapLimit) {
	for y := 0; y < mapLimits.maxY; y++ {
		for x := 0; x < mapLimits.maxX; x++ {
			if val, exists := grid[point{x, y}]; exists {
				fmt.Printf("%c", val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findAntennasCouples(antenna1 point, otherAntennas []point, antiNodes []point) []point {

	if len(otherAntennas) == 0 {
		return antiNodes
	}

	antiNode1, antiNode2 := getAntinodes(antenna1, otherAntennas[0])

	antiNodes = append(antiNodes, antiNode1, antiNode2)

	return findAntennasCouples(antenna1, otherAntennas[1:], antiNodes)
}

func getAntinodes(antenna1, antenna2 point) (antiNode1, antiNode2 point) {
	antiNode1.x = antenna1.x + (antenna1.x - antenna2.x)
	antiNode1.y = antenna1.y + (antenna1.y - antenna2.y)

	antiNode2.x = antenna2.x + (antenna2.x - antenna1.x)
	antiNode2.y = antenna2.y + (antenna2.y - antenna1.y)

	return antiNode1, antiNode2
}
