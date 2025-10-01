package main

import (
	"2018/day18"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	day18.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
