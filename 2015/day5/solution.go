package day5

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strings"
)

func SolveP1() {

	lines, err := utils.ReadFile("day5/input.txt")
	if err != nil {
		log.Fatalf("File could not be found")
	}

	niceWords := 0

	for _, line := range lines {
		if isNiceWord(line) {
			niceWords++
		}
	}

	fmt.Printf("Found %d nice words\n", niceWords)
}

func SolveP2() {

	lines, err := utils.ReadFile("day5/input.txt")
	if err != nil {
		log.Fatalf("File could not be found")
	}

	niceWords := 0

	for _, line := range lines {
		if isNiceWord2(line) {
			niceWords++
		}
	}

	fmt.Printf("Found %d nice words with the new rules\n", niceWords)
}

func isNiceWord(word string) bool {
	return contains3Vowels(word) && containsRepeatedLetters(word) && !containsBannedCombinations(word)
}

func isNiceWord2(word string) bool {
	return containsPairLetterNotOverlapping(word) && containsPairWithPause(word)
}

func contains3Vowels(text string) bool {
	const vowels = "aeiou"

	tot := 0

	for _, letter := range text {
		for _, vowel := range vowels {
			if letter == vowel {
				tot++
			}
		}
	}

	return tot >= 3
}

func containsRepeatedLetters(text string) bool {
	lastLetter := rune(text[0])

	for _, letter := range text[1:] {

		if letter == lastLetter {
			return true
		}

		lastLetter = letter
	}

	return false
}

func containsBannedCombinations(text string) bool {
	bannedCombinations := []string{"ab", "cd", "pq", "xy"}

	for _, combo := range bannedCombinations {
		if strings.Contains(text, combo) {
			return true
		}
	}

	return false
}

func containsPairLetterNotOverlapping(word string) bool {
	if len(word) < 4 {
		return false
	}

	pairMap := make(map[string]int)

	for i := 0; i < len(word)-1; i++ {
		pair := word[i : i+2] // Current pair of two letters

		if lastIndex, exists := pairMap[pair]; exists && lastIndex < i-1 {
			return true
		}
		pairMap[pair] = i
	}

	return false
}

func containsPairWithPause(word string) bool {
	for i := 0; i <= len(word)-3; i += 1 {
		if word[i] == word[i+2] {
			return true
		}
	}
	return false
}
