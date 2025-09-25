package main

import (
	"2018/day7"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	day7.SolveP2()

	elapsed := time.Since(now)

	fmt.Printf("Elapsed time: %d μs / %d ms\n", elapsed.Microseconds(), elapsed.Milliseconds())
}
