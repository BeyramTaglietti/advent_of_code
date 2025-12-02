import { expect, test } from "bun:test";
import { isValidID, isValidID2, solveP1, solveP2 } from "./solution";

test("p1 test", async () => {
  const input = [
    "11-22,95-115,998-1012,1188511880-1188511890,222220-222224",
    "1698522-1698528,446443-446449,38593856-38593862,565653-565659",
    "824824821-824824827,2121212118-2121212124",
  ];
  expect(solveP1(input)).toEqual(1227775554);
});

test("p2 test", async () => {
  const input = [
    "11-22,95-115,998-1012,1188511880-1188511890,222220-222224",
    "1698522-1698528,446443-446449,38593856-38593862,565653-565659",
    "824824821-824824827,2121212118-2121212124",
  ];
  expect(solveP2(input)).toEqual(4174379265);
});

test("is valid function", () => {
  expect(isValidID(12)).toBeTrue();
  expect(isValidID(92)).toBeTrue();
  expect(isValidID(11)).toBeFalse();
  expect(isValidID(1010)).toBeFalse();
  expect(isValidID(1188511885)).toBeFalse();
  expect(isValidID(38593859)).toBeFalse();
});

test("is valid 2 function", () => {
  expect(isValidID2(12)).toBeTrue();
  expect(isValidID2(110)).toBeTrue();
  expect(isValidID2(1010)).toBeFalse();
  expect(isValidID2(1188511885)).toBeFalse();
  expect(isValidID2(38593859)).toBeFalse();
  expect(isValidID2(565656)).toBeFalse();
  expect(isValidID2(2121212121)).toBeFalse();
  expect(isValidID2(824824824)).toBeFalse();
});
