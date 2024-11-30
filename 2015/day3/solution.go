package day3

import (
	"aoc2022/utils"
	"fmt"
	"slices"
)

type coordinate struct {
	x int
	y int
}

func SolveP1() {
	sleighPosition := coordinate{
		x: 0,
		y: 0,
	}

	lines := utils.ReadFile("day3/input.txt")

	var housesVisited []coordinate = []coordinate{
		sleighPosition,
	}

	for _, direction := range lines[0] {
		moveSleigh(direction, &sleighPosition)

		housesVisited = append(housesVisited, sleighPosition)
	}

	fmt.Printf("Santa visisted %d houses\n", countUniqueHouses(housesVisited))
}

func SolveP2() {
	var santasSleighPosition, roboSantaSleighPosition coordinate = coordinate{
		x: 0,
		y: 0,
	}, coordinate{
		x: 0,
		y: 0,
	}

	lines := utils.ReadFile("day3/input.txt")

	var housesVisitedBySanta []coordinate = []coordinate{
		santasSleighPosition,
	}

	var housesVisitedByRoboSanta []coordinate = []coordinate{
		roboSantaSleighPosition,
	}

	santasTurn := true
	for _, direction := range lines[0] {

		if santasTurn {
			moveSleigh(direction, &santasSleighPosition)
		} else {
			moveSleigh(direction, &roboSantaSleighPosition)
		}

		switch santasTurn {
		case true:
			housesVisitedBySanta = append(housesVisitedBySanta, santasSleighPosition)
		case false:
			housesVisitedByRoboSanta = append(housesVisitedByRoboSanta, roboSantaSleighPosition)
		}

		// switch turns
		santasTurn = !santasTurn
	}

	fmt.Printf("Santa visisted %d houses\n", countUniqueHouses(housesVisitedBySanta))
	fmt.Printf("Robo santa visited %d houses\n", countUniqueHouses(housesVisitedByRoboSanta))
	fmt.Printf("In total they visited %d houses\n", countUniqueHouses(slices.Concat(housesVisitedBySanta, housesVisitedByRoboSanta)))
}

func moveSleigh(direction rune, sleigh *coordinate) {
	switch direction {
	case '^':
		sleigh.y++
	case '>':
		sleigh.x++
	case 'v':
		sleigh.y--
	case '<':
		sleigh.x--
	}
}

func houseAlreadyVisited(houses []coordinate, house coordinate) bool {
	for _, h := range houses {
		if h == house {
			return true
		}
	}

	return false
}

func countUniqueHouses(houses []coordinate) int {
	uniqueHouses := []coordinate{}
	for _, house := range houses {
		if !houseAlreadyVisited(uniqueHouses, house) {
			uniqueHouses = append(uniqueHouses, house)
		}
	}

	return len(uniqueHouses)
}
