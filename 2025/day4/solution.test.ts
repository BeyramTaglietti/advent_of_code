import { expect, test } from "bun:test";
import { solveP1, solveP2 } from "./solution";

test("p1 test", async () => {
  const input: Array<string> = [
    "..@@.@@@@.",
    "@@@.@.@.@@",
    "@@@@@.@.@@",
    "@.@@@@..@.",
    "@@.@@@@.@@",
    ".@@@@@@@.@",
    ".@.@.@.@@@",
    "@.@@@.@@@@",
    ".@@@@@@@@.",
    "@.@.@@@.@.",
  ];
  expect(solveP1(input)).toEqual(13);
});

test("p2 test", async () => {
  const input: Array<string> = [
    "..@@.@@@@.",
    "@@@.@.@.@@",
    "@@@@@.@.@@",
    "@.@@@@..@.",
    "@@.@@@@.@@",
    ".@@@@@@@.@",
    ".@.@.@.@@@",
    "@.@@@.@@@@",
    ".@@@@@@@@.",
    "@.@.@@@.@.",
  ];
  expect(solveP2(input)).toEqual(43);
});
