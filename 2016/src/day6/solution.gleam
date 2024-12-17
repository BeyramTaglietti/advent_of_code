import gleam/dict
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import gleam/yielder
import simplifile

pub fn solve_p1() {
  simplifile.read("./src/day6/input.txt")
  |> result.unwrap("")
  |> string.split("\n")
  |> create_columns
  |> list.fold("", fn(acc, x) { acc |> string.append(x |> most_common_letter) })
  |> io.println
}

pub fn solve_p2() {
  simplifile.read("./src/day6/input.txt")
  |> result.unwrap("")
  |> string.split("\n")
  |> create_columns
  |> list.fold("", fn(acc, x) { acc |> string.append(x |> least_common_letter) })
  |> io.println
}

fn create_columns(rows: List(String)) -> List(String) {
  rows
  |> list.fold(dict.new(), fn(acc, enc_word) {
    let #(new_columns, _) =
      enc_word
      |> string.split("")
      |> list.fold(#(acc, 0), fn(acc2, letter) {
        let #(columns, index) = acc2

        let new_val =
          dict.get(columns, index)
          |> result.unwrap("")
          |> string.append(letter)
        #(dict.insert(columns, index, new_val), index + 1)
      })

    new_columns
  })
  |> dict.to_list
  |> list.map(fn(x) {
    let #(_, v) = x

    v
  })
}

fn most_common_letter(input: String) -> String {
  let #(_, most_frequent_letter) =
    input
    |> string.split("")
    |> list.fold(dict.new(), fn(acc, x) {
      let prev = dict.get(acc, x) |> result.unwrap(0)

      dict.insert(acc, x, prev + 1)
    })
    |> dict.fold(#(0, ""), fn(acc, k, v) {
      let #(max, _) = acc

      case v > max {
        True -> #(v, k)
        False -> acc
      }
    })

  most_frequent_letter
}

fn least_common_letter(input: String) -> String {
  let frequency_map =
    input
    |> string.split("")
    |> list.fold(dict.new(), fn(acc, x) {
      let prev = dict.get(acc, x) |> result.unwrap(0)

      dict.insert(acc, x, prev + 1)
    })

  let first_letter_frequency =
    frequency_map
    |> dict.to_list
    |> list.map(fn(x) {
      let #(_, f) = x
      f
    })
    |> yielder.from_list
    |> yielder.first
    |> result.unwrap(0)

  let #(_, least_frequent_letter) =
    frequency_map
    |> dict.fold(#(first_letter_frequency, ""), fn(acc, k, v) {
      let #(max, _) = acc

      case v < max {
        True -> #(v, k)
        False -> acc
      }
    })

  least_frequent_letter
}
