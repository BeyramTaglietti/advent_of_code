package main

import (
	"aoc2022/day2"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Run the solution
	// day1.SolveP1()
	// day1.SolveP2()
	day2.SolveP1()
	day2.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs\n", elapsed.Microseconds())
}
