package day8

import (
	"2018/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolveP1() {
	lines := utils.ReadFile("day8/input.txt")

	values := parseLicense(lines[0])

	sum := 0

	var dp func(start int) int

	dp = func(startIdx int) int {

		metadataEntries := values[startIdx+1]
		childrenAmount := values[startIdx]

		if childrenAmount == 0 {

			for i := 0; i < metadataEntries; i++ {
				sum += values[startIdx+2+i]
			}

			return 2 + metadataEntries
		} else {

			childrenLength := 2

			for i := 0; i < childrenAmount; i++ {
				childrenLength += dp(startIdx + childrenLength)
			}

			skip := startIdx + childrenLength
			for i := 0; i < metadataEntries; i++ {
				sum += values[skip+i]
			}

			return childrenLength + metadataEntries
		}
	}

	dp(0)
	fmt.Println("result:", sum)

}

func SolveP2() {
	lines := utils.ReadFile("day8/input.txt")

	values := parseLicense(lines[0])

	var dp func(start int) (int, int)

	dp = func(startIdx int) (int, int) {

		metadataEntriesLength := values[startIdx+1]
		childrenAmount := values[startIdx]

		if childrenAmount == 0 {

			nodeValue := 0

			for i := 0; i < metadataEntriesLength; i++ {
				nodeValue += values[startIdx+2+i]
			}

			return 2 + metadataEntriesLength, nodeValue
		} else {

			childrenLength := 2

			childrenValues := make([]int, childrenAmount)

			for i := 0; i < childrenAmount; i++ {
				childLength, childValue := dp(startIdx + childrenLength)
				childrenValues[i] = childValue
				childrenLength += childLength
			}

			metadataEntries := make([]int, metadataEntriesLength)

			skip := startIdx + childrenLength
			for i := 0; i < metadataEntriesLength; i++ {
				metadataEntries[i] = values[skip+i]
			}

			nodeValue := 0

			for _, entry := range metadataEntries {
				if entry >= 1 && entry <= childrenAmount {
					nodeValue += childrenValues[entry-1]
				}
			}

			return childrenLength + metadataEntriesLength, nodeValue
		}
	}

	_, rootNodeValue := dp(0)
	fmt.Println("result:", rootNodeValue)

}

func parseLicense(license string) []int {
	splits := strings.Split(license, " ")

	var licenseNumbers []int
	for _, split := range splits {
		number, _ := strconv.Atoi(split)
		licenseNumbers = append(licenseNumbers, number)
	}

	return licenseNumbers
}
