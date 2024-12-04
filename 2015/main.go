package main

import (
	"aoc2015/day11"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	day11.SolveP1()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
