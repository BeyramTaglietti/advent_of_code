package day2

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func SolveP1() {
	lines, err := utils.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	totalFeet := 0

	for _, line := range lines {
		lineArea, err := getPrismArea(line)
		if err != nil {
			log.Println("There was an error getting the area of line:", line)
			return
		}
		totalFeet += lineArea
	}

	fmt.Printf("Total area: %d feet\n", totalFeet)
}

func SolveP2() {
	lines, err := utils.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	totalRibbonArea := 0

	for _, line := range lines {
		if line != "" {
			tot, err := getRibbonArea(line)
			if err != nil {
				log.Println("There was an error getting the ribbon area of line", line)
			}
			totalRibbonArea += tot
		}
	}
	fmt.Printf("Ribbon area: %d feet\n", totalRibbonArea)

}

func getRibbonArea(expr string) (int, error) {
	var l, w, h int

	arr := strings.Split(expr, "x")
	l, err := strconv.Atoi(arr[0])
	w, err = strconv.Atoi(arr[1])
	h, err = strconv.Atoi(arr[2])

	if err != nil {
		return 0, err
	}

	numericArr := make([]int, len(arr), len(arr))
	for idx, run := range arr {
		num, err := strconv.Atoi(run)
		if err != nil {
			return 0, err
		}
		numericArr[idx] = num
	}

	first := slices.Min(numericArr)

	deleted := false
	filteredArr := make([]int, 0, len(arr)-1)
	for _, val := range numericArr {
		if val == first && deleted == false {
			deleted = true
			continue
		}

		filteredArr = append(filteredArr, val)
	}

	second := slices.Min(
		filteredArr,
	)

	return first*2 + second*2 + l*w*h, nil
}

func getPrismArea(expr string) (int, error) {
	var l, w, h int

	arr := strings.Split(expr, "x")
	l, err := strconv.Atoi(arr[0])
	w, err = strconv.Atoi(arr[1])
	h, err = strconv.Atoi(arr[2])

	if err != nil {
		return 0, err
	}

	area := 2*l*w + 2*w*h + 2*h*l
	extraArea := math.Min(float64(l*w), math.Min(float64(h*w), float64(l*h)))

	return area + int(extraArea), nil
}
