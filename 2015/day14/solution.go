package day14

import (
	"aoc2015/utils"
	"cmp"
	"fmt"
	"slices"
)

type horse struct {
	name         string
	speed        int
	endurance    int
	restRequired int

	runTime  int
	restTime int

	points   int
	distance int
}

func (h *horse) run1Second() {

	if h.runTime < h.endurance {
		h.runTime++

		h.distance += h.speed
		return
	}

	if h.restTime < h.restRequired {
		h.restTime++
		return
	}

	h.runTime = 1
	h.restTime = 0

	h.distance += h.speed
}

func SolveP1() {

	lines := utils.ReadFile("./day14/input.txt")

	horses := make([]horse, len(lines), len(lines))

	for idx, line := range lines {
		horses[idx] = parseLine(line)
	}

	raceDuration := 2503

	for s := 0; s < raceDuration; s++ {

		for idx := range horses {
			horses[idx].run1Second()
		}
	}

	for _, horse := range horses {
		fmt.Println("after", raceDuration, "seconds,", horse.name, "reached", horse.distance, "km")
	}
}

func SolveP2() {

	lines := utils.ReadFile("./day14/input.txt")

	horses := make([]horse, len(lines), len(lines))

	for idx, line := range lines {
		horses[idx] = parseLine(line)
	}

	raceDuration := 2503

	for s := 1; s <= raceDuration; s++ {

		max := -1

		for idx := range horses {
			horses[idx].run1Second()
			if horses[idx].distance > max {
				max = horses[idx].distance
			}
		}

		for idx := range horses {
			if horses[idx].distance == max {
				horses[idx].points++
			}
		}

	}

	fmt.Println("the horse with the most points has", slices.MaxFunc(horses, func(a horse, b horse) int {
		return cmp.Compare(a.points, b.points)
	}).points, "points")
}

func parseLine(line string) horse {
	var speed, endurance, rest int
	var name string

	fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds", &name, &speed, &endurance, &rest)

	return horse{
		speed:        speed,
		endurance:    endurance,
		restRequired: rest,
		name:         name,
	}
}
