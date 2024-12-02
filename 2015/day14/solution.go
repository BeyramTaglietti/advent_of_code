package day14

import (
	"fmt"
)

type horse struct {
	speed     int
	endurance int
	restTime  int
}

func (h horse) run(duration int) int {

	den := h.restTime + h.endurance
	quotient, remainder := duration/den, duration%den

	distance := quotient * h.speed * h.endurance

	if remainder > 0 {
		distance += (h.speed * h.endurance)
	}

	return distance
}

func SolveP1() {

	Comet := horse{
		speed:     14,
		endurance: 10,
		restTime:  127,
	}

	Dancer := horse{
		speed:     16,
		endurance: 11,
		restTime:  162,
	}

	raceDuration := 2503

	fmt.Println("after", raceDuration, "seconds, Comet reached", Comet.run(raceDuration), "km")
	fmt.Println("after", raceDuration, "seconds, Dancer reached", Dancer.run(raceDuration), "km")
}
