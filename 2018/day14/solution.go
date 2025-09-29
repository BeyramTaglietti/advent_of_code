package day14

import (
	"2018/utils"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	child *Node
	value int
}

func SolveP1() {
	lines := utils.ReadFile("day14/input.txt")
	wantToTry, _ := strconv.Atoi(lines[0])

	firstRecipe := &Node{value: 3}
	secondRecipe := &Node{value: 7}

	firstRecipe.child = secondRecipe
	secondRecipe.child = firstRecipe

	recipes := firstRecipe

	firstElf := firstRecipe
	secondElf := secondRecipe

	tail := secondRecipe

	recipesTried := 2

	for recipesTried < 10+wantToTry {
		newNodeStr := strconv.Itoa(firstElf.value + secondElf.value)
		recipesTried += len(newNodeStr)
		for _, letter := range newNodeStr {
			intValue, _ := strconv.Atoi(string(letter))
			newNode := Node{value: intValue}
			tail.child = &newNode
			tail = tail.child
		}
		tail.child = recipes

		firstElfMoves := firstElf.value
		secondElfMoves := secondElf.value
		for j := 0; j < firstElfMoves+1; j++ {
			firstElf = firstElf.child
		}
		for j := 0; j < secondElfMoves+1; j++ {
			secondElf = secondElf.child
		}
	}

	printRecipes(recipes)

	fmt.Println("result is:", take(wantToTry, 10, recipes))

}

func SolveP2() {
	lines := utils.ReadFile("day14/input.txt")
	wantToFind := lines[0]

	firstRecipe := &Node{value: 3}
	secondRecipe := &Node{value: 7}

	firstRecipe.child = secondRecipe
	secondRecipe.child = firstRecipe

	recipes := firstRecipe

	firstElf := firstRecipe
	secondElf := secondRecipe

	tail := secondRecipe

	recipesTried := 2

	equalityIndex := 0

	for {
		newNodeStr := strconv.Itoa(firstElf.value + secondElf.value)
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
			newNode := Node{value: intValue}
			tail.child = &newNode
			tail = tail.child
		}
		tail.child = recipes

		firstElfMoves := firstElf.value
		secondElfMoves := secondElf.value
		for j := 0; j < firstElfMoves+1; j++ {
			firstElf = firstElf.child
		}
		for j := 0; j < secondElfMoves+1; j++ {
			secondElf = secondElf.child
		}
	}

}

func take(after int, take int, recipes *Node) int {
	current := recipes
	result := strings.Builder{}

	for i := 0; i < after; i++ {
		current = current.child
	}

	for j := 0; j < take; j++ {
		result.WriteString(strconv.Itoa(current.value))
		current = current.child
	}

	resultValue, _ := strconv.Atoi(result.String())

	return resultValue
}

func printRecipes(recipes *Node) {
	current := recipes

	if current == nil {
		fmt.Println("nil")
		return
	}

	fmt.Printf("%d->", current.value)
	current = current.child

	for current != nil && current != recipes {
		fmt.Printf("%d->", current.value)
		current = current.child
	}

	fmt.Println("nil")
}
