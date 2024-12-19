import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import gleam/yielder
import simplifile

pub fn solve_p1() {
  let size = #(50, 6)

  let map_dict =
    list.range(0, size.0)
    |> list.fold(dict.new(), fn(acc, x) {
      list.range(0, size.1)
      |> list.fold(acc, fn(acc, y) { dict.insert(acc, #(x, y), False) })
    })

  simplifile.read("./src/day8/input.txt")
  |> result.unwrap("")
  |> string.split("\n")
  |> list.map(parse_line)
  |> list.fold(map_dict, fn(acc, x) {
    let new_map = run_cmd(acc, x, size)

    print_map_dict(new_map, size)
    io.println("")

    new_map
  })
  |> count_lit_pixels
  |> int.to_string
  |> string.append("Lit pixels: ", _)
  |> io.println
}

pub type Command {
  Create(width: Int, height: Int)
  RotateCol(x: Int, by: Int)
  RotateRow(x: Int, by: Int)
}

fn run_cmd(
  map: dict.Dict(#(Int, Int), Bool),
  cmd: Command,
  size: #(Int, Int),
) -> dict.Dict(#(Int, Int), Bool) {
  case cmd {
    Create(width, height) -> {
      list.range(0, width - 1)
      |> list.fold(map, fn(acc, x) {
        list.range(0, height - 1)
        |> list.fold(acc, fn(acc, y) { dict.insert(acc, #(x, y), True) })
      })
    }
    RotateCol(col, by) -> {
      let cols_to_move =
        map
        |> dict.fold(dict.new(), fn(acc, k, v) {
          let #(x, y) = k

          case x == col {
            True -> {
              case v {
                True -> {
                  let new_y = y + by

                  case new_y > size.1 {
                    True -> dict.insert(acc, #(x, new_y - size.1 - 1), True)
                    False -> dict.insert(acc, #(x, y + by), True)
                  }
                }
                False -> acc
              }
            }
            False -> acc
          }
        })

      dict.fold(map, map, fn(acc, k, v) {
        let #(x, y) = k

        case x == col {
          True -> {
            case
              cols_to_move
              |> dict.has_key(#(x, y))
            {
              True -> dict.insert(acc, #(x, y), True)
              False -> dict.insert(acc, #(x, y), False)
            }
          }
          False -> dict.insert(acc, #(x, y), v)
        }
      })
    }
    RotateRow(row, by) -> {
      let rows_to_move =
        map
        |> dict.fold(dict.new(), fn(acc, k, v) {
          let #(x, y) = k

          case y == row {
            True -> {
              case v {
                True -> {
                  let new_x = x + by

                  case new_x > size.0 {
                    True -> dict.insert(acc, #(new_x - size.0 - 1, y), True)
                    False -> dict.insert(acc, #(x + by, y), True)
                  }
                }
                False -> acc
              }
            }
            False -> acc
          }
        })

      dict.fold(map, map, fn(acc, k, v) {
        let #(x, y) = k

        case y == row {
          True -> {
            case
              rows_to_move
              |> dict.has_key(#(x, y))
            {
              True -> dict.insert(acc, #(x, y), True)
              False -> dict.insert(acc, #(x, y), False)
            }
          }
          False -> dict.insert(acc, #(x, y), v)
        }
      })
    }
  }
}

fn parse_line(line: String) -> Command {
  case line {
    "rect " <> size -> {
      let spl = size |> string.split("x")
      Create(
        spl
          |> yielder.from_list
          |> yielder.at(0)
          |> result.unwrap("?")
          |> int.parse
          |> result.unwrap(0),
        spl
          |> yielder.from_list
          |> yielder.at(1)
          |> result.unwrap("?")
          |> int.parse
          |> result.unwrap(0),
      )
    }
    "rotate column " <> cmd -> {
      let #(x, by) = parse_cmd(cmd)

      RotateCol(x, by)
    }
    "rotate row " <> cmd -> {
      let #(x, by) = parse_cmd(cmd)

      RotateRow(x, by)
    }
    _ -> panic as "invalid command"
  }
}

fn parse_cmd(line: String) -> #(Int, Int) {
  let spl = line |> string.split(" ") |> yielder.from_list

  let x =
    spl
    |> yielder.at(0)
    |> result.unwrap("")
    |> string.split("=")
    |> yielder.from_list
    |> yielder.at(1)
    |> result.unwrap("")
    |> int.parse
    |> result.unwrap(0)

  let by =
    spl
    |> yielder.at(2)
    |> result.unwrap("")
    |> int.parse
    |> result.unwrap(0)

  #(x, by)
}

fn print_map_dict(map: dict.Dict(#(Int, Int), Bool), size: #(Int, Int)) {
  list.range(0, size.1)
  |> list.map(fn(y) {
    list.range(0, size.0)
    |> list.map(fn(x) {
      let v =
        dict.get(map, #(x, y))
        |> result.unwrap(False)

      case v {
        True -> "#"
        False -> "."
      }
    })
    |> string.join("")
  })
  |> string.join("\n")
  |> io.println
}

fn count_lit_pixels(map: dict.Dict(#(Int, Int), Bool)) -> Int {
  map
  |> dict.fold(0, fn(acc, _, v) {
    case v {
      True -> acc + 1
      False -> acc
    }
  })
}
