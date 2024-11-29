package day7

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type operation string

const (
	AND        operation = "AND"
	OR         operation = "OR"
	LSHIFT     operation = "LSHIFT"
	RSHIFT     operation = "RSHIFT"
	NOT        operation = "NOT"
	ASSIGNMENT operation = "ASSIGNMENT"
)

type signal struct {
	from          []string
	fromOperation operation
}

func SolveP1() {
	lines, err := utils.ReadFile("day7/input.txt")
	if err != nil {
		log.Fatalf("File could not be found")
	}

	circuit := make(map[string]signal)

	for _, line := range lines {

		switch getOperation(line) {
		case ASSIGNMENT:
			from, to := parseAssignment(line)
			circuit[to] = signal{from: []string{from}, fromOperation: ASSIGNMENT}
		case AND:
			p1, p2, to := parseAND(line)
			circuit[to] = signal{from: []string{p1, p2}, fromOperation: AND}
		case OR:
			p1, p2, to := parseOR(line)
			circuit[to] = signal{from: []string{p1, p2}, fromOperation: OR}
		case LSHIFT:
			from, to, value := parseLSHIFT(line)
			circuit[to] = signal{from: []string{from, value}, fromOperation: LSHIFT}

		case RSHIFT:
			from, to, value := parseRSHIFT(line)
			circuit[to] = signal{from: []string{from, value}, fromOperation: RSHIFT}
		case NOT:
			p1, p2 := parseNOT(line)
			circuit[p2] = signal{from: []string{p1}, fromOperation: NOT}
		}
	}

	key := "a"
	fmt.Printf("value of key %s is %d\n", key, getValue(key, circuit))
	fmt.Println("finished")

}

func getValue(key string, circuit map[string]signal) uint16 {
	currValue := circuit[key]
	switch currValue.fromOperation {
	case ASSIGNMENT:
		isNum, val := isNumeric(currValue.from[0])
		if isNum {
			return val
		} else {
			val := getValue(currValue.from[0], circuit)
			circuit[key] = signal{from: []string{strconv.Itoa(int(val))}, fromOperation: ASSIGNMENT}
			return val
		}
	case AND:
		isNum1, val1 := isNumeric(currValue.from[0])
		isNum2, val2 := isNumeric(currValue.from[1])

		if isNum1 && isNum2 {
			return val1 & val2
		}

		if isNum1 {
			val := getValue(currValue.from[1], circuit)
			circuit[key] = signal{from: []string{strconv.Itoa(int(val1 & val))}, fromOperation: ASSIGNMENT}
			return val1 & val
		}

		if isNum2 {
			val := getValue(currValue.from[0], circuit)
			circuit[key] = signal{from: []string{strconv.Itoa(int(val & val2))}, fromOperation: ASSIGNMENT}
			return val & val2
		}

		val := getValue(currValue.from[0], circuit) & getValue(currValue.from[1], circuit)
		circuit[key] = signal{from: []string{strconv.Itoa(int(val))}, fromOperation: ASSIGNMENT}
		return val

	case OR:
		isNum1, val1 := isNumeric(currValue.from[0])
		isNum2, val2 := isNumeric(currValue.from[1])

		if isNum1 && isNum2 {
			return val1 | val2
		}

		if isNum1 {
			val := getValue(currValue.from[1], circuit)
			circuit[key] = signal{from: []string{strconv.Itoa(int(val1 | val))}, fromOperation: ASSIGNMENT}
			return val1 | val
		}

		if isNum2 {
			val := getValue(currValue.from[0], circuit)
			circuit[key] = signal{from: []string{strconv.Itoa(int(val | val2))}, fromOperation: ASSIGNMENT}
			return val | val2
		}

		val := getValue(currValue.from[0], circuit) | getValue(currValue.from[1], circuit)
		circuit[key] = signal{from: []string{strconv.Itoa(int(val))}, fromOperation: ASSIGNMENT}
		return val

	case LSHIFT:

		numVal, _ := strconv.ParseUint(currValue.from[1], 10, 16)
		return getValue(currValue.from[0], circuit) << uint16(numVal)

	case RSHIFT:

		numVal, _ := strconv.ParseUint(currValue.from[1], 10, 16)
		return getValue(currValue.from[0], circuit) >> uint16(numVal)

	case NOT:
		isNum, val := isNumeric(currValue.from[0])
		if isNum {
			return ^val
		}

		return ^getValue(currValue.from[0], circuit)
	}

	return 0
}

func getOperation(line string) operation {
	operations := []operation{
		AND,
		OR,
		LSHIFT,
		RSHIFT,
		NOT,
	}

	for _, operation := range operations {
		if strings.Contains(line, string(operation)) {
			return operation
		}
	}

	return ASSIGNMENT
}

func isNumeric(value string) (bool, uint16) {
	val, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return false, 0
	}

	return true, uint16(val)
}
