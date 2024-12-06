package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func MoveListElement[s ~[]E, E int | string](list s, from int, to int) s {
	if from < 0 || from >= len(list) {
		log.Fatalf("Invalid 'from' index: %d", from)
	}
	if to < 0 || to >= len(list) {
		log.Fatalf("Invalid 'to' index: %d", to)
	}

	element := list[from]
	list = append(list[:from], list[from+1:]...)
	list = append(list[:to], append([]E{element}, list[to:]...)...)

	return list
}
