package main

import (
	"aoc2015/day19"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	day19.SolveP1()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
