package day7

import (
	"2018/utils"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func SolveP1() {
	lines := utils.ReadFile("day7/input.txt")

	deps := make(map[string][]string)
	inDegree := make(map[string]int)

	for _, line := range lines {
		before, after := parseLine(line)

		deps[before] = append(deps[before], after)

		if _, ok := inDegree[before]; !ok {
			inDegree[before] = 0
		}

		inDegree[after]++
	}

	result := strings.Builder{}

	queue := priorityqueue.NewWith(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})

	for step, degree := range inDegree {
		if degree == 0 {
			queue.Enqueue(step)
		}
	}

	for !queue.Empty() {
		smallest, _ := queue.Dequeue()

		smallestString := smallest.(string)

		result.WriteString(smallestString)

		for _, child := range deps[smallestString] {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue.Enqueue(child)
			}
		}

	}

	fmt.Println("result", result.String())

}

func SolveP2() {
	lines := utils.ReadFile("day7/input.txt")

	deps := make(map[string][]string)
	inDegree := make(map[string]int)

	for _, line := range lines {
		before, after := parseLine(line)
		deps[before] = append(deps[before], after)
		if _, ok := inDegree[before]; !ok {
			inDegree[before] = 0
		}
		inDegree[after]++
	}

	ready := priorityqueue.NewWith(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	busy := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(Task).duration - b.(Task).duration
	})

	for step, degree := range inDegree {
		if degree == 0 {
			ready.Enqueue(step)
		}
	}

	free := 5
	t := 0
	result := strings.Builder{}

	for !ready.Empty() || !busy.Empty() {
		for free > 0 && !ready.Empty() {
			s, _ := ready.Dequeue()
			task := s.(string)
			busy.Enqueue(Task{
				name:     task,
				duration: t + taskDuration(task),
			})
			free--
		}

		if !busy.Empty() {
			e, _ := busy.Dequeue()
			j := e.(Task)
			t = j.duration
			result.WriteString(j.name)
			for _, c := range deps[j.name] {
				inDegree[c]--
				if inDegree[c] == 0 {
					ready.Enqueue(c)
				}
			}
			free++
		}
	}

	fmt.Println("finishing at", t, "result", result.String())
}

type Task struct {
	name     string
	duration int
}

func parseLine(line string) (before string, after string) {

	fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &before, &after)

	return before, after
}

func taskDuration(task string) int {
	return int(task[0]) - 64 + 60
}
