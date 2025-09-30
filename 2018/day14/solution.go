package day14

import (
	"2018/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("day14/input.txt")
	wantToTry, _ := strconv.Atoi(lines[0])

	recipes := []int{3, 7}
	firstElf := 0
	secondElf := 1

	for len(recipes) < 10+wantToTry {
		sum := recipes[firstElf] + recipes[secondElf]
		if sum >= 10 {
			recipes = append(recipes, sum/10)
		}
		recipes = append(recipes, sum%10)

		firstElf = (firstElf + 1 + recipes[firstElf]) % len(recipes)
		secondElf = (secondElf + 1 + recipes[secondElf]) % len(recipes)
	}

	fmt.Println("result is:", take(wantToTry, 10, recipes))

}

func SolveP2() {
	lines := utils.ReadFile("day14/input.txt")
	wantToFind := lines[0]

	recipes := []int{3, 7}
	firstElf := 0
	secondElf := 1

	recipesTried := 2

	equalityIndex := 0

	for {
		sum := recipes[firstElf] + recipes[secondElf]
		newNodeStr := strconv.Itoa(sum)
		for _, letter := range newNodeStr {
			recipesTried++
			if string(letter) == string(wantToFind[equalityIndex]) {
				equalityIndex++
			} else {
				equalityIndex = 0

				if string(letter) == string(wantToFind[0]) {
					equalityIndex++
				}
			}

			if equalityIndex == len(wantToFind) {
				fmt.Println("result is", recipesTried-equalityIndex)
				return
			}
			intValue, _ := strconv.Atoi(string(letter))
			recipes = append(recipes, intValue)
		}

		firstElf = (firstElf + 1 + recipes[firstElf]) % len(recipes)
		secondElf = (secondElf + 1 + recipes[secondElf]) % len(recipes)
	}

}

func take(after int, take int, recipes []int) int {
	b := strings.Builder{}
	for i := 0; i < take; i++ {
		b.WriteString(strconv.Itoa(recipes[after+i]))
	}

	resultValue, _ := strconv.Atoi(b.String())

	return resultValue
}
