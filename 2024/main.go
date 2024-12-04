package main

import (
	"aoc2024/day4"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	day4.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
