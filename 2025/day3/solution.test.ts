import { expect, test } from "bun:test";
import { solveP1, solveP2 } from "./solution";

test("p1 test", async () => {
  const input = [
    "987654321111111",
    "811111111111119",
    "234234234234278",
    "818181911112111",
  ];
  expect(solveP1(input)).toEqual(357);
});

test("p2 test", async () => {
  const input = [
    "987654321111111",
    "811111111111119",
    "234234234234278",
    "818181911112111",
  ];
  expect(solveP2(input)).toEqual(3121910778619);
});
