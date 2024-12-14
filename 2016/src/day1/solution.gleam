import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub type Direction {
  North
  South
  East
  West
}

pub type Rotation {
  R
  L
}

pub type Command {
  Command(rotation: Rotation, distance: Int)
}

pub type Point {
  Point(x: Int, y: Int, facing: Direction)
}

pub fn solve_p1() {
  let assert Ok(lines) = simplifile.read("./src/day1/input.txt")

  let point_reached =
    string.split(lines, ",")
    |> list.map(string.trim)
    |> list.fold(Point(0, 0, North), fn(acc, cmd) {
      let cmd = parse_cmd(cmd)
      let next_p =
        move(acc, cmd) |> list.last |> result.unwrap(Point(0, 0, North))
      next_p
    })

  int.absolute_value(point_reached.x) + int.absolute_value(point_reached.y)
  |> io.debug
}

pub fn solve_p2() {
  let assert Ok(lines) = simplifile.read("./src/day1/input.txt")

  let first_hq =
    string.split(lines, ",")
    |> list.map(string.trim)
    |> list.map(parse_cmd)
    |> find_first_hq(dict.new(), Point(0, 0, North), _)

  case first_hq {
    Ok(p) -> io.debug(int.absolute_value(p.x) + int.absolute_value(p.y))
    _ -> panic as "No HQ found"
  }
}

fn parse_cmd(s: String) -> Command {
  let rotation = case string.slice(s, 0, 1) {
    "R" -> R
    "L" -> L
    _ -> R
  }

  let distance =
    string.slice(s, 1, string.length(s))
    |> int.parse
    |> result.unwrap(0)

  Command(rotation: rotation, distance: distance)
}

fn move(p: Point, cmd: Command) -> List(Point) {
  case cmd.rotation {
    R -> {
      case p.facing {
        North ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x + x, p.y, East) })
        South ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x - x, p.y, West) })
        East ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x, p.y + x, South) })
        West ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x, p.y - x, North) })
      }
    }
    L -> {
      case p.facing {
        North ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x - x, p.y, West) })
        South ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x + x, p.y, East) })
        East ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x, p.y - x, North) })
        West ->
          list.range(1, cmd.distance)
          |> list.map(fn(x) { Point(p.x, p.y + x, South) })
      }
    }
  }
}

fn find_first_hq(
  visited_spots: dict.Dict(#(Int, Int), Bool),
  current_location: Point,
  commands: List(Command),
) -> Result(Point, Bool) {
  case list.is_empty(commands) {
    True -> Error(False)
    _ -> {
      let assert Ok(next_cmd) = list.first(commands)
      let next_points = move(current_location, next_cmd)

      let #(last_p, visited_locs, found_visited_twice_spot, visited_twice_spot) =
        list.fold(
          next_points,
          #(Point(0, 0, North), visited_spots, False, Point(0, 0, North)),
          fn(acc, p) {
            let #(_, visited_s, found, found_p) = acc
            case dict.has_key(visited_s, #(p.x, p.y)) {
              True -> {
                #(p, visited_s, True, p)
              }
              _ -> {
                let new_visited_spots =
                  dict.insert(visited_s, #(p.x, p.y), True)
                #(p, new_visited_spots, found, found_p)
              }
            }
          },
        )

      case found_visited_twice_spot {
        True -> Ok(visited_twice_spot)
        _ -> find_first_hq(visited_locs, last_p, list.drop(commands, 1))
      }
    }
  }
}
