package day11

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func Solve() {
	lines := utils.ReadFile("./day11/input.txt")

	values := strings.Split(lines[0], " ")

	cache := make(map[int]int)

	for _, v := range values {
		cache[utils.Atoi(v)]++
	}

	const blinks int = 75

	for i := 0; i < blinks; i++ {
		cache = blink(cache)
	}

	fmt.Println("after blinking", blinks, "times, the stones are:", countStones(cache))
}

func blink(currentCache map[int]int) map[int]int {
	newCache := make(map[int]int)

	for k, v := range currentCache {
		if k == 0 {
			newCache[1] += v
		} else if len(utils.Itoa(k))%2 == 0 {
			strVal := utils.Itoa(k)
			halfLen := len(strVal) / 2
			firstHalf := utils.Atoi(strVal[:halfLen])
			secondHalf := utils.Atoi(strVal[halfLen:])

			newCache[firstHalf] += v
			newCache[secondHalf] += v
		} else {
			newCache[k*2024] += v
		}
	}

	return newCache
}

func countStones(currentCache map[int]int) int {
	sum := 0
	for _, v := range currentCache {
		sum += v
	}

	return sum
}
