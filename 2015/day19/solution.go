package day19

import (
	"aoc2015/utils"
	"fmt"
)

type rulesMap map[string][]string

func SolveP1() {
	lines := utils.ReadFile("./day19/input.txt")

	rulesList := []string{}
	puzzle := ""
	for idx, line := range lines {

		if line == "" {
			puzzle = lines[idx+1]
			break
		}
		rulesList = append(rulesList, line)
	}

	rules := parseRules(rulesList)

	fmt.Println("createed distinct molecules", len(createMolecules(rules, puzzle)))
}

func parseRules(rulesList []string) rulesMap {
	rules := make(rulesMap)

	for _, line := range rulesList {
		var rule, transformation string
		fmt.Sscanf(line, "%s => %s", &rule, &transformation)

		rules[rule] = append(rules[rule], transformation)
	}

	return rules
}

func createMolecules(rules rulesMap, puzzle string) map[string]int {
	distinctMolecules := make(map[string]int)

	for rule, transformations := range rules {
		for charIdx := range puzzle {
			if matchesRule(puzzle[charIdx:], rule) {
				for _, transformation := range transformations {
					distinctMolecules[puzzle[:charIdx]+transformation+puzzle[charIdx+len(rule):]]++
				}
			}
		}
	}

	return distinctMolecules
}

func matchesRule(puzzle string, rule string) bool {
	if len(puzzle) < len(rule) {
		return false
	}

	return puzzle[:len(rule)] == rule
}

func matchesRuleBackwards(puzzle string, rule string) bool {
	if len(puzzle) < len(rule) {
		return false
	}

	return puzzle[len(puzzle)-len(rule):] == rule
}
