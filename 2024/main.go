package main

import (
	"aoc2024/day3"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	// day1.SolveP1()
	// day1.SolveP2()
	// day2.SolveP1()
	// day2.SolveP2()
	// day3.SolveP1()
	day3.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
