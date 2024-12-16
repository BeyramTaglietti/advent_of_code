package day16

import (
	"aoc2024/utils"
	"fmt"
	"maps"
	"slices"
	"sync"
)

type point struct {
	x int
	y int
}

type direction string

const (
	up    direction = "up"
	down  direction = "down"
	left  direction = "left"
	right direction = "right"
)

type path struct {
	current_position point
	facing           direction
	visited_points   map[point]bool
	score            int
}

func SolveP1() {
	lines := utils.ReadFile("./day16/input.txt")

	m, entrance_point, exit_point := parseMap(lines)

	valid_paths := traverse(m, []path{{current_position: entrance_point, visited_points: map[point]bool{entrance_point: true}, score: 1}}, exit_point, []path{})

	fmt.Println("valid paths found:", len(valid_paths))

	lowest_score_path := findLowestScorePath(valid_paths)
	fmt.Println("lowest score path:", lowest_score_path.score, "and it took", len(lowest_score_path.visited_points), "steps")

}

func SolveP2() {
	lines := utils.ReadFile("./day16/input.txt")

	m, entrance_point, exit_point := parseMap(lines)

	valid_paths := traverse(m, []path{{current_position: entrance_point, visited_points: map[point]bool{entrance_point: true}, score: 1}}, exit_point, []path{})

	fmt.Println("valid paths found:", len(valid_paths))

	lowest_score := findLowestScorePath(valid_paths)

	mostEfficientPaths := []path{}
	for _, p := range valid_paths {
		if p.score == lowest_score.score {
			mostEfficientPaths = append(mostEfficientPaths, p)
		}
	}

	fmt.Println("most efficient paths found:", len(mostEfficientPaths))

	fmt.Println("common tiles:", len(commonTiles(mostEfficientPaths)))
}

func parseMap(lines []string) (m map[point]bool, entrance, exit point) {

	m = make(map[point]bool)

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case 'S':
				entrance = point{x, y}
				m[point{x, y}] = false
			case 'E':
				exit = point{x, y}
				m[point{x, y}] = false
			case '#':
				m[point{x, y}] = true
			}
		}
	}
	return m, entrance, exit
}

func traverse(m map[point]bool, current_paths []path, exit_point point, valid_paths []path) []path {

	current_paths = cleanSimilarPaths(current_paths)

	if len(valid_paths) > 0 {
		lowestScore := findLowestScorePath(valid_paths).score
		current_paths = cleanInefficientPaths(current_paths, lowestScore)
	}

	if len(current_paths) == 0 {
		return valid_paths
	}

	newPaths := []path{}
	wg := sync.WaitGroup{}
	wg.Add(len(current_paths))
	mu := sync.Mutex{}
	for _, p := range current_paths {

		go func(p path) {
			defer wg.Done()
			next_possible_pos_scores := findNextPoints(m, p)

			next_possible_positions := []point{}

			for k := range next_possible_pos_scores {
				next_possible_positions = append(next_possible_positions, k)
			}
			if slices.Contains(next_possible_positions, exit_point) {
				mu.Lock()
				valid_paths = append(valid_paths, p)
				mu.Unlock()
			}

			for next_p, next_p_score := range next_possible_pos_scores {
				visited_p := maps.Clone(p.visited_points)
				visited_p[next_p] = true

				newPath := path{
					current_position: next_p,
					visited_points:   visited_p,
					facing:           next_p_score.facing,
					score:            p.score + next_p_score.score,
				}

				mu.Lock()
				newPaths = append(newPaths, newPath)
				mu.Unlock()
			}
		}(p)

	}

	wg.Wait()

	return traverse(m, newPaths, exit_point, valid_paths)
}

type nextPointScore struct {
	p      point
	score  int
	facing direction
}

func findNextPoints(m map[point]bool, current_path path) map[point]nextPointScore {

	var next_points []point

	possible_next_points := []point{
		{current_path.current_position.x, current_path.current_position.y - 1},
		{current_path.current_position.x, current_path.current_position.y + 1},
		{current_path.current_position.x - 1, current_path.current_position.y},
		{current_path.current_position.x + 1, current_path.current_position.y},
	}

	for _, p := range possible_next_points {
		if !m[p] && !current_path.visited_points[p] {
			next_points = append(next_points, p)
		}
	}

	find_new_facing := func(to_point point) direction {
		if to_point.x == current_path.current_position.x-1 {
			return left
		}
		if to_point.x == current_path.current_position.x+1 {
			return right
		}
		if to_point.y == current_path.current_position.y-1 {
			return up
		}
		if to_point.y == current_path.current_position.y+1 {
			return down
		}
		return current_path.facing
	}

	scores := make(map[point]nextPointScore)

	for _, p := range next_points {

		new_facing := find_new_facing(p)

		if new_facing != current_path.facing || current_path.facing == "" {
			scores[p] = nextPointScore{p, 1001, new_facing}
		} else {
			scores[p] = nextPointScore{p, 1, new_facing}
		}
	}

	return scores
}

func printPath(m map[point]bool, p path, mapWidth, mapHeight int) {
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if m[point{x, y}] {
				fmt.Print("#")
			} else if p.visited_points[point{x, y}] {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findLowestScorePath(paths []path) path {
	minScore := paths[0].score
	minPath := paths[0]
	for _, p := range paths {
		if p.score < minScore {
			minScore = p.score
			minPath = p
		}
	}

	return minPath
}

func cleanInefficientPaths(paths []path, lowest_score int) []path {
	newPaths := []path{}
	for _, p := range paths {
		if p.score <= lowest_score {
			newPaths = append(newPaths, p)
		}
	}

	return newPaths
}

func cleanSimilarPaths(paths []path) []path {
	type simple_path struct {
		current_position point
		score            int
	}
	newPaths := make(map[simple_path]path)
	for _, p := range paths {
		newPaths[simple_path{p.current_position, p.score}] = p
	}

	cleanedPaths := []path{}
	for k := range newPaths {
		cleanedPaths = append(cleanedPaths, newPaths[k])
	}

	return cleanedPaths
}

func commonTiles(valid_paths []path) []point {
	commonTiles := make(map[point]bool)
	for _, p := range valid_paths {
		for k := range p.visited_points {
			commonTiles[k] = true
		}
	}

	commonTilesList := []point{}
	for k := range commonTiles {
		commonTilesList = append(commonTilesList, k)
	}

	return commonTilesList
}
