package day9

import (
	"2018/utils"
	"fmt"
)

type marble struct {
	next  *marble
	prev  *marble
	value int
}

func SolveP1() {
	lines := utils.ReadFile("day9/input.txt")

	totalPlayersCount, totalMarblesCount := parseLine(lines[0])

	root := &marble{value: 0}
	root.next = root
	root.prev = root

	currentMarble := root

	scores := make([]int, totalPlayersCount)

	currentPlayerNumber := 0

	for i := 1; i <= totalMarblesCount; i++ {
		if i%23 == 0 {
			scores[currentPlayerNumber] += i
			for j := 0; j < 7; j++ {
				currentMarble = currentMarble.prev
			}
			removed := currentMarble
			scores[currentPlayerNumber] += removed.value
			removed.prev.next = removed.next
			removed.next.prev = removed.prev
			currentMarble = removed.next
		} else {
			currentMarble = currentMarble.next
			inserted := &marble{value: i}
			inserted.prev = currentMarble
			inserted.next = currentMarble.next
			currentMarble.next.prev = inserted
			currentMarble.next = inserted
			currentMarble = inserted
		}

		if currentPlayerNumber == totalPlayersCount-1 {
			currentPlayerNumber = 0
		} else {
			currentPlayerNumber++
		}
	}

	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	fmt.Println(max)

}

func SolveP2() {
	lines := utils.ReadFile("day9/input.txt")

	totalPlayersCount, totalMarblesCount := parseLine(lines[0])
	totalMarblesCount *= 100

	root := &marble{value: 0}
	root.next = root
	root.prev = root

	currentMarble := root

	scores := make([]int, totalPlayersCount)

	currentPlayerNumber := 0

	for i := 1; i <= totalMarblesCount; i++ {
		if i%23 == 0 {
			scores[currentPlayerNumber] += i
			for j := 0; j < 7; j++ {
				currentMarble = currentMarble.prev
			}
			removed := currentMarble
			scores[currentPlayerNumber] += removed.value
			removed.prev.next = removed.next
			removed.next.prev = removed.prev
			currentMarble = removed.next
		} else {
			currentMarble = currentMarble.next
			inserted := &marble{value: i}
			inserted.prev = currentMarble
			inserted.next = currentMarble.next
			currentMarble.next.prev = inserted
			currentMarble.next = inserted
			currentMarble = inserted
		}

		if currentPlayerNumber == totalPlayersCount-1 {
			currentPlayerNumber = 0
		} else {
			currentPlayerNumber++
		}
	}

	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	fmt.Println(max)
}

func parseLine(line string) (players int, marbles int) {
	fmt.Sscanf(line, "%d players; last marble is worth %d points", &players, &marbles)

	return players, marbles
}
