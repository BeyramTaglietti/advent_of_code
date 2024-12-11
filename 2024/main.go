package main

import (
	"aoc2024/day11"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	day11.Solve()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
