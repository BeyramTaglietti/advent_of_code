package day20

import (
	"fmt"
	"slices"
)

const puzzle int = 34000000

func SolveP1() {

	houseNumber := 1
	for {

		presents := 0

		for _, divisor := range findDivisors(houseNumber) {
			presents += divisor * 10
		}

		if presents >= puzzle {
			break
		}

		houseNumber++
	}

	fmt.Println("house number", houseNumber)

}

func SolveP2() {
	houseNumber := 1
	for {

		presents := 0

		for _, divisor := range findDivisors(houseNumber) {
			if houseNumber/divisor <= 50 {
				presents += divisor * 11
			}
		}

		if presents >= puzzle {
			break
		}

		houseNumber++
	}

	fmt.Println("house number", houseNumber)
}

func findDivisors(n int) []int {
	var divisors []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	slices.Sort(divisors)

	return divisors
}
