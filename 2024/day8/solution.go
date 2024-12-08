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

func createGrid(lines []string) (antennasMap map[rune][]point, grid map[point]rune) {
	antennasMap = make(map[rune][]point)
	grid = make(map[point]rune)

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
	return antennasMap, grid
}

func SolveP1() {

	lines := utils.ReadFile("./day8/input.txt")

	antennasMap, grid := createGrid(lines)

	mapLimits := mapLimit{
		maxX: len(lines[0]),
		maxY: len(lines),
	}

	totalAntennas := []point{}
	for _, v := range antennasMap {
		antiNodes := []point{}
		for idx, anten := range v {
			antiNodes = slices.Concat(antiNodes, findAntennasCouples(anten, v[idx+1:], []point{}, mapLimits, false))
		}

		// unique anti nodes
		antiNodes = slices.DeleteFunc(antiNodes, func(p point) bool {
			if slices.Contains(totalAntennas, p) {
				return true
			}
			return false
		})

		totalAntennas = slices.Concat(totalAntennas, antiNodes)

		for _, node := range antiNodes {
			if _, exists := grid[node]; !exists {
				grid[node] = '#'
			}
		}
	}

	printMap(grid, mapLimits)
	fmt.Println("\nAntinodes count", len(totalAntennas))
}

func SolveP2() {
	lines := utils.ReadFile("./day8/input.txt")

	antennasMap, grid := createGrid(lines)

	mapLimits := mapLimit{
		maxX: len(lines[0]),
		maxY: len(lines),
	}

	totalAntinodes := []point{}
	for _, v := range antennasMap {
		antiNodes := []point{}
		for idx, anten := range v {
			antiNodes = slices.Concat(antiNodes, findAntennasCouples(anten, v[idx+1:], []point{}, mapLimits, true))
		}

		antiNodes = slices.Concat(antiNodes, v)

		// unique anti nodes
		antiNodes = slices.DeleteFunc(antiNodes, func(p point) bool {
			if slices.Contains(totalAntinodes, p) {
				return true
			}
			return false
		})

		totalAntinodes = slices.Concat(totalAntinodes, antiNodes)
		for _, node := range antiNodes {
			if _, exists := grid[node]; !exists {
				grid[node] = '#'
			}
		}
	}

	printMap(grid, mapLimits)
	fmt.Println("\nAntinodes count", len(totalAntinodes))
}

func isOutOfBounds(p point, mapLimits mapLimit) bool {
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

func findAntennasCouples(antenna1 point, otherAntennas []point, antiNodes []point, mapLimits mapLimit, includeResonants bool) []point {

	if len(otherAntennas) == 0 {
		return antiNodes
	}

	antiNodes = slices.Concat(antiNodes, getAntinodes(antenna1, otherAntennas[0], mapLimits, includeResonants))

	return findAntennasCouples(antenna1, otherAntennas[1:], antiNodes, mapLimits, includeResonants)
}

func getAntinodes(antenna1, antenna2 point, mapLimits mapLimit, includeResonants bool) (antiNodes []point) {

	for i := 1; ; i++ {
		if !includeResonants && i > 1 {
			break
		}

		x1 := antenna1.x + ((antenna1.x - antenna2.x) * i)
		y1 := antenna1.y + ((antenna1.y - antenna2.y) * i)
		x2 := antenna2.x + ((antenna2.x - antenna1.x) * i)
		y2 := antenna2.y + ((antenna2.y - antenna1.y) * i)
		resonantAntenna1 := point{x: x1, y: y1}
		resonantAntenna2 := point{x: x2, y: y2}

		p1Valid := isOutOfBounds(resonantAntenna1, mapLimits)
		p2Valid := isOutOfBounds(resonantAntenna2, mapLimits)

		if !p1Valid && !p2Valid {
			break
		}

		if p1Valid {
			antiNodes = append(antiNodes, resonantAntenna1)
		}

		if p2Valid {
			antiNodes = append(antiNodes, resonantAntenna2)
		}

	}

	return antiNodes
}
