import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import gleam/yielder
import simplifile

pub fn solve_p1() {
  simplifile.read("./src/day9/input.txt")
  |> result.unwrap("")
  |> string.split("\n")
  |> list.fold(0, fn(acc, value) {
    let #(decompressed_value, _, _) =
      value
      |> string.split("")
      |> list.fold(#("", 0, 0), fn(acc, x) {
        let #(current_str, index, to_skip) = acc

        case to_skip > 0 {
          True -> #(current_str, index + 1, to_skip - 1)
          False ->
            case x {
              "(" -> {
                let #(repeat_len, repeat_times, marker_str_len) =
                  get_marker_value(
                    string.slice(value, index, string.length(value))
                    |> string.split(""),
                  )

                let marked_str =
                  string.slice(value, index + marker_str_len + 1, repeat_len)
                  |> exectute_marker(repeat_times)

                #(
                  current_str
                    |> string.append(marked_str),
                  index + 1,
                  marker_str_len + repeat_len,
                )
              }
              _ -> #(current_str |> string.append(x), index + 1, to_skip)
            }
        }
      })

    let decompressed_value_len = string.length(decompressed_value)
    acc + decompressed_value_len
  })
  |> int.to_string
  |> string.append("result: ", _)
  |> io.println
}

fn get_marker_value(l: List(String)) -> #(Int, Int, Int) {
  let #(x, _, index) =
    list.fold(l, #("", False, 0), fn(acc, x) {
      let #(val, skip, index) = acc

      case skip {
        True -> acc
        False -> {
          case x {
            ")" -> {
              #(val |> string.append(x), True, index)
            }
            _ -> #(val |> string.append(x), False, index + 1)
          }
        }
      }
    })

  let y =
    x
    |> string.replace(")", "")
    |> string.replace("(", "")
    |> string.split("x")
    |> list.map(fn(x) { x |> int.parse |> result.unwrap(0) })
    |> yielder.from_list

  #(
    y |> yielder.at(0) |> result.unwrap(0),
    y |> yielder.at(1) |> result.unwrap(0),
    index,
  )
}

fn exectute_marker(input: String, repeat_times: Int) -> String {
  list.range(0, repeat_times - 1)
  |> list.fold("", fn(acc, _) { acc |> string.append(input) })
}
