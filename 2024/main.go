package main

import (
	"aoc2024/day20"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	day20.SolveP1()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
