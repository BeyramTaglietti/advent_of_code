package main

import (
	"aoc2022/day6"
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
	// day3.SolveP2()
	// day4.SolveP1()
	// day4.SolveP2()
	// day5.SolveP1()
	// day5.SolveP2()
	day6.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs\n", elapsed.Microseconds())
}
