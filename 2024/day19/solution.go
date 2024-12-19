package day19

import (
	"aoc2024/utils"
	"fmt"
	"maps"
	"strings"
	"sync"
)

func SolveP1() {
	lines := utils.ReadFile("./day19/input.txt")

	towels := parseTowels(lines[0])
	designs := parseDesigns(lines[2:])

	possibleDesignsCount := 0
	wg := sync.WaitGroup{}
	wg.Add(len(designs))
	mu := sync.Mutex{}
	for _, design := range designs {
		go func(d string, t map[string]bool) {
			defer wg.Done()
			foundMatch := createDesign(d, t, findMatches(d, t, make(map[string]int)), 0)
			if foundMatch > 0 {
				mu.Lock()
				possibleDesignsCount++
				mu.Unlock()
			}
		}(design, towels)
	}

	wg.Wait()

	fmt.Println("we can create", possibleDesignsCount, "designs")
}

func SolveP2() {

	lines := utils.ReadFile("./day19/input.txt")

	towels := parseTowels(lines[0])
	designs := parseDesigns(lines[2:])

	wg := sync.WaitGroup{}
	wg.Add(len(designs))

	mu := sync.Mutex{}

	possibleDesignsCount := 0
	for _, design := range designs {

		go func(d string, t map[string]bool) {
			defer wg.Done()
			foundMatch := createDesign(d, t, findMatches(d, t, make(map[string]int)), 0)

			mu.Lock()
			possibleDesignsCount += foundMatch
			mu.Unlock()
		}(design, towels)

	}

	wg.Wait()

	fmt.Println("we can create these designs in", possibleDesignsCount, "ways")
}

func parseTowels(towels string) map[string]bool {
	list := strings.Split(towels, ", ")
	dict := make(map[string]bool)

	for _, l := range list {
		dict[l] = true
	}

	return dict
}

func parseDesigns(lines []string) []string {
	designs := make([]string, len(lines), len(lines))
	for idx, line := range lines {
		designs[idx] = line
	}
	return designs
}

func createDesign(design string, towels map[string]bool, currentMatches map[string]int, validsCount int) int {

	if len(currentMatches) == 0 {
		return validsCount
	}

	newMatches := make(map[string]int)

	for currentMatch, currentMatchV := range currentMatches {
		matchesFound := findMatches(design[len(currentMatch):], towels, make(map[string]int))

		if len(matchesFound) > 0 {
			for k := range matchesFound {
				newMatches[currentMatch+k] = newMatches[currentMatch+k] + currentMatchV
			}
		}
	}

	for k, v := range newMatches {
		if k == design {
			validsCount += v
			delete(newMatches, k)
		}
	}

	return createDesign(design, towels, newMatches, validsCount)
}

func findMatches(line string, towels map[string]bool, found map[string]int) map[string]int {

	if len(line) == 0 {
		return found
	}

	if _, ok := towels[line]; ok {
		newFoundDict := maps.Clone(found)
		newFoundDict[line] = 1
		return findMatches(line[:len(line)-1], towels, newFoundDict)
	} else {
		return findMatches(line[:len(line)-1], towels, found)
	}
}
