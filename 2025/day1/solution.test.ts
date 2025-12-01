import { expect, test } from "bun:test";
import { solveP1, solveP2 } from "./solution";

test("p1 test", async () => {
  const input = [
    "L68",
    "L30",
    "R48",
    "L5",
    "R60",
    "L55",
    "L1",
    "L99",
    "R14",
    "L82",
  ];
  expect(solveP1(input)).toEqual(3);
});

test("p2 test", async () => {
  const input = [
    "L68",
    "L30",
    "R48",
    "L5",
    "R60",
    "L55",
    "L1",
    "L99",
    "R14",
    "L82",
  ];
  expect(solveP2(input)).toEqual(6);
  expect(solveP2(["R1000"])).toEqual(10);
  expect(solveP2(["R50"])).toEqual(1);
  expect(solveP2(["L51"])).toEqual(1);
});
