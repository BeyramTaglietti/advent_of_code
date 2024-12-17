import gleam/io
import gleam/list
import gleam/result
import gleam/string
import gleam/yielder
import simplifile

pub fn solve_p1() {
  simplifile.read("./src/day7/input.txt")
  |> result.unwrap("")
  |> string.split("\n")
  |> list.map(fn(x) {
    let #(valids, invalids) = x |> parse_line

    let valids_contain_pattern =
      list.any(valids, fn(x) { x |> contains_abba_pattern })

    let invalids_contain_pattern =
      list.any(invalids, fn(x) { x |> contains_abba_pattern })

    valids_contain_pattern && !invalids_contain_pattern
  })
  |> list.filter(fn(x) { x == True })
  |> list.length
  |> io.debug
}

pub fn solve_p2() {
  simplifile.read("./src/day7/input.txt")
  |> result.unwrap("")
  |> string.split("\n")
  |> list.map(fn(x) {
    let #(valids, invalids) = x |> parse_line

    list.fold(valids, False, fn(acc, valid_input) {
      case acc {
        True -> True
        False -> {
          let valid_patterns_found =
            contains_aba_pattern(valid_input, list.new())

          list.fold(
            valid_patterns_found,
            False,
            fn(at_least_one_valied, valid_p_f) {
              case at_least_one_valied {
                True -> True
                False -> {
                  case contains_inverse_aba_pattern(valid_p_f, invalids) {
                    True -> True
                    False -> False
                  }
                }
              }
            },
          )
        }
      }
    })
  })
  |> list.filter(fn(x) { x == True })
  |> list.length
  |> io.debug
}

fn parse_line(line: String) -> #(List(String), List(String)) {
  let #(valids, invalids, _) =
    line
    |> string.split("[")
    |> list.map(fn(x) { x |> string.split("]") })
    |> list.flatten
    |> list.fold(#(list.new(), list.new(), 0), fn(acc, x) {
      let #(valids, invalids, index) = acc

      case index % 2 != 0 {
        True -> #(valids, list.prepend(invalids, x), index + 1)
        False -> #(valids |> list.prepend(x), invalids, index + 1)
      }
    })

  #(valids, invalids)
}

fn contains_abba_pattern(input: String) -> Bool {
  case string.length(input) < 4 {
    True -> False
    False -> {
      let y = input |> string.split("") |> yielder.from_list

      let #(l1, l2, l3, l4) = #(
        get_from_iterator(y, 0),
        get_from_iterator(y, 1),
        get_from_iterator(y, 2),
        get_from_iterator(y, 3),
      )

      case l1 == l4 && l2 == l3 && l1 != l2 {
        True -> {
          True
        }
        False ->
          contains_abba_pattern(
            string.split(input, "")
            |> list.drop(1)
            |> list.fold("", fn(acc, x) { acc |> string.append(x) }),
          )
      }
    }
  }
}

fn contains_aba_pattern(
  input: String,
  valid_patterns: List(String),
) -> List(String) {
  case string.length(input) < 3 {
    True -> valid_patterns
    False -> {
      let y = input |> string.split("") |> yielder.from_list

      let #(l1, l2, l3) = #(
        get_from_iterator(y, 0),
        get_from_iterator(y, 1),
        get_from_iterator(y, 2),
      )

      let pattern =
        l1
        |> string.append(l2)
        |> string.append(l3)

      let rest = input |> string.slice(1, string.length(input))

      case l1 == l3 && l1 != l2 {
        True -> {
          contains_aba_pattern(
            rest,
            valid_patterns
              |> list.prepend(pattern),
          )
        }
        False -> contains_aba_pattern(rest, valid_patterns)
      }
    }
  }
}

fn contains_inverse_aba_pattern(pattern: String, check: List(String)) -> Bool {
  let patterns_found =
    check
    |> list.map(fn(x) { contains_aba_pattern(x, list.new()) })
    |> list.flatten

  list.fold(patterns_found, False, fn(acc, x) {
    let p =
      pattern
      |> string.split("")
      |> yielder.from_list

    let pattern_to_check =
      ""
      |> string.append(get_from_iterator(p, 1))
      |> string.append(get_from_iterator(p, 0))
      |> string.append(get_from_iterator(p, 1))

    case acc {
      True -> True
      False -> {
        case pattern_to_check == x {
          True -> True
          False -> False
        }
      }
    }
  })
}

fn get_from_iterator(i: yielder.Yielder(String), index: Int) -> String {
  i |> yielder.at(index) |> result.unwrap("?")
}
