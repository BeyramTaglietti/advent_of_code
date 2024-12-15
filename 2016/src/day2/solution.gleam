import gleam/dict
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub type Point {
  Point(x: Int, y: Int)
}

pub type Direction {
  L
  R
  U
  D
}

// 1 2 3
// 4 5 6
// 7 8 9
pub fn solve_p1() {
  let assert Ok(lines) = simplifile.read("./src/day2/input.txt")

  let keypad =
    dict.from_list([
      #(Point(0, 0), "1"),
      #(Point(1, 0), "2"),
      #(Point(2, 0), "3"),
      #(Point(0, 1), "4"),
      #(Point(1, 1), "5"),
      #(Point(2, 1), "6"),
      #(Point(0, 2), "7"),
      #(Point(1, 2), "8"),
      #(Point(2, 2), "9"),
    ])

  let starting_position = Point(1, 1)

  let last_pos =
    list.fold(
      string.split(lines, "\n"),
      #(starting_position, ""),
      fn(acc, line) {
        let #(current_p, phone_n) = acc
        let next_p =
          parse_line(line)
          |> calculate_next_p(current_p, keypad)

        #(
          next_p,
          string.append(
            phone_n,
            dict.get(keypad, next_p)
              |> result.unwrap(""),
          ),
        )
      },
    )

  io.debug(last_pos)
}

//     1
//   2 3 4
// 5 6 7 8 9
//   A B C
//     D
pub fn solve_p2() {
  let assert Ok(lines) = simplifile.read("./src/day2/input.txt")

  let keypad =
    dict.from_list([
      #(Point(2, 0), "1"),
      #(Point(1, 1), "2"),
      #(Point(2, 1), "3"),
      #(Point(3, 1), "4"),
      #(Point(0, 2), "5"),
      #(Point(1, 2), "6"),
      #(Point(2, 2), "7"),
      #(Point(3, 2), "8"),
      #(Point(4, 2), "9"),
      #(Point(1, 3), "A"),
      #(Point(2, 3), "B"),
      #(Point(3, 3), "C"),
      #(Point(2, 4), "D"),
    ])

  let starting_position = Point(1, 1)

  let last_pos =
    list.fold(
      string.split(lines, "\n"),
      #(starting_position, ""),
      fn(acc, line) {
        let #(current_p, phone_n) = acc
        let next_p =
          parse_line(line)
          |> calculate_next_p(current_p, keypad)

        #(
          next_p,
          string.append(
            phone_n,
            dict.get(keypad, next_p)
              |> result.unwrap(""),
          ),
        )
      },
    )

  io.debug(last_pos)
}

fn parse_line(line: String) -> List(Direction) {
  string.split(line, "")
  |> list.fold(list.new(), fn(acc, c) {
    case c {
      "L" -> list.append(acc, [L])
      "R" -> list.append(acc, [R])
      "U" -> list.append(acc, [U])
      "D" -> list.append(acc, [D])
      _ -> acc
    }
  })
}

fn calculate_next_p(
  moves: List(Direction),
  current_p: Point,
  keypad: dict.Dict(Point, String),
) -> Point {
  let check_bounds = fn(next_p: Point, current_p: Point) {
    case dict.has_key(keypad, next_p) {
      True -> next_p
      False -> current_p
    }
  }

  list.fold(moves, current_p, fn(acc, move) {
    case move {
      L -> {
        Point(acc.x - 1, acc.y)
        |> check_bounds(acc)
      }
      R -> {
        Point(acc.x + 1, acc.y)
        |> check_bounds(acc)
      }
      U -> {
        Point(acc.x, acc.y - 1)
        |> check_bounds(acc)
      }
      D -> {
        Point(acc.x, acc.y + 1)
        |> check_bounds(acc)
      }
    }
  })
}
