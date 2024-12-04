package day11

import (
	"fmt"
	"slices"
)

func SolveP1() {
	password := "hxbxxyzz"

	fmt.Println("password", password, "is valid:", isPasswordValid(password))

	fmt.Println("next valid password:", bruteForcePassword(password))

}

func isPasswordValid(password string) bool {
	return contains3StraightLetters(password) && !containsBannedLetters(password) && containsNonOverlappingPairs(password)
}

func contains3StraightLetters(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1]-1 && password[i] == password[i+2]-2 {
			return true
		}
	}
	return false
}

func containsBannedLetters(password string) bool {
	bannedLetters := []rune{'i', 'o', 'l'}

	for _, letter := range password {
		if slices.Contains(bannedLetters, letter) {
			return true
		}
	}

	return false
}

func containsNonOverlappingPairs(password string) bool {

	pairs := 0

	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++
		}
	}

	return pairs >= 2
}

func bruteForcePassword(password string) string {
	for {
		password = next(password)
		if isPasswordValid(password) {
			return password
		}
	}
}

func next(s string) string {

	chars := []byte(s)

	// Start from the last character and move backward
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] < 'z' {
			chars[i]++
			return string(chars)
		}

		chars[i] = 'a'
	}

	return "a" + string(chars)
}
