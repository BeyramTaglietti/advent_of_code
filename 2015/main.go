package main

import (
	"aoc2015/day20"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	day20.SolveP1()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
