import gleam/bit_array
import gleam/crypto
import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import gleam/yielder

pub fn solve_p1() {
  let password = "abc"

  password
  |> find_5_zeros_hash(0, "", 8)
  |> io.debug
}

pub fn solve_p2() {
  let password = "abc"
  password
  |> find_5_zeroes_dict(0, dict.new(), 8)
  |> io.debug
}

fn find_5_zeros_hash(
  input: String,
  index: Int,
  current_password: String,
  required_length: Int,
) -> String {
  let new_hash =
    input
    |> string.append(
      index
      |> int.to_string,
    )
    |> hash_string

  case new_hash |> contains_leading_zeroes {
    True -> {
      let new_pass =
        string.append(
          current_password,
          new_hash
            |> string.split("")
            |> yielder.from_list
            |> yielder.at(5)
            |> result.unwrap("?"),
        )

      case string.length(new_pass) == required_length {
        True -> new_pass
        False -> find_5_zeros_hash(input, index + 1, new_pass, required_length)
      }
    }
    _ -> {
      find_5_zeros_hash(input, index + 1, current_password, required_length)
    }
  }
}

fn find_5_zeroes_dict(
  input: String,
  index: Int,
  current_password: dict.Dict(Int, String),
  required_length: Int,
) -> String {
  let hashed_str =
    input
    |> string.append(index |> int.to_string)
    |> hash_string

  case hashed_str |> contains_leading_zeroes {
    True -> {
      let iter = hashed_str |> string.split("") |> yielder.from_list
      let put_char = iter |> yielder.at(6) |> result.unwrap("?")
      let put_at_idx =
        iter
        |> yielder.at(5)
        |> result.unwrap("?")

      case
        put_at_idx |> int.parse |> result.is_ok
        && put_at_idx |> int.parse |> result.unwrap(0) < required_length
      {
        True -> {
          let parsed_idx = put_at_idx |> int.parse |> result.unwrap(0)
          case dict.has_key(current_password, parsed_idx) {
            True ->
              find_5_zeroes_dict(
                input,
                index + 1,
                current_password,
                required_length,
              )
            False -> {
              let new_dict = dict.insert(current_password, parsed_idx, put_char)

              case new_dict |> dict.to_list |> list.length == required_length {
                True ->
                  new_dict
                  |> dict.to_list
                  |> list.fold("", fn(acc, x) {
                    let #(_, char) = x
                    acc |> string.append(char)
                  })
                False -> {
                  find_5_zeroes_dict(
                    input,
                    index + 1,
                    new_dict,
                    required_length,
                  )
                }
              }
            }
          }
        }
        _ ->
          find_5_zeroes_dict(
            input,
            index + 1,
            current_password,
            required_length,
          )
      }
    }
    False ->
      find_5_zeroes_dict(input, index + 1, current_password, required_length)
  }
}

fn contains_leading_zeroes(input: String) -> Bool {
  case input {
    "00000" <> _ -> True
    _ -> False
  }
}

fn hash_string(input: String) {
  <<input:utf8>>
  |> crypto.hash(crypto.Md5, _)
  |> bit_array.base16_encode
}
