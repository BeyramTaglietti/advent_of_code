package day11

import (
	"fmt"
	"strconv"
)

type point struct {
	x, y int
}

func SolveP1() {

	serialNumber := 18

	maxFuel := 0
	maxFuelCell := point{0, 0}

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			v := getSquareFuel(y, x, serialNumber)

			if v > maxFuel {
				maxFuel = v
				maxFuelCell = point{y, x}
			}
		}
	}

	fmt.Println("result: ", maxFuelCell, maxFuel)
}

func (p point) getPowerLevel(serialNumber int) int {
	sum := 0

	rackId := p.x + 10

	sum += rackId
	sum *= p.y
	sum += serialNumber
	sum *= rackId

	hundredStr := strconv.Itoa(sum)
	if len(hundredStr) < 3 {
		return 0
	}
	hundredInt, _ := strconv.Atoi(string(hundredStr[len(hundredStr)-3]))

	return hundredInt - 5
}

func getSquareFuel(y, x, serialNumber int) int {
	sum := 0
	for iy := 0; iy < 3; iy++ {
		for ix := 0; ix < 3; ix++ {
			sum += point{y: iy + y, x: ix + x}.getPowerLevel(serialNumber)
		}
	}

	return sum
}
