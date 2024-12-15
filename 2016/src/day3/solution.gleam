import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub fn solve_p1() {
  let assert Ok(lines) = simplifile.read("./src/day3/input.txt")

  lines
  |> string.split("\n")
  |> list.fold(0, fn(acc, line) {
    let is_valid_triangle =
      line
      |> string.trim
      |> string.split(" ")
      |> list.filter(fn(v) { v != "" })
      |> list.map(fn(x) { x |> int.parse |> result.unwrap(0) })
      |> is_valid

    case is_valid_triangle {
      True -> acc + 1
      _ -> acc
    }
  })
  |> int.to_string
  |> string.append(" valid triangles found")
  |> io.println
}

pub fn solve_p2() {
  let assert Ok(lines) = simplifile.read("./src/day3/input.txt")

  lines
  |> string.split("\n")
  |> list.map(fn(x) { string.trim(x) })
  |> split_into_threes(list.new(), 0)
  |> list.map(fn(x) { x |> create_column_triangle })
  |> list.fold(list.new(), fn(acc, group) {
    let #(g1, g2, g3) = group

    list.prepend(acc, g1)
    |> list.prepend(g2)
    |> list.prepend(g3)
  })
  |> list.fold(0, fn(acc, group) {
    let is_valid_triangle =
      group
      |> list.map(fn(x) { x |> int.parse |> result.unwrap(0) })
      |> is_valid

    case is_valid_triangle {
      True -> acc + 1
      False -> acc
    }
  })
  |> int.to_string
  |> io.println
}

fn is_valid(sides: List(Int)) -> Bool {
  case sides {
    [s1, s2, s3] -> {
      s1 + s2 > s3 && s1 + s3 > s2 && s2 + s3 > s1
    }
    _ -> False
  }
}

fn split_into_threes(
  lines: List(String),
  groups: List(List(String)),
  index: Int,
) -> List(List(String)) {
  case list.is_empty(lines) {
    True -> groups
    False -> {
      split_into_threes(
        list.drop(lines, 3),
        list.append(groups, [list.take(lines, 3)]),
        index + 1,
      )
    }
  }
}

fn create_column_triangle(
  column: List(String),
) -> #(List(String), List(String), List(String)) {
  column
  |> list.fold(#(list.new(), list.new(), list.new()), fn(acc, row) {
    let #(r1, r2, r3) = acc
    case string.split(string.trim(row), " ") |> list.filter(fn(v) { v != "" }) {
      [s1, s2, s3] -> {
        #(list.append(r1, [s1]), list.append(r2, [s2]), list.append(r3, [s3]))
      }
      _ -> panic as "could not create column triangle"
    }
  })
}
