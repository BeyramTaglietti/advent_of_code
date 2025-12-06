import { expect, test } from "bun:test";
import { solveP1, solveP2 } from "./solution";

test("p1 test", async () => {
  const input: Array<string> = [
    "123 328  51 64",
    "45 64  387 23",
    "6 98  215 314",
    "*   +   *   +",
  ];
  expect(solveP1(input)).toEqual(4277556);
});

test("p2 test", async () => {
  const input: Array<string> = [
    "123 328  51 64 ",
    " 45 64  387 23 ",
    "  6 98  215 314",
    "*   +   *   +  ",
  ];
  expect(solveP2(input)).toEqual(3263827);
});
