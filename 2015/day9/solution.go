package day9

import (
	"aoc2022/utils"
	"fmt"
	"slices"
)

type destination struct {
	name     string
	distance int
}

type location struct {
	destinations []destination
}

func SolveP1() {
	lines := utils.ReadFile("day9/input.txt")

	locations := createPossibleRoutes(lines)

	distancesList := make([]int, 0)

	for k, v := range locations {
		for _, toLocation := range v.destinations {
			createDistancesList(k, locations, []string{toLocation.name}, 0, &distancesList)
		}
	}

	fmt.Println("Number of possible routes:", len(distancesList))
	fmt.Println("Shortest distance was:", slices.Min(distancesList))
}

func SolveP2() {
	lines := utils.ReadFile("day9/input.txt")

	locations := createPossibleRoutes(lines)

	distancesList := make([]int, 0)

	for k, v := range locations {
		for _, toLocation := range v.destinations {
			createDistancesList(k, locations, []string{toLocation.name}, 0, &distancesList)
		}
	}

	fmt.Println("Number of possible routes:", len(distancesList))
	fmt.Println("Longest distance was:", slices.Max(distancesList))
}

func parseLine(line string) (d1, d2 string, dist int) {
	var destination1, destination2 string
	var distance int

	fmt.Sscanf(line, "%s to %s = %d", &destination1, &destination2, &distance)

	return destination1, destination2, distance
}

func createPossibleRoutes(lines []string) map[string]location {
	locations := make(map[string]location)
	for _, line := range lines {
		destination1, destination2, distance := parseLine(line)

		loc1 := locations[destination1]
		loc2 := locations[destination2]

		loc1.destinations = append(loc1.destinations, destination{
			name:     destination2,
			distance: distance,
		})

		loc2.destinations = append(loc2.destinations, destination{
			name:     destination1,
			distance: distance,
		})

		locations[destination1] = loc1
		locations[destination2] = loc2
	}

	return locations
}

func createDistancesList(fromLocation string, wholeMap map[string]location, alreadyVisitedLocations []string, totalLength int, distancesList *[]int) {

	currLocation := wholeMap[fromLocation]

	toVisit := []destination{}
	for _, loc := range currLocation.destinations {
		if !slices.ContainsFunc(alreadyVisitedLocations, func(x string) bool {
			return x == loc.name
		}) {
			toVisit = append(toVisit, loc)
		}
	}

	if len(toVisit) > 0 {
		for _, dest := range toVisit {
			cannotVisit := append(alreadyVisitedLocations, fromLocation)
			createDistancesList(dest.name, wholeMap, cannotVisit, totalLength+dest.distance, distancesList)
		}
	} else {
		destIdx := slices.IndexFunc(currLocation.destinations, func(x destination) bool {
			return x.name == alreadyVisitedLocations[0]
		})
		dest := currLocation.destinations[destIdx]

		t := totalLength + dest.distance
		destination := alreadyVisitedLocations[0]
		a := append(alreadyVisitedLocations[1:], fromLocation, destination)

		for idx := range a {
			if idx == len(a)-1 {
				*distancesList = append(*distancesList, t)
				continue
			}
		}
	}
}

func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}
