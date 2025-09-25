package day3

import (
	"2018/utils"
	"fmt"
	"math"
)

func SolveP1() {
	lines := utils.ReadFile("day3/input.txt")

	var maxY, maxX int = 0, 0

	for _, line := range lines {
		claim := parseLine(line)

		maxY = int(math.Max(float64(maxY), float64(claim.y)+float64(claim.height)))
		maxX = int(math.Max(float64(maxX), float64(claim.x)+float64(claim.width)))
	}

	cloth := make([][]int, maxY)
	for y := 0; y < maxY; y++ {
		cloth[y] = make([]int, maxX)
	}

	intersectingFabric := 0

	for _, line := range lines {
		claim := parseLine(line)

		for y := claim.y; y < claim.height+claim.y; y++ {
			for x := claim.x; x < claim.width+claim.x; x++ {
				if cloth[y][x] == 1 {
					intersectingFabric++
				}
				cloth[y][x]++
			}
		}
	}

	fmt.Println("result:", intersectingFabric)
}

func SolveP2() {
	lines := utils.ReadFile("day3/input.txt")
	claims := make([]claim, len(lines))

	for i, line := range lines {
		claims[i] = parseLine(line)
	}

	overlapping := make(map[int]bool)

	for i := 0; i < len(claims); i++ {
		for j := i + 1; j < len(claims); j++ {
			if intersects(claims[i], claims[j]) {
				overlapping[claims[i].id] = true
				overlapping[claims[j].id] = true
			}
		}
	}

	for _, claim := range claims {
		if !overlapping[claim.id] {
			fmt.Println("result:", claim.id)
			return
		}
	}
}

func intersects(a, b claim) bool {
	return a.x < b.x+b.width && b.x < a.x+a.width &&
		a.y < b.y+b.height && b.y < a.y+a.height
}

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func parseLine(line string) claim {
	// #1 @ 1,3: 4x4

	var id, x, y, width, height int

	fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)

	return claim{
		id, x, y, width, height,
	}
}
