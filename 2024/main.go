package main

import (
	"aoc2024/day6"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	day6.SolveP1()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
