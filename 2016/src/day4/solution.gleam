import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub type Encryption {
  Encryption(check_sum: String, sector_id: Int, password: String)
}

pub fn solve_p1() {
  let assert Ok(lines) = simplifile.read("./src/day4/input.txt")

  lines
  |> string.split("\n")
  |> list.map(fn(x) { x |> parse_line })
  |> list.fold(0, fn(acc, x) {
    let is_valid =
      x
      |> get_letters_frequency
      |> dict.delete("-")
      |> is_valid_password(x.check_sum, _)

    case is_valid {
      True -> acc + x.sector_id
      _ -> acc
    }
  })
  |> int.to_string
  |> string.append("sector id's sum: ", _)
  |> io.println
}

pub fn solve_p2() {
  let assert Ok(lines) = simplifile.read("./src/day4/input.txt")

  lines
  |> string.split("\n")
  |> list.map(fn(x) { x |> parse_line })
  |> list.filter(fn(x) {
    x
    |> get_letters_frequency
    |> dict.delete("-")
    |> is_valid_password(x.check_sum, _)
  })
  |> list.map(fn(x) { x |> decrypt })
  |> list.filter(fn(x) {
    let #(decrypted_value, _) = x

    string.contains(decrypted_value, "object")
  })
  |> list.each(fn(x) { io.debug(x) })
}

fn parse_line(line: String) -> Encryption {
  let checksum =
    line
    |> string.split("[")
    |> list.last
    |> result.unwrap("")
    |> string.split("]")
    |> list.first
    |> result.unwrap("")

  let pass =
    line
    |> string.split("[")
    |> list.first
    |> result.unwrap("")
    |> string.split("-")

  let sector_id =
    pass
    |> list.last
    |> result.unwrap("")
    |> int.parse
    |> result.unwrap(0)

  let password =
    line
    |> string.split(sector_id |> int.to_string |> string.append("-", _))
    |> list.first
    |> result.unwrap("")

  Encryption(checksum, sector_id, password)
}

fn get_letters_frequency(enc: Encryption) -> dict.Dict(String, Int) {
  enc.password
  |> string.split("")
  |> list.fold(dict.new(), fn(acc, x) {
    let val =
      acc
      |> dict.get(x)
      |> result.unwrap(0)

    acc
    |> dict.insert(x, val + 1)
  })
}

fn valid_letters(letters: dict.Dict(String, Int)) -> List(String) {
  let required_frequency =
    letters
    |> dict.fold(0, fn(acc, _, v) {
      case v > acc {
        True -> v
        _ -> acc
      }
    })

  letters
  |> dict.filter(fn(_, freq) { freq == required_frequency })
  |> dict.to_list
  |> list.map(fn(x) {
    let #(letter, _) = x
    letter
  })
}

fn is_valid_password(
  check_sum: String,
  frequency_map: dict.Dict(String, Int),
) -> Bool {
  let #(valid, _) =
    check_sum
    |> string.split("")
    |> list.fold(#(True, frequency_map), fn(acc, x) {
      let #(curr_validity, freq_map) = acc
      case curr_validity {
        True -> {
          let is_valid =
            valid_letters(freq_map)
            |> list.fold("", fn(full_str, char) {
              full_str |> string.append(char)
            })
            |> string.contains(x)

          #(is_valid, freq_map |> dict.delete(x))
        }
        False -> #(False, freq_map)
      }
    })

  valid
}

fn decrypt(encrypted_value: Encryption) -> #(String, Int) {
  let #(alphabet, _) =
    "abcdefghijklmnopqrstuvwxyz"
    |> string.split("")
    |> list.fold(#(dict.new(), 0), fn(acc, letter) {
      let #(current_dict, current_idx) = acc

      #(current_dict |> dict.insert(current_idx, letter), current_idx + 1)
    })

  let get_alphabet_idx = fn(letter: String) {
    alphabet
    |> dict.fold(#(0, False), fn(acc, k, v) {
      let #(idx, found) = acc

      case found {
        True -> #(idx, found)
        _ -> {
          case letter == v {
            True -> {
              #(k, True)
            }
            False -> acc
          }
        }
      }
    })
  }

  let move_by = encrypted_value.sector_id % 26

  encrypted_value.password
  |> string.split("")
  |> list.fold(#("", 0), fn(acc, letter) {
    let #(current_pass, _) = acc

    let #(curr_idx, _) = get_alphabet_idx(letter)

    let move_to = case curr_idx + move_by > 25 {
      True -> curr_idx + move_by - 26
      False -> curr_idx + move_by
    }

    let next_letter = case letter == "-" {
      True -> " "
      False -> {
        dict.get(alphabet, move_to)
        |> result.unwrap("?")
      }
    }

    #(current_pass |> string.append(next_letter), encrypted_value.sector_id)
  })
}
