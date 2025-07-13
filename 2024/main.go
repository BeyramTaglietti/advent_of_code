package main

import (
	"aoc2024/day15"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	// day22.SolveP1()
	day15.SolveP1()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
