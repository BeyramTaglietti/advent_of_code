package main

import (
	"aoc2024/day22"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	// day22.SolveP1()
	day22.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
