import birl
import birl/duration
import day9/solution
import gleam/int
import gleam/io
import gleam/string

pub fn main() {
  let now = birl.now()
  solution.solve_p1()
  let end = birl.now()

  birl.difference(end, now)
  |> duration.blur_to(duration.MilliSecond)
  |> int.to_string
  |> string.append("Elapsed time: ", _)
  |> string.append("ms")
  |> io.println
}
