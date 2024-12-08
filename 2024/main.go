package main

import (
	"aoc2024/day8"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	day8.SolveP1()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
