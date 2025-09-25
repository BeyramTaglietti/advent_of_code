package day4

import (
	"2018/utils"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
)

func SolveP1() {
	lines := utils.ReadFile("day4/input.txt")

	logs := make([]log, len(lines))

	for idx, line := range lines {
		logs[idx] = parseLine(line)
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].date.Before(logs[j].date)
	})

	// current guard
	// fallsAsleepAt time.Time
	// wakesUpAt time.Time
	// guardSleepingCycle []int = for [0, 1, ..., 59] we put at guardSleepingCycle[i] the amount of times he was found asleep at that time
	// for i = fallsAsleepAt; i < wakesUpAt; i++

	// later on we cycle all the guards, get how much time they slept, and for the most sleeping one, we find the guardSleepingCycle max

	type status string

	const (
		sleeping status = "sleeping"
		awake    status = "awake"
	)

	type guard struct {
		id            int
		sleepingCycle []int
		status        status
		sleptAt       int
		awokeAt       int
	}

	createNewGuard := func(id int) guard {
		return guard{
			id:            id,
			status:        awake,
			sleepingCycle: make([]int, 60),
		}
	}

	guards := make(map[int]guard)

	var currentGuard guard

	for _, log := range logs {
		if log.action == begins {
			if g, ok := guards[log.guardId]; ok {
				currentGuard = g
			} else {
				currentGuard = createNewGuard(log.guardId)
				guards[currentGuard.id] = currentGuard
			}
		} else {
			switch log.action {
			case sleeps:
				currentGuard.sleptAt = log.date.Minute()
				currentGuard.status = sleeping
			case wakes:
				currentGuard.awokeAt = log.date.Minute()
				currentGuard.status = awake

				for i := currentGuard.awokeAt; i > currentGuard.sleptAt; i-- {
					currentGuard.sleepingCycle[i]++
				}
			}
		}
	}

	highestSleepingMinutes := 0
	highestSleepingGuard := guards[0]

	for _, guard := range guards {
		sleepingFor := 0
		for _, x := range guard.sleepingCycle {
			sleepingFor += x
		}

		if sleepingFor > highestSleepingMinutes {
			highestSleepingGuard = guard
			highestSleepingMinutes = sleepingFor
		}
	}

	mostCommonSleepingMinute := slices.Max(highestSleepingGuard.sleepingCycle)
	mostCommonSleepingMinuteIndex := slices.Index(highestSleepingGuard.sleepingCycle, mostCommonSleepingMinute) - 1

	fmt.Println("the result will be", highestSleepingGuard.id*mostCommonSleepingMinuteIndex)
}

func SolveP2() {
	lines := utils.ReadFile("day4/input.txt")

	logs := make([]log, len(lines))

	for idx, line := range lines {
		logs[idx] = parseLine(line)
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].date.Before(logs[j].date)
	})

	type status string

	const (
		sleeping status = "sleeping"
		awake    status = "awake"
	)

	type guard struct {
		id            int
		sleepingCycle []int
		status        status
		sleptAt       int
		awokeAt       int
	}

	createNewGuard := func(id int) guard {
		return guard{
			id:            id,
			status:        awake,
			sleepingCycle: make([]int, 60),
		}
	}

	guards := make(map[int]guard)

	var currentGuard guard

	for _, log := range logs {
		if log.action == begins {
			if g, ok := guards[log.guardId]; ok {
				currentGuard = g
			} else {
				currentGuard = createNewGuard(log.guardId)
				guards[currentGuard.id] = currentGuard
			}
		} else {
			switch log.action {
			case sleeps:
				currentGuard.sleptAt = log.date.Minute()
				currentGuard.status = sleeping
			case wakes:
				currentGuard.awokeAt = log.date.Minute()
				currentGuard.status = awake

				for i := currentGuard.awokeAt; i > currentGuard.sleptAt; i-- {
					currentGuard.sleepingCycle[i]++
				}
			}
		}
	}

	highestSleepingMinutes := 0
	highestSleepingGuard := guards[0]

	for _, guard := range guards {
		for _, x := range guard.sleepingCycle {
			if x > highestSleepingMinutes {
				highestSleepingMinutes = x
				highestSleepingGuard = guard
			}
		}
	}

	mostCommonSleepingMinute := slices.Max(highestSleepingGuard.sleepingCycle)
	mostCommonSleepingMinuteIndex := slices.Index(highestSleepingGuard.sleepingCycle, mostCommonSleepingMinute) - 1

	fmt.Println("the result will be", highestSleepingGuard.id*mostCommonSleepingMinuteIndex)
}

type action string

const (
	wakes  action = "wakes up"
	sleeps action = "falls asleep"
	begins action = "begins"
)

type log struct {
	date    time.Time
	guardId int
	action  action
}

func parseLine(line string) log {

	var year, month, day, hour, minute int

	var guardAction action
	var guardId int
	var guardIdAvailable = false

	fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)

	if strings.Contains(line, "wakes") {
		guardAction = wakes
	} else if strings.Contains(line, "falls") {
		guardAction = sleeps
	} else {
		guardIdAvailable = true
		fmt.Sscanf(line, "[%d-%d-%d %d:%d] Guard #%d %s", &year, &month, &day, &hour, &minute, &guardId)
	}

	logLine := log{
		date: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
	}

	if guardIdAvailable {
		logLine.guardId = guardId
		logLine.action = begins
	} else {
		logLine.action = guardAction
	}

	return logLine
}
