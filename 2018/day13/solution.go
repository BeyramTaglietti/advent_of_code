package day13

import (
	"2018/utils"
	"cmp"
	"fmt"
	"slices"
)

type direction string

const (
	up    direction = "^"
	down  direction = "v"
	left  direction = "<"
	right direction = ">"
)

type cart struct {
	going    direction
	nextTurn direction
}

func newCart(dir direction) cart {
	return cart{
		going:    dir,
		nextTurn: left,
	}
}

func (c cart) move() (int, int) {
	switch c.going {
	case up:
		return -1, 0
	case left:
		return 0, -1
	case right:
		return 0, 1
	case down:
		return 1, 0
	default:
		return 0, 0
	}
}

func (c *cart) turnIntersection() {
	switch c.nextTurn {
	case right:
		switch c.going {
		case down:
			c.going = left
		case up:
			c.going = right
		case right:
			c.going = down
		case left:
			c.going = up
		}
		c.nextTurn = left
	case up:
		c.nextTurn = right
	case left:
		switch c.going {
		case down:
			c.going = right
		case up:
			c.going = left
		case right:
			c.going = up
		case left:
			c.going = down
		}
		c.nextTurn = up
	}
}

func (c *cart) turnCorner(cornerTurn turn) {
	switch c.going {
	case left:
		if cornerTurn == forward {
			c.going = down
		} else {
			c.going = up
		}
	case right:
		if cornerTurn == forward {
			c.going = up
		} else {
			c.going = down
		}
	case up:
		if cornerTurn == forward {
			c.going = right
		} else {
			c.going = left
		}
	case down:
		if cornerTurn == forward {
			c.going = left
		} else {
			c.going = right
		}
	}
}

type orientation string

const (
	horizontal orientation = "-"
	vertical   orientation = "|"
)

type track struct {
	intersection bool
	corner       bool
	cornerTurn   turn
	orientation  orientation
	empty        bool
}

type turn string

const (
	forward  turn = "/"
	backward turn = "\\"
)

type point struct {
	y, x int
}

func SolveP1() {
	lines := utils.ReadFile("day13/input.txt")

	tracks := make(map[point]track)
	carts := make(map[point]cart)

	m := len(lines)
	n := len(lines[0])

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			currentPoint := point{y, x}
			switch lines[y][x] {
			case '>':
				carts[currentPoint] = newCart(right)
				tracks[currentPoint] = track{orientation: horizontal}
			case '<':
				carts[currentPoint] = newCart(left)
				tracks[currentPoint] = track{orientation: horizontal}
			case 'v':
				carts[currentPoint] = newCart(down)
				tracks[currentPoint] = track{orientation: vertical}
			case '^':
				carts[currentPoint] = newCart(up)
				tracks[currentPoint] = track{orientation: vertical}
			case '-':
				tracks[currentPoint] = track{orientation: horizontal}
			case '|':
				tracks[currentPoint] = track{orientation: vertical}
			case '+':
				tracks[currentPoint] = track{intersection: true}
			case '\\':
				tracks[currentPoint] = track{orientation: vertical, corner: true, cornerTurn: backward}
			case '/':
				tracks[currentPoint] = track{orientation: horizontal, corner: true, cornerTurn: forward}
			default:
				tracks[currentPoint] = track{empty: true}
			}
		}
	}

	for {
		positions := make([]point, 0, len(carts))
		for p := range carts {
			positions = append(positions, p)
		}
		slices.SortFunc(positions, func(a point, b point) int {
			if a.y != b.y {
				return cmp.Compare(a.y, b.y)
			}
			return cmp.Compare(a.x, b.x)
		})

		next := make(map[point]cart, len(carts))

		for _, pos := range positions {
			c, ok := carts[pos]
			if !ok {
				continue
			}
			delete(carts, pos)

			ny, nx := c.move()
			np := point{pos.y + ny, pos.x + nx}

			if _, ok := next[np]; ok {
				fmt.Println("collision at", np)
				return
			}
			if _, ok := carts[np]; ok {
				fmt.Println("collision at", np)
				return
			}

			t := tracks[np]
			if t.intersection {
				c.turnIntersection()
			} else if t.corner {
				c.turnCorner(t.cornerTurn)
			}

			next[np] = c
		}

		carts = next

	}

}

func SolveP2() {
	lines := utils.ReadFile("day13/input.txt")

	tracks := make(map[point]track)
	carts := make(map[point]cart)

	m := len(lines)
	n := len(lines[0])

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			currentPoint := point{y, x}
			switch lines[y][x] {
			case '>':
				carts[currentPoint] = newCart(right)
				tracks[currentPoint] = track{orientation: horizontal}
			case '<':
				carts[currentPoint] = newCart(left)
				tracks[currentPoint] = track{orientation: horizontal}
			case 'v':
				carts[currentPoint] = newCart(down)
				tracks[currentPoint] = track{orientation: vertical}
			case '^':
				carts[currentPoint] = newCart(up)
				tracks[currentPoint] = track{orientation: vertical}
			case '-':
				tracks[currentPoint] = track{orientation: horizontal}
			case '|':
				tracks[currentPoint] = track{orientation: vertical}
			case '+':
				tracks[currentPoint] = track{intersection: true}
			case '\\':
				tracks[currentPoint] = track{orientation: vertical, corner: true, cornerTurn: backward}
			case '/':
				tracks[currentPoint] = track{orientation: horizontal, corner: true, cornerTurn: forward}
			default:
				tracks[currentPoint] = track{empty: true}
			}
		}
	}

	for {

		if len(carts) == 1 {
			for p := range carts {
				fmt.Println("last cart at", p)
				return
			}
		}

		positions := make([]point, 0, len(carts))
		for p := range carts {
			positions = append(positions, p)
		}
		slices.SortFunc(positions, func(a point, b point) int {
			if a.y != b.y {
				return cmp.Compare(a.y, b.y)
			}
			return cmp.Compare(a.x, b.x)
		})

		next := make(map[point]cart, len(carts))

		for _, pos := range positions {
			c, ok := carts[pos]
			if !ok {
				continue
			}
			delete(carts, pos)

			ny, nx := c.move()
			np := point{pos.y + ny, pos.x + nx}

			if _, ok := next[np]; ok {
				delete(carts, np)
				delete(next, np)
				continue
			}
			if _, ok := carts[np]; ok {
				delete(carts, np)
				delete(next, np)
				continue
			}

			t := tracks[np]
			if t.intersection {
				c.turnIntersection()
			} else if t.corner {
				c.turnCorner(t.cornerTurn)
			}

			next[np] = c
		}

		carts = next

	}

}
