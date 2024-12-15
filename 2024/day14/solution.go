package day14

import (
	"aoc2024/utils"
	"fmt"
	"math"
)

type point struct {
	x, y int
}

type velocity struct {
	vx, vy int
}

type robot struct {
	point
	velocity
}

type quadrant struct {
	minX int
	maxX int
	minY int
	maxY int
}

func SolveP1() {
	lines := utils.ReadFile("./day14/input.txt")

	const limitX, limitY = 101, 103

	robots := parseRobots(lines)

	const seconds = 100

	for range seconds {
		for idx, r := range robots {
			robots[idx] = move(r, limitX, limitY)
		}

	}

	q1l, q2l, q3l, q4l := getQuadrants(limitX, limitY, robots)

	fmt.Println("Q1:", q1l)
	fmt.Println("Q2:", q2l)
	fmt.Println("Q3:", q3l)
	fmt.Println("Q4:", q4l)

	fmt.Println("Safety factor:", q1l*q2l*q3l*q4l)

}

func SolveP2() {
	lines := utils.ReadFile("./day14/input.txt")

	const limitX, limitY = 101, 103

	robots := parseRobots(lines)

	const seconds = 10000000

	for i := range seconds {
		for idx, r := range robots {
			robots[idx] = move(r, limitX, limitY)
		}

		if everyRobotIsAtUniquePosition(robots) {
			printGrid(limitX, limitY, robots)
			fmt.Println("Seconds:", i+1)
			break
		}

	}

}

func parseRobots(lines []string) []robot {

	robots := []robot{}

	for _, line := range lines {
		x, y, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		robots = append(robots, robot{point{x, y}, velocity{vx, vy}})
	}

	return robots
}

func printGrid(limitX, limitY int, robots []robot) {
	grid := make([][]bool, limitY)

	for i := 0; i < limitY; i++ {
		grid[i] = make([]bool, limitX)
	}

	for _, r := range robots {
		grid[r.y][r.x] = true
	}

	for y := 0; y < limitY; y++ {
		for x := 0; x < limitX; x++ {
			switch grid[y][x] {
			case true:
				fmt.Print("#")
			case false:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func move(r robot, limitX, limitY int) robot {

	if r.x+r.vx >= limitX {
		r.x = 0 + r.vx - (limitX - r.x)
	} else if r.x+r.vx < 0 {
		r.x = (limitX + r.x) + r.vx
	} else {
		r.x += r.vx
	}

	if r.y+r.vy >= limitY {
		r.y = 0 + r.vy - (limitY - r.y)
	} else if r.y+r.vy < 0 {
		r.y = (limitY + r.y) + r.vy
	} else {
		r.y += r.vy
	}

	return r
}

func getQuadrants(limitX, limitY int, robots []robot) (int, int, int, int) {
	floorX := int(math.Floor(float64(limitX) / 2))
	ceilX := int(math.Ceil(float64(limitX) / 2))
	floorY := int(math.Floor(float64(limitY) / 2))
	ceilY := int(math.Ceil(float64(limitY) / 2))

	q1, q2, q3, q4 :=
		quadrant{0, floorX, 0, floorY},
		quadrant{ceilX, limitX, 0, floorY},
		quadrant{0, floorX, ceilY, limitY},
		quadrant{ceilX, limitX, ceilY, limitY}

	q1l, q2l, q3l, q4l := 0, 0, 0, 0

	for _, r := range robots {
		if r.x >= q1.minX && r.x < q1.maxX && r.y >= q1.minY && r.y < q1.maxY {
			q1l++
		} else if r.x >= q2.minX && r.x < q2.maxX && r.y >= q2.minY && r.y < q2.maxY {
			q2l++
		} else if r.x >= q3.minX && r.x < q3.maxX && r.y >= q3.minY && r.y < q3.maxY {
			q3l++
		} else if r.x >= q4.minX && r.x < q4.maxX && r.y >= q4.minY && r.y < q4.maxY {
			q4l++
		}
	}

	return q1l, q2l, q3l, q4l
}

func everyRobotIsAtUniquePosition(robots []robot) bool {
	positions := make(map[point]bool)

	for _, r := range robots {
		if _, ok := positions[r.point]; ok {
			return false
		}
		positions[r.point] = true
	}

	return true
}
