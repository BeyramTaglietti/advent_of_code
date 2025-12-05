import { expect, test } from "bun:test";
import { solveP1, solveP2 } from "./solution";

test("p1 test", async () => {
  const input: Array<string> = [
    "3-5",
    "10-14",
    "16-20",
    "12-18",
    "",
    "1",
    "5",
    "8",
    "11",
    "17",
    "32",
  ];
  expect(solveP1(input)).toEqual(3);
});

test("p2 test", async () => {
  const input: Array<string> = [
    "3-5",
    "10-14",
    "16-20",
    "12-18",
    "",
    "1",
    "5",
    "8",
    "11",
    "17",
    "32",
  ];
  expect(solveP2(input)).toEqual(14);
});
