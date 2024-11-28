package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"sync"
)

const span = 10_000_000
const divisor = 1_000_000

func SolveP1() {
	secretKey := "yzbqklnj"

	values := []int{}

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	ranges := span / divisor
	wg.Add(ranges)

	for i := 1; i <= ranges; i++ {

		go func(idx int) {
			defer wg.Done()

			checkRange := span - (idx * divisor)

			for j := checkRange; j < checkRange+divisor; j++ {
				if isValidAdventCoin(getMD5Hash(secretKey+strconv.Itoa(j)), 5) {
					fmt.Println("found:", j)
					mutex.Lock()
					values = append(values, j)
					mutex.Unlock()
				}
			}
		}(i)

	}

	wg.Wait()

	if len(values) == 0 {
		log.Fatalf("No value that satisfies the constraint found")
		return
	}

	smallestNumberOfAll := slices.Min(values)
	fmt.Println("Smallest of all is", smallestNumberOfAll)

}

func SolveP2() {
	secretKey := "yzbqklnj"

	values := []int{}

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	ranges := span / divisor
	wg.Add(ranges)

	for i := 1; i <= ranges; i++ {

		go func(idx int) {
			defer wg.Done()

			checkRange := span - (idx * divisor)

			for j := checkRange; j < checkRange+divisor; j++ {
				if isValidAdventCoin(getMD5Hash(secretKey+strconv.Itoa(j)), 6) {
					fmt.Println("found:", j)
					mutex.Lock()
					values = append(values, j)
					mutex.Unlock()
				}
			}
		}(i)

	}

	wg.Wait()

	if len(values) == 0 {
		log.Fatalf("No value that satisfies the constraint found")
		return
	}

	smallestNumberOfAll := slices.Min(values)
	fmt.Println("Smallest of all is", smallestNumberOfAll)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func isValidAdventCoin(hash string, leadingZeros int) bool {
	return strings.HasPrefix(hash, strings.Repeat("0", leadingZeros))
}
