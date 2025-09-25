package day5

import (
	"2018/utils"
	"fmt"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("day5/input.txt")
	polymer := lines[0]

	polymer = reducePolymer(polymer)

	fmt.Println(len(polymer))
}

func SolveP2() {
	lines := utils.ReadFile("day5/input.txt")
	polymer := lines[0]
	shortestPolymer := polymer

	letters := make(map[rune]bool)

	for _, letter := range polymer {
		if letter >= 'a' && letter <= 'z' {
			letters[letter] = true
		}
	}

	for letter := range letters {
		newPolymer := strings.ReplaceAll(polymer, string(letter), "")
		newPolymer = strings.ReplaceAll(newPolymer, string(letter-32), "")
		newPolymer = reducePolymer(newPolymer)
		if len(newPolymer) < len(shortestPolymer) {
			shortestPolymer = newPolymer
		}
	}

	fmt.Println(len(shortestPolymer))
}

func reducePolymer(polymer string) string {
	for i := 0; i < len(polymer)-1; i++ {
		if polymer[i] != polymer[i+1] && (polymer[i]-polymer[i+1] == 32 || polymer[i+1]-polymer[i] == 32) {
			polymer = polymer[:i] + polymer[i+2:]
			i = -1
		}
	}
	return polymer
}
