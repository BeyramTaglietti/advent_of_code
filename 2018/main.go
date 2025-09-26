package main

import (
	"2018/day10"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	day10.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
