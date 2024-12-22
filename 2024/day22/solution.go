package day22

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func SolveP1() {

	lines := utils.ReadFile("./day22/input.txt")

	sum := 0

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	mu := sync.Mutex{}

	for _, line := range lines {

		go func(line string) {
			defer wg.Done()
			result, _ := calculateResult(utils.Atoi(line), 1, 2000, []int{}, map[string]int{})

			mu.Lock()
			sum += result
			mu.Unlock()
		}(line)
	}

	wg.Wait()

	fmt.Println("final result is", sum)

}

func SolveP2() {
	lines := utils.ReadFile("./day22/input.txt")

	full_sequences_map := map[string]int{}

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	mu := sync.Mutex{}

	for _, line := range lines {

		go func(line string) {
			defer wg.Done()
			_, sequences_map := calculateResult(utils.Atoi(line), 1, 2000, []int{}, map[string]int{})

			mu.Lock()
			for k, v := range sequences_map {
				full_sequences_map[k] = full_sequences_map[k] + v
			}
			mu.Unlock()

		}(line)
	}

	wg.Wait()

	best_sequence, bananas_total_amount := findBestSequence(full_sequences_map)

	fmt.Println("best sequence is", best_sequence, "with total amount of bananas", bananas_total_amount)
}

func getLastDigit(num int) int {
	num_str := strconv.Itoa(num)
	return utils.Atoi(string(num_str[len(num_str)-1]))
}

func calculateResult(secret_number, current_iteration, total_iterations int, previous_changes []int, sequences_map map[string]int) (res int, s_map map[string]int) {

	new_secret_number := prune(mix(secret_number*64, secret_number))
	new_secret_number = prune(mix(int(math.Floor(float64(new_secret_number)/32)), new_secret_number))
	new_secret_number = prune(mix(new_secret_number*2048, new_secret_number))

	previous_secret_number_last_digit := getLastDigit(secret_number)
	new_secret_number_last_digit := getLastDigit(new_secret_number)

	addPattern := func(pattern []int, num int) {
		sequence := listToString(pattern)

		if _, ok := sequences_map[sequence]; !ok {
			sequences_map[sequence] = num
		}
	}

	if len(previous_changes) != 4 {
		previous_changes = append(previous_changes, new_secret_number_last_digit-previous_secret_number_last_digit)
	} else {
		addPattern(previous_changes, previous_secret_number_last_digit)

		previous_changes = append(previous_changes[1:], new_secret_number_last_digit-previous_secret_number_last_digit)
	}

	if current_iteration == total_iterations {
		addPattern(previous_changes, new_secret_number_last_digit)
		return new_secret_number, sequences_map
	}

	return calculateResult(new_secret_number, current_iteration+1, total_iterations, previous_changes, sequences_map)
}

func prune(num int) int {
	const modulo_value int = 16777216

	return int(math.Mod(float64(num), float64(modulo_value)))
}

func mix(num1, num2 int) int {
	return num1 ^ num2
}

func findBestSequence(sequences_map map[string]int) (string, int) {
	max := 0
	best_sequence := ""

	for k, v := range sequences_map {
		if v > max {
			max = v
			best_sequence = k
		}
	}

	return best_sequence, max
}

func listToString(list []int) string {
	var str strings.Builder

	for _, n := range list {
		str.WriteString(strconv.Itoa(n))
	}

	return str.String()
}
