package day7

import "fmt"

func parseAssignment(line string) (string, string) {
	// 456 -> y

	var from, to string
	fmt.Sscanf(line, "%s -> %s", &from, &to)

	return from, to
}

func parseAND(line string) (string, string, string) {
	// x AND y -> d
	var p1, p2, to string
	fmt.Sscanf(line, "%s AND %s -> %s", &p1, &p2, &to)

	return p1, p2, to
}

func parseOR(line string) (string, string, string) {
	// x OR y -> e
	var p1, p2, to string
	fmt.Sscanf(line, "%s OR %s -> %s", &p1, &p2, &to)

	return p1, p2, to
}

func parseLSHIFT(line string) (string, string, string) {
	// x LSHIFT 2 -> f
	var from, to, value string
	fmt.Sscanf(line, "%s LSHIFT %s -> %s", &from, &value, &to)

	return from, to, value
}

func parseRSHIFT(line string) (string, string, string) {
	// x RSHIFT 2 -> f
	var from, to, value string
	fmt.Sscanf(line, "%s RSHIFT %s -> %s", &from, &value, &to)

	return from, to, value
}

func parseNOT(line string) (string, string) {
	// NOT y -> i
	var p1, p2 string
	fmt.Sscanf(line, "NOT %s -> %s", &p1, &p2)

	return p1, p2
}
