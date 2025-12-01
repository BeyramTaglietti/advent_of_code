import { expect, test } from "bun:test";
import { solveP1, solveP2 } from "./solution";

test("p1 test", async () => {
  expect(
    solveP1(["1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"])
  ).toEqual(142);
});

test("p2 test", async () => {
  expect(
    solveP2([
      "two1nine",
      "eightwothree",
      "abcone2threexyz",
      "xtwone3four",
      "4nineeightseven2",
      "zoneight234",
      "7pqrstsixteen",
    ])
  ).toEqual(281);
});
