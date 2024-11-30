package day6

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type command string

const (
	turn_on  command = "turn on"
	turn_off command = "turn off"
	toggle   command = "toggle"
)

type coordinate struct {
	x int
	y int
}

type basicLight struct {
	coordinate
	on bool
}

type advancedLight struct {
	coordinate
	brightness int
}

func SolveP1() {
	lines := utils.ReadFile("day6/input.txt")

	board := createBoard()
	for _, line := range lines {
		cmd, fromPosition, toPosition, err := getCmd(line)
		if err != nil {
			log.Fatal(err)
		}

		for idx, light := range board {
			if isLightIncludedInCmd(light.coordinate, fromPosition, toPosition) {
				workOnBasicLight(cmd, &board[idx])
			}
		}
	}

	fmt.Printf("There are %d lights turned on\n", countBasicLightsTurnedOn(board))
}

func SolveP2() {
	lines := utils.ReadFile("day6/input.txt")

	board := createAdvancedBoard()
	for _, line := range lines {
		cmd, fromPosition, toPosition, err := getCmd(line)
		if err != nil {
			log.Fatal(err)
		}

		for idx, light := range board {
			if isLightIncludedInCmd(light.coordinate, fromPosition, toPosition) {
				workOnAdvancedLight(cmd, &board[idx])
			}
		}
	}

	fmt.Printf("The total brightness of the advancedlights is %d\n", getTotalAdvancedLightsBrightness(board))
}

func getCmd(line string) (command, coordinate, coordinate, error) {
	commands := []command{turn_on, turn_off, toggle}

	pattern := `\b\d+,\d+\b`

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return "", coordinate{}, coordinate{}, err
	}
	matches := re.FindAllString(line, -1)
	if len(matches) != 2 {
		return "", coordinate{}, coordinate{}, err
	}

	fromX, err := strconv.Atoi(strings.Split(matches[0], ",")[0])
	fromY, err := strconv.Atoi(strings.Split(matches[0], ",")[1])
	toX, err := strconv.Atoi(strings.Split(matches[1], ",")[0])
	toY, err := strconv.Atoi(strings.Split(matches[1], ",")[1])

	fromPosition := coordinate{
		x: fromX,
		y: fromY,
	}

	toPosition := coordinate{
		x: toX,
		y: toY,
	}

	if err != nil {
		return "", fromPosition, toPosition, err
	}

	for _, cmd := range commands {
		if strings.Contains(line, string(cmd)) {
			return cmd, fromPosition, toPosition, nil
		}
	}

	return turn_on, fromPosition, toPosition, fmt.Errorf("Could not get the command at line %s\n", line)
}

func workOnBasicLight(cmd command, light *basicLight) {
	switch cmd {
	case turn_on:
		light.on = true
	case turn_off:
		light.on = false
	case toggle:
		light.on = !light.on
	}
}

func workOnAdvancedLight(cmd command, light *advancedLight) {
	switch cmd {
	case turn_on:
		light.brightness++
	case turn_off:
		if light.brightness > 0 {
			light.brightness--
		}
	case toggle:
		light.brightness += 2
	}
}

func isLightIncludedInCmd(lightPosition coordinate, fromPosition coordinate, toPosition coordinate) bool {
	if lightPosition.x >= fromPosition.x && lightPosition.x <= toPosition.x {
		if lightPosition.y >= fromPosition.y && lightPosition.y <= toPosition.y {
			return true
		}
	}

	return false
}

func createBoard() []basicLight {
	board := make([]basicLight, 0, 1000)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			light := basicLight{
				coordinate: coordinate{
					x: i,
					y: j,
				},
				on: false,
			}
			board = append(board, light)
		}
	}
	return board
}

func createAdvancedBoard() []advancedLight {
	board := make([]advancedLight, 0, 1000)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			advancedLight := advancedLight{
				coordinate: coordinate{
					x: i,
					y: j,
				},
				brightness: 0,
			}
			board = append(board, advancedLight)
		}
	}
	return board
}

func countBasicLightsTurnedOn(board []basicLight) int {
	counter := 0
	for _, light := range board {
		if light.on {
			counter++
		}
	}

	return counter
}

func getTotalAdvancedLightsBrightness(board []advancedLight) int {
	totalBrightness := 0
	for _, light := range board {
		totalBrightness += light.brightness
	}

	return totalBrightness
}
