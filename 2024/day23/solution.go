package day23

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strings"
)

type connection map[string][]string

func SolveP1() {

	lines := utils.ReadFile("./day23/input.txt")

	connections := make(connection)

	for _, line := range lines {
		c1, c2 := parseConnection(line)
		connections[c1] = append(connections[c1], c2)
		connections[c2] = append(connections[c2], c1)
	}

	lan_parties := findLanParties(connections)

	fmt.Println("found", countStartsWith(lan_parties, "t"), "parties starting with t")

}

func SolveP2() {

	lines := utils.ReadFile("./day23/input.txt")

	connections := make(connection)

	for _, line := range lines {
		c1, c2 := parseConnection(line)
		connections[c1] = append(connections[c1], c2)
		connections[c2] = append(connections[c2], c1)
	}

	lan_parties := findLanParties2(connections)

	valid_parties := make([]string, 0)

	for k := range lan_parties {
		if isValidParty(strings.Split(k, ","), connections) {
			valid_parties = append(valid_parties, k)
		}
	}

	largest_party_count := 0
	largest_party := ""

	for _, party := range valid_parties {
		if len(strings.Split(party, ",")) > largest_party_count {
			largest_party_count = len(strings.Split(party, ","))
			largest_party = party
		}
	}

	fmt.Println("largest party is", largest_party, "with", largest_party_count, "nodes")

}

func parseConnection(line string) (c1, c2 string) {
	c1 = line[:2]
	c2 = line[3:]
	return c1, c2
}

func findLanParties(connections connection) map[string]bool {
	parties := make(map[string]bool)

	for k, v := range connections {

		for _, connection := range v {
			common_node := findCommonConnections(k, connection, connections)
			if len(common_node) > 0 {
				for _, node := range common_node {
					parties[listToString([]string{k, connection, node})] = true
				}
			}
		}
	}

	return parties
}

func findLanParties2(connections connection) map[string]bool {
	parties := make(map[string]bool)

	for k, v := range connections {

		for _, connection := range v {
			common_nodes := findCommonConnections(k, connection, connections)
			parties[listToString(append([]string{k, connection}, common_nodes...))] = true
		}
	}

	return parties
}

func findCommonConnections(node string, compare_node string, connections connection) []string {
	common_connections := make([]string, 0)
	for _, v := range connections[node] {
		for _, connection := range connections[compare_node] {
			if connection == v {
				common_connections = append(common_connections, connection)
			}
		}
	}

	return common_connections
}

func listToString(list []string) string {
	slices.Sort(list)

	var str strings.Builder

	for idx, v := range list {
		if idx == len(list)-1 {
			str.WriteString(v)
		} else {
			str.WriteString(fmt.Sprintf("%v,", v))
		}
	}

	return str.String()
}

func countStartsWith(parties map[string]bool, start string) int {
	count := 0

	for k := range parties {
		connections := strings.Split(k, ",")

		for _, connection := range connections {
			if strings.HasPrefix(connection, start) {
				count++
				break
			}
		}
	}

	return count
}

// qp,tc,td,wh,yn
func isValidParty(party []string, connections connection) bool {
	for idx, node := range party {

		rest := slices.Concat(party[:idx], party[idx+1:])

		for _, connection := range rest {
			if !slices.Contains(connections[node], connection) {

				return false
			}
		}

	}

	return true
}
